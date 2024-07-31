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

	"github.com/mattn/go-runewidth"
)

//--------------------------------------------------------------------------------

type BorderStyle struct {
	Horizontal   string
	Vertical     string
	TopLeft      string
	TopRight     string
	BottomLeft   string
	BottomRight  string
	InnerLeft    string
	InnerRight   string
	TopMiddle    string
	InnerMiddle  string
	BottomMiddle string
}

var AsciiBorderStyle = BorderStyle{
	Horizontal:   "-",
	Vertical:     "|",
	TopLeft:      "+",
	TopRight:     "+",
	BottomLeft:   "+",
	BottomRight:  "+",
	InnerLeft:    "+",
	InnerRight:   "+",
	TopMiddle:    "+",
	InnerMiddle:  "+",
	BottomMiddle: "+",
}

var UnicodeBorderStyle = BorderStyle{
	Horizontal:   "\u2500",
	Vertical:     "\u2502",
	TopLeft:      "\u250C",
	TopRight:     "\u2510",
	BottomLeft:   "\u2514",
	BottomRight:  "\u2518",
	InnerLeft:    "\u251C",
	InnerRight:   "\u2524",
	TopMiddle:    "\u252C",
	InnerMiddle:  "\u253C",
	BottomMiddle: "\u2534",
}

//--------------------------------------------------------------------------------

type OptionFunc func(*Options)

type Options struct {
	Writer          io.Writer
	BorderStyle     BorderStyle
	Header          bool
	Padding         int
	TabWidth        int
	MaxColumnWidth  int
	MaxColumnWidths []int
	MaxWidth        int
	MaxHeight       int
	Row             int
	Column          int
}

var DefaultOptions = Options{BorderStyle: UnicodeBorderStyle, Padding: 1, TabWidth: 2}

//--------------------------------------------------------------------------------

func WithHeader(options *Options) { options.Header = true }

func WithBorderStyle(BorderStyle BorderStyle) OptionFunc {
	return func(options *Options) {
		options.BorderStyle = BorderStyle
	}
}

func WithPadding(padding int) OptionFunc {
	return func(options *Options) {
		options.Padding = padding
	}
}

func WithTabWidth(tabWidth int) OptionFunc {
	return func(options *Options) {
		options.TabWidth = tabWidth
	}
}

// max global column width
func WithMaxColumnWidth(maxColumnWidth int) OptionFunc {
	return func(options *Options) {
		options.MaxColumnWidth = maxColumnWidth
	}
}

// max individual column widths
func WithMaxColumnWidths(maxColumnWidths []int) OptionFunc {
	return func(options *Options) {
		options.MaxColumnWidths = maxColumnWidths
	}
}

func WithMaxWidth(maxWidth int) OptionFunc {
	return func(options *Options) {
		options.MaxWidth = maxWidth
	}
}

func WithMaxHeight(maxHeight int) OptionFunc {
	return func(options *Options) {
		options.MaxHeight = maxHeight
	}
}

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

//--------------------------------------------------------------------------------

func ParseOptions(optionFuncs ...OptionFunc) Options {
	//----------------------------------------
	options := DefaultOptions
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

func RuneCount(outputString string) int { return len([]rune(outputString)) }

//--------------------------------------------------------------------------------

func TruncateString(outputString string, maxWidth int) string {
	//----------------------------------------
	if maxWidth == 0 {
		//----------------------------------------
		return ""
		//----------------------------------------
	} else if runewidth.StringWidth(outputString) <= maxWidth {
		//----------------------------------------
		return outputString
		//----------------------------------------
	} else {
		//----------------------------------------
		outputRunes := []rune(outputString)
		outputRunesCount := 0
		outputWidth := 0
		//----------------------------------------
		for _, rune := range outputRunes {

			// if  runewidth.RuneWidth(rune)>1{outputRunes[outputRunesCount] = 0x20}

			if outputWidth+runewidth.RuneWidth(rune) > maxWidth {
				break
			}
			outputWidth += runewidth.RuneWidth(rune)
			outputRunesCount += 1
		}
		//----------------------------------------
		if maxWidth > runewidth.StringWidth(string(outputRunes[:outputRunesCount])) {
			outputRunes[outputRunesCount] = 0x20
			outputRunesCount += 1
		}
		//----------------------------------------
		return string(outputRunes[:outputRunesCount])
		//----------------------------------------
	}
	//----------------------------------------
}

//--------------------------------------------------------------------------------

func EscapeString(outputString string) string {
	//----------------------------------------
	outputRunes := []rune(outputString)
	//----------------------------------------
	for i := 0; i < len(outputRunes); i++ {
		rune := outputRunes[i]
		if rune < 0x20 || rune == 0x7F {
			outputRunes[i] = 0x20
		}
	}
	//----------------------------------------
	return string(outputRunes)
	//----------------------------------------
}

//--------------------------------------------------------------------------------
