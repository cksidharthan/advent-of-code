package elf

import (
	"fmt"
	"os"
	"strconv"
)

func GetInputFile(test bool, year, day, part int) (string, error) {
	if test {
		// get the root of the project
		err := os.Chdir("..")
		if err != nil {
			return "", err
		}
		root, err := os.Getwd()
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("%s/advent-%d/testfiles/day_%d_part_%d_test.txt", root, year, day, part), nil
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
