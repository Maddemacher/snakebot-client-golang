package utils

import "math"

func ManhattanDistance(x int, y int, targetX int, targetY int) int {
	return int(math.Abs(float64(x)-float64(targetX)) + math.Abs(float64(y)-float64(targetY)))
}
