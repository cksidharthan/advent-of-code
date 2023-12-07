package advent_2023

import (
	"testing"
)

func TestExampleDay01PartTwo(t *testing.T) {
	t.Parallel()

	result := Day01PartTwo(true)
	if result != 281 {
		t.Errorf("Expected 142, got %d", result)
	}
}
