package main

import (
	"encoding/json"
	"io/ioutil"
	"sort"
	"time"
)

type Download struct {
	Dtime       time.Time `json:"dtime"`
	MetricValue float64   `json:"metricValue"`
}

// File
const filename = "../inputs/1.json"

// Reads data from a JSON file and returns a slice of Download structures
func readData(filename string) []Download {
	// Read the file
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data
	var data []Download
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}

	return data, nil
}

// Converts bytes per second to megabits per second
func bytesToMegabits(bytes float64) float64 {
	return bytes * 8 / 1000000
}

// Calculate the average download speed
func average(values []float64) float64 {
	sum := 0.0
	for _, value := range values {
		sum += value
	}
	return sum / float64(len(values))
}

// Calculate the median download speed
func median(values []float64) float64 {
	length := len(values)
	if length == 0 {
		return 0
	}

	// Sort the values
	sort.Float64s(values)

	// Return the middle value
	if length%2 == 0 {
		return (values[length/2-1] + values[length/2]) / 2
	}
	return values[length/2]
}
