package domain

// CompanyStats defines the time series analysis output for a specific company
type CompanyStats struct {
	BeginAt int
	EndAt   int
	// Net Income
	NetIncomeHistory      []float64
	YoYNetIncomeChange    []float64
	AvgYoYNetIncomeChange float64
	// Total Sales
	TotalSalesHistory      []float64
	YoYTotalSalesChange    []float64
	AvgYoYTotalSalesChange float64
	// Total Cost
	TotalCostOfGoodsHistory      []float64
	YoYTotalCostOfGoodsChange    []float64
	AvgYoYTotalCostOfGoodsChange float64
	// Gross Profit Margin
	GrossProfitMarginHistory      []float64
	YoYGrossProfitMarginChange    []float64
	AvgYoYGrossProfitMarginChange float64
	// Assets to Liabilities
	AssetsToLiabilitiesHistory   []float64
	YoYAssetsToLiabilitiesChange []float64
	AvgAssetsToLiabilities       float64
	// Current Assets to Current Liabilities
	CurrentAssetsToCurrentLiabilitiesHistory   []float64
	YoYCurrentAssetsToCurrentLiabilitiesChange []float64
	AvgCurrentAssetsToCurrentLiabilities       float64
	// Current Assets to Liabilities
	CurrentAssetsToLiabilitiesHistory   []float64
	YoYCurrentAssetsToLiabilitiesChange []float64
	AvgCurrentAssetsToLiabilities       float64
	// Working Capital to Current Liabilities
	WorkingCapitalToCurrentLiabilitiesHistory   []float64
	YoYWorkingCapitalToCurrentLiabilitiesChange []float64
	AvgWorkingCapitalToCurrentLiabilities       float64
	// Working Capital to Liabilities
	WorkingCapitalToLiabilitiesHistory   []float64
	YoYWorkingCapitalToLiabilitiesChange []float64
	AvgWorkingCapitalToLiabilities       float64
	// Goodwill to Assets
	GoodwillToAssetsHistory   []float64
	YoYGoodwillToAssetsChange []float64
	AvgGoodwillToAssets       float64
	// Goodwill to Equity
	GoodwillToEquityHistory   []float64
	YoYGoodwillToEquityChange []float64
	AvgGoodwillToEquity       float64
	// Shares Outstanding
	SharesOutstandingHistory   []float64
	YoYSharesOutstandingChange []float64
	AvgSharesOutstanding       float64
	// EPS (considering splits)
	EPSHistory   []float64
	YoYEPSChange []float64
	AvgEPS       float64
	// Equity per Share (considering stock splits)
	EquityPerShareHistory   []float64
	YoYEquityPerShareChange []float64
	AvgEquityPerShare       float64
	// Tangible Assets per Share (considering stock splits)
	TangibleAssetsPerShareHistory   []float64
	YoYTangibleAssetsPerShareChange []float64
	AvgTangibleAssetsPerShare       float64
	// Liabilities per Share (considering stock splits)
	LiabilitiesPerShareHistory   []float64
	YoYLiabilitiesPerShareChange []float64
	AvgLiabilitiesPerShare       float64
	// Capitalization Structure
	DebtToEquityHistory      []float64
	YoYDebtToEquityChange    []float64
	AvgDebtToEquity          float64
	DebtToNetEquityHistory   []float64
	YoYDebtToNetEquityChange []float64
	AvgDebtToNetEquity       float64
	// Return on Equity
	ReturnOnEquityHistory   []float64
	YoYReturnOnEquityChange []float64
	AvgReturnOnEquity       float64
	// Return on Working Capital
	ReturnOnWorkingCapitalHistory   []float64
	YoYReturnOnWorkingCapitalChange []float64
	AvgReturnOnWorkingCapital       float64
	// Compound ROI (not considering goodwill)
	CompoundROIForPeriod float64
}
