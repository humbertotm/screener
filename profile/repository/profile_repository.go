package repository

import "mypath/domain"

type profileRepository struct {
	Client interface{} // This is to be replaced by the MongoDB client type
}

// NewProfileRepository returns a fresh profileRepository
func NewProfileRepository(Client interface{}) domain.ProfileRepository {
	return &profileRepository{Client}
}

// GetFullCIKList returns the full list of available CIKs in profiles collection in db
func (r *profileRepository) GetFullCIKList() ([]string, error) {}

// GetFullProfileForCIK returns the full financial profile for proviede CIK
// spanning available range of years
func (r *profileRepository) GetFullProfileForCIK(cik string) (*FullCompanyProfile, error) {}

// GetFullProfileForTicker returns the full financial profile for provided ticker
// spanning available range of years
func (r *profileRepository) GetFulProfileForTicker(ticker string) (*FullCompanyProfile, error) {}
