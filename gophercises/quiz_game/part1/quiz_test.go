package main

import (
	"bytes"
	"testing"
)

func TestRunQuiz(t *testing.T) {
	// Initialise the user input
	var stdin bytes.Buffer
	stdin.Write([]byte("10\n"))
	stdin.Write([]byte("2\n"))
	stdin.Write([]byte("11\n"))
	stdin.Write([]byte("3\n"))
	stdin.Write([]byte("14\n"))
	stdin.Write([]byte("4\n"))
	stdin.Write([]byte("5\n"))
	stdin.Write([]byte("6\n"))
	stdin.Write([]byte("5\n"))
	stdin.Write([]byte("6\n"))
	stdin.Write([]byte("6\n"))
	stdin.Write([]byte("7\n"))

	csvFilePath := "../problems.csv"
	result := runQuiz(&stdin, csvFilePath)
	if result != 12 {
		t.Errorf("Want 12 Got %d", result)
	}
}
