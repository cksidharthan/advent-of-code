package advent_2023

import (
	"fmt"
	"testing"
)

func TestExampleDay01PartOne(t *testing.T) {
	fmt.Println("Advent of Code Day 01 - Part 01!")
	result := Day01PartOne(true)
	if result != 142 {
		t.Errorf("Expected 142, got %d", result)
	}
}
