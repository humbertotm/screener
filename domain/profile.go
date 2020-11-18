package domain

// FinancialProfile defines the type for the map of ratio/measure - values for a company in a
// given year
type FinancialProfile map[string]float64

// YearlyProfile defines the type for a yearly profile as pulled from the Mongo DB collection
type YearlyProfile struct {
	ObjectID string
	Year     int
	CIK      string
	Ticker   string
	Profile  *FinancialProfile
}

// FullCompanyProfile defines a company's profile for a range of years
type FullCompanyProfile []*YearlyProfile

type ProfileRepository interface {
	GetFullCIKList() ([]string, error)
	GetFullProfileForCIK(cik string) (*FullCompanyProfile, error)
	GetFullProfileForTicker(ticker string) (*FullCompanyProfile, error)
}
