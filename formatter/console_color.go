package formatter

import (
	"strconv"
)

// 控制控制台颜色
const (
	CONSOLE_BLACK   = 30
	CONSOLE_RED     = 31
	CONSOLE_GREEN   = 32
	CONSOLE_YELLOW  = 33
	CONSOLE_BLUE    = 34
	CONSOLE_FUCHSIA = 35 // 紫红色
	CONSOLE_CYAN    = 36 // 青蓝色
	CONSOLE_WHITE   = 37

	CONSOLE_BG_BLACK   = 40
	CONSOLE_BG_RED     = 41
	CONSOLE_BG_GREEN   = 42
	CONSOLE_BG_YELLOW  = 43
	CONSOLE_BG_BLUE    = 44
	CONSOLE_BG_FUCHSIA = 45 // 紫红色
	CONSOLE_BG_CYAN    = 46 // 青蓝色
	CONSOLE_BG_WHITE   = 47
	CONSOLE_BG_TRANS   = 1
)

func ConsoleColor(color, bgColor int) string {
	return "\033[1;" + strconv.Itoa(color) + ";" + strconv.Itoa(bgColor) + "m"
}

func ConsoleColorReset() string {
	return "\033[0m"
}

func ConsoleText(text string, color, bgColor int) string {
	return ConsoleColor(color, bgColor) + text + ConsoleColorReset()
}
