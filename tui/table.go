/*

Copyright 2024, Tim Brockley. All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

*/

package tui

import (
	"bytes"
	"fmt"
	"strings"
	"text/tabwriter"
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

type TableOptions struct {
	Header   bool
	MaxWidth int
}

//--------------------------------------------------------------------------------

func RenderTable(rows [][]string, optionFuncs ...OptionFunc) string {
	//----------------------------------------
	var builder strings.Builder
	//----------------------------------------
	options := ParseOptions(optionFuncs...)
	//----------------------------------------
	if len(rows) == 0 {
		return topLeft + topRight + "\n" + bottomLeft + bottomRight + "\n"
	}
	//----------------------------------------
	columnWidths := make(map[int]int) // map used because max row length not know yet
	//----------------------------------------
	maxColumns := 0
	//----------------------------------------
	// get max columns and escape columns strings
	for rowIndex := range rows {
		//----------------------------------------
		for columnIndex := range rows[rowIndex] {
			//----------------------------------------
			if len(rows[rowIndex]) > maxColumns {
				maxColumns = len(rows[rowIndex])
			}
			//----------------------------------------
			rows[rowIndex][columnIndex] = EscapeString(rows[rowIndex][columnIndex])
			//----------------------------------------
			maxColumnWidth := 0
			//--------------------
			if options.MaxColumnWidth > 0 {
				maxColumnWidth = options.MaxColumnWidth
			}
			//--------------------
			if len(options.MaxColumnWidths) >= columnIndex+1 && options.MaxColumnWidths[columnIndex] > 0 {
				if maxColumnWidth == 0 || options.MaxColumnWidths[columnIndex] < maxColumnWidth {
					maxColumnWidth = options.MaxColumnWidths[columnIndex]
				}
			}
			//--------------------
			if maxColumnWidth > 0 && len(rows[rowIndex][columnIndex]) > maxColumnWidth {
				rows[rowIndex][columnIndex] = TruncateString(rows[rowIndex][columnIndex], maxColumnWidth)
			}
			//----------------------------------------
			if len(rows[rowIndex][columnIndex]) > columnWidths[columnIndex] {
				columnWidths[columnIndex] = len(rows[rowIndex][columnIndex])
			}
			//----------------------------------------
		}
		//----------------------------------------
	}
	//----------------------------------------
	if maxColumns == 0 {
		return topLeft + topRight + "\n" + bottomLeft + bottomRight + "\n"
	}
	//----------------------------------------
	// table top border
	builder.WriteString(topLeft)
	//----------------------------------------
	for columnIndex := 0; columnIndex < maxColumns; columnIndex++ {
		//----------------------------------------
		builder.WriteString(strings.Repeat(horizontal, columnWidths[columnIndex]))
		//----------------------------------------
		if columnIndex < maxColumns-1 {
			builder.WriteString(topMiddle)
		}
		//----------------------------------------
	}
	//----------------------------------------
	builder.WriteString(topRight + "\n")
	//----------------------------------------
	for rowIndex, row := range rows {
		//----------------------------------------
		// inner table border (if header = true)
		if rowIndex == 1 && options.Header {
			//----------------------------------------
			builder.WriteString(innerLeft)
			//----------------------------------------
			for columnIndex := 0; columnIndex < maxColumns; columnIndex++ {
				//----------------------------------------
				builder.WriteString(strings.Repeat(horizontal, columnWidths[columnIndex]))
				//----------------------------------------
				if columnIndex < maxColumns-1 {
					builder.WriteString(innerMiddle)
				}
				//----------------------------------------
			}
			//----------------------------------------
			builder.WriteString(innerRight + "\n")
			//----------------------------------------
		}
		//----------------------------------------
		// column values
		builder.WriteString(vertical)
		//----------------------------------------
		for columnIndex := 0; columnIndex < maxColumns; columnIndex++ {
			//----------------------------------------
			columnString := ""
			//----------------------------------------
			if columnIndex < len(row) {
				columnString = row[columnIndex]
			}
			//----------------------------------------
			if len(columnString) < columnWidths[columnIndex] {
				columnString = fmt.Sprintf("%s%s", columnString, strings.Repeat(" ", columnWidths[columnIndex]-len(columnString)))
			}
			//----------------------------------------
			builder.WriteString(columnString + vertical)
			//----------------------------------------
		}
		//----------------------------------------
		builder.WriteString("\n")
		//----------------------------------------
	}
	//----------------------------------------
	// table bottom border
	builder.WriteString(bottomLeft)
	//----------------------------------------
	for columnIndex := 0; columnIndex < maxColumns; columnIndex++ {
		//----------------------------------------
		builder.WriteString(strings.Repeat(horizontal, columnWidths[columnIndex]))
		//----------------------------------------
		if columnIndex < maxColumns-1 {
			builder.WriteString(bottomMiddle)
		}
		//----------------------------------------
	}
	//----------------------------------------
	builder.WriteString(bottomRight + "\n")
	//----------------------------------------
	if options.MaxWidth > 0 {
		//----------------------------------------
		lines := strings.Split(builder.String(), "\n")
		//----------------------------------------
		for lineIndex := range lines {
			lines[lineIndex] = TruncateString(lines[lineIndex], options.MaxWidth)
		}
		//----------------------------------------
		return strings.Join(lines, "\n")
		//----------------------------------------
	}
	//----------------------------------------
	if options.Writer != nil {
		fmt.Fprint(options.Writer, builder.String())
	}
	//----------------------------------------
	return builder.String()
	//----------------------------------------
}

//--------------------------------------------------------------------------------

func TabwriterTable(rows [][]string, optionFuncs ...OptionFunc) string {
	//----------------------------------------
	var buffer bytes.Buffer
	//----------------------------------------
	options := ParseOptions(optionFuncs...)
	//----------------------------------------
	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)
	//----------------------------------------
	for rowIndex, row := range rows {
		//----------------------------------------
		for columnIndex := range row {
			//----------------------------------------
			row[columnIndex] = EscapeString(row[columnIndex])
			//----------------------------------------
			maxColumnWidth := 0
			//--------------------
			if options.MaxColumnWidth > 0 {
				maxColumnWidth = options.MaxColumnWidth
			}
			//--------------------
			if len(options.MaxColumnWidths) >= columnIndex+1 && options.MaxColumnWidths[columnIndex] > 0 {
				if maxColumnWidth == 0 || options.MaxColumnWidths[columnIndex] < maxColumnWidth {
					maxColumnWidth = options.MaxColumnWidths[columnIndex]
				}
			}
			//--------------------
			if maxColumnWidth > 0 && len(row[columnIndex]) > maxColumnWidth {
				row[columnIndex] = TruncateString(row[columnIndex], maxColumnWidth)
			}
			//----------------------------------------
		}
		//----------------------------------------
		fmt.Fprintln(writer, strings.Join(row, "\t"))
		//----------------------------------------
		if rowIndex == 0 && options.Header {
			//----------------------------------------
			headerLines := make([]string, len(row))
			//----------------------------------------
			for columnIndex, column := range row {
				headerLines[columnIndex] = strings.Repeat("-", len(column))
			}
			//----------------------------------------
			fmt.Fprintln(writer, strings.Join(headerLines, "\t"))
			//----------------------------------------
		}
		//----------------------------------------
	}
	//----------------------------------------
	writer.Flush()
	//----------------------------------------
	if options.MaxWidth > 0 {
		//----------------------------------------
		lines := strings.Split(buffer.String(), "\n")
		//----------------------------------------
		for lineIndex := range lines {
			lines[lineIndex] = TruncateString(lines[lineIndex], options.MaxWidth)
		}
		//----------------------------------------
		return strings.Join(lines, "\n")
		//----------------------------------------
	}
	//----------------------------------------
	if options.Writer != nil {
		fmt.Fprint(options.Writer, buffer.String())
	}
	//----------------------------------------
	return buffer.String()
	//----------------------------------------
}

//--------------------------------------------------------------------------------
