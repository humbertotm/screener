package domain

import "context"

// ProfileRepository defines the interface employed to interact with the profiles db
type ProfileRepository interface {
	GetFullCIKList(ctx context.Context) ([]interface{}, error)
	GetFullProfileForCIK(ctx context.Context, cik string) (FullCompanyProfile, error)
	GetFullProfileForTicker(ctx context.Context, ticker string) (FullCompanyProfile, error)
}

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

// FullCompanyProfile defines a company's profile for a range of years
type FullCompanyProfile []*YearlyProfile
