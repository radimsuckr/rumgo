package rules

import (
	"fmt"
	"testing"
)

func TestXPathContainsGetTypeReturnsCorrectType(t *testing.T) {
	rule := NewXPathContainsRule("abc", "xyz")

	ruletype := rule.GetType()
	expected := RuleTypeXPathContains

	if ruletype != RuleType(expected) {
		t.Errorf("rule type should be %s but is %s", expected, ruletype)
	}
}

func TestXPathContainsEvaluateContainsText(t *testing.T) {
	id := "secret"
	path := "//p[@id=\"" + id + "\"]"
	text := "hello"
	test_str := fmt.Sprintf("<p id=\"%s\">%s</p>", id, text)
	rule := NewXPathContainsRule(path, text)

	got, err := rule.Evaluate(&test_str)
	want := true

	if err != nil {
		t.Fatal("unexpected error")
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestXPathContainsEvaluateDoesNotContainText(t *testing.T) {
	id := "secret"
	path := "//p[@id=\"" + id + "\"]"
	text := "hello"
	test_str := fmt.Sprintf("<p id=\"%s\">%s</p>", id, "madeupvalue")
	rule := NewXPathContainsRule(path, text)

	got, err := rule.Evaluate(&test_str)
	want := false

	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestXPathContainsEvaluateDoesNotContainTextWithWrongXPath(t *testing.T) {
	id := "secret"
	path := "//p[@id=\"" + id + "\"]"
	text := "hello"
	test_str := fmt.Sprintf("<p id=\"%s\">%s</p>", "madeupvalue", text)
	rule := NewXPathContainsRule(path, text)

	got, err := rule.Evaluate(&test_str)
	want := false

	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}
