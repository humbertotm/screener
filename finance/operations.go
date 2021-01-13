package finance

func ComputeAdjustedPerShareHistory(perShareHistory, stockSplitHistory []float64) []float64 {
	if len(perShareHistory) != len(stockSplitHistory) {
		return []float64{}
	}

	adjustedHistory := make([]float64, len(perShareHistory), len(perShareHistory))
	stockSplitFactor := 1.0

	for i, eps := range perShareHistory {
		if stockSplitHistory[i] != 0 {
			stockSplitFactor = stockSplitFactor * stockSplitHistory[i]
		}

		adjustedHistory[i] = eps * stockSplitFactor
	}

	return adjustedHistory
}

func ComputeCashFlowsPV(cashFlows []float64, baseYear int) float64 {

	presentValues := make([]float64, len(cashFlows), len(cashFlows))
	var totalPV float64

	for i, flow := range cashFlows {
		presentValues[i] = ComputePV(flow, baseYear, baseYear+i)
	}

	for _, pv := range presentValues {
		totalPV = totalPV + pv
	}

	return totalPV
}

func ComputePV(fcfValue float64, baseYear, flowYear int) float64 {
	if flowYear == baseYear {
		return fcfValue
	} else {
		inflationRate := retrieveInfRate(flowYear)
		discountedValue := fcfValue * (1 + inflationRate)
		return ComputePV(discountedValue, baseYear, flowYear-1)
	}
}

func retrieveInfRate(year int) float64 {
	baseYear := 2009
	inflationRatesFor2010s := make([]float64, 12, 12)

	return inflationRatesFor2010s[year-baseYear]
}
