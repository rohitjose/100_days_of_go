package main

import (
	"bytes"
	"testing"
)

func TestRunQuiz(t *testing.T) {
	csvFilePath := "../problems.csv"
	tests := []struct {
		name  string
		input []byte
		want  int
	}{
		{
			name:  "all_valid",
			input: []byte("10\n2\n11\n3\n14\n4\n5\n6\n5\n6\n6\n7\n"),
			want:  12,
		},
		{
			name:  "one_invalid",
			input: []byte("9\n2\n11\n3\n14\n4\n5\n6\n5\n6\n6\n7\n"),
			want:  11,
		},
		{
			name:  "all_invalid",
			input: []byte("9\n1\n1\n1\n1\n1\n1\n1\n1\n1\n1\n1\n"),
			want:  0,
		},
	}
	for _, tc := range tests {
		var stdin bytes.Buffer
		stdin.Write(tc.input)
		result := runQuiz(&stdin, csvFilePath)
		if result != tc.want {
			t.Errorf("Want %d Got %d", tc.want, result)
		}
	}
}
