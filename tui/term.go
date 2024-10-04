/*

Copyright 2024, Tim Brockley. All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

*/

package tui

import (
	"os"

	"golang.org/x/term"
)

//--------------------------------------------------------------------------------

/*

	returns terminal size

	2024-07-01: worked ok when run in non-testing environment but failed in testing

*/

func TermSize(FD ...int) (width, height int, err error) {
	//----------------------------------------
	if len(FD) > 0 {
		return term.GetSize(FD[0])
	} else {
		return term.GetSize(int(os.Stdin.Fd()))
	}
	//----------------------------------------
}

//--------------------------------------------------------------------------------
