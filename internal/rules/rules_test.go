package rules

import "testing"

func TestRuleTypeContainsHasExpectedValue(t *testing.T) {
	expected := "contains"

	if RuleTypeContains != expected {
		t.Errorf("RuleTypeContains has value %s instead of %s", RuleTypeContains, expected)
	}
}

func TestRuleTypeNotContainsHasExpectedValue(t *testing.T) {
	var expected RuleType = "!contains"

	if val := GetNotType(RuleTypeContains); val != expected {
		t.Errorf("Inverted RuleTypeContains has value %s instead of %s", val, expected)
	}
}

func TestRuleTypeXPathContainsHasExpectedValue(t *testing.T) {
	expected := "xpath-contains"

	if RuleTypeXPathContains != expected {
		t.Errorf("RuleTypeXPathContains has value %s instead of %s", RuleTypeXPathContains, expected)
	}
}
