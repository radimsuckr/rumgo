package rules

import "strings"

type ContainsRule struct {
	Value string
}

func NewContainsRule(value string) *ContainsRule {
	return &ContainsRule{
		Value: value,
	}
}

func (rule ContainsRule) Evaluate(content *string) (*bool, error) {
	result := strings.Contains(*content, rule.Value)
	return &result, nil
}

func (ContainsRule) GetType() RuleType {
	return RuleTypeContains
}
