package ruleSet

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type mongoRuleSetRepository struct {
	mongoClient       *mongo.Client
	ruleSetCollection *mongo.Collection
}

func NewMongoRepository(hostname, port string) (*mongoRuleSetRepository, error) {
	client, err := connectToMongo(hostname, port)
	if err != nil {
		return nil, err
	}

	ruleSetCollection := client.Database("validationService").Collection("ruleSets")
	repository := mongoRuleSetRepository{client, ruleSetCollection}

	return &repository, nil
}

func (r mongoRuleSetRepository) Create(ctx context.Context, ruleSet RuleSet) error {
	_, err := r.ruleSetCollection.InsertOne(ctx, ruleSet)

	if err != nil {
		return errors.New("failed to create rule set in repository")
	}

	return nil
}

func (r mongoRuleSetRepository) GetById(ctx context.Context, entityId string, ruleSetId string) (RuleSet, error) {
	var ruleSet RuleSet

	err := r.ruleSetCollection.FindOne(context.TODO(), bson.D{
		bson.E{
			Key:   "entityId",
			Value: entityId,
		},
		bson.E{
			Key:   "id",
			Value: ruleSetId,
		},
	}).Decode(&ruleSet)

	if err != nil {
		return ruleSet, errors.New("error while getting rule set by id")
	}

	return ruleSet, nil
}

func (r mongoRuleSetRepository) ListByEntityId(ctx context.Context, entityId string) ([]RuleSet, error) {
	cursor, err := r.ruleSetCollection.Find(ctx, bson.D{
		bson.E{
			Key:   "entityId",
			Value: entityId,
		},
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

func (r mongoRuleSetRepository) Replace(ctx context.Context, entityId string, ruleSet RuleSet) (bool, error) {
	var replaced bool

	result, err := r.ruleSetCollection.ReplaceOne(ctx, bson.D{
		bson.E{
			Key:   "entityId",
			Value: entityId,
		},
		bson.E{
			Key:   "id",
			Value: ruleSet.Id,
		},
	}, ruleSet)

	if err != nil {
		return replaced, errors.New("error while replacing rule set")
	}

	if result.ModifiedCount > 0 {
		replaced = true
	}

	return replaced, nil
}

func (r mongoRuleSetRepository) Delete(ctx context.Context, entityId string, ruleSetId string) (bool, error) {
	var deleted bool

	result, err := r.ruleSetCollection.DeleteMany(ctx, bson.D{
		bson.E{
			Key:   "entityId",
			Value: entityId,
		},
		bson.E{
			Key:   "id",
			Value: ruleSetId,
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

func connectToMongo(hostname, port string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+hostname+":"+port))
	if err != nil {
		return nil, errors.New("error while establishing connection to MongoDB")
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, errors.New("error while ensuring connection to MongoDB")
	}

	return client, nil
}
