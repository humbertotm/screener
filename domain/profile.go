package domain

import (
	"context"

	"screener.com/math"
	"screener.com/finance"
)

const (
	minYrsOfData = 5
)

var descriptorList = []string{
	"net_income",
	"total_sales",
	"total_cost",
	"gross_profit_margin",
	"assets_to_liabilities",
	"current_assets_to_current_liabilities",
	"current_assets_to_liabilities",
	"working_capital_to_current_liabilities",
	"working_capital_to_liabilities",
	"goodwill_to_assets",
	"goodwill_to_equity",
	"shares_outstanding",
	"stock_split_ratio",
	"eps",
	"equity_per_share",
	"tangible_assets_per_share",
	"liabilities_per_share",
	"debt_to_equity",
	"debt_to_net_equity",
	"return_on_equity",
	"return_on_working_capital",
}

// ProfileRepository defines the interface employed to interact with the profiles db
type ProfileRepository interface {
	GetFullCIKList(ctx context.Context) (*[]interface{}, error)
	GetFullProfileForCIK(ctx context.Context, cik string) (*FullCompanyProfile, error)
	GetFullProfileForTicker(ctx context.Context, ticker string) (*FullCompanyProfile, error)
}

// FinancialProfile is employed to unmarshall the financial data contained in
// YearlyProfile.Profile
// A type *float64 is employed for values in map as it is important to know a value is null
// as opposed to using its zero value.
type FinancialProfile map[string]*float64

// YearlyProfile defines the structure of the documents pulled out of the profiles collection
type YearlyProfile struct {
	CIK     string           `json:"cik"`
	Ticker  string           `json:"ticker"`
	Year    int              `json:"year"`
	Profile FinancialProfile `json:"profile"`
}

// FullCompanyProfile defines a company's profile for a range of years
type FullCompanyProfile []YearlyProfile

// HasEnoughData checks if the referenced FullCompanyProfile has at least minYrsOfData
func (p *FullCompanyProfile) HasEnoughData() bool {
	profLen := len(*p)

	for i := profLen - 1; i > (profLen - 1 - minYrsOfData); i-- {
		if len((*p)[i].Profile) == 0 {
			return false
		}
	}

	return true
}

// Purge modifies the referenced FullCompanyProfile in place leaving only consecutive
// YearlyProfiles containing data. Assumes elements have been preorded ascendingly based on
// year
func (p *FullCompanyProfile) Purge() {
	profLen := len(*p)
	for i := profLen - minYrsOfData - 2; i >= 0; i-- {
		if len((*p)[i].Profile) == 0 {
			purgedProfile := (*p)[i+1:]
			p = &purgedProfile
			return
		}
	}

	return
}

// ExtractHistoryFor...
// TODO: Seems like I'll need to return a []float64 instead of []*float64.
// Figure out how to deal with zero value and its interpretation in this case.
func (p *FullCompanyProfile) ExtractHistoryFor(key string) []float64 {
	var history []float64

	for i := 0; i < len(*p); i++ {
		history = append(history, *((*p)[i].Profile[key]))
	}

	return history
}

