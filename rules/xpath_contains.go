package rules

import (
	"strings"

	"github.com/antchfx/htmlquery"
)

type XPathContains struct {
	Path  string
	Value string
}

func NewXPathContainsRule(path string, value string) *XPathContains {
	return &XPathContains{
		Path:  path,
		Value: value,
	}
}

func (rule XPathContains) Evaluate(content *string) (*bool, error) {
	doc, doc_err := htmlquery.Parse(strings.NewReader(*content))
	if doc_err != nil {
		return nil, doc_err
	}

	elements, elements_error := htmlquery.QueryAll(doc, rule.Path)
	if elements_error != nil {
		return nil, elements_error
	}

	contains := false
	for _, el := range elements {
		if strings.Contains(htmlquery.InnerText(el), rule.Value) {
			contains = true
		}
	}

	return &contains, nil
}

func (XPathContains) GetType() RuleType {
	return RuleTypeXPathContains
}
