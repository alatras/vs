package ruleSet

type stubRuleSetRepository struct {
	cache map[string][]RuleSet
}

func NewStubRuleSetRepository() (*stubRuleSetRepository, error) {
	r := &stubRuleSetRepository{
		cache: make(map[string][]RuleSet),
	}

	err := r.reloadCache()

	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r *stubRuleSetRepository) ListForEntity(entity string) []RuleSet {
	if rules, ok := r.cache[entity]; ok {
		return rules
	}

	return nil
}

func (r *stubRuleSetRepository) reloadCache() error {
	entity1Rules, err := r.fetchCacheForEntity("1")

	if err != nil {
		return err
	}

	entity2Rules, err := r.fetchCacheForEntity("2")

	if err != nil {
		return err
	}

	r.cache = map[string][]RuleSet{
		"1": entity1Rules,
		"2": entity2Rules,
	}

	return nil
}

func (r *stubRuleSetRepository) fetchCacheForEntity(entity string) ([]RuleSet, error) {
	if entity == "1" {
		r, err := New("Is greater than 5 and less than 5000", "1", Block, []Metadata{
			{
				Key:      "amount",
				Operator: "<",
				Value:    "5000",
			},
			{
				Key:      "amount",
				Operator: ">",
				Value:    "5",
			},
		})

		if err != nil {
			return nil, err
		}

		return []RuleSet{r}, nil
	}

	if entity == "2" {
		r, err := New("Is greater than 500 and less than 1000", "2", Tag, []Metadata{
			{
				Key:      "amount",
				Operator: "<",
				Value:    "1000",
			},
			{
				Key:      "amount",
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
