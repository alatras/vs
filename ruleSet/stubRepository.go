package ruleSet

import (
	"bitbucket.verifone.com/validation-service/entity"
	"bitbucket.verifone.com/validation-service/ruleSet/rule"
)

type stubRuleSetRepository struct {
	cache map[entity.Id][]RuleSet
}

func NewStubRepository() (*stubRuleSetRepository, error) {
	r := &stubRuleSetRepository{
		cache: make(map[entity.Id][]RuleSet),
	}

	err := r.reloadCache()

	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r *stubRuleSetRepository) ListForEntity(entityId entity.Id) ([]RuleSet, error) {
	if rules, ok := r.cache[entityId]; ok {
		return rules, nil
	}

	return nil, nil
}

func (r *stubRuleSetRepository) reloadCache() error {
	org1Rules, err := r.fetchCacheForOrganization("1")

	if err != nil {
		return err
	}

	org2Rules, err := r.fetchCacheForOrganization("2")

	if err != nil {
		return err
	}

	r.cache = map[entity.Id][]RuleSet{
		"1": org1Rules,
		"2": org2Rules,
	}

	return nil
}

func (r *stubRuleSetRepository) fetchCacheForOrganization(organization string) ([]RuleSet, error) {
	if organization == "1" {
		r, err := New("1", "Is greater than 5 and less than 5000", Block, []rule.Metadata{
			{
				Property: "amount",
				Operator: "<",
				Value:    "5000",
			},
			{
				Property: "amount",
				Operator: ">",
				Value:    "5",
			},
		})

		if err != nil {
			return nil, err
		}

		return []RuleSet{r}, nil
	}

	if organization == "2" {
		r, err := New("2", "Is greater than 500 and less than 1000", Tag, []rule.Metadata{
			{
				Property: "amount",
				Operator: "<",
				Value:    "1000",
			},
			{
				Property: "amount",
				Operator: ">",
				Value:    "500",
			},
		})

		if err != nil {
			return nil, err
		}

		return []RuleSet{r}, nil
	}

	return []RuleSet{}, nil
}
