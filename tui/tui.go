/*

Copyright 2024, Tim Brockley. All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

*/

package tui

import (
	"fmt"
	"io"
	"os"
	"strings"
)

//--------------------------------------------------------------------------------

type OptionFunc func(*Options)

type Options struct {
	Writer          io.Writer
	Header          bool
	MaxWidth        int
	MaxColumnWidth  int
	MaxColumnWidths []int
}

//--------------------------------------------------------------------------------

func WithOutput(writer io.Writer) OptionFunc {
	return func(options *Options) {
		if writer == nil {
			options.Writer = os.Stdout
		} else {
			options.Writer = writer
		}
	}
}

func WithStdout(options *Options) { options.Writer = os.Stdout }

func WithHeader(options *Options) { options.Header = true }

func WithMaxTableWidth(maxWidth int) OptionFunc {
	return func(options *Options) {
		options.MaxWidth = maxWidth
	}
}

func WithMaxColumnWidth(maxColumnWidth int) OptionFunc {
	return func(options *Options) {
		options.MaxColumnWidth = maxColumnWidth
	}
}

func WithMaxColumnWidths(maxColumnWidths []int) OptionFunc {
	return func(options *Options) {
		options.MaxColumnWidths = maxColumnWidths
	}
}

//--------------------------------------------------------------------------------

func ParseOptions(optionFuncs ...OptionFunc) Options {
	//----------------------------------------
	options := Options{}
	for _, fn := range optionFuncs {
		fn(&options)
	}
	return options
	//----------------------------------------
}

//--------------------------------------------------------------------------------

func Render(outputString string, optionFuncs ...OptionFunc) string {
	//----------------------------------------
	options := ParseOptions(optionFuncs...)
	//----------------------------------------
	if options.MaxWidth > 0 {
		outputString = TruncateString(outputString, options.MaxWidth)
	}
	//----------------------------------------
	if options.Writer != nil {
		fmt.Fprint(options.Writer, outputString)
	}
	//----------------------------------------
	return outputString
	//----------------------------------------
}

//--------------------------------------------------------------------------------

func TruncateString(outputString string, maxWidth int) string {
	if maxWidth > 0 && len([]rune(outputString)) > maxWidth {
		return string([]rune(outputString)[:maxWidth])
	}
	return outputString
}

//--------------------------------------------------------------------------------

func EscapeString(outputString string) string {
	//----------------------------------------
	replacer := strings.NewReplacer(
		"\x5C", "\\\\", // \x5C = backslash
		"\x09", "\\t", // \x09 = tab
		"\x0A", "\\n", // \x0A = newline
		"\x0D", "\\r", // \x0D = carriage return
	// 	"\x22", "\\q", // \x22 = double quotes
	// 	"\x27", "\\a", // \x27 = apostrophe
	// 	"\x60", "\\g", // \x60 = grave accent
	)
	outputString = replacer.Replace(outputString)
	// --------------------------------------------------------------------------------
	escapedOutputString := ""
	// ----------
	for i := 0; i < len(outputString); i++ {
		charByte := outputString[i]
		if charByte >= 0x20 && charByte < 0x7F {
			escapedOutputString += string(charByte)
		} else {
			escapedOutputString += fmt.Sprintf("\\x%02X", charByte)
		}
	}
	//----------------------------------------
	return escapedOutputString
	//----------------------------------------
}

//--------------------------------------------------------------------------------
