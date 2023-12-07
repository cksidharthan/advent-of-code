package advent_2023

import (
	"fmt"
	"testing"
)

func TestExampleDay02PartTwo(t *testing.T) {
	fmt.Println("Advent of Code Day 01 - Part 02!")
	result := Day02PartTwo(true)
	if result != 2286 {
		t.Errorf("Expected 2286, got %d", result)
	}
}
