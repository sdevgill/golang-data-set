package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
)

const (
	// InputFilePath - constant for input file path
	InputFilePath = "./inputs/1.json"
	// MegaBitMultiplier - constant for Megabits/second conversion
	MegaBitMultiplier = 8
)

// DataPoint represents a single data point with a metric value and a timestamp
type DataPoint struct {
	MetricValue float64 `json:"metricValue"`
	Dtime       string  `json:"dtime"`
}

// Convert to Megabits per second
func convertToMegabits(metricValue float64) float64 {
	megabits := (metricValue / 1000000) * MegaBitMultiplier
	return megabits
}

// Main function
func main() {
	// Read the input data from the file
	inputData, err := readInput(InputFilePath)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}

	// Sort the data by date
	sortedInputData, err := sortData(inputData)
	if err != nil {
		fmt.Println("Error sorting data:", err)
		os.Exit(1)
	}

	// Calculate the min, max, median and average values
	min, max, median, avg, err := calculateStatistics(sortedInputData)
	if err != nil {
		fmt.Println("Error calculating statistics:", err)
		os.Exit(1)
	}

	// Convert min, max, median, and avg values to Megabits
	minMegabits := convertToMegabits(min)
	maxMegabits := convertToMegabits(max)
	medianMegabits := convertToMegabits(median)
	avgMegabits := convertToMegabits(avg)

	// Print the results
	printResults(sortedInputData, minMegabits, maxMegabits, medianMegabits, avgMegabits)
}

// Helper functions

// Read input from a file
func readInput(inputPath string) ([]DataPoint, error) {
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

// Sort the data by timestamps
func sortData(inputData []DataPoint) ([]DataPoint, error) {
	sort.Slice(inputData, func(i, j int) bool { return inputData[i].Dtime < inputData[j].Dtime })
	return inputData, nil
}

func calculateStatistics(inputData []DataPoint) (float64, float64, float64, float64, error) {
	var sum float64

	// Initialize the min and max values with the first data point
	min := inputData[0].MetricValue
	max := inputData[0].MetricValue

	// Iterate through the data to find the min, max, and sum values
	for _, datapoint := range inputData {
		sum += datapoint.MetricValue
		if datapoint.MetricValue < min {
			min = datapoint.MetricValue
		}
		if datapoint.MetricValue > max {
			max = datapoint.MetricValue
		}
	}

	// Calculate the average value
	n := len(inputData)
	avg := sum / float64(n)

	// Create a slice of metric values and sort it
	metricValues := make([]float64, len(inputData))
	for i, datapoint := range inputData {
		metricValues[i] = datapoint.MetricValue
	}
	sort.Float64s(metricValues)

	// Calculate the median value
	medianIndex := int(n / 2)
	var median float64
	// If there is an even number of data points, calculate the median as the average of the two middle values
	if n%2 == 0 {
		median = (metricValues[medianIndex-1] + metricValues[medianIndex]) / 2
	} else {
		// If there is an odd number of data points, the median is the middle value
		median = metricValues[medianIndex]
	}
	median = math.Round(median*100) / 100

	return min, max, median, avg, nil
}

// Print the results
func printResults(data []DataPoint, min float64, max float64, median float64, avg float64) {
	fmt.Printf("\nMetric Analyser v1.0.0\n")
	fmt.Printf("=========================\n")
	fmt.Println("\nPeriod checked:")
	fmt.Println("\n	From:", data[0].Dtime)
	fmt.Println("	To:", data[len(data)-1].Dtime)
	fmt.Println("\nStatistics:")
	fmt.Println("\n	Unit: Megabits per second")
	fmt.Printf("\n	Average: %.2f Megabits\n", avg)
	fmt.Printf("	Min: %.2f Megabits\n", min)
	fmt.Printf("	Max: %.2f Megabits\n", max)
	fmt.Printf("	Median: %.2f Megabits\n", median)
	fmt.Println()
}
