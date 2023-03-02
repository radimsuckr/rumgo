package rules

import "strings"

type ContainsRule struct {
	Value string
}

func (rule ContainsRule) Evaluate(content *string) (*bool, error) {
	result := strings.Contains(*content, rule.Value)
	return &result, nil
}

func (ContainsRule) GetType() RuleType {
	return CONTAINS_RULE_TYPE
}
