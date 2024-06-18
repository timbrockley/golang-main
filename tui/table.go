/*

Copyright 2024, Tim Brockley. All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

*/

package table

import (
	"fmt"
	"strings"
)

//--------------------------------------------------------------------------------

const (
	horizontal   = "\u2500"
	vertical     = "\u2502"
	topLeft      = "\u250C"
	topRight     = "\u2510"
	bottomLeft   = "\u2514"
	bottomRight  = "\u2518"
	innerLeft    = "\u251C"
	innerRight   = "\u2524"
	topMiddle    = "\u252C"
	innerMiddle  = "\u253C"
	bottomMiddle = "\u2534"
)

//--------------------------------------------------------------------------------

func RenderTable(rows [][]string, header bool) string {
	//----------------------------------------
	var builder strings.Builder
	//----------------------------------------
	if len(rows) == 0 {
		return topLeft + topRight + "\n" + bottomLeft + bottomRight + "\n"
	}
	//----------------------------------------
	columnWidths := make(map[int]int) // map used because max row length not know yet
	maxColumns := 0
	//----------------------------------------
	// get max columns and escape columns strings
	for rowIndex := range rows {
		for columnIndex := range rows[rowIndex] {
			if len(rows[rowIndex]) > maxColumns {
				maxColumns = len(rows[rowIndex])
			}
			rows[rowIndex][columnIndex] = EscapeString(rows[rowIndex][columnIndex])
			if len(rows[rowIndex][columnIndex]) > columnWidths[columnIndex] {
				columnWidths[columnIndex] = len(rows[rowIndex][columnIndex])
			}
		}
	}
	//----------------------------------------
	if maxColumns == 0 {
		return topLeft + topRight + "\n" + bottomLeft + bottomRight + "\n"
	}
	//----------------------------------------
	// table top border
	builder.WriteString(topLeft)
	for columnIndex := 0; columnIndex < maxColumns; columnIndex++ {
		builder.WriteString(strings.Repeat(horizontal, columnWidths[columnIndex]))
		if columnIndex < maxColumns-1 {
			builder.WriteString(topMiddle)
		}
	}
	builder.WriteString(topRight + "\n")
	//----------------------------------------
	for rowIndex, row := range rows {
		//----------------------------------------
		// inner table border (if header = true)
		if rowIndex == 1 && header {
			//----------------------------------------
			builder.WriteString(innerLeft)
			for columnIndex := 0; columnIndex < maxColumns; columnIndex++ {
				builder.WriteString(strings.Repeat(horizontal, columnWidths[columnIndex]))
				if columnIndex < maxColumns-1 {
					builder.WriteString(innerMiddle)
				}
			}
			builder.WriteString(innerRight + "\n")
			//----------------------------------------
		}
		//----------------------------------------
		// column values
		builder.WriteString(vertical)
		for columnIndex := 0; columnIndex < maxColumns; columnIndex++ {
			columnString := ""
			if columnIndex < len(row) {
				columnString = row[columnIndex]
			}
			if len(columnString) < columnWidths[columnIndex] {
				columnString = fmt.Sprintf("%s%s", columnString, strings.Repeat(" ", columnWidths[columnIndex]-len(columnString)))
			}
			builder.WriteString(columnString + vertical)
		}
		builder.WriteString("\n")
		//----------------------------------------
	}
	//----------------------------------------
	// table bottom border
	builder.WriteString(bottomLeft)
	for columnIndex := 0; columnIndex < maxColumns; columnIndex++ {
		builder.WriteString(strings.Repeat(horizontal, columnWidths[columnIndex]))
		if columnIndex < maxColumns-1 {
			builder.WriteString(bottomMiddle)
		}
	}
	builder.WriteString(bottomRight + "\n")
	//----------------------------------------
	return builder.String()
	//----------------------------------------
}

//--------------------------------------------------------------------------------

func EscapeString(stringValue string) string {
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
	stringValue = replacer.Replace(stringValue)
	// --------------------------------------------------------------------------------
	escapedStringValue := ""
	// ----------
	for i := 0; i < len(stringValue); i++ {
		charByte := stringValue[i]
		if charByte >= 0x20 && charByte < 0x7F {
			escapedStringValue += string(charByte)
		} else {
			escapedStringValue += fmt.Sprintf("\\x%02X", charByte)
		}
	}
	//----------------------------------------
	return escapedStringValue
	//----------------------------------------
}

//--------------------------------------------------------------------------------
