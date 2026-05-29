package models

type Asset struct {
	Type      string
	Value     string
	Source    string
	Metadata  map[string]string
	RiskScore int
}

type Result struct {
	Assets []Asset
	Errors []error
	Source string
}
