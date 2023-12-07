package advent_2023

import (
	"fmt"
	"testing"
)

func TestExampleDay01PartTwo(t *testing.T) {
	fmt.Println("Advent of Code Day 01 - Part 02!")
	result := Day01PartTwo(true)
	if result != 281 {
		t.Errorf("Expected 142, got %d", result)
	}
}
