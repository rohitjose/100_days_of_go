package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// Read input arguments
	csvFilePath := flag.String("file", "../problems.csv", "Questions file path")
	flag.Parse()

	// Read the csv file for the quiz problems
	// Open the CSV file
	problemsFile, err := os.Open(*csvFilePath)
	if err != nil {
		log.Fatalln("Could not open the CSV file", err)
	}

	// Parse the file
	r := csv.NewReader(problemsFile)
	rans := 0

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

		// Read Input
		fmt.Printf("What is %s ? ", q[0])
		reader := bufio.NewReader(os.Stdin)
		ans, _ := reader.ReadString('\n')
		// Trim the endline character
		cans := strings.Trim(ans, " \n")

		if cans == q[1] {
			rans++
		}
	}

	// Results
	fmt.Printf("You got %d answers right!!!", rans)
}
