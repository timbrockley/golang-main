/*

Copyright 2024, Tim Brockley. All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

*/

package tui

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

//------------------------------------------------------------

const (
	Black   = 30
	Red     = 31
	Green   = 32
	Yellow  = 33
	Blue    = 34
	Magenta = 35
	Cyan    = 36
	White   = 37
	Default = 39

	BrightBlack   = 90
	BrightRed     = 91
	BrightGreen   = 92
	BrightYellow  = 93
	BrightBlue    = 94
	BrightMagenta = 95
	BrightCyan    = 96
	BrightWhite   = 97
	BrightDefault = 99

	Bold      = 1
	Dim       = 2
	Underline = 4
	Blink     = 5
	Reverse   = 7
	Hide      = 8

	BoldOff      = 21
	DimOff       = 22
	UnderlineOff = 24
	BlinkOff     = 25
	ReverseOff   = 27
	HideOff      = 28
)

//------------------------------------------------------------

// carriage return
func CR(OptionsMap ...map[string]any) string {
	return Render("\r", OptionsMap...)
}

// line feed
func LF(OptionsMap ...map[string]any) string {
	return Render("\n", OptionsMap...)
}

// carriage return + line feed
func CRLF(OptionsMap ...map[string]any) string {
	return Render("\r\n", OptionsMap...)
}

//------------------------------------------------------------

func CursorUp(n int, OptionsMap ...map[string]any) string {
	return Render(fmt.Sprintf("\033[%dA", n), OptionsMap...)
}

func CursorDown(n int, OptionsMap ...map[string]any) string {
	return Render(fmt.Sprintf("\033[%dB", n), OptionsMap...)
}

func CursorRight(n int, OptionsMap ...map[string]any) string {
	return Render(fmt.Sprintf("\033[%dC", n), OptionsMap...)
}

func CursorLeft(n int, OptionsMap ...map[string]any) string {
	return Render(fmt.Sprintf("\033[%dD", n), OptionsMap...)
}

//------------------------------------------------------------

func CursorHome(OptionsMap ...map[string]any) string {
	return Render("\033[H", OptionsMap...)
}

func CursorMove(row int, col int, OptionsMap ...map[string]any) string {
	return Render(fmt.Sprintf("\033[%d;%dH", row, col), OptionsMap...)
}

func CursorSave(OptionsMap ...map[string]any) string {
	return Render("\033[s", OptionsMap...)
}

func CursorRestore(OptionsMap ...map[string]any) string {
	return Render("\033[u", OptionsMap...)
}

//------------------------------------------------------------

func CursorShow(OptionsMap ...map[string]any) string {
	return Render("\033[?25h", OptionsMap...)
}

func CursorHide(OptionsMap ...map[string]any) string {
	return Render("\033[?25l", OptionsMap...)
}

//------------------------------------------------------------

/*

	outputs to then reads from stdin

	2024-07-01: worked ok when run in non-testing environment but failed in testing

*/

func CursorRowCol() (int, int, error) {
	//----------------------------------------
	var oldState *term.State
	var err error
	var row, col int
	//----------------------------------------
	// put terminal into raw mode to read individual key presses
	oldState, err = term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return row, col, err
	}
	//----------------------------------------
	// escape sequence to request row and col
	_, err = fmt.Print("\033[6n")
	if err != nil {
		return row, col, err
	}
	//----------------------------------------
	// read row and col from stdin
	_, err = fmt.Scanf("\033[%d;%dR", &row, &col)
	if err != nil {
		return row, col, err
	}
	//----------------------------------------
	// restore old terminal state
	err = term.Restore(int(os.Stdin.Fd()), oldState)
	if err != nil {
		return row, col, err
	}
	//----------------------------------------
	return row, col, nil
	//----------------------------------------
}

//------------------------------------------------------------

func ScrollUp(n int, OptionsMap ...map[string]any) string {
	return Render(fmt.Sprintf("\033[%dS", n), OptionsMap...)
}

func ScrollDown(n int, OptionsMap ...map[string]any) string {
	return Render(fmt.Sprintf("\033[%dT", n), OptionsMap...)
}

//------------------------------------------------------------

func ClearScrollbackBuffer(OptionsMap ...map[string]any) string {
	return Render("\033[3J", OptionsMap...)
}

// clear visible screen only (not scrollback buffer)
func ClearWindow(OptionsMap ...map[string]any) string {
	return Render("\033[2J\033[H", OptionsMap...)
}

// Clear screen and scrollback buffer
func ClearScreen(OptionsMap ...map[string]any) string {
	return Render("\033[2J\033[3J\033[H", OptionsMap...)
}

// clear line and return cursor to first column
func ClearLine(OptionsMap ...map[string]any) string {
	return Render("\033[2K", OptionsMap...)
}

//------------------------------------------------------------

// Enable Alternative Screen
func AlternativeScreenEnable(OptionsMap ...map[string]any) string {
	return Render("\033[?1049h", OptionsMap...)
}

// Disable Alternative Screen
func AlternativeScreenDisable(OptionsMap ...map[string]any) string {
	return Render("\033[?1049l", OptionsMap...)
}

//------------------------------------------------------------

func Colour(effectValue byte, Background ...bool) string {
	if len(Background) > 0 && Background[0] {
		effectValue += 10
	}
	return Render(fmt.Sprintf("\033[%dm", effectValue))
}

func Colour256(colourValue byte, Background ...bool) string {
	effectValue := 38
	if len(Background) > 0 && Background[0] {
		effectValue = 48
	}
	return Render(fmt.Sprintf("\033[%d;5;%dm", effectValue, colourValue))
}

func ColourRGB(R, G, B byte, Background ...bool) string {
	effectValue := 38
	if len(Background) > 0 && Background[0] {
		effectValue = 48
	}
	return Render(fmt.Sprintf("\033[%d;2;%d;%d;%dm", effectValue, R, G, B))
}

func Effect(effectValue byte) string {
	return Render(fmt.Sprintf("\033[%dm", effectValue))
}

func Reset(OptionsMap ...map[string]any) string {
	return Render("\033[0m", OptionsMap...)
}

//------------------------------------------------------------
