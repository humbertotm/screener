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

// CompanyStats defines the time series analysis output for a specific company
type CompanyStats struct{}

// FullCompanyProfile defines a company's profile for a range of years
type FullCompanyProfile []*YearlyProfile

// ProfileService defines the business use cases for this application
type ProfileService interface {
	FindCandidateCompanies() ([]string, error)
	GetStatsForCIK(cik string) (*CompanyStats, error)
}

// ProfileRepository defines the interface employed to interact with the profiles db
type ProfileRepository interface {
	GetFullCIKList() ([]string, error)
	GetFullProfileForCIK(cik string) (*FullCompanyProfile, error)
	GetFullProfileForTicker(ticker string) (*FullCompanyProfile, error)
}

// IsAnalyzable determines if a FullCompanyProfile has enough data to proceed with analysis
func (p *FullCompanyProfile) IsAnalyzable() bool {}

// HasBeenProfitableForNPastYears checks if company has been invariably profitable during the
// last n past consecutive years
func (p *FullCompanyProfile) HasBeenProfitableForNPastYears(n int) bool {}

// CalculateRegressionForDescriptor computes the best fitting line for specified descriptor
func (p *FullCompanyProfile) CalculateRegressionForDescriptor(descriptor string) (alpha, beta float64, err error) {
}

// CalculateEPSCompoundGrowthRate returns the annualized compound rate for EPS
func (p *FullCompanyProfile) CalculateEPSCompoundGrowthRate() (float64, error) {}

// CalculateAverageForDescriptor calculates the mean value for specified descriptor
func (p *FullCompanyProfile) CalculateAverageForDescriptor(descriptor string) (float64, error) {}
