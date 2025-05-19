//------------------------------------------------------------

package crypto

import (
	"bytes"
	"encoding/base64"
	"testing"
)

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

//--------------------------------------------------
// Encrypt
//--------------------------------------------------

func TestEncrypt(t *testing.T) {
	//------------------------------------------------------------
	var err error
	var cipherBytes []byte
	//------------------------------------------------------------
	dataString := "test1234"
	dataBytes := []byte(dataString)
	//------------------------------------------------------------
	keyBase64String := "uGMs769fAJVJhonxf7q3gXYkRWKimax/vRpZ3JKHaME="
	keyBytes, _ := base64.StdEncoding.DecodeString(keyBase64String)
	//------------------------------------------------------------
	ivBase64String := "vuKXL035Y3vwRJ1F"
	ivBytes, _ := base64.StdEncoding.DecodeString(ivBase64String)
	//------------------------------------------------------------
	cipherBytes, err = EncryptBytes(dataBytes, keyBytes, ivBytes)
	//------------------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		expected := []byte{9, 117, 130, 78, 219, 246, 196, 168, 246, 72, 56, 196, 163, 21, 88, 128, 134, 121, 118, 92, 35, 191, 247, 209}

		if !bytes.Equal(cipherBytes, expected) {
			t.Errorf("cipherBytes = %v but expected %v", cipherBytes, expected)
		}
	}
	//--------------------------------------------------
}

//--------------------------------------------------
// Decrypt
//--------------------------------------------------

func TestDecrypt(t *testing.T) {
	//------------------------------------------------------------
	var err error
	var plainBytes []byte
	//------------------------------------------------------------
	expectedString := "test1234"
	expectedBytes := []byte(expectedString)
	//------------------------------------------------------------
	keyBase64String := "uGMs769fAJVJhonxf7q3gXYkRWKimax/vRpZ3JKHaME="
	keyBytes, _ := base64.StdEncoding.DecodeString(keyBase64String)
	//------------------------------------------------------------
	ivBase64String := "vuKXL035Y3vwRJ1F"
	ivBytes, _ := base64.StdEncoding.DecodeString(ivBase64String)
	//------------------------------------------------------------
	cipherBytes := []byte{9, 117, 130, 78, 219, 246, 196, 168, 246, 72, 56, 196, 163, 21, 88, 128, 134, 121, 118, 92, 35, 191, 247, 209}
	//------------------------------------------------------------
	plainBytes, err = DecryptBytes(cipherBytes, keyBytes, ivBytes)
	//------------------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		if !bytes.Equal(plainBytes, expectedBytes) {
			t.Errorf("plainBytes = %s but expected %s", plainBytes, expectedBytes)
		}
	}
	//--------------------------------------------------
}

//--------------------------------------------------
// EncryptString / DecryptString
//--------------------------------------------------

func TestEncryptDecryptString(t *testing.T) {
	//------------------------------------------------------------
	var err error
	var encryptedBase64, decryptedString string
	//------------------------------------------------------------
	dataString := "test1234"
	//------------------------------------------------------------
	keyBase64String := "uGMs769fAJVJhonxf7q3gXYkRWKimax/vRpZ3JKHaME="
	keyBytes, _ := base64.StdEncoding.DecodeString(keyBase64String)
	//------------------------------------------------------------
	encryptedBase64, err = EncryptString(dataString, keyBytes)
	//------------------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		//--------------------------------------------------
		decryptedString, err = DecryptString(encryptedBase64, keyBytes)
		//------------------------------------------------------------
		if err != nil {
			t.Error(err)
		} else {
			//--------------------------------------------------
			if decryptedString != dataString {
				t.Errorf("decryptedString = %s but expected %s", decryptedString, dataString)
			}
			//--------------------------------------------------
		}
		//--------------------------------------------------
	}
	//--------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------
