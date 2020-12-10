package main

import (
	"log"

	filereader "github.com/ultran0ise/shared"
)

func main() {
	filePath := "q01_input.txt"
	result, _ := Find2020(filePath)
	log.Println("Q 01 result:", result)
}

// Find2020 finds two integers which have sum of 2020 then product them
func Find2020(filePath string) (int64, error) {
	allValues, _ := filereader.IntFromReader(filePath)
	for i := 0; i < len(allValues); i++ {
		currentValue := allValues[i]
		diffValue := 2020 - currentValue

		for idx := range allValues {
			if allValues[idx] != diffValue {
				continue
			}
			return (allValues[idx] * currentValue), nil
		}
	}
	return 0, nil
}
