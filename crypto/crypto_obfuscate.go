/*

Copyright 2023-2025, Tim Brockley. All rights reserved.

This software is licensed under the MIT License.

*/

package crypto

import (
	"errors"
	"strings"

	"github.com/timbrockley/golang-main/conv"
)

//----------------------------------------------------------------------
//######################################################################
//----------------------------------------------------------------------

type Option func(*Options)

type Options struct {
	Encoding string
	MixChars bool
}

var DefaultOptions = Options{
	Encoding: "",
	MixChars: true,
}

//----------------------------------------------------------------------

func WithEncoding(encoding string) Option {
	return func(options *Options) { options.Encoding = encoding }
}

func WithMixChars(mixChars bool) Option {
	return func(options *Options) { options.MixChars = mixChars }
}

//----------------------------------------------------------------------

func NewOptions(options ...Option) Options {
	obfuscationOptions := DefaultOptions
	for _, optionFunc := range options {
		optionFunc(&obfuscationOptions)
	}
	return obfuscationOptions
}

//----------------------------------------------------------------------
//######################################################################
//----------------------------------------------------------------------

//------------------------------------------------------------
// ObfuscateXOR
//------------------------------------------------------------

func ObfuscateXOR(dataString string, value byte) (string, error) {
	//------------------------------------------------------------
	if value == 0 {
		return "", errors.New("value should be an integer between 1 and 255")
	}
	//------------------------------------------------------------
	if dataString == "" {
		return dataString, nil
	}
	//------------------------------------------------------------
	dataLength := len(dataString)
	//------------------------------------------------------------
	outputBytes := make([]byte, dataLength)
	//------------------------------------------------------------
	// for loop uses bytes instead of range which would use UTF-8 runes
	for index := 0; index < len(dataString); index += 1 {
		outputBytes[index] = dataString[index] ^ value
	}
	//------------------------------------------------------------
	return string(outputBytes), nil
	//------------------------------------------------------------
}

//------------------------------------------------------------

//------------------------------------------------------------
// ObfuscateXOREncode
//------------------------------------------------------------

