package ruleSet

import (
	"context"
	"sync"
)

type StubRuleSetRepository struct {
	cache map[string]map[string]RuleSet
	lock  *sync.RWMutex
}

func NewStubRepository() (*StubRuleSetRepository, error) {
	return &StubRuleSetRepository{
		cache: make(map[string]map[string]RuleSet),
		lock:  &sync.RWMutex{},
	}, nil
}

func (r *StubRuleSetRepository) Create(ctx context.Context, ruleSet RuleSet) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	if _, ok := r.cache[ruleSet.EntityId]; !ok {
		r.cache[ruleSet.EntityId] = make(map[string]RuleSet)
	}
	r.cache[ruleSet.EntityId][ruleSet.Id] = ruleSet

	return nil
}

func (r *StubRuleSetRepository) GetById(ctx context.Context, entityId string, ruleSetId string) (*RuleSet, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	entityMap, ok := r.cache[entityId]

	if !ok {
		return nil, nil
	}

	if cachedRuleSet, ok := entityMap[ruleSetId]; ok {
		return &cachedRuleSet, nil
	} else {
		return nil, nil
	}
}

func (r *StubRuleSetRepository) ListByEntityId(ctx context.Context, entityId string) ([]RuleSet, error) {
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

func (r *StubRuleSetRepository) ListByEntityIds(ctx context.Context, entityIds ...string) ([]RuleSet, error) {
	var ruleSets []RuleSet

	r.lock.RLock()
	defer r.lock.RUnlock()

	for _, entityId := range entityIds {
		if cachedRuleSetMap, ok := r.cache[entityId]; ok {
			for _, ruleSet := range cachedRuleSetMap {
				ruleSets = append(ruleSets, ruleSet)
			}
		}
	}

	return ruleSets, nil
}

func (r *StubRuleSetRepository) Replace(ctx context.Context, entityId string, ruleSet RuleSet) (bool, error) {
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

func (r *StubRuleSetRepository) Delete(ctx context.Context, entityId string, ruleSetIds ...string) (bool, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	entityMap, ok := r.cache[entityId]

	if !ok {
		return false, nil
	}

	var ruleSetsExist bool

	for _, ruleSetId := range ruleSetIds {
		_, ruleSetsExist = entityMap[ruleSetId]
	}

	if !ruleSetsExist {
		return false, nil
	}

	for _, ruleSetId := range ruleSetIds {
		delete(entityMap, ruleSetId)
	}

	return true, nil
}
