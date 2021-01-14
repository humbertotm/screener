package domain

import "context"

// ScreenerHandler defines the interface for the CLI command handler
type ScreenerHandler interface {
	GetStatsForCIK(ctx context.Context, cik string) error
	GetStatsForTicker(ctx context.Context, ticker string) error
}

// ScreenerService defines the business use cases for this application
type ScreenerService interface {
	FindCandidateCompanies(ctx context.Context, criteria *ScreeningCriteria) ([]string, error)
	GetStatsForCIK(ctx context.Context, cik string) (*CompanyStats, error)
	GetStatsForTicker(ctx context.Context, ticker string) (*CompanyStats, error)
}

// ScreeningCriteria defines the whole set of possible criteria combinations to be employed
// for screening purposes
type ScreeningCriteria struct{}
