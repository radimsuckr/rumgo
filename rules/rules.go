package rules

type RuleType string

const (
	RuleTypeContains      = "contains"
	RuleTypeNotContains   = "not-contains"
	RuleTypeXPathContains = "xpath-contains"
)

type Rule interface {
	Evaluate(content *string) (bool, error)
	GetType() RuleType
}
