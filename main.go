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
	InputFilePath = "./inputs/2.json"
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

	// Get the statistics
	output := printResults(sortedInputData, minMegabits, maxMegabits, medianMegabits, avgMegabits)

	// Create the output file
	outputFile, err := os.Create("./outputs/output.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	// Write the results to the output file
	fmt.Fprint(outputFile, output)
}

// Helper functions

// Read input from a file
func readInput(inputPath string) ([]DataPoint, error) {
	// Check if the input file exists
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("Error: input file does not exist at path %s", inputPath)
	}
	if fi, _ := os.Stat(inputPath); fi.IsDir() {
		return nil, fmt.Errorf("Error: input path %s is a directory, expected a file", inputPath)
	}

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
	if len(data) == 0 {
		return nil, fmt.Errorf("Error: input file is empty")
	}

	return data, nil
}

// Sort the data by timestamps
func sortData(inputData []DataPoint) ([]DataPoint, error) {
	sort.Slice(inputData, func(i, j int) bool { return inputData[i].Dtime < inputData[j].Dtime })
	return inputData, nil
}

func calculateStatistics(inputData []DataPoint) (float64, float64, float64, float64, error) {
	// Check if input data is not empty
	if len(inputData) == 0 {
		return 0, 0, 0, 0, fmt.Errorf("Error: input data is empty")
	}

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
func printResults(data []DataPoint, min float64, max float64, median float64, avg float64) string {
	output := fmt.Sprintf("Metric Analyser v1.0.0\n")
	output += fmt.Sprintf("=========================\n")
	output += fmt.Sprintf("\nPeriod checked:\n")
	output += fmt.Sprintf("\n	From: %s\n", data[0].Dtime)
	output += fmt.Sprintf("	To: %s\n", data[len(data)-1].Dtime)
	output += fmt.Sprintf("\nStatistics:\n")
	output += fmt.Sprintf("\n	Unit: Megabits per second\n")
	output += fmt.Sprintf("\n	Average: %.2f \n", avg)
	output += fmt.Sprintf("	Min: %.2f \n", min)
	output += fmt.Sprintf("	Max: %.2f \n", max)
	output += fmt.Sprintf("	Median: %.2f \n", median)
	return output
}
