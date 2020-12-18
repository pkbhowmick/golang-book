package math

import "testing"

type testPair struct {
	values []float64
	average float64
	min float64
	max float64
}

var tests = []testPair {
	{[]float64{1,2},1.5,1,2},
	{[]float64{-3,0,6},1,-3,6},
	{[]float64{-2,-3},-2.5,-3,-2},
}

func TestAverage(t *testing.T) {
	for _,pair := range tests{
		avg := Average(pair.values)
		if avg != pair.average {
			t.Error(
				"For", pair.values,
				"Expected", pair.average,
				"got", avg,
				)
		}
	}
}

func TestMin(t *testing.T) {
	for _,pair := range tests{
		min := Min(pair.values)
		if min != pair.min {
			t.Error(
				"For", pair.values,
				"Expected", pair.min,
				"got", min,
			)
		}
	}
}

func TestMax(t *testing.T) {
	for _,pair := range tests{
		max := Max(pair.values)
		if max != pair.max {
			t.Error(
				"For", pair.values,
				"Expected", pair.max,
				"got", max,
			)
		}
	}
}