package tests

import (
	"testing"

	"../utils"
)

func TestManhattanDistance(t *testing.T) {
	value := utils.ManhattanDistance(0, -1, 10, -3)

	if value != 12 {
		t.Error("Expected 12 got ", value)
	}
}
