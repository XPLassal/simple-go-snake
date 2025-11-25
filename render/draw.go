package render

import (
	"strings"
)

func DrawBg(haveApple bool, isDarkBg bool, isSnake bool, isHeadOfSnake bool) string {
	returnTemplate := func(isDark bool, s string) string {
		if isDark {
			return BgDarkGreen + s + Reset
		} else {
			return BgGreen + s + Reset
		}
	}

	var valueForReturn string

	if isSnake {
		if isHeadOfSnake {
			valueForReturn = "🐍"
		} else {
			valueForReturn = "🟢"
		}
	} else if haveApple {
		valueForReturn = "🍎"
	} else {
		valueForReturn = "  "
	}

	return returnTemplate(isDarkBg, valueForReturn)
}

const BgForBorder = BgWhite

func DrawBordersForX() string {
	return BgForBorder + "  " + Reset
}

func DrawBordersForY(size int) string {
	return strings.Repeat(BgForBorder+"  ", size+2) + Reset
}
