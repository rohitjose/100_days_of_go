package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	// Read the csv file for the quiz problems
	// Open the CSV file
	problemsFile, err := os.Open("../problems.csv")
	if err != nil {
		log.Fatalln("Could not open the CSV file", err)
	}

	// Parse the file
	r := csv.NewReader(problemsFile)

	// Iterate through the records
	for {
		// Read each record from csv file
		q, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Question: %s Answer: %s\n", q[0], q[1])
	}
}
