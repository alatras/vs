package test

import (
	"bitbucket.verifone.com/validation-service/ruleSet"
	"bitbucket.verifone.com/validation-service/ruleSet/rule"
)

const (
	MalformedParametersErrorMessage = "At least one parameter is invalid. Examine the details property for more information. Invalid parameters are listed and prefixed accordingly: body for parameters submitted in the request's body, query for parameters appended to the request's URL, and params for templated parameters of the request's URL."
	UnexpectedErrorMessage          = "Unexpected error: if the error persists, please contact an administrator, quoting the code and timestamp of this error"
	NotFoundErrorMessage            = `The requested resource, or one of its sub-resources, can't be 
found. If the submitted query is valid, this error is likely to be caused by a problem with a nested 
resource that has been deleted or modified. Check the details property for additional insights.`
)

var (
	MockRuleSet = ruleSet.New("12345", "Test", ruleSet.Block, []rule.Metadata{
		{
			Property: "amount",
			Operator: ">=",
			Value:    "1000",
		},
	})
)
