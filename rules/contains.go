package rules

type ContainsRule struct {
	Value string
}

func (ContainsRule) Evaluate() bool {
	return true
}

func (ContainsRule) GetType() RuleType {
	return CONTAINS_RULE_TYPE
}
