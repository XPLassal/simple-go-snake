package render

import (
	"strings"
)

func DrawBg(sb *strings.Builder, haveApple, isDarkBg, isSnake, isHeadOfSnake bool) {
	if isDarkBg {
		sb.WriteString(BgDarkGreen)
	} else {
		sb.WriteString(BgGreen)
	}

	if isSnake {
		if isHeadOfSnake {
			sb.WriteString("🐍")
		} else {
			sb.WriteString("🟢")
		}
	} else if haveApple {
		sb.WriteString("🍎")
	} else {
		sb.WriteString("  ")
	}
	sb.WriteString(Reset)
}

const BgForBorder = BgWhite

func DrawBordersForX(sb *strings.Builder) {
	sb.WriteString(BgForBorder)
	sb.WriteString("  ")
	sb.WriteString(Reset)
}

func DrawBordersForY(sb *strings.Builder, size int) {
	for i := 0; i < size+2; i++ {
		sb.WriteString(BgForBorder)
		sb.WriteString("  ")
	}
	sb.WriteString(Reset)
}
