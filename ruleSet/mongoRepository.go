package ruleSet

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"time"
)

type MongoRuleSetRepository struct {
	mongoClient       *mongo.Client
	ruleSetCollection *mongo.Collection
}

func NewMongoRepository(url string, dbName string) (*MongoRuleSetRepository, error) {
	client, err := connectToMongo(url)

	if err != nil {
		return nil, err
	}

	ruleSetCollection := client.Database(dbName).Collection("ruleSets")
	repository := MongoRuleSetRepository{client, ruleSetCollection}

	err = createIndexes(ruleSetCollection)

	if err != nil {
		return nil, err
	}

	return &repository, nil
}

func (r MongoRuleSetRepository) Create(ctx context.Context, ruleSet RuleSet) error {
	_, err := r.ruleSetCollection.InsertOne(ctx, ruleSet)

	if err != nil {
		return errors.New("failed to create rule set in repository")
	}

	return nil
}

func (r MongoRuleSetRepository) GetById(ctx context.Context, entityId string, ruleSetId string) (*RuleSet, error) {
	var ruleSet RuleSet

	err := r.ruleSetCollection.FindOne(context.TODO(), bson.M{
		"entityId": entityId,
		"id":       ruleSetId,
	}).Decode(&ruleSet)

	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("error while getting rule set by id: %v", err)
	}

	return &ruleSet, nil
}

func (r MongoRuleSetRepository) ListByEntityId(ctx context.Context, entityId string) ([]RuleSet, error) {
	cursor, err := r.ruleSetCollection.Find(ctx, bson.M{
		"entityId": entityId,
	})

	if err != nil {
		return nil, errors.New("error while listing rule sets by entity id")
	}

	var ruleSets []RuleSet

	for cursor.Next(ctx) {
		var ruleSet RuleSet
		err := cursor.Decode(&ruleSet)
		if err != nil {
			return nil, errors.New("error while decoding rule set from db")
		}
		ruleSets = append(ruleSets, ruleSet)
	}

	return ruleSets, nil
}

func (r MongoRuleSetRepository) ListByEntityIds(ctx context.Context, entityIds ...string) ([]RuleSet, error) {
	cursor, err := r.ruleSetCollection.Find(ctx, bson.M{
		"entityId": bson.M{
			"$in": entityIds,
		},
	})

	if err != nil {
		return nil, errors.New("error while listing rule sets by entity ids")
	}

	var ruleSets []RuleSet

	for cursor.Next(ctx) {
		var ruleSet RuleSet
		err := cursor.Decode(&ruleSet)
		if err != nil {
			return nil, errors.New("error while decoding rule set from db")
		}
		ruleSets = append(ruleSets, ruleSet)
	}

	return ruleSets, nil
}

func (r MongoRuleSetRepository) Replace(ctx context.Context, entityId string, ruleSet RuleSet) (bool, error) {
	var replaced bool

	result, err := r.ruleSetCollection.ReplaceOne(ctx, bson.M{
		"entityId": entityId,
		"id":       ruleSet.Id,
	}, ruleSet)

	if err != nil {
		return replaced, errors.New("error while replacing rule set")
	}

	if result.ModifiedCount > 0 {
		replaced = true
	}

	return replaced, nil
}

func (r MongoRuleSetRepository) Delete(ctx context.Context, entityId string, ruleSetIds ...string) (bool, error) {
	var deleted bool

	result, err := r.ruleSetCollection.DeleteMany(ctx, bson.M{
		"entityId": entityId,
		"id": bson.M{
			"$in": ruleSetIds,
		},
	})

	if err != nil {
		return deleted, errors.New("error while deleting rule set from db")
	}

	if result.DeletedCount > 0 {
		deleted = true
	}

	return deleted, nil
}

func connectToMongo(url string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))

	if err != nil {
		return nil, errors.New("error while establishing connection to MongoDB")
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, errors.New("error while ensuring connection to MongoDB")
	}

	return client, nil
}

func createIndexes(collection *mongo.Collection) error {
	indexOptions := options.Index()
	indexOptions.SetBackground(true)
	indexOptions.SetName("entity_index")

	_, err := collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bsonx.Doc{{Key: "entity", Value: bsonx.Int32(1)}},
		Options: indexOptions,
	})

	return err
}
