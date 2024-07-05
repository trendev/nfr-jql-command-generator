package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

type Labels struct {
	Labels []string `json:"labels"`
}

func main() {
	// Read the JSON file
	jf, err := os.Open("labels.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jf.Close()

	byteValue, err := io.ReadAll(jf)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	var labels Labels
	err = json.Unmarshal(byteValue, &labels)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Sort the labels alphabetically
	sort.Strings(labels.Labels)

	// Create a CSV f
	f, err := os.Create("NFR_JQL_Scripts.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	// Write the header
	w.Write([]string{"Label", "JQL Command"})

	// Iterate over the labels and generate the JQL commands
	for _, l := range labels.Labels {
		jql := fmt.Sprintf("labels = %q ORDER BY created DESC", l)
		w.Write([]string{l, jql})
	}

	fmt.Println("CSV file successfully created.")
}
