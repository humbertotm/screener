package finance

func ComputeAdjustedEPSHistory(epsHistory, stockSplitHistory []float64) []float64 {
	if len(epsHistory) != len(stockSplitHistory) {
		return []float64{}
	}

	adjustedHistory := make([]float64, len(epsHistory), len(epsHistory))
	stockSplitFactor := 1.0

	for i, eps := range epsHistory {
		if stockSplitHistory[i] != 0 {
			stockSplitFactor = stockSplitFactor * stockSplitHistory[i]
		}

		adjustedHistory[i] = eps * stockSplitFactor
	}

	return adjustedHistory
}

func ComputePVOfCashFlows(dividendFlows []float64, netEquity float64) float64 {
	baseYear := 2009
	presentValues := make([]float64, len(dividendFlows), len(dividendFlows))
	var totalPV float64

	for i, flow := range dividendFlows {
		presentValues[i] = ComputePV(flow, baseYear, baseYear+i)
	}

	for _, pv := range presentValues {
		totalPV = totalPV + pv
	}

	return totalPV + ComputePV(netEquity, baseYear, 2020)
}

func ComputePV(fcfValue float64, baseYear, flowYear int) float64 {
	if flowYear == baseYear {
		return fcfValue
	} else {
		inflationRate := retrieveInfRate(flowYear)
		nextValue := fcfValue * (1 + inflationRate)
		return ComputePV(nextValue, baseYear, flowYear-1)
	}
}

func retrieveInfRate(year int) float64 {
	inflationRatesFor2010s := make([]float64, 12, 12)

	return inflationRatesFor2010s[year]
}
