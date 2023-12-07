package main

import (
	"fmt"
	advent2023 "github.com/cksidharthan/advent-of-code/advent-2023"
)

func main() {
	fmt.Println("Advent of Code!")

	result := advent2023.Day01PartOne(false)
	fmt.Println("Advent 2023 - Day 01 - Part 01 - Result: ", result)

	result = advent2023.Day01PartTwo(false)
	fmt.Println("Advent 2023 - Day 01 - Part 02 - Result: ", result)

	result = advent2023.Day02PartOne(false, 12, 14, 13)
	fmt.Println("Advent 2023 - Day 02 - Part 01 - Result: ", result)
}
