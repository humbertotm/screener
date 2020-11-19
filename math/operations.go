package math

// ComputeLinearRegression computes the best fitting line for provided set of data points.
// Return value represents a line as y = alpha + beta*x
func ComputeLinearRegression(dataPoints []float64) (alpha, beta float64, err error) {}

// ComputeAverage computes the mean value for provided set of data points
func ComputeAverage(dataPoints []float64) (float64, error) {}

// ComputeCompoundingRate computes a compounding rate for delta between final and initial
// value in a specified period count
func ComputeCompoundingRate(initial, final float64, periodCount int) (float64, error)
