package rules

type Not struct {
	InnerRule Rule
}

func NewNot(rule Rule) Not {
	return Not{
		InnerRule: rule,
	}
}

func (rule Not) Evaluate(content *string) (bool, error) {
	result, err := rule.InnerRule.Evaluate(content)
	if err != nil {
		return false, err
	}
	return !result, nil
}

func GetNotType(ruleType RuleType) RuleType {
	return "!" + ruleType
}

func (rule Not) GetType() RuleType {
	return GetNotType(rule.InnerRule.GetType())
}
