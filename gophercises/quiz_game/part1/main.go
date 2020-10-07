package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Read input arguments
	csvFilePath := flag.String("file", "../problems.csv", "Questions file path")
	flag.Parse()

	rans := runQuiz(os.Stdin, *csvFilePath)

	// Results
	fmt.Printf("You got %d answers right!!!", rans)
}
