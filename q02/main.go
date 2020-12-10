package main

import (
	"log"
	"strconv"
	"strings"

	filereader "github.com/ultran0ise/shared"
)

func main() {
	filePath := "q02_input.txt"
	result, _ := CheckRules(filePath)
	log.Println("Q 02 result:", result)
}

// CheckRules ...
func CheckRules(filePath string) (int64, error) {
	mathesCount := 0
	data, _ := filereader.StrFromReader(filePath)
	for i := 0; i < len(data); i++ {
		ruleSet := strings.Split(data[i], " ")

		charRepeatCount := strings.Split(ruleSet[0], "-")
		repeatLow, _ := strconv.ParseInt(charRepeatCount[0], 10, 32)
		repeatHigh, _ := strconv.ParseInt(charRepeatCount[1], 10, 32)

		repeatChar := strings.Trim(ruleSet[1], ":")
		password := ruleSet[2]

		ruleCount := int64(strings.Count(password, repeatChar))
		if ruleCount <= repeatHigh && ruleCount >= repeatLow {
			mathesCount++
		}
	}
	return int64(mathesCount), nil
}
