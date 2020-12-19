package math

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeAverage(t *testing.T) {
	testCases := map[string]struct {
		dataPoints  []float64
		expectedAvg float64
		expectedErr error
	}{
		"Correct calculation": {
			dataPoints:  []float64{1.5, 2.0, 3.8, 0.5},
			expectedAvg: 1.95,
			expectedErr: nil,
		},
		"Empty collection of data points -- Error": {
			dataPoints:  []float64{},
			expectedAvg: 0.0,
			expectedErr: fmt.Errorf("Must provide a non empty collection of data points"),
		},
	}

	for _, tc := range testCases {
		avg, err := ComputeAverage(tc.dataPoints)
		if tc.expectedErr != nil {
			assert.NotNil(t, err)
			assert.Equal(t, tc.expectedErr.Error(), err.Error())
		} else {
			assert.Nil(t, err)
			assert.Equal(t, tc.expectedAvg, avg)
		}
	}
}

func TestComputeCompoundingRate(t *testing.T) {
	testCases := map[string]struct {
		inital       float64
		final        float64
		periodCount  int
		expectedRate float64
		expectedErr  error
	}{
		"Correct calculation -- Positive rate": {
			inital:       110.5,
			final:        562.9325,
			periodCount:  11,
			expectedRate: 0.15952809731470885,
			expectedErr:  nil,
		},
		"Correct calculation -- Negative rate": {
			inital:       562.9325,
			final:        110.5,
			periodCount:  11,
			expectedRate: -0.1375801911865282,
			expectedErr:  nil,
		},
		"Period count is zero -- Error": {
			inital:       110.5,
			final:        562.9325,
			periodCount:  0,
			expectedRate: 0,
			expectedErr:  fmt.Errorf("Invalid period count"),
		},
	}

	for _, tc := range testCases {
		r, err := ComputeCompoundingRate(tc.inital, tc.final, tc.periodCount)
		if tc.expectedErr != nil {
			assert.NotNil(t, err)
			assert.Equal(t, tc.expectedErr.Error(), err.Error())
		} else {
			assert.Nil(t, err)
			assert.Equal(t, tc.expectedRate, r)
		}
	}
}

func TestComputeLinearRegression(t *testing.T) {
	// yearRange := []float64{2009, 2010, 2011, 2012, 2013, 2014, 2015, 2016, 2017, 2018, 2019, 2020}
	xs := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	testCases := map[string]struct {
		xs            []float64
		ys            []float64
		expectedAlpha float64
		expectedBeta  float64
		expectedErr   error
	}{
		"Simple case -- Slope of 1, through origin": {
			xs:            xs,
			ys:            xs,
			expectedAlpha: 0,
			expectedBeta:  1,
			expectedErr:   nil,
		},
		"Simple case -- Slope of -1": {
			xs:            []float64{11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			ys:            []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			expectedAlpha: 11,
			expectedBeta:  -1,
			expectedErr:   nil,
		},
		"Simple case -- Slope of 1 with offset": {
			xs:            xs,
			ys:            []float64{3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14},
			expectedAlpha: 3,
			expectedBeta:  1,
			expectedErr:   nil,
		},
		"Real use case simulation -- No negative values": {
			xs: xs,
			ys: []float64{
				258429681.3580,
				238529388.234,
				259804859.123,
				271002345.222,
				275003245.693,
				293475934.2345,
				273984385.8348,
				263457035.5832,
				281750235.3845,
				305384693.5834,
				302385482.5,
				320485932.658,
			},
			expectedAlpha: 2.4833356842505512e+08,
			expectedBeta:  5.510460580420281e+06,
			expectedErr:   nil,
		},
		"Real use case simulation -- Mixed positive and negative values": {
			xs: xs,
			ys: []float64{
				258429681.3580,
				238529388.234,
				259804859.123,
				-18002345.222,
				-5003245.693,
				183475934.2345,
				243984385.8348,
				263457035.5832,
				281750235.3845,
				305384693.5834,
				282385482.5,
				320485932.658,
			},
			expectedAlpha: 1.4556258032938847e+08,
			expectedBeta:  1.3150470812511189e+07,
			expectedErr:   nil,
		},
		"Wide variation": {
			xs: xs,
			ys: []float64{
				10.5,
				-5,
				-8.5,
				7.5,
				5.5,
				-12.3,
				15,
				4.5,
				-15.8,
				-0.1,
				20.5,
				-26.4,
			},
			expectedAlpha: 3.651282051282051,
			expectedBeta:  -0.7335664335664336,
			expectedErr:   nil,
		},
		"Narrow variation": {
			xs: xs,
			ys: []float64{
				5.4,
				4.2,
				6.3,
				3.5,
				5,
				5.3,
				2.3,
				6.7,
				4.8,
				4.7,
				5.2,
				3.6,
			},
			expectedAlpha: 5.007692307692308,
			expectedBeta:  -0.04685314685314686,
			expectedErr:   nil,
		},
		"Invalid inputs": {
			xs:            xs,
			ys:            []float64{1, 2, 3},
			expectedAlpha: 5.007692307692308,
			expectedBeta:  -0.04685314685314686,
			expectedErr:   fmt.Errorf("Xs and Ys must be the same length"),
		},
	}

	for _, tc := range testCases {
		a, b, err := ComputeLinearRegression(tc.xs, tc.ys)
		if tc.expectedErr != nil {
			assert.NotNil(t, err)
			assert.Equal(t, tc.expectedErr.Error(), err.Error())
		} else {
			assert.Equal(t, tc.expectedAlpha, a, "Unexpected alpha")
			assert.Equal(t, tc.expectedBeta, b, "Unexpected beta")
		}
	}
}
