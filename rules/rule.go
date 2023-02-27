package rules

type RuleType string

const (
	CONTAINS_RULE_TYPE       = "contains"
	XPATH_CONTAINS_RULE_TYPE = "xpath-contains"
)

type Rule interface {
	Evaluate() bool
	GetType() RuleType
}
