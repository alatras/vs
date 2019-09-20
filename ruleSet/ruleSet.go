package ruleSet

import (
	"bitbucket.verifone.com/validation-service/ruleSet/rule"
	"bitbucket.verifone.com/validation-service/transaction"
	"context"
	"github.com/google/uuid"
)

type Action int

const (
	Pass Action = iota
	Block
	Tag
)

type RuleSet struct {
	Id           string          `json:"id" bson:"id"`
	EntityId     string          `json:"entityId" bson:"entityId"`
	Action       Action          `bson:"action"`
	Name         string          `json:"name" bson:"name"`
	RuleMetadata []rule.Metadata `json:"rules" bson:"validationRuleMetadata"`
}

type Repository interface {
	Create(ctx context.Context, ruleSet RuleSet) error
	GetById(ctx context.Context, entityId string, ruleSetId string) (RuleSet, error)
	ListByEntityId(ctx context.Context, entityId string) ([]RuleSet, error)
	Replace(ctx context.Context, entityId string, ruleSet RuleSet) (bool, error)
	Delete(ctx context.Context, entityId string, ruleSetId string) (bool, error)
}

func New(entityId string, name string, action Action, metadata []rule.Metadata) (RuleSet, error) {
	ruleSet := RuleSet{
		Id:           uuid.New().String(),
		EntityId:     entityId,
		Name:         name,
		Action:       action,
		RuleMetadata: metadata,
	}

	return ruleSet, nil
}

func (ruleSet RuleSet) Matches(trx transaction.Transaction) (Action, error) {
	if len(ruleSet.RuleMetadata) == 0 {
		return Pass, nil
	}

	for _, metadata := range ruleSet.RuleMetadata {
		validator, err := rule.NewValidator(metadata)
		if err != nil {
			return Pass, err
		}

		if !validator.Validate(trx) {
			return Pass, nil
		}
	}

	return ruleSet.Action, nil
}
