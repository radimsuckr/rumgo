package rules

type NotContainsRule struct {
	Rule ContainsRule
}

func NewNotContainsRule(value string) NotContainsRule {
	return NotContainsRule{
		Rule: NewContainsRule(value),
	}
}

func (rule NotContainsRule) Evaluate(content *string) (bool, error) {
	result, err := rule.Rule.Evaluate(content)
	if err != nil {
		return false, err
	}
	return !result, nil
}

func (NotContainsRule) GetType() RuleType {
	return RuleTypeNotContains
}
