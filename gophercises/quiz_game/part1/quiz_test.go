package main

import (
	"bytes"
	"testing"
)

func TestRunQuiz(t *testing.T) {
	// Initialise the user input
	var stdin bytes.Buffer
	stdin.Write([]byte("10\n2\n11\n3\n14\n4\n5\n6\n5\n6\n6\n7\n"))

	csvFilePath := "../problems.csv"
	result := runQuiz(&stdin, csvFilePath)
	if result != 12 {
		t.Errorf("Want 12 Got %d", result)
	}
}
