package terminal

import (
	"fmt"
)

// HideCursor hides the terminal cursor
func HideCursor() {
	fmt.Print("\033[?25l")
}

// ShowCursor shows the terminal cursor
func ShowCursor() {
	fmt.Print("\033[?25h")
}

// MoveCursorHome moves cursor to top-left position
func MoveCursorHome() {
	fmt.Print("\033[H")
}

// ClearScreen clears the entire screen
func ClearScreen() {
	fmt.Print("\033[2J\033[H")
}