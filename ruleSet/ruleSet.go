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
	Id             string          `json:"id" bson:"id"`
	EntityId       string          `json:"entityId" bson:"entityId"`
	Action         Action          `bson:"action"`
	Name           string          `json:"name" bson:"name"`
	RuleMetadata   []rule.Metadata `json:"rules" bson:"validationRuleMetadata"`
	ruleValidators []rule.Validator
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
		Id:             uuid.New().String(),
		EntityId:       entityId,
		Action:         action,
		Name:           name,
		RuleMetadata:   metadata,
		ruleValidators: []rule.Validator{},
	}

	for _, m := range metadata {
		v, err := rule.NewValidator(m)
		if err != nil {
			return ruleSet, err
		}
		ruleSet.ruleValidators = append(ruleSet.ruleValidators, v)
	}

	return ruleSet, nil
}

func (r RuleSet) IsMatch(trx transaction.Transaction) Action {
	for _, validator := range r.ruleValidators {
		if !validator.IsMatch(trx) {
			return Pass
		}
	}

	return r.Action
}
