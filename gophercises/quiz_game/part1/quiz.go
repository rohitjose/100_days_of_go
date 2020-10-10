package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func runQuiz(stdin io.Reader, path string) int {

	scanner := bufio.NewScanner(stdin)
	// Read the csv file for the quiz problems
	// Open the CSV file
	problemsFile, err := os.Open(path)
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
		// Read input
		var ans string
		if scanner.Scan() {
			ans = scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		// Trim the endline character
		cans := strings.Trim(ans, " \n")

		if cans == q[1] {
			rans++
		}
	}

	return rans
}
