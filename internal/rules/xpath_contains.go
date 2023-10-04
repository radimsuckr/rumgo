package rules

import (
	"strings"

	"github.com/antchfx/htmlquery"
)

type xPathContains struct {
	Path  string
	Value string
}

func newXPathContainsRule(path string, value string) *xPathContains {
	return &xPathContains{
		Path:  path,
		Value: value,
	}
}

func (rule xPathContains) Evaluate(content *string) (bool, error) {
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
			return contains, nil
		}
	}

	return contains, nil
}

func (xPathContains) GetType() string {
	return ruleTypeXPathContains
}
