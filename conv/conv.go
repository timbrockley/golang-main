/*

Copyright 2023-2024, Tim Brockley. All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

*/

package conv

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mtraver/base91"
)

const BASE_CHARSET = "!#%&()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[]^_abcdefghijklmnopqrstuvwxyz"

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

//------------------------------------------------------------
// Base_encode
//------------------------------------------------------------

func Base_encode(dataString string) string {
	//------------------------------------------------------------
	if dataString == "" {
		return dataString
	}
	//------------------------------------------------------------
	paddingLength := 0
	if len(dataString)%4 > 0 {
		paddingLength = 4 - len(dataString)%4
	}
	//------------------------------------------------------------
	outputLength := ((len(dataString) + paddingLength) / 4 * 5) - paddingLength
	outputBytes := make([]byte, outputLength)
	outputIndex := 0
	//------------------------------------------------------------
	for dataIndex := 0; dataIndex < len(dataString); dataIndex += 4 {
		//---------------------------------------------------
		var b0, b1, b2, b3 byte
		//---------------------------------------------------
		b0 = dataString[dataIndex]
		if dataIndex+1 < len(dataString) {
			b1 = dataString[dataIndex+1]
		}
		if dataIndex+2 < len(dataString) {
			b2 = dataString[dataIndex+2]
		}
		if dataIndex+3 < len(dataString) {
			b3 = dataString[dataIndex+3]
		}
		//---------------------------------------------------
		charCodeSum := int(b0)<<24 | int(b1)<<16 | int(b2)<<8 | int(b3)
		//---------------------------------------------------
		if charCodeSum == 0 {
			//---------------------------------------------------
			OutputToByteSlice(&outputBytes, outputIndex, 33)
			OutputToByteSlice(&outputBytes, outputIndex+1, 33)
			OutputToByteSlice(&outputBytes, outputIndex+2, 33)
			OutputToByteSlice(&outputBytes, outputIndex+3, 33)
			OutputToByteSlice(&outputBytes, outputIndex+4, 33)
			//---------------------------------------------------
		} else {
			//---------------------------------------------------
			for subIndex := 4; subIndex >= 0; subIndex -= 1 {
				value := charCodeSum % 85
				charCodeSum = (charCodeSum - value) / 85
				OutputToByteSlice(&outputBytes, outputIndex+subIndex, BASE_CHARSET[value])
			}
			//---------------------------------------------------
		}
		//---------------------------------------------------
		outputIndex += 5
		//---------------------------------------------------
	}
	//------------------------------------------------------------
	return string(outputBytes)
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Base_decode
//------------------------------------------------------------

func Base_decode(dataString string) (string, error) {
	//------------------------------------------------------------
	if dataString == "" {
		return dataString, nil
	}
	//------------------------------------------------------------
	paddingLength := 0
	if len(dataString)%5 > 0 {
		paddingLength = 5 - len(dataString)%5
	}
	//------------------------------------------------------------
	outputLength := ((len(dataString) + paddingLength) / 5 * 4) - paddingLength
	outputBytes := make([]byte, outputLength)
	outputIndex := 0
	//------------------------------------------------------------
	for dataIndex := 0; dataIndex < len(dataString); dataIndex += 5 {
		//---------------------------------------------------
		b0, b1, b2, b3, b4 := 84, 84, 84, 84, 84
		//---------------------------------------------------
		b0 = strings.IndexByte(BASE_CHARSET, dataString[dataIndex])
		if dataIndex+1 < len(dataString) {
			b1 = strings.IndexByte(BASE_CHARSET, dataString[dataIndex+1])
		}
		if dataIndex+2 < len(dataString) {
			b2 = strings.IndexByte(BASE_CHARSET, dataString[dataIndex+2])
		}
		if dataIndex+3 < len(dataString) {
			b3 = strings.IndexByte(BASE_CHARSET, dataString[dataIndex+3])
		}
		if dataIndex+4 < len(dataString) {
			b4 = strings.IndexByte(BASE_CHARSET, dataString[dataIndex+4])
		}
		//---------------------------------------------------
		if b0 == -1 || b1 == -1 || b2 == -1 || b3 == -1 || b4 == -1 {
			return "", fmt.Errorf("data contains one or more invalid characters")
		}
		//---------------------------------------------------
		decodedChunk := 52200625*b0 + 614125*b1 + 7225*b2 + 85*b3 + b4
		//---------------------------------------------------
		OutputToByteSlice(&outputBytes, outputIndex, byte(decodedChunk>>24))
		OutputToByteSlice(&outputBytes, outputIndex+1, byte(decodedChunk>>16))
		OutputToByteSlice(&outputBytes, outputIndex+2, byte(decodedChunk>>8))
		OutputToByteSlice(&outputBytes, outputIndex+3, byte(decodedChunk))
		//---------------------------------------------------
		outputIndex += 4
		//---------------------------------------------------
	}
	//------------------------------------------------------------
	return string(outputBytes), nil
	//------------------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

//------------------------------------------------------------
// Base64_encode
//------------------------------------------------------------

func Base64_encode(dataString string) string {
	//------------------------------------------------------------
	if dataString == "" {
		return dataString
	}
	//------------------------------------------------------------
	return base64.StdEncoding.EncodeToString([]byte(dataString))
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Base64_decode
//------------------------------------------------------------

func Base64_decode(dataString string) (string, error) {
	//------------------------------------------------------------
	var err error
	var dataBytes []byte
	//------------------------------------------------------------
	if dataString == "" {
		return dataString, err
	}
	//------------------------------------------------------------
	dataBytes, err = base64.StdEncoding.DecodeString(dataString)
	if err != nil {
		dataBytes = []byte{}
	}
	//------------------------------------------------------------
	return string(dataBytes), err
	//------------------------------------------------------------
}

//------------------------------------------------------------
// BaseEncodeByte
//------------------------------------------------------------

func OutputToByteSlice(slice *[]byte, index int, value byte) bool {
	if index >= 0 && index < len(*slice) {
		(*slice)[index] = value
		return true
	}
	return false
}

//------------------------------------------------------------
// Base64url_encode
//------------------------------------------------------------

func Base64url_encode(dataString string) string {
	//------------------------------------------------------------
	if dataString == "" {
		return dataString
	}
	//------------------------------------------------------------
	dataString = Base64_encode(dataString)
	//------------------------------------------------------------
	replacer := strings.NewReplacer(
		"+", "-",
		"/", "_",
		"=", "",
	)
	//------------------------------------------------------------
	return replacer.Replace(dataString)
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Base64url_decode
//------------------------------------------------------------

func Base64url_decode(dataString string) (string, error) {
	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	if dataString == "" {
		return dataString, err
	}
	//------------------------------------------------------------
	replacer := strings.NewReplacer(
		"-", "+",
		"_", "/",
	)
	//------------------------------------------------------------
	dataString = replacer.Replace(dataString)
	//------------------------------------------------------------
	switch len(dataString) % 4 { // Pad with trailing '='s
	case 2:
		dataString += "==" // 2 pad chars
	case 3:
		dataString += "=" // 1 pad char
	}
	//------------------------------------------------------------
	return Base64_decode(dataString)
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Base64_Base64url
//------------------------------------------------------------

func Base64_Base64url(dataString string) string {
	//------------------------------------------------------------
	if dataString == "" {
		return dataString
	}
	//------------------------------------------------------------
	replacer := strings.NewReplacer(
		"+", "-",
		"/", "_",
		"=", "",
	)
	//------------------------------------------------------------
	return replacer.Replace(dataString)
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Base64url_Base64
//------------------------------------------------------------

func Base64url_Base64(dataString string) string {
	//------------------------------------------------------------
	if dataString == "" {
		return dataString
	}
	//------------------------------------------------------------
	replacer := strings.NewReplacer(
		"-", "+",
		"_", "/",
	)
	//------------------------------------------------------------
	dataString = replacer.Replace(dataString)
	//------------------------------------------------------------
	switch len(dataString) % 4 { // Pad with trailing '='s
	case 2:
		dataString += "==" // 2 pad chars
	case 3:
		dataString += "=" // 1 pad char
	}
	//------------------------------------------------------------
	return dataString
	//------------------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

//------------------------------------------------------------
// Base91_encode
//------------------------------------------------------------

func Base91_encode(dataString string, escapeBool bool) string {
	//------------------------------------------------------------
	if dataString == "" {
		return dataString
	}
	//------------------------------------------------------------
	dataString = base91.StdEncoding.EncodeToString([]byte(dataString))
	//------------------------------------------------------------
	if escapeBool {
		//------------------------------------------------------------
		replacer := strings.NewReplacer(
			"\x22", "-q",
			"\x24", "-d",
			"\x60", "-g",
		)
		//------------------------------------------------------------
		dataString = replacer.Replace(dataString)
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
	return dataString
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Base91_decode
//------------------------------------------------------------

func Base91_decode(dataString string, unescapeBool bool) (string, error) {
	//------------------------------------------------------------
	var err error
	var dataBytes []byte
	//------------------------------------------------------------
	if dataString == "" {
		return dataString, err
	}
	//------------------------------------------------------------
	if unescapeBool {
		//------------------------------------------------------------
		replacer := strings.NewReplacer(
			"-g", "\x60",
			"-d", "\x24",
			"-q", "\x22",
		)
		//------------------------------------------------------------
		dataString = replacer.Replace(dataString)
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
	dataBytes, err = base91.StdEncoding.DecodeString(dataString)
	if err != nil {
		dataBytes = []byte{}
	}
	//------------------------------------------------------------
	return string(dataBytes), err
	//------------------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

//------------------------------------------------------------
// JSON_Marshal => json encodes input into "bytes" without escaping html characters
//------------------------------------------------------------

func JSON_Marshal(input interface{}) ([]byte, error) {
	//------------------------------------------------------------
	var err error
	var encodeBuffer bytes.Buffer
	//------------------------------------------------------------
	encoder := json.NewEncoder(&encodeBuffer)
	encoder.SetEscapeHTML(false)
	//------------------------------------------------------------
	err = encoder.Encode(input)
	if err != nil {
		return nil, err
	}
	//------------------------------------------------------------
	return bytes.TrimRight(encodeBuffer.Bytes(), "\n"), err
	//------------------------------------------------------------
}

//------------------------------------------------------------
// JSON_MarshalIndent => json encodes input into "bytes" without escaping html characters
//------------------------------------------------------------

func JSON_MarshalIndent(input interface{}, prefix string, indent string) ([]byte, error) {
	//------------------------------------------------------------
	var err error
	var encodeBuffer bytes.Buffer
	var indentBuffer bytes.Buffer
	//------------------------------------------------------------
	encoder := json.NewEncoder(&encodeBuffer)
	encoder.SetEscapeHTML(false)
	//------------------------------------------------------------
	err = encoder.Encode(input)
	if err != nil {
		return nil, err
	}
	//------------------------------------------------------------
	err = json.Indent(&indentBuffer, bytes.TrimRight(encodeBuffer.Bytes(), "\n"), prefix, indent)
	if err != nil {
		return nil, err
	}
	//------------------------------------------------------------
	return indentBuffer.Bytes(), err
	//------------------------------------------------------------
}

//------------------------------------------------------------
// JSON_encode
//------------------------------------------------------------

func JSON_encode(jsonInterface interface{}) (string, error) {
	//------------------------------------------------------------
	var err error
	var jsonBytes []byte
	//------------------------------------------------------------
	jsonBytes, err = JSON_Marshal(jsonInterface)
	if err != nil {
		jsonBytes = []byte{}
	}
	//------------------------------------------------------------
	return string(jsonBytes), err
	//------------------------------------------------------------
}

//------------------------------------------------------------
// JSON_decode
//------------------------------------------------------------

func JSON_decode(jsonString string) (interface{}, error) {
	//------------------------------------------------------------
	var err error
	var jsonInterface interface{}
	//------------------------------------------------------------
	err = json.Unmarshal([]byte(jsonString), &jsonInterface)
	if err != nil {
		jsonInterface = nil
	}
	//------------------------------------------------------------
	return jsonInterface, err
	//------------------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------
