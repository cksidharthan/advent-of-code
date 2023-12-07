package advent_2023

import (
	"testing"
)

func TestExampleDay01PartOne(t *testing.T) {
	t.Parallel()

	result := Day01PartOne(true)
	if result != 142 {
		t.Errorf("Expected 142, got %d", result)
	}
}
