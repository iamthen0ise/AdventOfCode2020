package filereader

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// IntFromReader Formats data from Reader into array of int64
func IntFromReader(filePath string) ([]int64, error) {
	var lines []int64

	f, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		curVal, _ := strconv.ParseInt(scanner.Text(), 10, 32)
		lines = append(lines, curVal)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// StrFromReader Formats data from Reader into array of strings
func StrFromReader(filePath string) ([]string, error) {
	var lines []string

	f, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		curVal := scanner.Text()
		lines = append(lines, curVal)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
