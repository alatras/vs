package test

import (
	"bitbucket.verifone.com/validation-service/app/listDescendantsRuleSet"
	"bitbucket.verifone.com/validation-service/entityService"
	"bitbucket.verifone.com/validation-service/logger"
	"context"
	"testing"
)

func Test_App_ListDescendantsRuleSet_EntityIdFormatIncorrect(t *testing.T) {
	log := logger.NewStubLogger()
	repo := stubRepository{}
	entityService := stubEntityService{
		err: entityService.EntityIdFormatIncorrect,
	}

	app := listDescendantsRuleSet.NewListDescendantsRuleSet(log, &repo, &entityService)

	_, err := app.Execute(
		context.TODO(),
		"123",
	)

	if !err.HasError() {
		t.Error("listing descendants RuleSet succeeded but should fail with entity id format incorrect error")
	} else if !err.Is(listDescendantsRuleSet.EntityIdFormatIncorrectErr) {
		t.Errorf("listing descendants RuleSet failed but not with entity id format incorrect error: %v", err)
	}
}
