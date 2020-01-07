package test

import (
	"bitbucket.verifone.com/validation-service/app/listAncestorsRuleSet"
	"bitbucket.verifone.com/validation-service/entityService"
	"bitbucket.verifone.com/validation-service/logger"
	"context"
	"testing"
)

func Test_App_ListAncestorsRuleSet_EntityIdFormatIncorrect(t *testing.T) {
	log := logger.NewStubLogger()
	repo := stubRepository{}
	entityService := stubEntityService{
		err: entityService.EntityIdFormatIncorrect,
	}

	app := listAncestorsRuleSet.NewListAncestorsRuleSet(log, &repo, &entityService)

	_, err := app.Execute(
		context.TODO(),
		"123",
	)

	if !err.HasError() {
		t.Error("listing ancestors RuleSet succeeded but should fail with entity id format incorrect error")
	} else if !err.Is(listAncestorsRuleSet.EntityIdFormatIncorrectErr) {
		t.Errorf("listing ancestors RuleSet failed but not with entity id format incorrect error: %v", err)
	}
}
