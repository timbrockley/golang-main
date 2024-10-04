/*

Copyright 2024, Tim Brockley. All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

*/

package tui

import (
	"fmt"
	"strings"

	"github.com/mattn/go-runewidth"
)

//--------------------------------------------------------------------------------

func RenderWindow(rows []string, OptionsMap ...map[string]any) string {
	//----------------------------------------
	var builder strings.Builder
	//----------------------------------------
	options := ParseOptions(OptionsMap...)
	//----------------------------------------
	maxWidth := 0
	//----------------------------------------
	for rowIndex := range rows {
		//--------------------
		if !options.RawMode {
			rows[rowIndex] = EscapeString(rows[rowIndex])
		}
		//--------------------
		if runewidth.StringWidth(rows[rowIndex]) > maxWidth {
			maxWidth = runewidth.StringWidth(rows[rowIndex])
		}
		//--------------------
	}
	//----------------------------------------
	if options.MaxWidth == 0 {
		maxWidth += options.Padding * 2
		if options.Border {
			maxWidth += 2
		}
	} else if options.MaxWidth != maxWidth {
		maxWidth = options.MaxWidth
	}
	//----------------------------------------
	rowWidth := maxWidth
	if options.Border {
		rowWidth -= 2
	}
	//----------------------------------------
	if options.Padding > 0 {
		rowWidth -= options.Padding * 2
	}
	if rowWidth < 0 {
		rowWidth = 0
	}
	//----------------------------------------
	paddingHorizontal := ""
	paddingSpace := ""
	if options.Padding > 0 && maxWidth >= (2+options.Padding*2) {
		paddingHorizontal = strings.Repeat(options.BorderStyle.Horizontal, options.Padding)
		paddingSpace = strings.Repeat(" ", options.Padding)
	}
	//----------------------------------------
	maxHeight := 0
	rowHeight := 0
	//----------------------------------------
	if options.MaxHeight > 0 {
		maxHeight = options.MaxHeight
		if options.Border {
			rowHeight = maxHeight - 2
		}
		if rowHeight < 0 {
			rowHeight = 0
		}
	} else {
		maxHeight = len(rows)
		rowHeight = maxHeight
	}
	//----------------------------------------
	row := options.Row
	column := options.Column
	//--------------------
	if row == 0 {
		row = 1
	}
	if column == 0 {
		column = 1
	}
	//----------------------------------------
	rowNumber := row
	columnNumber := column
	//----------------------------------------
	top := options.BorderStyle.TopLeft + paddingHorizontal + strings.Repeat(options.BorderStyle.Horizontal, rowWidth) + paddingHorizontal + options.BorderStyle.TopRight
	bottom := options.BorderStyle.BottomLeft + paddingHorizontal + strings.Repeat(options.BorderStyle.Horizontal, rowWidth) + paddingHorizontal + options.BorderStyle.BottomRight
	left := options.BorderStyle.Vertical
	right := options.BorderStyle.Vertical
	//----------------------------------------
	builder.WriteString(CursorMove(rowNumber, column) + Reset())
	//----------------------------------------
	if options.Border {
		builder.WriteString(top)
		rowNumber += 1
	}
	//----------------------------------------
	for index := 0; index < rowHeight; index++ {
		//----------------------------------------
		rowString := ""
		//----------------------------------------
		if len(rows) > index {
			//--------------------
			rowString = rows[index]
			//--------------------
			if runewidth.StringWidth(rowString) > rowWidth {
				rowString = TruncateString(rows[index], rowWidth)
			}
			//--------------------
		}
		//----------------------------------------
		if runewidth.StringWidth(rowString) < rowWidth {
			rowString = fmt.Sprintf("%s%s", rowString, strings.Repeat(" ", rowWidth-runewidth.StringWidth(rowString)))
		}
		//----------------------------------------
		builder.WriteString(CursorMove(rowNumber, columnNumber) + Reset())
		//----------------------------------------
		if options.Border {
			builder.WriteString(left)
		}
		//--------------------
		if options.Padding > 0 {
			builder.WriteString(paddingSpace)
		}
		//--------------------
		builder.WriteString(rowString + Reset())
		//--------------------
		if options.Padding > 0 || options.Border {
			//--------------------
			rightColumnNumber := columnNumber + maxWidth
			if options.Padding > 0 {
				rightColumnNumber -= options.Padding
			}
			if options.Border {
				rightColumnNumber -= 1
			}
			builder.WriteString(CursorMove(rowNumber, rightColumnNumber))
			//--------------------
			if options.Padding > 0 {
				builder.WriteString(paddingSpace)
			}
			//--------------------
			if options.Border {
				builder.WriteString(right)
			}
			//--------------------
		}
		//----------------------------------------
		rowNumber += 1
		//----------------------------------------
	}
	//----------------------------------------
	if options.Border {
		builder.WriteString(CursorMove(rowNumber, columnNumber) + Reset() + bottom)
		// rowNumber += 1
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
