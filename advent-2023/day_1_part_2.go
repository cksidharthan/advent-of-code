package advent_2023

import (
	"fmt"
	"github.com/cksidharthan/advent-of-code/elf"
	"log"
	"strings"
)

func Day01PartTwo(test bool) int {
	var file string
	var err error

	if test {
		file, err = elf.GetInputFile(true, 2023, 1, 2)
	} else {
		file, err = elf.GetInputFile(false, 2023, 1, 2)
	}

	if err != nil {
		log.Fatal(err)
	}

	contents, err := elf.GetContentsFromFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lineNumList := make([]int, 0)

	for _, line := range contents {
		firstDigit := getFirstDigit(line)
		lastDigit := getLastDigit(line)

		if firstDigit == lastDigit {
			lineNum := elf.ConvertStringToInt(fmt.Sprintf("%d%d", firstDigit, lastDigit))
			lineNumList = append(lineNumList, lineNum)

			continue
		}

		lineNum := elf.ConvertStringToInt(fmt.Sprintf("%d%d", firstDigit, lastDigit))
		lineNumList = append(lineNumList, lineNum)
	}

	// sum all the ints in the list
	var sum int
	for _, num := range lineNumList {
		sum += num
	}

	return sum
}

func getFirstDigit(s string) int {
	// start with the first character in the string, append it to tempstring, if the temp string is present in the map, return the value, if the current character is a difit, return the digit.
	tempString := ""
	for _, char := range s {
		if char >= '0' && char <= '9' {
			return elf.ConvertStringToInt(string(char))
		}
		tempString += string(char)
		for key, val := range elf.DigitsMap {
			if strings.Contains(tempString, key) {
				return val
			}
		}
	}

	return 0
}

func getLastDigit(s string) int {
	// start with the last character in the string, append it to tempstring, if the temp string is present in the map, return the value, if the current character is a difit, return the digit.
	tempString := ""
	for i := len(s) - 1; i >= 0; i-- {
		char := s[i]
		if char >= '0' && char <= '9' {
			return elf.ConvertStringToInt(string(char))
		}
		tempString = string(char) + tempString
		for key, val := range elf.DigitsMap {
			if strings.Contains(tempString, key) {
				return val
			}
		}
	}

	return 0
}
