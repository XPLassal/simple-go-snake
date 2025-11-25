package render

import "fmt"

const (
	Reset = "\033[0m"

	Bold = "\033[1m"

	Black   = "\033[30m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"

	BrightBlack   = "\033[90m"
	BrightRed     = "\033[91m"
	BrightGreen   = "\033[92m"
	BrightYellow  = "\033[93m"
	BrightBlue    = "\033[94m"
	BrightMagenta = "\033[95m"
	BrightCyan    = "\033[96m"
	BrightWhite   = "\033[97m"

	BgBlack   = "\033[40m"
	BgRed     = "\033[41m"
	BgGreen   = "\033[42m"
	BgYellow  = "\033[43m"
	BgBlue    = "\033[44m"
	BgMagenta = "\033[45m"
	BgCyan    = "\033[46m"
	BgWhite   = "\033[47m"

	BgBrightBlack   = "\033[100m"
	BgBrightRed     = "\033[101m"
	BgBrightGreen   = "\033[102m"
	BgBrightYellow  = "\033[103m"
	BgBrightBlue    = "\033[104m"
	BgBrightMagenta = "\033[105m"
	BgBrightCyan    = "\033[106m"
	BgBrightWhite   = "\033[107m"

	BgDarkGreen   = "\033[48;5;22m"  // Темно-зеленый
	BgDarkRed     = "\033[48;5;52m"  // Темно-красный
	BgDarkBlue    = "\033[48;5;17m"  // Темно-синий
	BgDarkYellow  = "\033[48;5;94m"  // Темно-желтый (охровый)
	BgDarkMagenta = "\033[48;5;53m"  // Темно-пурпурный
	BgDarkCyan    = "\033[48;5;30m"  // Темно-бирюзовый
	BgGray        = "\033[48;5;236m" // Темно-серый
	BgLightGray   = "\033[48;5;246m" // Светло-серый
)

func ClearConsole() {
	fmt.Print("\033[H\033[2J")
}
