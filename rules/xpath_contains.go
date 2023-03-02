package rules

import (
	"strings"

	"github.com/antchfx/htmlquery"
)

type XPathContains struct {
	Path  string
	Value string
}

func (rule XPathContains) Evaluate(content *string) bool {
	doc, doc_err := htmlquery.Parse(strings.NewReader(*content))
	if doc_err != nil {
		// TODO: handle errors
	}

	elements, elements_error := htmlquery.QueryAll(doc, rule.Path)
	if elements_error != nil {
		// TODO: handle errors
	}

	contains := false
	for _, el := range elements {
		if strings.Contains(htmlquery.InnerText(el), rule.Value) {
			contains = true
		}
	}

	return contains
}

func (XPathContains) GetType() RuleType {
	return XPATH_CONTAINS_RULE_TYPE
}
