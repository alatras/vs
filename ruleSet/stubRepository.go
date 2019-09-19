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

func (r *stubRuleSetRepository) ListForOrganization(organization string) []RuleSet {
	if rules, ok := r.cache[organization]; ok {
		return rules
	}

	return nil
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

	r.cache = map[string][]RuleSet{
		"1": org1Rules,
		"2": org2Rules,
	}

	return nil
}

func (r *stubRuleSetRepository) fetchCacheForOrganization(organization string) ([]RuleSet, error) {
	if organization == "1" {
		r, err := New("Is greater than 5 and less than 5000", Block, []Metadata{
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

	if organization == "2" {
		r, err := New("Is greater than 500 and less than 1000", Tag, []Metadata{
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
