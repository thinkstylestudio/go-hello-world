package main

import (
	"fmt"
)

// create a struct to represent each of the timeseries threshold pairs
type TimeSeriesThreshold struct {
	PointInTime float64
	Threshold   float64
}

// given the collection/array/slice of structs (pairs) and a given threshold return a array of floats
func processThreshold(pairs []TimeSeriesThreshold, threshold float64) []float64 {
	var output []float64
	// go through and append to output timestamps that are above the threshold
	for _, pair := range pairs {
		if pair.Threshold >= threshold {
			output = append(output, pair.PointInTime)
		}
	}
	// assuming out of [10, 30]
	return output
}

func processStartEndPairings(pairs []TimeSeriesThreshold, threshold float64) []float64 {
	var output []float64

	aboveThreshold := false

	// go through and figure the beginning threshold and ending threshold pairs
	for index, pair := range pairs {

		if !aboveThreshold && pair.Threshold >= threshold {
			output = append(output, pair.PointInTime)
			aboveThreshold = true
		}
		if aboveThreshold && pair.Threshold < threshold {
			output = append(output, pairs[index-1].PointInTime)
			aboveThreshold = false
		}

	}

	if aboveThreshold {
		output = append(output, pairs[len(pairs)-1].PointInTime)
	}

	// assuming out of [[10, 10],[30,40]]
	return output
}

func main() {

	// Input PointInTime: [[10,0.9],[20,0.1],[30,0.8], [40, 0.9]]
	// Input 0.8
	// Output [[10, 30]]

	// Create a slice
	pairs := []TimeSeriesThreshold{
		{PointInTime: 10, Threshold: 0.9},
		{PointInTime: 20, Threshold: 0.1},
		{PointInTime: 30, Threshold: 0.8},
		{PointInTime: 40, Threshold: 0.9},
	}
	threshold := 0.8

	startEndingResult := processStartEndPairings(pairs, threshold)
	fmt.Println(startEndingResult)

	thresholdResult := processThreshold(pairs, threshold)
	fmt.Println(thresholdResult)

}
