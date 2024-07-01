/*

Copyright 2024, Tim Brockley. All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

*/

package tui

import (
	"fmt"
)

//--------------------------------------------------------------------------------

func CR(optionFuncs ...OptionFunc) string { // carriage return
	return ReturnOutput("\r", optionFuncs...)
}

func LF(optionFuncs ...OptionFunc) string { // line feed
	return ReturnOutput("\n", optionFuncs...)
}

func CRLF(optionFuncs ...OptionFunc) string {
	return ReturnOutput("\r\n", optionFuncs...)
}

//--------------------------------------------------------------------------------

func CursorUp(n int, optionFuncs ...OptionFunc) string {
	return ReturnOutput(fmt.Sprintf("\033[%dA", n), optionFuncs...)
}

func CursorDown(n int, optionFuncs ...OptionFunc) string {
	return ReturnOutput(fmt.Sprintf("\033[%dB", n), optionFuncs...)
}

func CursorRight(n int, optionFuncs ...OptionFunc) string {
	return ReturnOutput(fmt.Sprintf("\033[%dC", n), optionFuncs...)
}

func CursorLeft(n int, optionFuncs ...OptionFunc) string {
	return ReturnOutput(fmt.Sprintf("\033[%dD", n), optionFuncs...)
}

//--------------------------------------------------------------------------------

func CursorHome(optionFuncs ...OptionFunc) string {
	return ReturnOutput("\033[H", optionFuncs...)
}

func CursorMove(row int, col int, optionFuncs ...OptionFunc) string {
	return ReturnOutput(fmt.Sprintf("\033[%d;%dH", row, col), optionFuncs...)
}
func CursorSave(optionFuncs ...OptionFunc) string {
	return ReturnOutput("\033[s", optionFuncs...)
}

func CursorRestore(optionFuncs ...OptionFunc) string {
	return ReturnOutput("\033[u", optionFuncs...)
}

//--------------------------------------------------------------------------------

func CursorShow(optionFuncs ...OptionFunc) string {
	return ReturnOutput("\033[?25h", optionFuncs...)
}

func CursorHide(optionFuncs ...OptionFunc) string {
	return ReturnOutput("\033[?25l", optionFuncs...)
}

//--------------------------------------------------------------------------------

func ScrollUp(n int, optionFuncs ...OptionFunc) string {
	return ReturnOutput(fmt.Sprintf("\033[%dS", n), optionFuncs...)
}

func ScrollDown(n int, optionFuncs ...OptionFunc) string {
	return ReturnOutput(fmt.Sprintf("\033[%dT", n), optionFuncs...)
}

//--------------------------------------------------------------------------------

func ClearScrollbackBuffer(optionFuncs ...OptionFunc) string {
	return ReturnOutput("\033[3J", optionFuncs...)
}

// clear visible screen only (not scrollback buffer)
func ClearWindow(optionFuncs ...OptionFunc) string {
	return ReturnOutput("\033[2J", optionFuncs...)
}

// Clear screen and scrollback buffer
func ClearScreen(optionFuncs ...OptionFunc) string {
	return ReturnOutput("\033[2J\033[3J", optionFuncs...)
}

func ClearLine(optionFuncs ...OptionFunc) string {
	return ReturnOutput("\033[2K", optionFuncs...)
}

//--------------------------------------------------------------------------------
