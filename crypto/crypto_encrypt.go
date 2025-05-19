/*

Copyright 2025, Tim Brockley. All rights reserved.

This software is licensed under the MIT License.

*/

package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
)

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

//------------------------------------------------------------
// EncryptBytes
//------------------------------------------------------------

func EncryptBytes(dataBytes []byte, keyBytes []byte, ivBytes []byte) ([]byte, error) {
	//------------------------------------------------------------
	var err error
	//----------------------------------------
	var cipherBlock cipher.Block
	var cipherAEAD cipher.AEAD
	//----------------------------------------
	var cipherBytes []byte
	//------------------------------------------------------------
	cipherBlock, err = aes.NewCipher(keyBytes)
	//----------------------------------------
	if err == nil {
		//----------------------------------------
		cipherAEAD, err = cipher.NewGCM(cipherBlock)
		//----------------------------------------
		if err == nil {
			//----------------------------------------
			cipherBytes = cipherAEAD.Seal(nil, ivBytes, dataBytes, nil)
			//----------------------------------------
		}
		//----------------------------------------
	}
	//------------------------------------------------------------
	return cipherBytes, err
	//------------------------------------------------------------
}

//------------------------------------------------------------
// DecryptBytes
//------------------------------------------------------------

func DecryptBytes(cipherBytes []byte, keyBytes []byte, ivBytes []byte) ([]byte, error) {
	//------------------------------------------------------------
	var err error
	//----------------------------------------
	var cipherBlock cipher.Block
	var cipherAEAD cipher.AEAD
	//----------------------------------------
	var dataBytes []byte
	//------------------------------------------------------------
	cipherBlock, err = aes.NewCipher(keyBytes)
	//----------------------------------------
	if err == nil {
		//----------------------------------------
		cipherAEAD, err = cipher.NewGCM(cipherBlock)
		//----------------------------------------
		if err == nil {
			//----------------------------------------
			dataBytes, err = cipherAEAD.Open(nil, ivBytes, cipherBytes, nil)
			//----------------------------------------
		}
		//----------------------------------------
	}
	//------------------------------------------------------------
	return dataBytes, err
	//------------------------------------------------------------
}

//------------------------------------------------------------
// EncryptString
//------------------------------------------------------------

func EncryptString(dataString string, keyBytes []byte) (string, error) {
	//------------------------------------------------------------
	var err error
	var ivBytes, cipherBytes, outputBytes []byte
	//------------------------------------------------------------
	ivBytes, err = GenerateIV()
	if err != nil {
		return "", err
	}
	//------------------------------------------------------------
	cipherBytes, err = EncryptBytes([]byte(dataString), keyBytes, ivBytes)
	if err != nil {
		return "", err
	}
	//------------------------------------------------------------
	outputBytes = append(ivBytes, cipherBytes...)
	//------------------------------------------------------------
	return base64.StdEncoding.EncodeToString(outputBytes), nil

	//------------------------------------------------------------
}

//------------------------------------------------------------
// DecryptString
//------------------------------------------------------------

func DecryptString(encryptedBase64 string, keyBytes []byte) (string, error) {
	//------------------------------------------------------------
	var err error
	var ivBytes, base64DecodedBytes, cipherBytes, dataBytes []byte
	//------------------------------------------------------------
	base64DecodedBytes, err = base64.StdEncoding.DecodeString(encryptedBase64)
	if err != nil {
		return "", err
	}
	//------------------------------------------------------------
	if len(base64DecodedBytes) < 13 {
		return "", errors.New("invalid data length")
	}
	//------------------------------------------------------------
	ivBytes = base64DecodedBytes[:12]
	cipherBytes = base64DecodedBytes[12:]
	//------------------------------------------------------------
	dataBytes, err = DecryptBytes(cipherBytes, keyBytes, ivBytes)
	if err != nil {
		return "", err
	}
	//------------------------------------------------------------
	return string(dataBytes), nil
	//------------------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

//------------------------------------------------------------
// GenerateKey
//------------------------------------------------------------

func GenerateKey() ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	return key, err
}

//------------------------------------------------------------
// GenerateIV
//------------------------------------------------------------

func GenerateIV() ([]byte, error) {
	key := make([]byte, 12)
	_, err := rand.Read(key)
	return key, err
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------
