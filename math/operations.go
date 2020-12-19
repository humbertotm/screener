package math

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/stat"
)

// ComputeLinearRegression computes the best fitting line for provided set of x & y data
// points.
// Return value represents a line as y = alpha + beta*x
// Assumes data points are properly sorted representing points in a cartesian plane.
func ComputeLinearRegression(xs, ys []float64) (float64, float64, error) {
	var weights []float64
	origin := false

	if len(xs) != len(ys) {
		return 0, 0, fmt.Errorf("Xs and Ys must be the same length")
	}

	alpha, beta := stat.LinearRegression(xs, ys, weights, origin)
	return alpha, beta, nil
}

// ComputeAverage computes the mean value for provided set of data points
func ComputeAverage(dataPoints []float64) (float64, error) {
	if !(len(dataPoints) > 0) {
		return 0.0, fmt.Errorf("Must provide a non empty collection of data points")
	}

	return stat.Mean(dataPoints, nil), nil
}

// ComputeCompoundingRate computes a compounding rate for delta between final and initial
// value in a specified period count
func ComputeCompoundingRate(initial, final float64, periodCount int) (float64, error) {
	if !(periodCount > 0) {
		return 0, fmt.Errorf("Invalid period count")
	}

	base := final / initial
	exp := 1.0 / float64(periodCount)

	return math.Pow(base, exp) - 1.0, nil
}
