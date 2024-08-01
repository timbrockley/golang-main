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

func RenderWindow(rows []string, optionFuncs ...OptionFunc) string {
	//----------------------------------------
	var builder strings.Builder
	//----------------------------------------
	options := ParseOptions(optionFuncs...)
	//----------------------------------------
	maxWidth := 0
	//----------------------------------------
	for rowIndex := range rows {
		//--------------------
		rows[rowIndex] = EscapeString(rows[rowIndex])
		//--------------------
		if runewidth.StringWidth(rows[rowIndex]) > maxWidth {
			maxWidth = runewidth.StringWidth(rows[rowIndex])
		}
		//--------------------
	}
	//----------------------------------------
	if options.MaxWidth > 0 && options.MaxWidth != maxWidth {
		maxWidth = options.MaxWidth
	}
	//----------------------------------------
	rowWidth := maxWidth - 2
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
		rowHeight = maxHeight - 2
		if rowHeight < 0 {
			rowHeight = 0
		}
	} else {
		maxHeight = len(rows)
		rowHeight = maxHeight
	}
	//----------------------------------------
	top := options.BorderStyle.TopLeft + paddingHorizontal + strings.Repeat(options.BorderStyle.Horizontal, rowWidth) + paddingHorizontal + options.BorderStyle.TopRight
	bottom := options.BorderStyle.BottomLeft + paddingHorizontal + strings.Repeat(options.BorderStyle.Horizontal, rowWidth) + paddingHorizontal + options.BorderStyle.BottomRight
	//----------------------------------------
	builder.WriteString(CursorMove(1, 1) + top)
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
		rowString = options.BorderStyle.Vertical + paddingSpace + rowString + paddingSpace + options.BorderStyle.Vertical
		//----------------------------------------
		builder.WriteString(CursorMove(index+2, 1) + rowString)
		//----------------------------------------
	}
	//----------------------------------------
	builder.WriteString(CursorMove(rowHeight+2, 1) + bottom)
	//----------------------------------------
	if options.Writer != nil {
		fmt.Fprint(options.Writer, builder.String())
	}
	//----------------------------------------
	return builder.String()
	//----------------------------------------
}

//--------------------------------------------------------------------------------
