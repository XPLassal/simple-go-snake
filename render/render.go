package render

import (
	"strconv"
	"strings"

	. "github.com/XPLassal/simple-snake-on-go/structs"
)

func RenderField(numberOfColumns int, apple *Apple, snake *Snake, sb *strings.Builder) {
	sb.WriteString("\033[H")

	sb.WriteString(Bold)
	sb.WriteString("Your Score: ")
	sb.WriteString(strconv.Itoa(snake.GetLen()))
	sb.WriteString(Reset)
	sb.WriteRune('\n')

	DrawBordersForY(sb, numberOfColumns)
	sb.WriteRune('\n')

	var isDark bool
	for y := 0; y < numberOfColumns; y++ {
		DrawBordersForX(sb)

		for x := 0; x < numberOfColumns; x++ {
			coords := SetCoordinates(x, y)
			isHeadOfSnake, isHaveSnake := snake.Contains(coords)
			isHaveApple := apple.Contains(coords)

			DrawBg(sb, isHaveApple, isDark, isHaveSnake, isHeadOfSnake)

			isDark = !isDark
		}

		DrawBordersForX(sb)
		sb.WriteRune('\n')
	}

	DrawBordersForY(sb, numberOfColumns)
	sb.WriteString("\nTo exit, press q.\n")
}
