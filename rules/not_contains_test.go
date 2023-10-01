package rules

import "testing"

func TestNotContainsGetTypeReturnsCorrectType(t *testing.T) {
	rule := NewNotContainsRule("random")

	ruletype := rule.GetType()
	expected := RuleTypeNotContains

	if ruletype != RuleType(expected) {
		t.Errorf("rule type should be %s but is %s", expected, ruletype)
	}
}

func TestNotContainsEvaluateContainsText(t *testing.T) {
	rule := NewNotContainsRule("golang")
	text := "the best language of all time is golang, that's for sure"

	got, _ := rule.Evaluate(&text)
	want := true

	if got != false {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestNotContainsEvaluateDoesNotContainText(t *testing.T) {
	rule := NewNotContainsRule("golang")
	text := "the best language of all time is rust, that's for sure"

	got, _ := rule.Evaluate(&text)
	want := true

	if got != true {
		t.Errorf("got %t, wanted %t", got, want)
	}
}
