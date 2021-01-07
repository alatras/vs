package ruleSet

import (
	"context"
	"fmt"
	"time"

	appd "validation-service/appdynamics"
	"validation-service/enums/appdBackend"
	"validation-service/enums/contextKey"
	"validation-service/logger"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

var retryableCodes = []int32{11600, 11602, 10107, 13435, 13436, 189, 91, 7, 6, 89, 9001, 262}

const networkErrorLabel = "NetworkError"

type MongoRuleSetRepository struct {
	mongoClient       *mongo.Client
	ruleSetCollection *mongo.Collection
	logger            *logger.Logger
	retryDelay        time.Duration
}

func NewMongoRepository(url string, dbName string, retryDelay time.Duration, logger *logger.Logger) (*MongoRuleSetRepository, error) {
	client, err := connectToMongo(url, logger)

	if err != nil {
		return nil, err
	}

	ruleSetCollection := client.Database(dbName).Collection("ruleSets")
	repository := MongoRuleSetRepository{
		mongoClient:       client,
		ruleSetCollection: ruleSetCollection,
		logger:            logger.Scoped("MongoRuleSetRepository"),
		retryDelay:        retryDelay,
	}

	err = createIndexes(ruleSetCollection)

	if err != nil {
		return nil, err
	}

	return &repository, nil
}

func (r MongoRuleSetRepository) Create(ctx context.Context, ruleSet RuleSet) error {
	var businessTransaction appd.BtHandle

	if businessTransactionUid, ok := ctx.Value(contextKey.BusinessTransaction).(string); ok {
		businessTransaction = appd.GetBT(businessTransactionUid)
	}

	exitCallHandle := appd.StartExitcall(businessTransaction, string(appdBackend.MongoDB))
	_ = appd.SetExitcallDetails(exitCallHandle, "InsertOne")

	defer appd.EndExitcall(exitCallHandle)

	_, err := r.ruleSetCollection.InsertOne(ctx, ruleSet)

	if err != nil {
		return fmt.Errorf("failed to create rule set in repository: %w", err)
	}

	return nil
}

func (r MongoRuleSetRepository) GetById(ctx context.Context, entityId string, ruleSetId string) (*RuleSet, error) {
	query := bson.M{
		"entityId": entityId,
		"id":       ruleSetId,
	}

	var businessTransaction appd.BtHandle

	if businessTransactionUid, ok := ctx.Value(contextKey.BusinessTransaction).(string); ok {
		businessTransaction = appd.GetBT(businessTransactionUid)
	}

	exitCallHandle := appd.StartExitcall(businessTransaction, string(appdBackend.MongoDB))
	_ = appd.SetExitcallDetails(exitCallHandle, fmt.Sprintf("FindOne: %s", query))
	defer appd.EndExitcall(exitCallHandle)

	var ruleSet RuleSet

	var err error

	retryErr := backOffRetryWithContext(ctx, r.retryDelay, func() (retry bool) {
		err = r.ruleSetCollection.FindOne(ctx, query).Decode(&ruleSet)
		return isRetryableError(err)
	})

	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	if retryErr != nil && err == nil {
		err = retryErr
	}

	if err != nil {
		return nil, fmt.Errorf("error while getting rule set by id: %v", err)
	}

	return &ruleSet, nil
}

func (r MongoRuleSetRepository) ListByEntityIds(ctx context.Context, entityIds ...string) ([]RuleSet, error) {
	query := bson.M{
		"entityId": bson.M{
			"$in": entityIds,
		},
	}

	var businessTransaction appd.BtHandle

	if businessTransactionUid, ok := ctx.Value(contextKey.BusinessTransaction).(string); ok {
		businessTransaction = appd.GetBT(businessTransactionUid)
	}

	exitCallHandle := appd.StartExitcall(businessTransaction, string(appdBackend.MongoDB))
	_ = appd.SetExitcallDetails(exitCallHandle, fmt.Sprintf("Find: %s", query))
	defer appd.EndExitcall(exitCallHandle)

	var err error
	var cursor *mongo.Cursor

	retryErr := backOffRetryWithContext(ctx, r.retryDelay, func() (retry bool) {
		cursor, err = r.ruleSetCollection.Find(ctx, query)
		return isRetryableError(err)
	})

	if retryErr != nil && err == nil {
		err = retryErr
	}

	if err != nil {
		return nil, fmt.Errorf("error while listing rule sets by entity ids: %w", err)
	}

	var ruleSets []RuleSet

	for cursor.Next(ctx) {
		var ruleSet RuleSet
		err := cursor.Decode(&ruleSet)

		if err != nil {
			return nil, fmt.Errorf("error while decoding rule set from db: %w", err)
		}
		ruleSets = append(ruleSets, ruleSet)
	}

	return ruleSets, nil
}

func (r MongoRuleSetRepository) Replace(ctx context.Context, entityId string, ruleSet RuleSet) (bool, error) {
	query := bson.M{
		"entityId": entityId,
		"id":       ruleSet.Id,
	}

	var businessTransaction appd.BtHandle

	if businessTransactionUid, ok := ctx.Value(contextKey.BusinessTransaction).(string); ok {
		businessTransaction = appd.GetBT(businessTransactionUid)
	}

	exitCallHandle := appd.StartExitcall(businessTransaction, string(appdBackend.MongoDB))
	_ = appd.SetExitcallDetails(exitCallHandle, fmt.Sprintf("ReplaceOne: %s", query))
	defer appd.EndExitcall(exitCallHandle)

	var replaced bool

	result, err := r.ruleSetCollection.ReplaceOne(ctx, query, ruleSet)

	if err != nil {
		return replaced, fmt.Errorf("error while replacing rule set: %w", err)
	}

	if result.ModifiedCount > 0 {
		replaced = true
	}

	return replaced, nil
}

func (r MongoRuleSetRepository) Delete(ctx context.Context, entityId string, ruleSetIds ...string) (bool, error) {
	query := bson.M{
		"entityId": entityId,
		"id": bson.M{
			"$in": ruleSetIds,
		},
	}

	var businessTransaction appd.BtHandle

	if businessTransactionUid, ok := ctx.Value(contextKey.BusinessTransaction).(string); ok {
		businessTransaction = appd.GetBT(businessTransactionUid)
	}

	exitCallHandle := appd.StartExitcall(businessTransaction, string(appdBackend.MongoDB))
	_ = appd.SetExitcallDetails(exitCallHandle, fmt.Sprintf("DeleteMany: %s", query))
	defer appd.EndExitcall(exitCallHandle)

	var deleted bool

	result, err := r.ruleSetCollection.DeleteMany(ctx, query)

	if err != nil {
		return deleted, fmt.Errorf("error while deleting rule set from db: %w", err)
	}

	if result.DeletedCount > 0 {
		deleted = true
	}

	return deleted, nil
}

func connectToMongo(url string, logger *logger.Logger) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url).SetRetryWrites(true))

	if err != nil {
		logger.Error.WithError(err).Error("failed to establish MongoDB connection")
		return nil, fmt.Errorf("error while establishing connection to MongoDB: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		logger.Error.WithError(err).Error("failed to ping MongoDB")
		return nil, fmt.Errorf("error while ensuring connection to MongoDB: %w", err)
	}

	return client, nil
}

