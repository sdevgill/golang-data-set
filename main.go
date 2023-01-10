package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

// DataPoint represents a single data point with a metric value and a timestamp
type DataPoint struct {
	MetricValue float64 `json:"metricValue"`
	Dtime       string  `json:"dtime"`
}

// Main function
func main() {
	// Read the input data from the file
	data, err := readInput("./inputs/1.json")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}

	// Sort the data by date
	sortedData, err := sortData(data)
	if err != nil {
		fmt.Println("Error sorting data:", err)
		os.Exit(1)
	}

	// Calculate the min, max, median and average values
	min, max, median, avg, err := calculateStatistics(sortedData)
	if err != nil {
		fmt.Println("Error calculating statistics:", err)
		os.Exit(1)
	}

	// Print the results
	printResults(sortedData, min, max, median, avg)
}

// Helper functions

func readInput(inputPath string) ([]DataPoint, error) {
	// Read the contents of the input file
	bytes, err := ioutil.ReadFile(inputPath)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON into a slice of DataPoint objects
	var data []DataPoint
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func sortData(data []DataPoint) ([]DataPoint, error) {
	// Sort the data by timestamps
	sort.Slice(data, func(i, j int) bool { return data[i].Dtime < data[j].Dtime })
	return data, nil
}

func calculateStatistics(data []DataPoint) (float64, float64, float64, float64, error) {
	var sum float64

	// Initialize the min and max values with the first data point
	min := data[0].MetricValue
	max := data[0].MetricValue

	// Iterate through the data to find the min, max, and sum values
	for _, datapoint := range data {
		sum += datapoint.MetricValue
		if datapoint.MetricValue < min {
			min = datapoint.MetricValue
		}
		if datapoint.MetricValue > max {
			max = datapoint.MetricValue
		}
	}

	// Calculate the average value
	n := len(data)
	avg := sum / float64(n)

	// Calculate the median value
	medianIndex := int(n / 2)
	var median float64

	// If there is an even number of data points, calculate the median as the average of the two middle values
	if n%2 == 0 {
		median = (data[medianIndex-1].MetricValue + data[medianIndex].MetricValue) / 2
	} else {
		// If there is an odd number of data points, the median is the middle value
		median = data[medianIndex].MetricValue
	}

	return min, max, median, avg, nil
}

// Print the results
func printResults(data []DataPoint, min float64, max float64, median float64, avg float64) {
	fmt.Printf("\nMetric Analyser v1.0.0\n")
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
	fmt.Println()
}
