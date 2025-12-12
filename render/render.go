package render

import (
	"strconv"
	"strings"

	. "github.com/XPLassal/simple-go-snake/structs"
)

func RenderField(numberOfColumns int, apple *Apple, snake *Snake, sb *strings.Builder, withEmoji bool) {
	sb.WriteString("\033[H")

	sb.WriteString(Bold)
	sb.WriteString("Your Score: ")
	sb.WriteString(strconv.Itoa(snake.GetLen()))
	sb.WriteString(Reset)
	sb.WriteRune('\n')

	DrawBordersForY(sb, numberOfColumns)
	sb.WriteRune('\n')

	var isDark bool
	for y := range numberOfColumns {
		DrawBordersForX(sb)

		for x := range numberOfColumns {
			coords := SetCoordinates(x, y)
			isHeadOfSnake, isHaveSnake := snake.Contains(coords)
			isHaveApple := apple.Contains(coords)

			DrawBg(sb, isHaveApple, isDark, isHaveSnake, isHeadOfSnake, withEmoji)

			isDark = !isDark
		}

		DrawBordersForX(sb)
		sb.WriteRune('\n')
	}

	DrawBordersForY(sb, numberOfColumns)
	sb.WriteRune('\n')
	keys := []string{"q: Exit", "c: Config", "p: Pause"}
	for _, v := range keys {
		sb.WriteString(v)
		sb.WriteByte(' ')
	}
	sb.WriteRune('\n')
}
