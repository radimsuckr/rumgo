package rules

type RuleType string

const (
	RuleTypeContains      = "contains"
	RuleTypeXPathContains = "xpath-contains"
)

type Rule interface {
	Evaluate(content *string) (*bool, error)
	GetType() RuleType
}
