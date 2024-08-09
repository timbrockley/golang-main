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

	"github.com/mattn/go-runewidth"
)

//--------------------------------------------------------------------------------

func RenderTable(rows [][]string, OptionsMap ...map[string]any) string {
	//----------------------------------------
	var builder strings.Builder
	//----------------------------------------
	options := ParseOptions(OptionsMap...)
	//----------------------------------------
	if len(rows) == 0 {
		return options.BorderStyle.TopLeft + options.BorderStyle.TopRight + "\n" + options.BorderStyle.BottomLeft + options.BorderStyle.BottomRight + "\n"
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
			if maxColumnWidth > 0 && runewidth.StringWidth(rows[rowIndex][columnIndex]) > maxColumnWidth {
				rows[rowIndex][columnIndex] = TruncateString(rows[rowIndex][columnIndex], maxColumnWidth)
			}
			//----------------------------------------
			if runewidth.StringWidth(rows[rowIndex][columnIndex]) > columnWidths[columnIndex] {
				columnWidths[columnIndex] = runewidth.StringWidth(rows[rowIndex][columnIndex])
			}
			//----------------------------------------
		}
		//----------------------------------------
	}
	//----------------------------------------
	if maxColumns == 0 {
		return options.BorderStyle.TopLeft + options.BorderStyle.TopRight + "\n" + options.BorderStyle.BottomLeft + options.BorderStyle.BottomRight + "\n"
	}
	//----------------------------------------
	// table top border
	builder.WriteString(options.BorderStyle.TopLeft)
	//----------------------------------------
	for columnIndex := 0; columnIndex < maxColumns; columnIndex++ {
		//----------------------------------------
		builder.WriteString(strings.Repeat(options.BorderStyle.Horizontal, columnWidths[columnIndex]+options.Padding*2))
		//----------------------------------------
		if columnIndex < maxColumns-1 {
			builder.WriteString(options.BorderStyle.TopMiddle)
		}
		//----------------------------------------
	}
	//----------------------------------------
	builder.WriteString(options.BorderStyle.TopRight + "\n")
	//----------------------------------------
	for rowIndex, row := range rows {
		//----------------------------------------
		// inner table border (if header = true)
		if rowIndex == 1 && options.Header {
			//----------------------------------------
			builder.WriteString(options.BorderStyle.InnerLeft)
			//----------------------------------------
			for columnIndex := 0; columnIndex < maxColumns; columnIndex++ {
				//----------------------------------------
				builder.WriteString(strings.Repeat(options.BorderStyle.Horizontal, columnWidths[columnIndex]+options.Padding*2))
				//----------------------------------------
				if columnIndex < maxColumns-1 {
					builder.WriteString(options.BorderStyle.InnerMiddle)
				}
				//----------------------------------------
			}
			//----------------------------------------
			builder.WriteString(options.BorderStyle.InnerRight + "\n")
			//----------------------------------------
		}
		//----------------------------------------
		// column values
		builder.WriteString(options.BorderStyle.Vertical)
		//----------------------------------------
		for columnIndex := 0; columnIndex < maxColumns; columnIndex++ {
			//----------------------------------------
			columnString := ""
			//----------------------------------------
			if columnIndex < len(row) {
				columnString = row[columnIndex]
			}
			//----------------------------------------
			if runewidth.StringWidth(columnString) < columnWidths[columnIndex] {
				columnString = fmt.Sprintf("%s%s", columnString, strings.Repeat(" ", columnWidths[columnIndex]-runewidth.StringWidth(columnString)))
			}
			//----------------------------------------
			builder.WriteString(strings.Repeat(" ", options.Padding))
			builder.WriteString(columnString)
			builder.WriteString(strings.Repeat(" ", options.Padding))
			builder.WriteString(options.BorderStyle.Vertical)
			//----------------------------------------
		}
		//----------------------------------------
		builder.WriteString("\n")
		//----------------------------------------
	}
	//----------------------------------------
	// table bottom border
	builder.WriteString(options.BorderStyle.BottomLeft)
	//----------------------------------------
	for columnIndex := 0; columnIndex < maxColumns; columnIndex++ {
		//----------------------------------------
		builder.WriteString(strings.Repeat(options.BorderStyle.Horizontal, columnWidths[columnIndex]+options.Padding*2))
		//----------------------------------------
		if columnIndex < maxColumns-1 {
			builder.WriteString(options.BorderStyle.BottomMiddle)
		}
		//----------------------------------------
	}
	//----------------------------------------
	builder.WriteString(options.BorderStyle.BottomRight + "\n")
	//----------------------------------------
	outputString := builder.String()
	//----------------------------------------
	if options.MaxWidth > 0 {
		//----------------------------------------
		lines := strings.Split(outputString, "\n")
		//----------------------------------------
		for lineIndex := range lines {
			//--------------------
			stringWidth := runewidth.StringWidth(lines[lineIndex])
			//--------------------
			if stringWidth > options.MaxWidth {
				//--------------------
				runes := []rune(lines[lineIndex])
				//--------------------
				eol := ""
				eolLength := 1
				if options.MaxWidth >= 2 && len(runes) >= 2 && options.Padding > 0 {
					eol += string([]rune(runes)[len(runes)-2])
					eolLength += 1
				}
				eol += string([]rune(runes)[len(runes)-1])
				//--------------------
				lines[lineIndex] = TruncateString(lines[lineIndex], options.MaxWidth-eolLength) + eol
				//--------------------
			}
			//--------------------
		}
		//----------------------------------------
		outputString = strings.Join(lines, "\n")
		//----------------------------------------
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

func TabwriterTable(rows [][]string, OptionsMap ...map[string]any) string {
	//----------------------------------------
	var buffer bytes.Buffer
	//----------------------------------------
	options := ParseOptions(OptionsMap...)
	//----------------------------------------
	writer := tabwriter.NewWriter(&buffer, 0, 0, options.TabWidth, ' ', 0)
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
			if maxColumnWidth > 0 && runewidth.StringWidth(row[columnIndex]) > maxColumnWidth {
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
				headerLines[columnIndex] = strings.Repeat("-", runewidth.StringWidth(column))
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
	outputString := buffer.String()
	//----------------------------------------
	if options.MaxWidth > 0 {
		//----------------------------------------
		lines := strings.Split(outputString, "\n")
		//----------------------------------------
		for lineIndex := range lines {
			lines[lineIndex] = TruncateString(lines[lineIndex], options.MaxWidth)
		}
		//----------------------------------------
		outputString = strings.Join(lines, "\n")
		//----------------------------------------
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
