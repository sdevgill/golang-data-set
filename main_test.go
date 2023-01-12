package main

import (
	"testing"
)

func TestConvertToMegabits(t *testing.T) {
	testCases := []struct {
		metricValue      float64
		expectedMegabits float64
	}{
		{1000000, 8},
		{2000000, 16},
		{3000000, 24},
	}

	for _, tc := range testCases {
		megabits := convertToMegabits(tc.metricValue)
		if megabits != tc.expectedMegabits {
			t.Errorf("convertToMegabits(%f) = %f, expected %f", tc.metricValue, megabits, tc.expectedMegabits)
		}
	}
}
