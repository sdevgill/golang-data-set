package main

import (
	"encoding/json"
	"io/ioutil"
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
