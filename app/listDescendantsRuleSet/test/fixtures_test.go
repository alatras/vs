package test

import (
	"bitbucket.verifone.com/validation-service/ruleSet"
	"bitbucket.verifone.com/validation-service/ruleSet/rule"
)

var mockRuleSets = [1]ruleSet.RuleSet{ruleSet.New(
	"12345",
	"Test",
	ruleSet.Block,
	[]rule.Metadata{
		{
			Property: "amount",
			Operator: ">=",
			Value:    "1000",
		},
	},
)}

type stubEntityService struct {
	err error
}

func (s *stubEntityService) Ping() error {
	return nil
}

func (s *stubEntityService) GetAncestorsOf(entityId string) ([]string, error) {
	return []string{entityId}, s.err
}

func (s *stubEntityService) GetDescendantsOf(entityId string) ([]string, error) {
	return []string{entityId}, s.err
}
