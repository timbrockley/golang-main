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

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

//------------------------------------------------------------
// ObfuscateV5
//------------------------------------------------------------

func ObfuscateV5(dataString string) string {
	//------------------------------------------------------------
	if dataString == "" {
		return dataString
	}
	//------------------------------------------------------------
	var dataBytes, strBytes, strXBytes []byte
	var charByte byte
	var length, half int
	//------------------------------------------------------------
	dataBytes = []byte(dataString)
	//------------------------------------------------------------
	if len(dataBytes) >= 4 {
		//--------------------
		length = len(dataBytes)
		half = int(length / 2)
		//--------------------
		if half%2 != 0 {
			half -= 1
			length = half * 2
			strXBytes = dataBytes[length:]
			dataBytes = dataBytes[0:length]
		}
		//--------------------
		for i1 := 0; i1 < len(dataBytes); i1 += 1 {
			//--------------------
			if i1%2 != 0 {
				if i1 < half {
					charByte = dataBytes[i1+half]
				} else {
					charByte = dataBytes[i1-half]
				}
			} else {
				charByte = dataBytes[i1]
			}
			//--------------------
			strBytes = append(strBytes, SlideByteV5(charByte))
			//--------------------
		}

	} else {
		strXBytes = dataBytes
	}
	//------------------------------------------------------------
	if len(strXBytes) > 0 {
		for i2 := 0; i2 < len(strXBytes); i2 += 1 {
			strBytes = append(strBytes, SlideByteV5((strXBytes[i2])))
		}
	}
	//------------------------------------------------------------
	return string(strBytes)
	//------------------------------------------------------------
}

//------------------------------------------------------------
// SlideByteV5
//------------------------------------------------------------

func SlideByteV5(charByte byte) byte {
	//------------------------------------------------------------
	if charByte <= 0x1F {
		return 0x1F - charByte
	} else if charByte <= 0x7E {
		return 0x7E - (charByte - 0x20)
	} else if charByte >= 0x80 {
		return 0xFF - (charByte - 0x80)
	}
	//------------------------------------------------------------
	return charByte
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
	//------------------------------------------------------------
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
	//------------------------------------------------------------
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

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

//------------------------------------------------------------
// ObfuscateV4
//------------------------------------------------------------

func ObfuscateV4(dataString string, MixChars ...bool) string {
	//------------------------------------------------------------
	if dataString == "" {
		return dataString
	}
	//------------------------------------------------------------
	var dataBytes, strBytes, strXBytes []byte
	var charByte byte
	var length, half int
	//------------------------------------------------------------
	dataBytes = []byte(dataString)
	//------------------------------------------------------------
	mixChars := true
	//------------------------------------------------------------
	if len(MixChars) > 0 {
		mixChars = MixChars[0]
	}
	//------------------------------------------------------------
	if mixChars && len(dataBytes) >= 4 {
		//--------------------
		length = len(dataBytes)
		half = int(length / 2)
		//--------------------
		if half%2 != 0 {
			half -= 1
			length = half * 2
			strXBytes = dataBytes[length:]
			dataBytes = dataBytes[0:length]
		}
		//--------------------
		for i1 := 0; i1 < len(dataBytes); i1 += 1 {
			//--------------------
			if i1%2 != 0 {
				if i1 < half {
					charByte = dataBytes[i1+half]
				} else {
					charByte = dataBytes[i1-half]
				}
			} else {
				charByte = dataBytes[i1]
			}
			//--------------------
			strBytes = append(strBytes, SlideByteV4(charByte))
			//--------------------
		}

	} else {
		strXBytes = dataBytes
	}
	//------------------------------------------------------------
	if len(strXBytes) > 0 {
		for i2 := 0; i2 < len(strXBytes); i2 += 1 {
			strBytes = append(strBytes, SlideByteV4((strXBytes[i2])))
		}
	}
	//------------------------------------------------------------
	return string(strBytes)
	//------------------------------------------------------------
}

//------------------------------------------------------------
// SlideByteV4
//------------------------------------------------------------

func SlideByteV4(charByte byte) byte {
	//------------------------------------------------------------
	if charByte >= 0x20 && charByte <= 0x7E {
		charByte = 0x9E - charByte
	}
	//------------------------------------------------------------
	return charByte
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateV4Encode
//------------------------------------------------------------

func ObfuscateV4Encode(dataString string, Encoding ...string) string {
	//------------------------------------------------------------
	if dataString == "" {
		return dataString
	}
	//------------------------------------------------------------
	encoding := ""
	//------------------------------------------------------------
	if len(Encoding) > 0 {
		encoding = Encoding[0]
	}
	//------------------------------------------------------------
	dataString = ObfuscateV4(dataString)
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

func ObfuscateV4Decode(dataString string, Encoding ...string) (string, error) {
	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	if dataString == "" {
		return dataString, nil
	}
	//------------------------------------------------------------
	encoding := ""
	//------------------------------------------------------------
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
	return ObfuscateV4(dataString), nil
	//------------------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

//------------------------------------------------------------
// ObfuscateV0
//------------------------------------------------------------

func ObfuscateV0(dataString string) string {
	//------------------------------------------------------------
	if dataString == "" {
		return dataString
	}
	//------------------------------------------------------------
	var dataBytes, strBytes []byte
	//------------------------------------------------------------
	dataBytes = []byte(dataString)
	//------------------------------------------------------------
	if len(dataBytes) > 0 {
		for i1 := 0; i1 < len(dataBytes); i1 += 1 {
			strBytes = append(strBytes, SlideByteV5((dataBytes[i1])))
		}
	}
	//------------------------------------------------------------
	return string(strBytes)
	//------------------------------------------------------------
}

//------------------------------------------------------------
// SlideByteV0
//------------------------------------------------------------

func SlideByteV0(charByte byte) byte {
	//------------------------------------------------------------
	if charByte <= 0x1F {
		return 0x1F - charByte
	} else if charByte <= 0x7E {
		return 0x7E - (charByte - 0x20)
	} else if charByte >= 0x80 {
		return 0xFF - (charByte - 0x80)
	}
	//------------------------------------------------------------
	return charByte
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
	//------------------------------------------------------------
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
	//------------------------------------------------------------
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

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

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
	dataBytes := []byte(dataString)
	//------------------------------------------------------------
	if len(dataBytes) > 0 {
		for i := range dataBytes {
			dataBytes[i] ^= value
		}
	}
	//------------------------------------------------------------
	return string(dataBytes), nil
	//------------------------------------------------------------
}

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
	//------------------------------------------------------------
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
	//------------------------------------------------------------
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

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------
