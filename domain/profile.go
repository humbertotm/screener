package domain

import (
	"context"

	"screener.com/finance"
	"screener.com/math"
	"screener.com/utils"
)

const (
	minYrsOfData = 5
)

// to be added
// total_cost
// working_capital_to_current_liabilities
// working_capital_to_total_liabilities
// goodwill_to_equity
// equity_per_share
// net_equity_per_share
// tangible_assets_per_share
// liabilities_per_share
// return-on-equity
// return-on-working-capital

// (:research-expense-to-revenue
//  :current-assets-to-current-liabilities
//  :working-capital
//  :comprehensive-stocks-outstanding
//  :capital-expenditures
//  :total-liabilities
//  :stock-options-exercised
//  :return-on-working-capital
//  :goodwill-to-total-assets
//  :return-on-equity
//  :operating-income
//  :depreciation
//  :total-equity
//  :debt-to-net-equity
//  :net-profit-margin
//  :tangible-assets
//  :debt-to-equity
//  :total-sales
//  :goodwill
//  :stock-options-granted
//  :research-and-development-expense
//  :stock-split-ratio
//  :inventory
//  :free-cash-flow
//  :accounts-payable
//  :current-assets
//  :operational-profit-margin
//  :long-term-debt
//  :net-equity
//  :dividends-paid-to-net-income
//  :total-tangible-assets-to-total-liabilities
//  :current-liabilities
//  :stock-repurchase-payment
//  :common-stock-outstanding
//  :net-income
//  :dividend-payment
//  :total-assets
//  :dividends-per-share-paid
//  :current-assets-to-total-liabilities
//  :accounts-payable-to-current-assets
//  :diluted-eps
//  :eps)

