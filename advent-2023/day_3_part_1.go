package advent_2023

import (
	"fmt"
	"github.com/cksidharthan/advent-of-code/elf"
	"log"
)

func Day03PartOne(test bool) int {
	var file string
	var err error

	if test {
		file, err = elf.GetInputFile(true, 2023, 3, 1)
	} else {
		file, err = elf.GetInputFile(false, 2023, 3, 1)
	}

	if err != nil {
		log.Fatal(err)
	}

	contents, err := elf.GetContentsFromFile(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(contents)

	return 0
}
