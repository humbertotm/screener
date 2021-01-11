package domain

import "context"

// ScreenerService defines the business use cases for this application
type ScreenerService interface {
	FindCandidateCompanies(ctx context.Context, criteria *ScreeningCriteria) ([]string, error)
	GetStatsForCIK(ctx context.Context, cik string) (*CompanyStats, error)
	GetStatsForTicker(ctx context.Context, ticker string) (*CompanyStats, error)
}

// ScreeningCriteria defines the whole set of possible criteria combinations to be employed
// for screening purposes
type ScreeningCriteria struct{}
