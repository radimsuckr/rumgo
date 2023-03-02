package rules

import "strings"

type ContainsRule struct {
	Value string
}

func (rule ContainsRule) Evaluate(content *string) bool {
	return strings.Contains(*content, rule.Value)
}

func (ContainsRule) GetType() RuleType {
	return CONTAINS_RULE_TYPE
}
