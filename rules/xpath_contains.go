package rules

type XPathContains struct {
	Path  string
	Value string
}

func (XPathContains) Evaluate() bool {
	return true
}

func (XPathContains) GetType() string {
	return "xpath-contains"
}
