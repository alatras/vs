package ruleSet

import (
	"context"
	"sync"
)

type stubRuleSetRepository struct {
	cache map[string]map[string]RuleSet
	lock  *sync.RWMutex
}

func NewStubRepository() (*stubRuleSetRepository, error) {
	return &stubRuleSetRepository{
		cache: make(map[string]map[string]RuleSet),
		lock:  &sync.RWMutex{},
	}, nil
}

func (r *stubRuleSetRepository) Create(ctx context.Context, ruleSet RuleSet) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	if _, ok := r.cache[ruleSet.EntityId]; !ok {
		r.cache[ruleSet.EntityId] = make(map[string]RuleSet)
	}
	r.cache[ruleSet.EntityId][ruleSet.Id] = ruleSet

	return nil
}

func (r *stubRuleSetRepository) GetById(ctx context.Context, entityId string, ruleSetId string) (RuleSet, error) {
	var ruleSet RuleSet

	r.lock.RLock()
	defer r.lock.RUnlock()

	entityMap, ok := r.cache[entityId]
	if !ok {
		return ruleSet, nil
	}

	if cachedRuleSet, ok := entityMap[ruleSetId]; ok {
		ruleSet = cachedRuleSet
	}

	return ruleSet, nil
}

func (r *stubRuleSetRepository) ListByEntityId(ctx context.Context, entityId string) ([]RuleSet, error) {
	var ruleSets []RuleSet

	r.lock.RLock()
	defer r.lock.RUnlock()

	if cachedRuleSetMap, ok := r.cache[entityId]; ok {
		for _, ruleSet := range cachedRuleSetMap {
			ruleSets = append(ruleSets, ruleSet)
		}
	}

	return ruleSets, nil
}

func (r *stubRuleSetRepository) Replace(ctx context.Context, entityId string, ruleSet RuleSet) (bool, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	entityMap, ok := r.cache[entityId]

	if !ok {
		return false, nil
	}

	_, ok = entityMap[ruleSet.Id]

	if !ok {
		return false, nil
	}

	entityMap[ruleSet.Id] = ruleSet

	return true, nil
}

func (r *stubRuleSetRepository) Delete(ctx context.Context, entityId string, ruleSetId string) (bool, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	entityMap, ok := r.cache[entityId]

	if !ok {
		return false, nil
	}

	_, ok = entityMap[ruleSetId]

	if !ok {
		return false, nil
	}

	delete(entityMap, ruleSetId)

	return true, nil
}
