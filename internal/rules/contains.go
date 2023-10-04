package rules

import "strings"

type containsRule struct {
	Value string
}

func newContainsRule(value string) containsRule {
	return containsRule{
		Value: value,
	}
}

func (rule containsRule) Evaluate(content *string) (bool, error) {
	result := strings.Contains(*content, rule.Value)
	return result, nil
}

func (containsRule) GetType() string {
	return ruleTypeContains
}
