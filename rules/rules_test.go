package rules

import "testing"

func TestRuleTypeContainsHasExpectedValue(t *testing.T) {
	expected := "contains"

	if RuleTypeContains != expected {
		t.Errorf("RuleTypeContains has value %s instead of %s", RuleTypeContains, expected)
	}
}

func TestRuleTypeNotContainsHasExpectedValue(t *testing.T) {
	expected := "not-contains"

	if RuleTypeNotContains != expected {
		t.Errorf("RuleTypeNotContains has value %s instead of %s", RuleTypeNotContains, expected)
	}
}

func TestRuleTypeXPathContainsHasExpectedValue(t *testing.T) {
	expected := "xpath-contains"

	if RuleTypeXPathContains != expected {
		t.Errorf("RuleTypeXPathContains has value %s instead of %s", RuleTypeXPathContains, expected)
	}
}
