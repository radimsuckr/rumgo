package rules

import "testing"

func TestContainsGetTypeReturnsCorrectType(t *testing.T) {
	rule := NewContainsRule("random")

	ruletype := rule.GetType()
	expected := RuleTypeContains

	if ruletype != RuleType(expected) {
		t.Errorf("rule type should be %s but is %s", expected, ruletype)
	}
}

func TestContainsEvaluateContainsText(t *testing.T) {
	rule := NewContainsRule("golang")
	text := "the best language of all time is golang, that's for sure"

	got, _ := rule.Evaluate(&text)
	want := true

	if got != true {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestContainsEvaluateDoesNotContainText(t *testing.T) {
	rule := NewContainsRule("golang")
	text := "the best language of all time is rust, that's for sure"

	got, _ := rule.Evaluate(&text)
	want := true

	if got != false {
		t.Errorf("got %t, wanted %t", got, want)
	}
}
