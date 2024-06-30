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

func ClearLine(optionFuncs ...OptionFunc) string {
	return ReturnOutput("\033[2K", optionFuncs...)
}

//--------------------------------------------------------------------------------

func CursorUp(n int, optionFuncs ...OptionFunc) string {
	return ReturnOutput(fmt.Sprintf("\033[%dA", n), optionFuncs...)
}

func CursorDown(n int, optionFuncs ...OptionFunc) string {
	return ReturnOutput(fmt.Sprintf("\033[%dB", n), optionFuncs...)
}

//--------------------------------------------------------------------------------

func CursorHide(optionFuncs ...OptionFunc) string {
	return ReturnOutput("\033[?25l", optionFuncs...)
}

func CursorShow(optionFuncs ...OptionFunc) string {
	return ReturnOutput("\033[?25h", optionFuncs...)
}

//--------------------------------------------------------------------------------
