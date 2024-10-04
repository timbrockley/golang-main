/*

Copyright 2024, Tim Brockley. All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

*/

package tui

import (
	"fmt"
	"io"
	"reflect"

	"github.com/mattn/go-runewidth"
)

//--------------------------------------------------------------------------------

type Options struct {
	Header          bool
	Padding         int
	TabWidth        int
	MaxColumnWidth  int
	MaxColumnWidths []int
	MaxWidth        int
	MaxHeight       int
	Row             int
	Column          int
	Border          bool
	BorderStyle     BorderStyle
	RawMode         bool
	Writer          io.Writer
}

var DefaultOptions = Options{BorderStyle: UnicodeBorderStyle, Border: true, TabWidth: 2}

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

func Render(outputString string, OptionsMap ...map[string]any) string {
	//----------------------------------------
	options := ParseOptions(OptionsMap...)
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

func ParseOptions(OptionsMap ...map[string]any) Options {
	//----------------------------------------
	if len(OptionsMap) == 0 {
		return DefaultOptions
	}
	optionsMap := OptionsMap[0]
	//----------------------------------------
	defaultOptions := DefaultOptions
	reflectStruct := reflect.ValueOf(&defaultOptions).Elem()
	//----------------------------------------
	fields := reflect.VisibleFields(reflect.TypeOf(Options{}))
	for _, reflectField := range fields {
		if value, ok := optionsMap[reflectField.Name]; ok {
			fieldValue := reflectStruct.FieldByName(reflectField.Name)
			if fieldValue.IsValid() && fieldValue.CanSet() {
				reflectValue := reflect.ValueOf(value)
				if reflectValue.Type().AssignableTo(fieldValue.Type()) {
					fieldValue.Set(reflectValue)
				}
			}
		}
	}
	//----------------------------------------
	return defaultOptions
	//----------------------------------------
}

//--------------------------------------------------------------------------------
