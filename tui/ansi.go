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

// carriage return
func CR(optionFuncs ...OptionFunc) string {
	return ReturnOutput("\r", optionFuncs...)
}

// line feed
func LF(optionFuncs ...OptionFunc) string {
	return ReturnOutput("\n", optionFuncs...)
}

// carriage return + line feed
func CRLF(optionFuncs ...OptionFunc) string {
	return ReturnOutput("\r\n", optionFuncs...)
}

//------------------------------------------------------------

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

//------------------------------------------------------------

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

//------------------------------------------------------------

func CursorShow(optionFuncs ...OptionFunc) string {
	return ReturnOutput("\033[?25h", optionFuncs...)
}

func CursorHide(optionFuncs ...OptionFunc) string {
	return ReturnOutput("\033[?25l", optionFuncs...)
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

func ScrollUp(n int, optionFuncs ...OptionFunc) string {
	return ReturnOutput(fmt.Sprintf("\033[%dS", n), optionFuncs...)
}

func ScrollDown(n int, optionFuncs ...OptionFunc) string {
	return ReturnOutput(fmt.Sprintf("\033[%dT", n), optionFuncs...)
}

//------------------------------------------------------------

func ClearScrollbackBuffer(optionFuncs ...OptionFunc) string {
	return ReturnOutput("\033[3J", optionFuncs...)
}

// clear visible screen only (not scrollback buffer)
func ClearWindow(optionFuncs ...OptionFunc) string {
	return ReturnOutput("\033[2J\033[H", optionFuncs...)
}

// Clear screen and scrollback buffer
func ClearScreen(optionFuncs ...OptionFunc) string {
	return ReturnOutput("\033[2J\033[3J\033[H", optionFuncs...)
}

func ClearLine(optionFuncs ...OptionFunc) string {
	return ReturnOutput("\033[2K\r", optionFuncs...)
}

//------------------------------------------------------------
