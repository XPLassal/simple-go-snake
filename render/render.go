package render

import (
	"fmt"
	"strconv"
	"strings"

	. "github.com/XPLassal/simple-snake-on-go/structs"
)

func RenderField(numberOfColumns int, apple *Apple, snake *Snake) {
	ClearConsole()

	var field strings.Builder
	var isHaveApple, isHaveSnake, isHeadOfSnake, isDark bool

	field.WriteString(Bold + "Your Score: " + strconv.Itoa(snake.GetLen()) + Reset + "\n")
	field.WriteString(DrawBordersForY(numberOfColumns) + "\n")

	for y := range numberOfColumns {
		field.WriteString(DrawBordersForX())
		for x := range numberOfColumns {
			coords := SetCoordinates(x, y)
			isHeadOfSnake, isHaveSnake = snake.Contains(coords)
			isHaveApple = apple.Contains(coords)
			field.WriteString(DrawBg(isHaveApple, isDark, isHaveSnake, isHeadOfSnake))
			isDark = !isDark
		}
		field.WriteString(DrawBordersForX() + "\n")
	}

	field.WriteString(DrawBordersForY(numberOfColumns))
	fmt.Println(field.String(), "\nTo exit, press q.")
}
