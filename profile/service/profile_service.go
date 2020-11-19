package service

import "mypath/domain"

type profileService struct {
	domain.ProfileRepository
}

// NewProfileService returns a fresh profileService
func NewProfileService(profileRepo domain.ProfileRepository) domain.ProfileService {
	return &profileService{profileRepo}
}

// FindCandidateCompanies returns a list of ciks for companies successfully passing the set
// of rules that define them as attractive companies in paper
func (s *profileService) FindCandidateCompanies() ([]string, error) {}

// GetStatsForCIK returns full detail on the stats of a particular company as denoted by its
// cik
func (s *profileService) GetStatsForCIK(cik string) (*domain.CompanyStats, error) {}
