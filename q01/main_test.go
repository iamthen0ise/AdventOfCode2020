package main

import (
	"testing"

	filereader "github.com/ultran0ise/shared"
)

func TestUrlToLines(t *testing.T) {
	filePath := "q01_test_input.txt"
	res, _ := filereader.IntFromReader(filePath)

	acc := 0
	for i := range res {
		acc += i
	}
	if acc != 2020 {
		t.Errorf("Sum for Q01 not 2020")
	}
}
