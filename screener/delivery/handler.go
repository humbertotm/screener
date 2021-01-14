package delivery

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"screener.com/domain"
	"screener.com/screener/service"
)

type screenerHandler struct {
	screenerService domain.ScreenerService
}

func NewScreenerHandler(client *mongo.Client) domain.ScreenerHandler {
	return &screenerHandler{
		screenerService: service.NewScreenerService(client),
	}
}

func (h *screenerHandler) GetStatsForCIK(ctx context.Context, cik string) error {
	companyStats, err := h.screenerService.GetStatsForCIK(ctx, cik)
	if err != nil {
		return err
	}

	json, err := json.MarshalIndent(companyStats, "", "\t")
	if err != nil {
		return err
	}

	fmt.Printf("Stats for CIK: %s\n%s\n", cik, string(json))

	return nil
}

func (h *screenerHandler) GetStatsForTicker(ctx context.Context, ticker string) error {
	return nil
}
