package rules

import "testing"

func TestRuleTypeContainsHasExpectedValue(t *testing.T) {
	expected := "contains"

	if RuleTypeContains != expected {
		t.Errorf("RuleTypeContains has value %s instead of %s", RuleTypeContains, expected)
	}
}

func TestRuleTypeXPathContainsHasExpectedValue(t *testing.T) {
	expected := "xpath-contains"

	if RuleTypeXPathContains != expected {
		t.Errorf("RuleTypeXPathContains has value %s instead of %s", RuleTypeXPathContains, expected)
	}
}