func ObfuscateXOREncode(dataString string, value byte, Encoding ...string) (string, error) {
	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	if value == 0 {
		return "", errors.New("value should be an integer between 1 and 255")
	}
	//------------------------------------------------------------
	if dataString == "" {
		return dataString, nil
	}
	//------------------------------------------------------------
	encoding := ""
	if len(Encoding) > 0 {
		encoding = Encoding[0]
	}
	//------------------------------------------------------------
	dataString, err = ObfuscateXOR(dataString, value)
	//------------------------------------------------------------
	if err != nil {
		return "", err
	}
	//------------------------------------------------------------
	switch encoding {
	case "base":
		return conv.Base_encode(dataString), nil
	case "base64":
		return conv.Base64_encode(dataString), nil
	case "base64url":
		return conv.Base64url_encode(dataString), nil
	case "base91":
		return conv.Base91_encode(dataString, true), nil
	case "hex":
		return conv.Hex_encode(dataString), nil
	default:
		replacer := strings.NewReplacer(
			"-", "--",
			"\x09", "-t", // tab
			"\x0A", "-n", // new line
			"\x0D", "-r", // carriage return
			"\x20", "-s", // space
			"\x22", "-q", // double quote
			"\x24", "-d", // dollar sign
			"\x27", "-a", // apostrophy
			"\x5C", "-b", // backslash
			"\x60", "-g", // grave accent
		)
		return replacer.Replace(dataString), nil
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateXORDecode
//------------------------------------------------------------

func ObfuscateXORDecode(dataString string, value byte, Encoding ...string) (string, error) {
	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	if dataString == "" {
		return dataString, nil
	}
	//------------------------------------------------------------
	encoding := ""
	if len(Encoding) > 0 {
		encoding = Encoding[0]
	}
	//------------------------------------------------------------
	switch encoding {
	case "base":
		dataString, err = conv.Base_decode(dataString)
	case "base64":
		dataString, err = conv.Base64_decode(dataString)
	case "base64url":
		dataString, err = conv.Base64url_decode(dataString)
	case "base91":
		dataString, err = conv.Base91_decode(dataString, true)
	case "hex":
		dataString, err = conv.Hex_decode(dataString)
	default:
		replacer := strings.NewReplacer(
			"--", "-SUB",
			"-g", "\x60", // grave accent
			"-b", "\x5C", // backslash
			"-a", "\x27", // apostrophy
			"-d", "\x24", // dollar sign
			"-q", "\x22", // double quote
			"-s", "\x20", // space
			"-r", "\x0D", // carriage return
			"-n", "\x0A", // new line
			"-t", "\x09", // tab
		)
		dataString = replacer.Replace(dataString)
		dataString = strings.ReplaceAll(dataString, "-SUB", "-")
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
	if err != nil {
		return "", err
	}
	//------------------------------------------------------------
	return ObfuscateXOR(dataString, value)
	//------------------------------------------------------------
}

//----------------------------------------------------------------------
//######################################################################
//----------------------------------------------------------------------

//------------------------------------------------------------
// ObfuscateV0
//------------------------------------------------------------

func ObfuscateV0(dataString string) string {
	//------------------------------------------------------------
	if dataString == "" {
		return dataString
	}
	//------------------------------------------------------------
	outputBytes := make([]byte, len(dataString))
	//------------------------------------------------------------
	// for loop uses bytes instead of range which would use UTF-8 runes
	for index := 0; index < len(dataString); index += 1 {
		outputBytes[index] = SlideByteV0(dataString[index])
	}
	//------------------------------------------------------------
	return string(outputBytes)
	//------------------------------------------------------------
}

//------------------------------------------------------------
// SlideByteV0
//------------------------------------------------------------

func SlideByteV0(charByte byte) byte {
	//------------------------------------------------------------
	switch {
	case charByte <= 0x1F:
		return 0x1F - charByte
	case charByte <= 0x7E:
		return 0x7E - (charByte - 0x20)
	case charByte == 0x7F:
		return charByte
	default:
		return 0xFF - (charByte - 0x80)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateV0Encode
//------------------------------------------------------------

func ObfuscateV0Encode(dataString string, Encoding ...string) string {
	//------------------------------------------------------------
	if dataString == "" {
		return dataString
	}
	//------------------------------------------------------------
	encoding := ""
	if len(Encoding) > 0 {
		encoding = Encoding[0]
	}
	//------------------------------------------------------------
	dataString = ObfuscateV0(dataString)
	//------------------------------------------------------------
	switch encoding {
	case "base":
		return conv.Base_encode(dataString)
	case "base64":
		return conv.Base64_encode(dataString)
	case "base64url":
		return conv.Base64url_encode(dataString)
	case "base91":
		return conv.Base91_encode(dataString, true)
	case "hex":
		return conv.Hex_encode(dataString)
	default:
		replacer := strings.NewReplacer(
			"-", "--",
			"\x09", "-t", // tab
			"\x0A", "-n", // new line
			"\x0D", "-r", // carriage return
			"\x20", "-s", // space
			"\x22", "-q", // double quote
			"\x24", "-d", // dollar sign
			"\x27", "-a", // apostrophy
			"\x5C", "-b", // backslash
			"\x60", "-g", // grave accent
		)
		return replacer.Replace(dataString)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateV0Decode
//------------------------------------------------------------

func ObfuscateV0Decode(dataString string, Encoding ...string) (string, error) {
	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	if dataString == "" {
		return dataString, nil
	}
	//------------------------------------------------------------
	encoding := ""
	if len(Encoding) > 0 {
		encoding = Encoding[0]
	}
	//------------------------------------------------------------
	switch encoding {
	case "base":
		dataString, err = conv.Base_decode(dataString)
	case "base64":
		dataString, err = conv.Base64_decode(dataString)
	case "base64url":
		dataString, err = conv.Base64url_decode(dataString)
	case "base91":
		dataString, err = conv.Base91_decode(dataString, true)
	case "hex":
		dataString, err = conv.Hex_decode(dataString)
	default:
		replacer := strings.NewReplacer(
			"--", "-SUB",
			"-g", "\x60", // grave accent
			"-b", "\x5C", // backslash
			"-a", "\x27", // apostrophy
			"-d", "\x24", // dollar sign
			"-q", "\x22", // double quote
			"-s", "\x20", // space
			"-r", "\x0D", // carriage return
			"-n", "\x0A", // new line
			"-t", "\x09", // tab
		)
		dataString = replacer.Replace(dataString)
		dataString = strings.ReplaceAll(dataString, "-SUB", "-")
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
	if err != nil {
		return "", err
	}
	//------------------------------------------------------------
	return ObfuscateV0(dataString), nil
	//------------------------------------------------------------
}

//----------------------------------------------------------------------
//######################################################################
//----------------------------------------------------------------------

//------------------------------------------------------------
// ObfuscateV4
//------------------------------------------------------------

// dataString
//
//	[WithMixChars]
func ObfuscateV4(dataString string, Options ...Option) string {
	//------------------------------------------------------------
	if dataString == "" {
		return dataString
	}
	//------------------------------------------------------------
	options := NewOptions(Options...)
	//------------------------------------------------------------
	dataLength := len(dataString)
	//------------------------------------------------------------
	outputBytes := make([]byte, dataLength)
	//------------------------------------------------------------
	if !options.MixChars || dataLength < 4 {
		// for loop uses bytes instead of range which would use UTF-8 runes
		for index := 0; index < len(dataString); index += 1 {
			outputBytes[index] = SlideByteV4(dataString[index])
		}
		return string(outputBytes)
	}
	//------------------------------------------------------------
	mixedLength := dataLength
	mixedHalf := mixedLength / 2
	//------------------------------------------------------------
	if mixedHalf%2 != 0 {
		mixedHalf -= 1
		mixedLength = mixedHalf * 2
	}
	//------------------------------------------------------------
	// for loop uses bytes instead of range which would use UTF-8 runes
	for index := 0; index < dataLength; index++ {
		if index < mixedLength && index%2 != 0 {
			if index < mixedHalf {
				outputBytes[index+mixedHalf] = SlideByteV4(dataString[index])
			} else {
				outputBytes[index-mixedHalf] = SlideByteV4(dataString[index])
			}
		} else {
			outputBytes[index] = SlideByteV4(dataString[index])
		}
	}
	//------------------------------------------------------------
	return string(outputBytes)
	//------------------------------------------------------------
}

//------------------------------------------------------------
// SlideByteV4
//------------------------------------------------------------

func SlideByteV4(charByte byte) byte {
	//------------------------------------------------------------
	switch {
	case charByte >= 0x20 && charByte <= 0x7E:
		return 0x7E - (charByte - 0x20)
	default:
		return charByte
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateV4Encode
//------------------------------------------------------------

// dataString
//
//	[WithEncoding]
//	[WithMixChars]
func ObfuscateV4Encode(dataString string, Options ...Option) string {
	//------------------------------------------------------------
	if dataString == "" {
		return dataString
	}
	//------------------------------------------------------------
	options := NewOptions(Options...)
	//------------------------------------------------------------
	dataString = ObfuscateV4(dataString, Options...)
	//------------------------------------------------------------
	switch options.Encoding {
	case "base":
		return conv.Base_encode(dataString)
	case "base64":
		return conv.Base64_encode(dataString)
	case "base64url":
		return conv.Base64url_encode(dataString)
	case "base91":
		return conv.Base91_encode(dataString, true)
	case "hex":
		return conv.Hex_encode(dataString)
	default:
		replacer := strings.NewReplacer(
			"\\", "\\\\",
			"\x09", "\\t", // tab
			"\x0A", "\\n", // new line
			"\x0D", "\\r", // carriage return
			"\x22", "\\q", // double quote
			"\x27", "\\a", // apostrophe
			"\x60", "\\g", // grave accent
		)
		return replacer.Replace(dataString)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateV4Decode
//------------------------------------------------------------

// dataString
//
//	[WithEncoding]
//	[WithMixChars]
func ObfuscateV4Decode(dataString string, Options ...Option) (string, error) {
	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	if dataString == "" {
		return dataString, nil
	}
	//------------------------------------------------------------
	options := NewOptions(Options...)
	//------------------------------------------------------------
	switch options.Encoding {
	case "base":
		dataString, err = conv.Base_decode(dataString)
	case "base64":
		dataString, err = conv.Base64_decode(dataString)
	case "base64url":
		dataString, err = conv.Base64url_decode(dataString)
	case "base91":
		dataString, err = conv.Base91_decode(dataString, true)
	case "hex":
		dataString, err = conv.Hex_decode(dataString)
	default:
		replacer := strings.NewReplacer(
			"\\\\", "\\SUB",
			"\\g", "\x60", // grave accent
			"\\a", "\x27", // apostrophy
			"\\q", "\x22", // double quote
			"\\r", "\x0D", // carriage return
			"\\n", "\x0A", // new line
			"\\t", "\x09", // tab
		)
		dataString = replacer.Replace(dataString)
		dataString = strings.ReplaceAll(dataString, "\\SUB", "\\")
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
	if err != nil {
		return "", err
	}
	//------------------------------------------------------------
	return ObfuscateV4(dataString, Options...), nil
	//------------------------------------------------------------
}

//----------------------------------------------------------------------
//######################################################################
//----------------------------------------------------------------------

//------------------------------------------------------------
// ObfuscateV5
//------------------------------------------------------------

func ObfuscateV5(dataString string) string {
	//------------------------------------------------------------
	if dataString == "" {
		return dataString
	}
	//------------------------------------------------------------
	dataLength := len(dataString)
	//------------------------------------------------------------
	outputBytes := make([]byte, dataLength)
	//------------------------------------------------------------
	if dataLength < 4 {
		// for loop uses bytes instead of range which would use UTF-8 runes
		for index := 0; index < len(dataString); index += 1 {
			outputBytes[index] = SlideByteV5(dataString[index])
		}
		return string(outputBytes)
	}
	//------------------------------------------------------------
	mixedLength := dataLength
	mixedHalf := mixedLength / 2
	//------------------------------------------------------------
	if mixedHalf%2 != 0 {
		mixedHalf -= 1
		mixedLength = mixedHalf * 2
	}
	//------------------------------------------------------------
	// for loop uses bytes instead of range which would use UTF-8 runes
	for index := 0; index < dataLength; index++ {
		if index < mixedLength && index%2 != 0 {
			if index < mixedHalf {
				outputBytes[index+mixedHalf] = SlideByteV5(dataString[index])
			} else {
				outputBytes[index-mixedHalf] = SlideByteV5(dataString[index])
			}
		} else {
			outputBytes[index] = SlideByteV5(dataString[index])
		}
	}
	//------------------------------------------------------------
	return string(outputBytes)
	//------------------------------------------------------------
}

//------------------------------------------------------------
// SlideByteV5
//------------------------------------------------------------

func SlideByteV5(charByte byte) byte {
	//------------------------------------------------------------
	switch {
	case charByte <= 0x1F:
		return 0x1F - charByte
	case charByte <= 0x7E:
		return 0x7E - (charByte - 0x20)
	case charByte == 0x7F:
		return charByte
	default:
		return 0xFF - (charByte - 0x80)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateV5Encode
//------------------------------------------------------------

func ObfuscateV5Encode(dataString string, Encoding ...string) string {
	//------------------------------------------------------------
	if dataString == "" {
		return dataString
	}
	//------------------------------------------------------------
	encoding := ""
	if len(Encoding) > 0 {
		encoding = Encoding[0]
	}
	//------------------------------------------------------------
	dataString = ObfuscateV5(dataString)
	//------------------------------------------------------------
	switch encoding {
	case "base":
		return conv.Base_encode(dataString)
	case "base64":
		return conv.Base64_encode(dataString)
	case "base64url":
		return conv.Base64url_encode(dataString)
	case "base91":
		return conv.Base91_encode(dataString, true)
	case "hex":
		return conv.Hex_encode(dataString)
	default:
		replacer := strings.NewReplacer(
			"-", "--",
			"\x09", "-t", // tab
			"\x0A", "-n", // new line
			"\x0D", "-r", // carriage return
			"\x20", "-s", // space
			"\x22", "-q", // double quote
			"\x24", "-d", // dollar sign
			"\x27", "-a", // apostrophy
			"\x5C", "-b", // backslash
			"\x60", "-g", // grave accent
		)
		return replacer.Replace(dataString)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateV5Decode
//------------------------------------------------------------

func ObfuscateV5Decode(dataString string, Encoding ...string) (string, error) {
	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	if dataString == "" {
		return dataString, nil
	}
	//------------------------------------------------------------
	encoding := ""
	if len(Encoding) > 0 {
		encoding = Encoding[0]
	}
	//------------------------------------------------------------
	switch encoding {
	case "base":
		dataString, err = conv.Base_decode(dataString)
	case "base64":
		dataString, err = conv.Base64_decode(dataString)
	case "base64url":
		dataString, err = conv.Base64url_decode(dataString)
	case "base91":
		dataString, err = conv.Base91_decode(dataString, true)
	case "hex":
		dataString, err = conv.Hex_decode(dataString)
	default:
		replacer := strings.NewReplacer(
			"--", "-SUB",
			"-g", "\x60", // grave accent
			"-b", "\x5C", // backslash
			"-a", "\x27", // apostrophy
			"-d", "\x24", // dollar sign
			"-q", "\x22", // double quote
			"-s", "\x20", // space
			"-r", "\x0D", // carriage return
			"-n", "\x0A", // new line
			"-t", "\x09", // tab
		)
		dataString = replacer.Replace(dataString)
		dataString = strings.ReplaceAll(dataString, "-SUB", "-")
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
	if err != nil {
		return "", err
	}
	//------------------------------------------------------------
	return ObfuscateV5(dataString), nil
	//------------------------------------------------------------
}

//----------------------------------------------------------------------
//######################################################################
//----------------------------------------------------------------------
