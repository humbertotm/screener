package domain

// FinancialProfile is employed to unmarshall the financial data contained in
// YearlyProfile.Profile
// A type *float64 is employed for values in map as it is important to know a value is null
// as opposed to using its zero value.
type FinancialProfile map[string]*float64

// YearlyProfile defines the structure of the documents pulled out of the profiles collection
type YearlyProfile struct {
	CIK     string           `json:"cik"`
	Ticker  string           `json:"ticker"`
	Year    int              `json:"year"`
	Profile FinancialProfile `json:"profile"`
}