func (r MongoRuleSetRepository) Ping(ctx context.Context) error {
	var businessTransaction appd.BtHandle

	if businessTransactionUid, ok := ctx.Value(contextKey.BusinessTransaction).(string); ok {
		businessTransaction = appd.GetBT(businessTransactionUid)
	}

	exitCallHandle := appd.StartExitcall(businessTransaction, string(appdBackend.MongoDB))
	_ = appd.SetExitcallDetails(exitCallHandle, "Ping")
	defer appd.EndExitcall(exitCallHandle)

	log := r.logger.Output.WithFields(logrus.Fields{
		"method": "Ping",
	})

	errorLog := r.logger.Error.WithFields(logrus.Fields{
		"method": "Ping",
	})

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	err := r.mongoClient.Ping(ctx, nil)

	if err != nil {
		errorLog.WithError(err).Error("mongo is not healthy")
		return fmt.Errorf("cannot connect to mongo: %w", err)
	}

	log.Trace("mongo is healthy")

	return nil
}

func createIndexes(collection *mongo.Collection) error {
	indexOptions := options.Index()
	indexOptions.SetName("entity_index")

	_, err := collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bsonx.Doc{{Key: "entity", Value: bsonx.Int32(1)}},
		Options: indexOptions,
	})

	return err
}

func isRetryableError(err error) bool {
	if err == nil {
		return false
	}

	cmdErr, ok := err.(mongo.CommandError)

	if !ok {
		return false
	}

	if cmdErr.HasErrorLabel(networkErrorLabel) {
		return true
	}

	for _, code := range retryableCodes {
		if cmdErr.Code == code {
			return true
		}
	}

	return false
}
