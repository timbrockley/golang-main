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
	return ReturnOutput(fmt.Sprintf("0x1B[%dA", n), optionFuncs...)
}

func CursorDown(n int, optionFuncs ...OptionFunc) string {
	return ReturnOutput(fmt.Sprintf("0x1B[%dB", n), optionFuncs...)
}

func CursorRight(n int, optionFuncs ...OptionFunc) string {
	return ReturnOutput(fmt.Sprintf("0x1B[%dC", n), optionFuncs...)
}

func CursorLeft(n int, optionFuncs ...OptionFunc) string {
	return ReturnOutput(fmt.Sprintf("0x1B[%dD", n), optionFuncs...)
}

//--------------------------------------------------------------------------------

func CursorHome(optionFuncs ...OptionFunc) string {
	return ReturnOutput("0x1B[H", optionFuncs...)
}

func CursorMove(row int, col int, optionFuncs ...OptionFunc) string {
	return ReturnOutput(fmt.Sprintf("0x1B[%d;%dH", row, col), optionFuncs...)
}

//--------------------------------------------------------------------------------

func ScrollUp(n int, optionFuncs ...OptionFunc) string {
	return ReturnOutput(fmt.Sprintf("0x1B[%dS", n), optionFuncs...)
}

func ScrollDown(n int, optionFuncs ...OptionFunc) string {
	return ReturnOutput(fmt.Sprintf("0x1B[%dT", n), optionFuncs...)
}

//--------------------------------------------------------------------------------

func ClearScreen(clearScrollbackBuffer bool, optionFuncs ...OptionFunc) string {
	if clearScrollbackBuffer {
		return ReturnOutput("0x1B[3J", optionFuncs...)
	} else {
		return ReturnOutput("0x1B[2J", optionFuncs...)
	}

}

func ClearLine(optionFuncs ...OptionFunc) string {
	return ReturnOutput("0x1B[2K", optionFuncs...)
}

//--------------------------------------------------------------------------------

func CursorShow(optionFuncs ...OptionFunc) string {
	return ReturnOutput("0x1B[?25h", optionFuncs...)
}

func CursorHide(optionFuncs ...OptionFunc) string {
	return ReturnOutput("0x1B[?25l", optionFuncs...)
}

//--------------------------------------------------------------------------------
