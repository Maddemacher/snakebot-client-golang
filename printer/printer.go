package printer

import (
	"fmt"
	"strings"

	"github.com/fatih/color"

	"../common"
	"../communication"
)

var colors = []color.Attribute{color.FgRed, color.FgBlue, color.FgMagenta, color.FgCyan, color.FgCyan}
var snakeIds = make(common.Ids, 0)

func GetSnakeColor(id common.Id) color.Attribute {
	hasColor, pos := snakeIds.Contains(id)

	if hasColor {
		return colors[pos]
	}

	snakeIds = append(snakeIds, id)
	return colors[len(snakeIds)-1]
}

func GetTileContent(m communication.Map, row int, column int) (string, color.Attribute) {

	coord := row*m.Width + column

	c, _ := m.FoodPositions.Contains(coord)
	if c {
		return "F", color.FgGreen
	}

	c, _ = m.ObstaclePositions.Contains(coord)
	if c {
		return "O", color.FgYellow
	}

	for _, snakeInfo := range m.SnakeInfos {
		if len(snakeInfo.Positions) > 0 && coord == snakeInfo.Positions[0] {
			return "@", GetSnakeColor(snakeInfo.Id)
		}

		c, _ = snakeInfo.Positions.Contains(coord)
		if c {
			return "#", GetSnakeColor(snakeInfo.Id)
		}
	}

	return " ", color.FgWhite
}

func PrintMap(m communication.Map) {
	fmt.Println(strings.Repeat("-", m.Width))
	for r := 0; r < m.Height; r++ {
		for c := 0; c < m.Width; c++ {
			content, col := GetTileContent(m, r, c)

			color.Set(col)
			fmt.Print(content)
		}

		color.Unset()
		fmt.Print("\n")
	}

	fmt.Println(strings.Repeat("-", m.Width))
}
