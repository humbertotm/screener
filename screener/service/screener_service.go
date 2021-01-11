package service

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"screener.com/domain"
	"screener.com/profile/repository/mongodb"
)

type screenerService struct {
	profileRepository domain.ProfileRepository
}

func NewScreenerService(client *mongo.Client) domain.ScreenerService {
	return &screenerService{
		profileRepository: mongodb.NewProfileRepository(client),
	}
}

// FindCandidateCompanies returns a list of ciks for companies successfully passing the set
// of rules that define them as attractive companies in paper
func (s *screenerService) FindCandidateCompanies(ctx context.Context, criteria *domain.ScreeningCriteria) ([]string, error) {
	return []string{}, nil
}

func (s *screenerService) GetStatsForCIK(ctx context.Context, cik string) (*domain.CompanyStats, error) {
	profile, err := s.profileRepository.GetFullProfileForCIK(ctx, cik)
	if err != nil {
		return nil, err
	}

	// Return err if not enough information exists to compute company stats
	if ok := profile.HasEnoughData(); !ok {
		return nil, fmt.Errorf("Not enough data for analysis")
	}

	stats, err := profile.ComputeCompanyStats()
	if err != nil {
		return nil, err
	}

	return stats, nil
}

func (s *screenerService) GetStatsForTicker(ctx context.Context, ticker string) (*domain.CompanyStats, error) {
	profile, err := s.profileRepository.GetFullProfileForTicker(ctx, ticker)
	if err != nil {
		return nil, err
	}

	// Return err if not enough information exists to compute company stats
	if ok := profile.HasEnoughData(); !ok {
		return nil, fmt.Errorf("Not enough data for analysis")
	}

	stats, err := profile.ComputeCompanyStats()
	if err != nil {
		return nil, err
	}

	return stats, nil
}
