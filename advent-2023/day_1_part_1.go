package advent_2023

import (
	"github.com/cksidharthan/advent-of-code/elf"
	"log"
)

func Day01PartOne(test bool) int {
	var file string
	var err error

	file, err = elf.GetInputFile(test, 2023, 1, 1)

	if err != nil {
		log.Fatal(err)
	}

	contents, err := elf.GetContentsFromFile(file)
	if err != nil {
		log.Fatal(err)
	}

	firstAndLastIntList := make([]int, 0)
	for _, line := range contents {
		lineIntString := getIntegersFromString(line)
		firstAndLastIntList = append(firstAndLastIntList, getFirstAndLastIntsFromString(lineIntString))
	}

	// sum all the ints in the list
	var sum int
	for _, num := range firstAndLastIntList {
		sum += num
	}

	return sum
}

func getIntegersFromString(line string) string {
	var resultString string
	for _, char := range line {
		if char >= '0' && char <= '9' {
			resultString += string(char)
		}
	}

	return resultString
}

func getFirstAndLastIntsFromString(line string) int {
	firstInt := line[0]
	lastInt := line[len(line)-1]

	resultString := string(firstInt) + string(lastInt)
	result := elf.ConvertStringToInt(resultString)

	return result
}
