package rules

import "testing"

func TeststringContainsHasExpectedValue(t *testing.T) {
	expected := "contains"

	if ruleTypeContains != expected {
		t.Errorf("stringContains has value %s instead of %s", ruleTypeContains, expected)
	}
}

func TeststringNotContainsHasExpectedValue(t *testing.T) {
	var expected string = "!contains"

	if val := getNotType(ruleTypeContains); val != expected {
		t.Errorf("Inverted stringContains has value %s instead of %s", val, expected)
	}
}

func TeststringXPathContainsHasExpectedValue(t *testing.T) {
	expected := "xpath-contains"

	if ruleTypeXPathContains != expected {
		t.Errorf("stringXPathContains has value %s instead of %s", ruleTypeXPathContains, expected)
	}
}
