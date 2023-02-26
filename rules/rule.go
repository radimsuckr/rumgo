package rules

type Rule interface {
	Evaluate() bool
	GetType() string
}
