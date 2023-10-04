// Package rules implements all available rules for response handling
package rules

const (
	ruleTypeContains      string = "contains"
	ruleTypeXPathContains string = "xpath-contains"
)

type rule interface {
	Evaluate(content *string) (bool, error)
	GetType() string
}
