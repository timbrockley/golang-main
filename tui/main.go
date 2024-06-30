/*

Copyright 2024, Tim Brockley. All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

*/

package tui

import "fmt"

//--------------------------------------------------------------------------------

type OptionFunc func(*Options)

type Options struct {
	UseStdOut bool
	Header    bool
	MaxWidth  int
}

//--------------------------------------------------------------------------------

func Header(options *Options) { options.Header = true }

func MaxWidth(maxWidth int) OptionFunc {
	return func(Options *Options) {
		Options.MaxWidth = maxWidth
	}
}

func Stdout(options *Options) { options.UseStdOut = true }

//--------------------------------------------------------------------------------

func ReturnOutput(outputString string, optionFuncs ...OptionFunc) string {
	//----------------------------------------
	options := ParseOptions(optionFuncs...)
	//----------------------------------------
	if options.MaxWidth > 0 {
		outputString = TruncateString(outputString, options.MaxWidth)
	}
	//----------------------------------------
	if options.UseStdOut {
		fmt.Print(outputString)
	}
	//----------------------------------------
	return outputString
	//----------------------------------------
}

//--------------------------------------------------------------------------------

func ParseOptions(optionFuncs ...OptionFunc) Options {
	//----------------------------------------
	options := Options{}
	for _, fn := range optionFuncs {
		fn(&options)
	}
	//----------------------------------------
	return options
	//----------------------------------------
}

//--------------------------------------------------------------------------------