func (p *FullCompanyProfile) ComputeCompanyStats() (*CompanyStats, error) {
	// prof := *p

	// beginAt := prof[0].Year
	// endAt := prof[len(prof)-1].Year
	// netIncomeHistory := p.ExtractHistoryFor("net_income")

	var historiesMap map[string][]float64
	var changeRateHistoriesMap map[string][]float64
	var avgMap map[string]float64

	// TODO: adjust all per share stats to be computed as described below
	for _, key := range descriptorList {
		if key == "eps" {
			epsHistory := p.ExtractHistoryFor("eps")
			stockSplitHistory := p.ExtractHistoryFor("stock_split_ratio")
			adjustedEPSHistory := finance.ComputeAdjustedEPSHistory(epsHistory, stockSplitHistory)
			historiesMap[key] = adjustedEPSHistory
			changeRateHistoriesMap[key] = math.ComputeChangeRateHistory(adjustedEPSHistory)
			avg, err := math.ComputeAverage(adjustedEPSHistory)
			if err != nil {
				return nil, err
			}

			avgMap[key] = avg
		} else {
		history := p.ExtractHistoryFor(key)
		historiesMap[key] = history
		changeRateHistoriesMap[key] = math.ComputeChangeRateHistory(history)

		avg, err := math.ComputeAverage(history)
		if err != nil {
			return nil, err
		}

			avgMap[key] = avg
		}

	}

	netEquityBaseYr := historiesMap["net_equity"][0]
	pvDividendsPerShare := finance.ComputePVOfCashFlows(historiesMap["dividends_per_share"])
	netEquityLastYr := historiesMap["net_equity_per_share"][len(historiesMap["net_equity_per_share"]) - 1]
	pvNetEquityLastYr := ComputePV(netEquityLastYr, 2009, 2020)
	roi := math.ComputeCompoundingRate(netEquityBaseYr, pvDividendsPerShare + pvNetEquityLastYr, len(historiesMap["net_equity_per_share"]))

	stats := CompanyStats{
		BeginAt:          (*p)[0].Year,
		EndAt:            (*p)[len(*p)-1].Year,
		NetIncomeHistory: historiesMap["net_income"],
		YoYNetIncomeChange: changeRateHistoriesMap["net_income"],
		AvgYoYNetIncomeChange: avgMap["net_income"],
		TotalSalesHistory: historiesMap["total_sales"],
		YoYTotalSalesChange: changeRateHistoriesMap["total_sales"],
		AvgYoYTotalSalesChange: avgMap["total_sales"],
		TotalCostOfGoodsHistory: historiesMap["total_cost"],
		YoYTotalCostOfGoodsChange: changeRateHistoriesMap["total_cost"],
		AvgYoYTotalCostOfGoodsChange: avgMap["total_cost"],
		GrossProfitMarginHistory: historiesMap["gross_profit_margin"],
		YoYGrossProfitMarginChange: changeRateHistoriesMap["gross_profit_margin"],
		AvgYoYGrossProfitMarginChange: avgMap["gross_profit_margin"],
		AssetsToLiabilitiesHistory: historiesMap["assets_to_liabilities"],
		YoYAssetsToLiabilitiesChange: changeRateHistoriesMap["assets_to_liabilities"],
		AvgAssetsToLiabilities: avgMap["assets_to_liabilities"],
		CurrentAssetsToCurrentLiabilitiesHistory: historiesMap["current_assets_to_current_liabilities"],
		AvgCurrentAssetsToCurrentLiabilities: avgMap["current_assets_to_current_liabilities"],
		YoYCurrentAssetsToCurrentLiabilitiesChange: changeRateHistoriesMap["current_assets_to_current_liabilities"],
		CurrentAssetsToLiabilitiesHistory: historiesMap["current_assets_to_liabilities"],
		YoYCurrentAssetsToLiabilitiesChange: changeRateHistoriesMap["current_assets_to_liabilities"],
		AvgCurrentAssetsToLiabilities: avgMap["current_assets_to_liabilities"],
		WorkingCapitalToCurrentLiabilitiesHistory: historiesMap["working_capital_to_current_liabilities"],
		YoYWorkingCapitalToCurrentLiabilitiesChange: changeRateHistoriesMap["working_capital_to_current_liabilities"],
		AvgWorkingCapitalToCurrentLiabilities: avgMap["working_capital_to_current_liabilities"],
		WorkingCapitalToLiabilitiesHistory: historiesMap["working_capital_to_liabilities"],
		YoYWorkingCapitalToLiabilitiesChange: changeRateHistoriesMap["working_capital_to_liabilities"],
		AvgWorkingCapitalToLiabilities: avgMap["working_capital_to_liabilities"],
		GoodwillToAssetsHistory: historiesMap["goodwill_to_assets"],
		YoYGoodwillToAssetsChange: changeRateHistoriesMap["goodwill_to_assets"],
		AvgGoodwillToAssets: avgMap["goodwill_to_assets"],
		GoodwillToEquityHistory: historiesMap["goodwill_to_equity"],
		YoYGoodwillToEquityChange: changeRateHistoriesMap["goodwill_to_equity"],
		AvgGoodwillToEquity: avgMap["goodwill_to_equity"],
		SharesOutstandingHistory: historiesMap["shares_outstanding"],
		YoYSharesOutstandingChange: changeRateHistoriesMap["shares_outstanding"],
		AvgSharesOutstanding: avgMap["shares_outstanding"],
		EPSHistory: historiesMap["eps"],
		YoYEPSChange: changeRateHistoriesMap["eps"],
		AvgEPS: avgMap["eps"],
		EquityPerShareHistory: historiesMap["equity_per_share"],
		YoYEquityPerShareChange: changeRateHistoriesMap["equity_per_share"],
		AvgEquityPerShare: avgMap["equity_per_share"],
		TangibleAssetsPerShareHistory: historiesMap["tangible_assets_per_share"],
		YoYTangibleAssetsPerShareChange: changeRateHistoriesMap["tangible_assets_per_share"],
		AvgTangibleAssetsPerShare: avgMap["tangible_assets_per_share"],
		LiabilitiesPerShareHistory: historiesMap["liabilities_per_share"],
		YoYLiabilitiesPerShareChange: changeRateHistoriesMap["liabilities_per_share"],
		AvgLiabilitiesPerShare: avgMap["liabilities_per_share"],
		DebtToEquityHistory: historiesMap["debt_to_equity"],
		YoYDebtToEquityChange: changeRateHistoriesMap["debt_to_equity"],
		AvgDebtToEquity: avgMap["debt_to_equity"],
		DebtToNetEquityHistory: historiesMap["debt_to_net_equity"],
		YoYDebtToNetEquityChange: changeRateHistoriesMap["debt_to_net_equity"],
		AvgDebtToNetEquity: avgMap["debt_to_net_equity"],
		ReturnOnEquityHistory: historiesMap["return_on_equity"],
		YoYReturnOnEquityChange: changeRateHistoriesMap["return_on_equity"],
		AvgReturnOnEquity: avgMap["return_on_equity"],
		ReturnOnWorkingCapitalHistory: historiesMap["return_on_working_capital"],
		YoYReturnOnWorkingCapitalChange: changeRateHistoriesMap["return_on_working_capital"],
		AvgReturnOnWorkingCapital: avgMap["return_on_working_capital"],
		CompoundROIForPeriod: math.ComputeCompoundingRate(historiesMap["net_equity"][0], finance.ComputePVOfCashFlows(historiesMap["dividends_per_share"], historiesMap["net_equity_per_share"][len(historiesMap["net_equity_per_share"]) - 1]))
	}

	return &stats, nil
}


