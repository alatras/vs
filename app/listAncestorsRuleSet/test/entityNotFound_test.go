package test

import (
	"bitbucket.verifone.com/validation-service/app/listAncestorsRuleSet"
	"bitbucket.verifone.com/validation-service/entityService"
	"bitbucket.verifone.com/validation-service/logger"
	"context"
	"testing"
)

func Test_App_ListAncestorsRuleSet_EntityNotFound(t *testing.T) {
	log := logger.NewStubLogger()
	repo := stubRepository{}
	entityService := stubEntityService{
		err: entityService.EntityNotFound,
	}

	app := listAncestorsRuleSet.NewListAncestorsRuleSet(log, &repo, &entityService)

	_, err := app.Execute(
		context.TODO(),
		"123",
	)

	if !err.HasError() {
		t.Error("listing ancestors RuleSet succeeded but should fail with not found error")
	} else if !err.Is(listAncestorsRuleSet.EntityIdNotFoundErr) {
		t.Errorf("listing ancestors RuleSet failed but not with not found error: %v", err)
	}
}
