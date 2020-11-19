package service

import "mypath/domain"

type screenerService struct {
	domain.ProfileRepository
}

// NewScreenerService returns a fresh screenerService
func NewScreenerService(profileRepo domain.ProfileRepository) domain.ScreenerService {
	return &screenerService{profileRepo}
}

// FindCandidateCompanies returns a list of ciks for companies successfully passing the set
// of rules that define them as attractive companies in paper
func (s *screenerService) FindCandidateCompanies() ([]string, error) {}

// GetStatsForCIK returns full detail on the stats of a particular company as denoted by its
// cik
func (s *screenerService) GetStatsForCIK(cik string) (*domain.CompanyStats, error) {}
