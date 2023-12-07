package advent_2023

import (
	"fmt"
	"testing"
)

func TestExampleDay02PartOne(t *testing.T) {
	fmt.Println("Advent of Code Day 01 - Part 02!")
	result := Day02PartOne(true, 12, 14, 13)
	if result != 8 {
		t.Errorf("Expected 8, got %d", result)
	}
}
