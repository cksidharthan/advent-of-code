package elf

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var DigitsMap = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func GetInputFile(test bool, year, day, part int) (string, error) {
	if test {
		return fmt.Sprintf("testfiles/day_%d_part_%d_test.txt", day, part), nil
	} else {
		root, err := os.Getwd()
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("%s/advent-%d/testfiles/day_%d_part_%d.txt", root, year, day, part), nil
	}
}

func GetContentsFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var contents []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		contents = append(contents, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return contents, nil
}

func ConvertStringToInt(input string) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error converting string to int")
	}

	return result
}
