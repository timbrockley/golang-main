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
	if options.MaxWidth > 0 && options.MaxWidth < maxWidth {
		maxWidth = options.MaxWidth
	}
	//----------------------------------------
	rowWidth := maxWidth
	if maxWidth >= 2 {
		rowWidth = maxWidth - 2
	}
	//----------------------------------------
	maxHeight := 0
	//----------------------------------------
	if options.MaxHeight == 0 {
		maxHeight = len(rows)
	} else {
		maxHeight = options.MaxHeight
	}
	//----------------------------------------
	rowHeight := maxHeight
	if maxHeight >= 2 {
		rowHeight = maxHeight - 2
	}
	//----------------------------------------
	top := options.BorderStyle.TopLeft + strings.Repeat(options.BorderStyle.Horizontal, rowWidth) + options.BorderStyle.TopRight
	bottom := options.BorderStyle.BottomLeft + strings.Repeat(options.BorderStyle.Horizontal, rowWidth) + options.BorderStyle.BottomRight
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
		rowString = options.BorderStyle.Vertical + rowString + options.BorderStyle.Vertical
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
