package rules

type not struct {
	InnerRule rule
}

func newNot(innerRule rule) not {
	return not{
		InnerRule: innerRule,
	}
}

func (rule not) Evaluate(content *string) (bool, error) {
	result, err := rule.InnerRule.Evaluate(content)
	if err != nil {
		return false, err
	}
	return !result, nil
}

func getNotType(ruleType string) string {
	return "!" + ruleType
}

func (rule not) GetType() string {
	return getNotType(rule.InnerRule.GetType())
}