var descriptorList = []string{
	"net_income",
	"total_sales",
	"total_cost",
	"operational_profit_margin",
	"total_tangible_assets_to_total_liabilities",
	"current_assets_to_current_liabilities",
	"current_assets_to_total_liabilities",
	"working_capital_to_current_liabilities",
	"working_capital_to_total_liabilities",
	"goodwill_to_total_assets",
	"goodwill_to_equity",
	"comprehensive_stocks_outstanding",
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

var perShareDescriptorList = []string{
	"tangible_assets_per_share",
	"liabilities_per_share",
	"eps",
	"equity_per_share",
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
type FinancialProfile map[string]float64

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

	for i := profLen - 1; i > (profLen - minYrsOfData - 1); i-- {
		if len((*p)[i].Profile) == 0 {
			return false
		}
	}

	return true
}

// Purge modifies the referenced FullCompanyProfile in place leaving only consecutive
// YearlyProfiles containing data. Assumes elements have been preorded ascendingly based on
// year
func (p *FullCompanyProfile) purge() {
	for i := len(*p) - 1; i >= 0; i-- {
		if len((*p)[i].Profile) == 0 {
			*p = (*p)[i+1:]
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
		history = append(history, (*p)[i].Profile[key])
	}

	return history
}

func (p *FullCompanyProfile) ComputeCompanyStats() (*CompanyStats, error) {
	historiesMap := make(map[string][]float64)
	changeRateHistoriesMap := make(map[string][]float64)
	avgMap := make(map[string]float64)
	baseYear := (*p)[0].Year
	finalYear := (*p)[len(*p)-1].Year

	p.purge()
	stockSplitHistory := p.ExtractHistoryFor("stock_split_ratio")

	// Populate data source maps
	for _, key := range descriptorList {
		if utils.IndexOf(perShareDescriptorList, key) >= 0 {
			history := p.ExtractHistoryFor(key)
			adjustedHistory := finance.ComputeAdjustedPerShareHistory(history, stockSplitHistory)
			historiesMap[key] = adjustedHistory
			changeRateHistoriesMap[key] = math.ComputeChangeRateHistory(adjustedHistory)
			avg, err := math.ComputeAverage(adjustedHistory)
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

	netEquityBaseYr := historiesMap["net_equity_per_share"][0]
	netEquityLastYr := historiesMap["net_equity_per_share"][finalYear-baseYear]
	pvDividendsPerShare := finance.ComputeCashFlowsPV(historiesMap["dividends_per_share"], baseYear)

	pvNetEquityFinalYr := finance.ComputePV(netEquityLastYr, baseYear, finalYear)
	roi, err := math.ComputeCompoundingRate(netEquityBaseYr, pvDividendsPerShare+pvNetEquityFinalYr, finalYear-baseYear)
	if err != nil {
		return nil, err
	}

	stats := CompanyStats{
		BeginAt:                                     baseYear,
		EndAt:                                       finalYear,
		NetIncomeHistory:                            historiesMap["net_income"],
		YoYNetIncomeChange:                          changeRateHistoriesMap["net_income"],
		AvgYoYNetIncomeChange:                       avgMap["net_income"],
		TotalSalesHistory:                           historiesMap["total_sales"],
		YoYTotalSalesChange:                         changeRateHistoriesMap["total_sales"],
		AvgYoYTotalSalesChange:                      avgMap["total_sales"],
		TotalCostOfGoodsHistory:                     historiesMap["total_cost"],
		YoYTotalCostOfGoodsChange:                   changeRateHistoriesMap["total_cost"],
		AvgYoYTotalCostOfGoodsChange:                avgMap["total_cost"],
		GrossProfitMarginHistory:                    historiesMap["gross_profit_margin"],
		YoYGrossProfitMarginChange:                  changeRateHistoriesMap["gross_profit_margin"],
		AvgYoYGrossProfitMarginChange:               avgMap["gross_profit_margin"],
		AssetsToLiabilitiesHistory:                  historiesMap["assets_to_liabilities"],
		YoYAssetsToLiabilitiesChange:                changeRateHistoriesMap["assets_to_liabilities"],
		AvgAssetsToLiabilities:                      avgMap["assets_to_liabilities"],
		CurrentAssetsToCurrentLiabilitiesHistory:    historiesMap["current_assets_to_current_liabilities"],
		AvgCurrentAssetsToCurrentLiabilities:        avgMap["current_assets_to_current_liabilities"],
		YoYCurrentAssetsToCurrentLiabilitiesChange:  changeRateHistoriesMap["current_assets_to_current_liabilities"],
		CurrentAssetsToLiabilitiesHistory:           historiesMap["current_assets_to_liabilities"],
		YoYCurrentAssetsToLiabilitiesChange:         changeRateHistoriesMap["current_assets_to_liabilities"],
		AvgCurrentAssetsToLiabilities:               avgMap["current_assets_to_liabilities"],
		WorkingCapitalToCurrentLiabilitiesHistory:   historiesMap["working_capital_to_current_liabilities"],
		YoYWorkingCapitalToCurrentLiabilitiesChange: changeRateHistoriesMap["working_capital_to_current_liabilities"],
		AvgWorkingCapitalToCurrentLiabilities:       avgMap["working_capital_to_current_liabilities"],
		WorkingCapitalToLiabilitiesHistory:          historiesMap["working_capital_to_liabilities"],
		YoYWorkingCapitalToLiabilitiesChange:        changeRateHistoriesMap["working_capital_to_liabilities"],
		AvgWorkingCapitalToLiabilities:              avgMap["working_capital_to_liabilities"],
		GoodwillToAssetsHistory:                     historiesMap["goodwill_to_assets"],
		YoYGoodwillToAssetsChange:                   changeRateHistoriesMap["goodwill_to_assets"],
		AvgGoodwillToAssets:                         avgMap["goodwill_to_assets"],
		GoodwillToEquityHistory:                     historiesMap["goodwill_to_equity"],
		YoYGoodwillToEquityChange:                   changeRateHistoriesMap["goodwill_to_equity"],
		AvgGoodwillToEquity:                         avgMap["goodwill_to_equity"],
		SharesOutstandingHistory:                    historiesMap["shares_outstanding"],
		YoYSharesOutstandingChange:                  changeRateHistoriesMap["shares_outstanding"],
		AvgSharesOutstanding:                        avgMap["shares_outstanding"],
		EPSHistory:                                  historiesMap["eps"],
		YoYEPSChange:                                changeRateHistoriesMap["eps"],
		AvgEPS:                                      avgMap["eps"],
		EquityPerShareHistory:                       historiesMap["equity_per_share"],
		YoYEquityPerShareChange:                     changeRateHistoriesMap["equity_per_share"],
		AvgEquityPerShare:                           avgMap["equity_per_share"],
		TangibleAssetsPerShareHistory:               historiesMap["tangible_assets_per_share"],
		YoYTangibleAssetsPerShareChange:             changeRateHistoriesMap["tangible_assets_per_share"],
		AvgTangibleAssetsPerShare:                   avgMap["tangible_assets_per_share"],
		LiabilitiesPerShareHistory:                  historiesMap["liabilities_per_share"],
		YoYLiabilitiesPerShareChange:                changeRateHistoriesMap["liabilities_per_share"],
		AvgLiabilitiesPerShare:                      avgMap["liabilities_per_share"],
		DebtToEquityHistory:                         historiesMap["debt_to_equity"],
		YoYDebtToEquityChange:                       changeRateHistoriesMap["debt_to_equity"],
		AvgDebtToEquity:                             avgMap["debt_to_equity"],
		DebtToNetEquityHistory:                      historiesMap["debt_to_net_equity"],
		YoYDebtToNetEquityChange:                    changeRateHistoriesMap["debt_to_net_equity"],
		AvgDebtToNetEquity:                          avgMap["debt_to_net_equity"],
		ReturnOnEquityHistory:                       historiesMap["return_on_equity"],
		YoYReturnOnEquityChange:                     changeRateHistoriesMap["return_on_equity"],
		AvgReturnOnEquity:                           avgMap["return_on_equity"],
		ReturnOnWorkingCapitalHistory:               historiesMap["return_on_working_capital"],
		YoYReturnOnWorkingCapitalChange:             changeRateHistoriesMap["return_on_working_capital"],
		AvgReturnOnWorkingCapital:                   avgMap["return_on_working_capital"],
		CompoundROIForPeriod:                        roi,
	}

	return &stats, nil
}
