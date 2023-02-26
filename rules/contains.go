package rules

type ContainsRule struct {
	Value string
}

func (ContainsRule) Evaluate() bool {
	return true
}

func (ContainsRule) GetType() string {
	return "contains"
}
