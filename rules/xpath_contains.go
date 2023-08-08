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

func (rule XPathContains) Evaluate(content *string) (bool, error) {
	doc, err := htmlquery.Parse(strings.NewReader(*content))
	if err != nil {
		return false, err
	}

	elements, err := htmlquery.QueryAll(doc, rule.Path)
	if err != nil {
		return false, err
	}

	contains := false
	for _, el := range elements {
		if strings.Contains(htmlquery.InnerText(el), rule.Value) {
			contains = true
		}
	}

	return contains, nil
}

func (XPathContains) GetType() RuleType {
	return RuleTypeXPathContains
}
