package rules

import (
	"fmt"
	"testing"
)

func TestXPathContainsGetTypeReturnsCorrectType(t *testing.T) {
	rule := newXPathContainsRule("abc", "xyz")

	ruletype := rule.GetType()
	expected := ruleTypeXPathContains

	if ruletype != string(expected) {
		t.Errorf("rule type should be %s but is %s", expected, ruletype)
	}
}

func TestXPathContainsEvaluateContainsText(t *testing.T) {
	id := "secret"
	path := "//p[@id=\"" + id + "\"]"
	text := "hello"
	testStr := fmt.Sprintf("<p id=\"%s\">%s</p>", id, text)
	rule := newXPathContainsRule(path, text)

	got, err := rule.Evaluate(&testStr)
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
	testStr := fmt.Sprintf("<p id=\"%s\">%s</p>", id, "madeupvalue")
	rule := newXPathContainsRule(path, text)

	got, err := rule.Evaluate(&testStr)
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
	testStr := fmt.Sprintf("<p id=\"%s\">%s</p>", "madeupvalue", text)
	rule := newXPathContainsRule(path, text)

	got, err := rule.Evaluate(&testStr)
	want := false

	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}
