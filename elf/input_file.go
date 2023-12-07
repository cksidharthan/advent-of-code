package elf

import (
	"fmt"
	"os"
	"strconv"
)

func GetInputFile(test bool, day, part int) (string, error) {
	// get the root of the project
	root, err := os.Getwd()
	if err != nil {
		return "", err
	}

	if test {
		return fmt.Sprintf("%s/testfiles/day_%d_part_%d_test.txt", root, day, part), nil
	} else {
		return fmt.Sprintf("%s/testfiles/day_%d_part_%d.txt", root, day, part), nil
	}
}

func GetContentsFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var contents []string
	for {
		var line string
		_, err := fmt.Fscanf(file, "%s\n", &line)
		if err != nil {
			break
		}
		contents = append(contents, line)
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
