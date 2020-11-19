package domain

// ScreenerService defines the business use cases for this application
type ScreenerService interface {
	FindCandidateCompanies() ([]string, error)
	GetStatsForCIK(cik string) (*CompanyStats, error)
}
