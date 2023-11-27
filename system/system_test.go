//------------------------------------------------------------

package system

import (
	"fmt"
	"os"
	"regexp"
	"testing"
)

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

//------------------------------------------------------------
// ConvertToString
//------------------------------------------------------------

func TestConvertToString(t *testing.T) {
	//------------------------------------------------------------
	stringVal := ConvertToString("stringVal")
	//------------------------------------------------------------
	if stringVal != "stringVal" {

		t.Errorf("ConvertToString(\"stringVal\") = %q but should = %q", stringVal, "stringVal")
	}
	//------------------------------------------------------------
	bytesVal := ConvertToString([]byte("bytesVal"))
	//------------------------------------------------------------
	if string(bytesVal) != "bytesVal" {

		t.Errorf("ConvertToString([]byte(\"bytesVal\")) = %q but should = %q", bytesVal, []byte("bytesVal"))
	}
	//------------------------------------------------------------
	intVal := ConvertToString(123)
	//------------------------------------------------------------
	if intVal != "123" {

		t.Errorf("ConvertToString(123) = %q but should = %q", intVal, "123")
	}
	//------------------------------------------------------------
	floatVal := ConvertToString(123.456)
	//------------------------------------------------------------
	if floatVal != "123.456" {

		t.Errorf("ConvertToString(123.456) = %q but should = %q", floatVal, "123.456")
	}
	//------------------------------------------------------------
	boolVal1 := ConvertToString(true)
	//------------------------------------------------------------
	if boolVal1 != "true" {

		t.Error("ConvertToString(true) = \"false\" but should = \"true\"")
	}
	//------------------------------------------------------------
	boolVal2 := ConvertToString(false)
	//------------------------------------------------------------
	if boolVal2 != "false" {

		t.Error("ConvertToString(false) = \"true\" but should = \"false\"")
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ConvertToBytes
//------------------------------------------------------------

func TestConvertToBytes(t *testing.T) {
	//------------------------------------------------------------
	stringVal := ConvertToBytes("stringVal")
	//------------------------------------------------------------
	if string(stringVal) != "stringVal" {

		t.Errorf("ConvertToBytes(\"stringVal\") = %q but should = %q", stringVal, []byte("stringVal"))
	}
	//------------------------------------------------------------
	bytesVal := ConvertToBytes([]byte("bytesVal"))
	//------------------------------------------------------------
	if string(bytesVal) != "bytesVal" {

		t.Errorf("ConvertToBytes([]byte(\"bytesVal\")) = %v but should = %v", bytesVal, []byte("bytesVal"))
	}
	//------------------------------------------------------------
	intVal1 := ConvertToBytes(123)
	//------------------------------------------------------------
	// if fmt.Sprintf("%#v", intVal1) != "[]byte{0x7b}" {

	// 	t.Errorf("ConvertToBytes(123) = %v but should = %v", intVal1, []byte{123})
	// }
	if fmt.Sprintf("%#v", intVal1) != "[]byte{}" {

		t.Errorf("ConvertToBytes(123) = %v but should = %v", intVal1, []byte{})
	}
	//------------------------------------------------------------
	intVal2 := ConvertToBytes(257)
	//------------------------------------------------------------
	// if fmt.Sprintf("%#v", intVal2) != "[]byte{0x1}" {

	// 	t.Errorf("ConvertToBytes(257) = %v but should = %v", intVal2, []byte{1})
	// }
	if fmt.Sprintf("%#v", intVal2) != "[]byte{}" {

		t.Errorf("ConvertToBytes(257) = %v but should = %v", intVal2, []byte{})
	}
	//------------------------------------------------------------
	floatVal := ConvertToBytes(123.456)
	//------------------------------------------------------------
	if fmt.Sprintf("%#v", floatVal) != "[]byte{}" {

		t.Errorf("ConvertToBytes(123.456) = %v but should = %v", floatVal, []byte{})
	}
	//------------------------------------------------------------
	boolVal1 := ConvertToBytes(true)
	//------------------------------------------------------------
	if fmt.Sprintf("%#v", boolVal1) != "[]byte{}" {

		t.Errorf("ConvertToBytes(true) = %v but should = %v", boolVal1, []byte{})
	}
	//------------------------------------------------------------
	boolVal2 := ConvertToBytes(false)
	//------------------------------------------------------------
	if fmt.Sprintf("%#v", boolVal2) != "[]byte{}" {

		t.Errorf("ConvertToBytes(false) = %v but should = %v", boolVal2, []byte{})
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ConvertToInt
//------------------------------------------------------------

func TestConvertToInt(t *testing.T) {
	//------------------------------------------------------------
	stringVal1 := ConvertToInt("stringVal")
	//------------------------------------------------------------
	if stringVal1 != 0 {

		t.Errorf("ConvertToInt(\"stringVal\") = %d but should = %d", stringVal1, 0)
	}
	//------------------------------------------------------------
	stringVal2 := ConvertToInt("123")
	//------------------------------------------------------------
	if stringVal2 != 123 {

		t.Errorf("ConvertToInt(\"123\") = %d but should = %d", stringVal2, 123)
	}
	//------------------------------------------------------------
	stringVal3 := ConvertToInt("123.456")
	//------------------------------------------------------------
	if stringVal3 != 123 {

		t.Errorf("ConvertToInt(\"123.456\") = %d but should = %d", stringVal3, 123)
	}
	//------------------------------------------------------------
	bytesVal := ConvertToInt("bytesVal")
	//------------------------------------------------------------
	if bytesVal != 0 {

		t.Errorf("ConvertToInt([]byte(\"bytesVal\")) = %d but should = %d", bytesVal, 0)
	}
	//------------------------------------------------------------
	intVal := ConvertToInt(123)
	//------------------------------------------------------------
	if intVal != 123 {

		t.Errorf("ConvertToInt(123) = %d but should = %d", intVal, 123)
	}
	//------------------------------------------------------------
	floatVal := ConvertToInt(123.456)
	//------------------------------------------------------------
	if floatVal != 123 {

		t.Errorf("ConvertToInt(123.456) = %d but should = %d", floatVal, 123)
	}
	//------------------------------------------------------------
	boolVal1 := ConvertToInt(true)
	//------------------------------------------------------------
	if boolVal1 != 1 {

		t.Errorf("ConvertToInt(true) = %d but should = %d", boolVal1, 1)
	}
	//------------------------------------------------------------
	boolVal2 := ConvertToInt(false)
	//------------------------------------------------------------
	if boolVal2 != 0 {

		t.Errorf("ConvertToInt(false) = %d but should = %d", boolVal2, 0)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ConvertToFloat
//------------------------------------------------------------

func TestConvertToFloat(t *testing.T) {
	//------------------------------------------------------------
	stringVal1 := ConvertToFloat("stringVal")
	//------------------------------------------------------------
	if stringVal1 != 0 {

		t.Errorf("ConvertToFloat(\"stringVal\") = %v but should = %v", stringVal1, 0)
	}
	//------------------------------------------------------------
	stringVal2 := ConvertToFloat("123")
	//------------------------------------------------------------
	if stringVal2 != 123 {

		t.Errorf("ConvertToFloat(\"123\") = %v but should = %v", stringVal2, 123)
	}
	//------------------------------------------------------------
	stringVal3 := ConvertToFloat("123.456")
	//------------------------------------------------------------
	if stringVal3 != 123.456 {

		t.Errorf("ConvertToFloat(\"123.456\") = %v but should = %v", stringVal3, 123.456)
	}
	//------------------------------------------------------------
	bytesVal := ConvertToFloat("bytesVal")
	//------------------------------------------------------------
	if bytesVal != 0 {

		t.Errorf("ConvertToFloat([]byte(\"bytesVal\")) = %v but should = %v", bytesVal, 0)
	}
	//------------------------------------------------------------
	intVal := ConvertToFloat(123)
	//------------------------------------------------------------
	if intVal != 123 {

		t.Errorf("ConvertToFloat(\"intVal\") = %v but should = %v", intVal, 123)
	}
	//------------------------------------------------------------
	floatVal := ConvertToFloat(123.456)
	//------------------------------------------------------------
	if floatVal != 123.456 {

		t.Errorf("ConvertToFloat(\"floatVal\") = %v but should = %v", floatVal, 123.456)
	}
	//------------------------------------------------------------
	boolVal1 := ConvertToFloat(true)
	//------------------------------------------------------------
	if boolVal1 != 1 {

		t.Errorf("ConvertToFloat(true) = %v but should = %v", boolVal1, 1)
	}
	//------------------------------------------------------------
	boolVal2 := ConvertToFloat(false)
	//------------------------------------------------------------
	if boolVal2 != 0 {

		t.Errorf("ConvertToFloat(false) = %v but should = %v", boolVal2, 0)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ConvertToBool
//------------------------------------------------------------

func TestConvertToBool(t *testing.T) {
	//------------------------------------------------------------
	stringVal1 := ConvertToBool("")
	//------------------------------------------------------------
	if stringVal1 {

		t.Errorf("ConvertToBool(\"\") = %v but should = %v", stringVal1, false)
	}
	//------------------------------------------------------------
	stringVal2 := ConvertToBool("123")
	//------------------------------------------------------------
	if !stringVal2 {

		t.Errorf("ConvertToBool(123) = %v but should = %v", stringVal2, true)
	}
	//------------------------------------------------------------
	stringVal3 := ConvertToBool("true")
	//------------------------------------------------------------
	if !stringVal3 {

		t.Errorf("ConvertToBool(\"true\") = %v but should = %v", stringVal3, true)
	}
	//------------------------------------------------------------
	stringVal4 := ConvertToBool("false")
	//------------------------------------------------------------
	if stringVal4 {

		t.Errorf("ConvertToBool(\"false\") = %v but should = %v", stringVal4, false)
	}
	//------------------------------------------------------------
	stringVal5 := ConvertToBool("true")
	//------------------------------------------------------------
	if !stringVal5 {

		t.Errorf("ConvertToBool(\"true\") = %v but should = %v", stringVal5, true)
	}
	//------------------------------------------------------------
	bytesVal := ConvertToBool([]byte{})
	//------------------------------------------------------------
	if bytesVal {

		t.Errorf("ConvertToBool([]byte(\"bytesVal\")) = %v but should = %v", bytesVal, false)
	}
	//------------------------------------------------------------
	intVal := ConvertToBool(123)
	//------------------------------------------------------------
	if !intVal {

		t.Errorf("ConvertToBool(123) = %v but should = %v", intVal, true)
	}
	//------------------------------------------------------------
	floatVal := ConvertToBool(123.456)
	//------------------------------------------------------------
	if !floatVal {

		t.Errorf("ConvertToBool(123.456) = %v but should = %v", floatVal, true)
	}
	//------------------------------------------------------------
	boolVal1 := ConvertToBool(true)
	//------------------------------------------------------------
	if !boolVal1 {

		t.Error("ConvertToBool(true) = false but should = true")
	}
	//------------------------------------------------------------
	boolVal2 := ConvertToBool(false)
	//------------------------------------------------------------
	if boolVal2 {

		t.Error("ConvertToBool(false) = true but should = false")
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// GetHostname
//------------------------------------------------------------

func TestGetHostname(t *testing.T) {
	//------------------------------------------------------------
	result := GetHostname()
	//------------------------------------------------------------
	if result == "" {

		t.Errorf("hostname = %q", result)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// GetLocalIPs
//------------------------------------------------------------

func TestGetLocalIPs(t *testing.T) {
	//------------------------------------------------------------
	result, err := GetLocalIPs()
	//------------------------------------------------------------
	if err != nil {

		t.Error(err)
	}
	if result == nil {

		t.Errorf("LocalIPs = %q", result)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// GetLocalIP
//------------------------------------------------------------

func TestGetLocalIPAddr(t *testing.T) {
	//------------------------------------------------------------
	result := GetLocalIPAddr()
	//------------------------------------------------------------
	if result == "" {

		t.Errorf("LocalIPs = %q", result)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// GetOS
//------------------------------------------------------------

func TestGetOS(t *testing.T) {
	//------------------------------------------------------------
	result := GetHostname()
	//------------------------------------------------------------
	if result == "" {

		t.Errorf("hostname = %q", result)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// GetOS
//------------------------------------------------------------

func TestCLIParams(t *testing.T) {
	//------------------------------------------------------------
	result := CLIParams()
	resultType := fmt.Sprintf("%T", result)
	//------------------------------------------------------------
	if resultType != "[]string" {

		t.Errorf("result type = %q but should = %q", resultType, "[]string")

	} else {

		if len(result) <= 0 {

			t.Error("result does not contain any elements")

		} else if resultType != "[]string" || result[0] != os.Args[0] {

			t.Errorf("result[0] = %q but should = %q", result[0], os.Args[0])
		}
	}
	//------------------------------------------------------------
}

func TestCLIParam(t *testing.T) {
	//------------------------------------------------------------
	result := CLIParam(0)
	resultType := fmt.Sprintf("%T", result)
	//------------------------------------------------------------
	if resultType != "string" {
		t.Errorf("result type = %q but should = %q", resultType, "string")
	}
	//------------------------------------------------------------
	if result != os.Args[0] {
		t.Errorf("result = %q but should = %q", result, os.Args[0])
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// GetENV
//------------------------------------------------------------

func TestGetENV(t *testing.T) {
	//------------------------------------------------------------
	os.Setenv("GOLANG_TEST", "TEST1234")
	//----------
	result := GetENV("GOLANG_TEST")
	//------------------------------------------------------------
	if result != "TEST1234" {

		t.Errorf("GOLANG_TEST = %q", result)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// GetENVs
//------------------------------------------------------------

func TestGetENVs(t *testing.T) {
	//------------------------------------------------------------
	result := GetENVs()
	//------------------------------------------------------------
	if result == nil {
		t.Errorf("GetENVs = %v", result)
	}
	//----------
	if os.Getenv("GOROOT") == "" {
		t.Error("error getting environment variables")
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// SetENV
//------------------------------------------------------------

func TestSetENV(t *testing.T) {
	//------------------------------------------------------------
	SetENV("GOLANG_SETENV_TEST", "SETENV_TEST1234")
	//----------
	result := os.Getenv("GOLANG_SETENV_TEST")
	//------------------------------------------------------------
	if result != "SETENV_TEST1234" {

		t.Errorf("GOLANG_SETENV_TEST = %q", result)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// SetENVs
//------------------------------------------------------------

func TestSetENVs(t *testing.T) {
	//------------------------------------------------------------
	err := SetENVs(map[string]string{"GOLANG_SETENV_TEST1": "SETENV_TEST1", "GOLANG_SETENV_TEST2": "SETENV_TEST2"})
	//----------
	result1 := os.Getenv("GOLANG_SETENV_TEST1")
	result2 := os.Getenv("GOLANG_SETENV_TEST2")
	//------------------------------------------------------------
	if err != nil {
		//----------
		t.Error(err)
		//----------
	} else {
		//----------
		if result1 != "SETENV_TEST1" {

			t.Errorf("GOLANG_SETENV_TEST1 = %q", result1)
		}
		//----------
		if result2 != "SETENV_TEST2" {

			t.Errorf("GOLANG_SETENV_TEST2 = %q", result2)
		}
		//----------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// LoadENVs
//------------------------------------------------------------

func TestLoadENVs(t *testing.T) {
	//------------------------------------------------------------
	err := LoadENVs()
	//----------
	if err != nil {

		t.Error(err)
	}
	//----------
	if os.Getenv("_SYSTEM_TEST") != "_SYSTEM_TEST_VALUE" {

		t.Error("error getting environment variables")
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

func UUIDIsValid(uuidString string) bool {
	//------------------------------------------------------------
	match, err := regexp.MatchString(`^[0-9a-fA-F]{8}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{12}$`, uuidString)
	if err != nil {

		return false
	}
	//------------------------------------------------------------
	return match
	//------------------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------
