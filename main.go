package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

type DataPoint struct {
	MetricValue float64 `json:"metricValue"`
	Dtime       string  `json:"dtime"`
}

func main() {
	// Read the input data
	bytes, err := ioutil.ReadFile("./inputs/1.json")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}

	// Unmarshal the JSON into a slice of DataPoint objects
	var data []DataPoint
	if err := json.Unmarshal(bytes, &data); err != nil {
		fmt.Println("Error parsing JSON:", err)
		os.Exit(1)
	}

	// Sort the data by date
	sort.Slice(data, func(i, j int) bool { return data[i].Dtime < data[j].Dtime })

	// Calculate the min, max, median and average
	var sum float64

	min := data[0].MetricValue
	max := data[0].MetricValue
	for _, datapoint := range data {
		sum += datapoint.MetricValue
		if datapoint.MetricValue < min {
			min = datapoint.MetricValue
		}
		if datapoint.MetricValue > max {
			max = datapoint.MetricValue
		}
	}

	n := len(data)
	avg := sum / float64(n)
	medianIndex := int(n / 2)
	var median float64
	if n%2 == 0 {
		median = (data[medianIndex-1].MetricValue + data[medianIndex].MetricValue) / 2
	} else {
		median = data[medianIndex].MetricValue
	}

	// Print the results
	fmt.Printf("Metric Analyser v1.0.0\n")
	fmt.Printf("=========================\n")
	fmt.Println("\nPeriod checked:")
	fmt.Println("\n  From:", data[0].Dtime)
	fmt.Println("  To:", data[len(data)-1].Dtime)
	fmt.Println("\nStatistics:")
	fmt.Println("\n    Unit: Megabits per second")
	fmt.Printf("\n    Min: %.2f\n", min/100000)
	fmt.Printf("    Max: %.2f\n", max/100000)
	fmt.Printf("    Median: %.2f\n", median/100000)
	fmt.Printf("    Average: %.2f\n", avg/100000)
}
