//------------------------------------------------------------

package system

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"testing"
)

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

//------------------------------------------------------------
// ToString
//------------------------------------------------------------

func TestToString(t *testing.T) {
	//------------------------------------------------------------
	stringVal := ToString("stringVal")
	//------------------------------------------------------------
	if stringVal != "stringVal" {
		t.Errorf("ToString(\"stringVal\") = %q but should = %q", stringVal, "stringVal")
	}
	//------------------------------------------------------------
	bytesVal := ToString([]byte("bytesVal"))
	//------------------------------------------------------------
	if string(bytesVal) != "bytesVal" {
		t.Errorf("ToString([]byte(\"bytesVal\")) = %q but should = %q", bytesVal, []byte("bytesVal"))
	}
	//------------------------------------------------------------
	intVal := ToString(123)
	//------------------------------------------------------------
	if intVal != "123" {
		t.Errorf("ToString(123) = %q but should = %q", intVal, "123")
	}
	//------------------------------------------------------------
	floatVal := ToString(123.456)
	//------------------------------------------------------------
	if floatVal != "123.456" {
		t.Errorf("ToString(123.456) = %q but should = %q", floatVal, "123.456")
	}
	//------------------------------------------------------------
	boolVal1 := ToString(true)
	//------------------------------------------------------------
	if boolVal1 != "true" {
		t.Error("ToString(true) = \"false\" but should = \"true\"")
	}
	//------------------------------------------------------------
	boolVal2 := ToString(false)
	//------------------------------------------------------------
	if boolVal2 != "false" {
		t.Error("ToString(false) = \"true\" but should = \"false\"")
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ToBytes
//------------------------------------------------------------

func TestToBytes(t *testing.T) {
	//------------------------------------------------------------
	bytesVal := ToBytes([]byte{0x31, 0x32, 0x33})
	//------------------------------------------------------------
	if fmt.Sprintf("%#v", bytesVal) != "[]byte{0x31, 0x32, 0x33}" {
		t.Errorf("ToBytes(\"stringVal\") = %q but should = %q", bytesVal, []byte{0x31, 0x32, 0x33})
	}
	//------------------------------------------------------------
	stringVal := ToBytes("stringVal")
	//------------------------------------------------------------
	if string(stringVal) != "stringVal" {

		t.Errorf("ToBytes(\"stringVal\") = %q but should = %q", stringVal, []byte("stringVal"))
	}
	//------------------------------------------------------------
	intVal := ToBytes(123)
	//------------------------------------------------------------
	if fmt.Sprintf("%#v", intVal) != "[]byte{0x31, 0x32, 0x33}" {
		t.Errorf("ToBytes(123) = %v but should = %v", intVal, []byte{0x31, 0x32, 0x33})
	}
	//------------------------------------------------------------
	floatVal := ToBytes(123.456)
	//------------------------------------------------------------
	if fmt.Sprintf("%#v", floatVal) != "[]byte{0x31, 0x32, 0x33, 0x2e, 0x34, 0x35, 0x36}" {
		t.Errorf("ToBytes(123.456) = %v but should = %v", floatVal, []byte{0x31, 0x32, 0x33, 0x2e, 0x34, 0x35, 0x36})
	}
	//------------------------------------------------------------
	boolVal := ToBytes(true)
	//------------------------------------------------------------
	if fmt.Sprintf("%#v", boolVal) != "[]byte{}" {
		t.Errorf("ToBytes(true) = %v but should = %v", boolVal, []byte{})
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ToInt
//------------------------------------------------------------

func TestToInt(t *testing.T) {
	//------------------------------------------------------------
	stringVal1 := ToInt("stringVal")
	//------------------------------------------------------------
	if stringVal1 != 0 {
		t.Errorf("ToInt(\"stringVal\") = %d but should = %d", stringVal1, 0)
	}
	//------------------------------------------------------------
	stringVal2 := ToInt("123")
	//------------------------------------------------------------
	if stringVal2 != 123 {
		t.Errorf("ToInt(\"123\") = %d but should = %d", stringVal2, 123)
	}
	//------------------------------------------------------------
	stringVal3 := ToInt("123.456")
	//------------------------------------------------------------
	if stringVal3 != 123 {
		t.Errorf("ToInt(\"123.456\") = %d but should = %d", stringVal3, 123)
	}
	//------------------------------------------------------------
	bytesVal := ToInt("bytesVal")
	//------------------------------------------------------------
	if bytesVal != 0 {
		t.Errorf("ToInt([]byte(\"bytesVal\")) = %d but should = %d", bytesVal, 0)
	}
	//------------------------------------------------------------
	intVal := ToInt(123)
	//------------------------------------------------------------
	if intVal != 123 {
		t.Errorf("ToInt(123) = %d but should = %d", intVal, 123)
	}
	//------------------------------------------------------------
	floatVal := ToInt(123.456)
	//------------------------------------------------------------
	if floatVal != 123 {
		t.Errorf("ToInt(123.456) = %d but should = %d", floatVal, 123)
	}
	//------------------------------------------------------------
	boolVal1 := ToInt(true)
	//------------------------------------------------------------
	if boolVal1 != 1 {
		t.Errorf("ToInt(true) = %d but should = %d", boolVal1, 1)
	}
	//------------------------------------------------------------
	boolVal2 := ToInt(false)
	//------------------------------------------------------------
	if boolVal2 != 0 {
		t.Errorf("ToInt(false) = %d but should = %d", boolVal2, 0)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ToFloat
//------------------------------------------------------------

func TestToFloat(t *testing.T) {
	//------------------------------------------------------------
	stringVal1 := ToFloat("stringVal")
	//------------------------------------------------------------
	if stringVal1 != 0 {
		t.Errorf("ToFloat(\"stringVal\") = %v but should = %v", stringVal1, 0)
	}
	//------------------------------------------------------------
	stringVal2 := ToFloat("123")
	//------------------------------------------------------------
	if stringVal2 != 123 {
		t.Errorf("ToFloat(\"123\") = %v but should = %v", stringVal2, 123)
	}
	//------------------------------------------------------------
	stringVal3 := ToFloat("123.456")
	//------------------------------------------------------------
	if stringVal3 != 123.456 {
		t.Errorf("ToFloat(\"123.456\") = %v but should = %v", stringVal3, 123.456)
	}
	//------------------------------------------------------------
	bytesVal := ToFloat("bytesVal")
	//------------------------------------------------------------
	if bytesVal != 0 {
		t.Errorf("ToFloat([]byte(\"bytesVal\")) = %v but should = %v", bytesVal, 0)
	}
	//------------------------------------------------------------
	intVal := ToFloat(123)
	//------------------------------------------------------------
	if intVal != 123 {
		t.Errorf("ToFloat(\"intVal\") = %v but should = %v", intVal, 123)
	}
	//------------------------------------------------------------
	floatVal := ToFloat(123.456)
	//------------------------------------------------------------
	if floatVal != 123.456 {
		t.Errorf("ToFloat(\"floatVal\") = %v but should = %v", floatVal, 123.456)
	}
	//------------------------------------------------------------
	boolVal1 := ToFloat(true)
	//------------------------------------------------------------
	if boolVal1 != 1 {
		t.Errorf("ToFloat(true) = %v but should = %v", boolVal1, 1)
	}
	//------------------------------------------------------------
	boolVal2 := ToFloat(false)
	//------------------------------------------------------------
	if boolVal2 != 0 {
		t.Errorf("ToFloat(false) = %v but should = %v", boolVal2, 0)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ToBool
//------------------------------------------------------------

func TestToBool(t *testing.T) {
	//------------------------------------------------------------
	stringVal1 := ToBool("")
	//------------------------------------------------------------
	if stringVal1 {
		t.Errorf("ToBool(\"\") = %v but should = %v", stringVal1, false)
	}
	//------------------------------------------------------------
	stringVal2 := ToBool("123")
	//------------------------------------------------------------
	if !stringVal2 {
		t.Errorf("ToBool(123) = %v but should = %v", stringVal2, true)
	}
	//------------------------------------------------------------
	stringVal3 := ToBool("true")
	//------------------------------------------------------------
	if !stringVal3 {
		t.Errorf("ToBool(\"true\") = %v but should = %v", stringVal3, true)
	}
	//------------------------------------------------------------
	stringVal4 := ToBool("false")
	//------------------------------------------------------------
	if stringVal4 {
		t.Errorf("ToBool(\"false\") = %v but should = %v", stringVal4, false)
	}
	//------------------------------------------------------------
	stringVal5 := ToBool("true")
	//------------------------------------------------------------
	if !stringVal5 {
		t.Errorf("ToBool(\"true\") = %v but should = %v", stringVal5, true)
	}
	//------------------------------------------------------------
	bytesVal := ToBool([]byte{})
	//------------------------------------------------------------
	if bytesVal {
		t.Errorf("ToBool([]byte(\"bytesVal\")) = %v but should = %v", bytesVal, false)
	}
	//------------------------------------------------------------
	intVal := ToBool(123)
	//------------------------------------------------------------
	if !intVal {
		t.Errorf("ToBool(123) = %v but should = %v", intVal, true)
	}
	//------------------------------------------------------------
	floatVal := ToBool(123.456)
	//------------------------------------------------------------
	if !floatVal {
		t.Errorf("ToBool(123.456) = %v but should = %v", floatVal, true)
	}
	//------------------------------------------------------------
	boolVal1 := ToBool(true)
	//------------------------------------------------------------
	if !boolVal1 {
		t.Error("ToBool(true) = false but should = true")
	}
	//------------------------------------------------------------
	boolVal2 := ToBool(false)
	//------------------------------------------------------------
	if boolVal2 {
		t.Error("ToBool(false) = true but should = false")
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ToStringSlice
//------------------------------------------------------------

func TestToStringSlice(t *testing.T) {
	//------------------------------------------------------------
	stringSlice := ToStringSlice([]string{"1", "2", "3"})
	//------------------------------------------------------------
	if fmt.Sprintf("%#v", stringSlice) != `[]string{"1", "2", "3"}` {
		t.Errorf(`ToStringSlice([]string{"1", "2", "3"}) = %v but should = %v`, fmt.Sprintf("%#v", stringSlice), `[]string{"1", "2", "3"}`)
	}
	//------------------------------------------------------------
	stringVal := ToStringSlice("")
	//------------------------------------------------------------
	if fmt.Sprintf("%#v", stringVal) != "[]string{}" {
		t.Errorf("ToStringSlice([]string{}) = %v but should = %v", fmt.Sprintf("%#v", stringVal), "[]string{}")
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ToIntSlice
//------------------------------------------------------------

func TestToIntSlice(t *testing.T) {
	//------------------------------------------------------------
	intSlice := ToIntSlice([]int{1, 2, 3})
	//------------------------------------------------------------
	if fmt.Sprintf("%#v", intSlice) != `[]int{1, 2, 3}` {
		t.Errorf(`ToIntSlice([]int{1, 2, 3}) = %v but should = %v`, fmt.Sprintf("%#v", intSlice), `[]int{1, 2, 3}`)
	}
	//------------------------------------------------------------
	intVal := ToIntSlice("")
	//------------------------------------------------------------
	if fmt.Sprintf("%#v", intVal) != "[]int{}" {
		t.Errorf("ToIntSlice([]int{}) = %v but should = %v", fmt.Sprintf("%#v", intVal), "[]int{}")
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ToWriter
//------------------------------------------------------------

func TestToWriter(t *testing.T) {
	//------------------------------------------------------------
	writerVal := ToWriter(io.Discard)
	//------------------------------------------------------------
	if writerVal == nil || fmt.Sprintf("%#v", writerVal) != "io.discard{}" {
		t.Errorf("ToWriter(io.Discard) = %v but should = %v", fmt.Sprintf("%#v", writerVal), "io.discard{}")
	}
	//------------------------------------------------------------
	stringVal := ToWriter("")
	//------------------------------------------------------------
	// if stringVal != nil {
	// 	t.Errorf("ToWriter(\"\") = %v but should = %v", stringVal, nil)
	// }
	if stringVal == nil || fmt.Sprintf("%#v", stringVal) != "io.discard{}" {
		t.Errorf("ToWriter(\"\") = %v but should = %v", stringVal, "io.discard{}")
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// CopyMap
//------------------------------------------------------------

func TestCopyMap(t *testing.T) {
	//------------------------------------------------------------
	map1 := map[string]any{"IntVal": 1}
	map2 := CopyMap(map1)
	map2["IntVal"] = 2
	//------------------------------------------------------------
	if map1["IntVal"] != 1 {
		t.Errorf(`map1["IntVal"] = %d but should = %d`, map1["IntVal"], 1)
	}
	if map2["IntVal"] != 2 {
		t.Errorf(`map2["IntVal"] = %d but should = %d`, map2["IntVal"], 2)
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
	//--------------------
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
	//--------------------
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
	//--------------------
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
	//--------------------
	result1 := os.Getenv("GOLANG_SETENV_TEST1")
	result2 := os.Getenv("GOLANG_SETENV_TEST2")
	//------------------------------------------------------------
	if err != nil {
		//--------------------
		t.Error(err)
		//--------------------
	} else {
		//--------------------
		if result1 != "SETENV_TEST1" {

			t.Errorf("GOLANG_SETENV_TEST1 = %q", result1)
		}
		//--------------------
		if result2 != "SETENV_TEST2" {

			t.Errorf("GOLANG_SETENV_TEST2 = %q", result2)
		}
		//--------------------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// LoadENVs
//------------------------------------------------------------

func TestLoadENVs(t *testing.T) {
	//------------------------------------------------------------
	err := LoadENVs()
	//--------------------
	if err != nil {

		t.Error(err)
	}
	//--------------------
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
