package advent_2023

import (
	"fmt"
	"testing"
)

func TestExampleDay03PartOne(t *testing.T) {
	fmt.Println("Advent of Code Day 03 - Part 01!")
	result := Day03PartOne(true)
	if result != 4361 {
		t.Errorf("Expected 4361, got %d", result)
	}
}
