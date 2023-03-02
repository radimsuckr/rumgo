package rules

type XPathContains struct {
	Path  string
	Value string
}

func (XPathContains) Evaluate(content *string) bool {
	return true
}

func (XPathContains) GetType() RuleType {
	return XPATH_CONTAINS_RULE_TYPE
}
