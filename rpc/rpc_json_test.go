//--------------------------------------------------------------------------------

package rpc

import (
	"fmt"
	"testing"
)

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------

//--------------------------------------------------------------------------------
// autoInc.ID
//--------------------------------------------------------------------------------

func Test_autoInc_ID(t *testing.T) {
	//--------------------------------------------------
	var auto1 autoInc
	//--------------------------------------------------
	newID := auto1.ID()
	//--------------------------------------------------
	if newID != 1 {
		//----------------------------------------
		t.Errorf("newID = %v but should = 1", newID)
		//----------------------------------------
	}
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------

//--------------------------------------------------------------------------------
// isString
//--------------------------------------------------------------------------------

func Test_isString(t *testing.T) {
	//--------------------------------------------------
	result1 := isString(0)
	//----------------------------------------
	if result1 != false {
		t.Errorf("result1 = %v but should = %v", result1, false)
	}
	//--------------------------------------------------
	result2 := isString("string")
	//----------------------------------------
	if result2 != true {
		t.Errorf("result2 = %v but should = %v", result2, true)
	}
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// isNumber
//--------------------------------------------------------------------------------

func Test_isNumber(t *testing.T) {
	//--------------------------------------------------
	result1 := isNumber("0")
	//----------------------------------------
	if result1 != false {
		t.Errorf("result1 = %v but should = %v", result1, false)
	}
	//--------------------------------------------------
	result2 := isNumber(0)
	//----------------------------------------
	if result2 != true {
		t.Errorf("result2 = %v but should = %v", result2, true)
	}
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// isArray
//--------------------------------------------------------------------------------

func Test_isArray(t *testing.T) {
	//--------------------------------------------------
	result1 := isArray(0)
	//----------------------------------------
	if result1 != false {
		t.Errorf("result1 = %v but should = %v", result1, false)
	}
	//--------------------------------------------------
	result2 := isArray([]any{})
	//----------------------------------------
	if result2 != true {
		t.Errorf("result2 = %v but should = %v", result2, true)
	}
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// isObject
//--------------------------------------------------------------------------------

func Test_isObject(t *testing.T) {
	//--------------------------------------------------
	result1 := isObject(0)
	//----------------------------------------
	if result1 != false {
		t.Errorf("result1 = %v but should = %v", result1, false)
	}
	//--------------------------------------------------
	result2 := isObject(map[string]any{})
	//----------------------------------------
	if result2 != true {
		t.Errorf("result2 = %v but should = %v", result2, true)
	}
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------

//--------------------------------------------------------------------------------
// RPC_encode_json
//--------------------------------------------------------------------------------

func TestRPC_encode_json(t *testing.T) {
	//--------------------------------------------------
	JSON := map[string]any{"result": true}
	EXPECTED_RESULT := `{"result":true}`
	//----------------------------------------
	result, err := RPC_encode_json(JSON)
	//--------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		if result != EXPECTED_RESULT {
			t.Errorf("result = %q but should = %q", result, EXPECTED_RESULT)
		}
	}
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------

//--------------------------------------------------------------------------------
// RPC_encode_json_request
//--------------------------------------------------------------------------------

func TestRPC_encode_json_request(t *testing.T) {
	//--------------------------------------------------
	REQUEST := map[string]any{"method": "echo"}
	EXPECTED_RESULT := `{"method":"echo"}`
	//----------------------------------------
	result, err := RPC_encode_json_request(REQUEST)
	//--------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		if result != EXPECTED_RESULT {
			t.Errorf("result = %q but should = %q", result, EXPECTED_RESULT)
		}
	}
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_encode_json_response
//--------------------------------------------------------------------------------

func TestRPC_encode_json_response(t *testing.T) {
	//--------------------------------------------------
	RESPONSE := map[string]any{"result": true}
	EXPECTED_RESULT := `{"result":true}`
	//----------------------------------------
	result, err := RPC_encode_json_response(RESPONSE)
	//--------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		if result != EXPECTED_RESULT {
			t.Errorf("result = %q but should = %q", result, EXPECTED_RESULT)
		}
	}
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_encode_json_result_response
//--------------------------------------------------------------------------------

func TestRPC_encode_json_result_response(t *testing.T) {
	//--------------------------------------------------
	RESULT := true
	EXPECTED_RESULT := `{"result":true}`
	//----------------------------------------
	result, err := RPC_encode_json_result_response(RESULT)
	//--------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		if result != EXPECTED_RESULT {
			t.Errorf("result = %q but should = %q", result, EXPECTED_RESULT)
		}
	}
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_encode_json_response
//--------------------------------------------------------------------------------

func TestRPC_encode_json_error_response(t *testing.T) {
	//--------------------------------------------------
	ERROR := "something went wrong"
	EXPECTED_RESULT := `{"error":"something went wrong"}`
	//----------------------------------------
	result, err := RPC_encode_json_error_response(ERROR)
	//--------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		if result != EXPECTED_RESULT {
			t.Errorf("result = %q but should = %q", result, EXPECTED_RESULT)
		}
	}
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------

//--------------------------------------------------------------------------------
// RPC_encode_jsonrpc_response
//--------------------------------------------------------------------------------

func TestRPC_encode_jsonrpc_response(t *testing.T) {
	//--------------------------------------------------
	ID := auto.ID()
	REQUEST_MAP := map[string]any{"id": ID}
	//----------
	EXPECTED_RESULT := fmt.Sprintf(`{"id":%d,"jsonrpc":"2.0"}`, ID)
	//----------------------------------------
	jsonrpc := RPC_jsonrpcStruct{RequestMap: REQUEST_MAP}
	//----------------------------------------
	result, err := jsonrpc.RPC_encode_jsonrpc_response()
	//----------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		if result != EXPECTED_RESULT {
			t.Errorf("result = %q but should = %q", result, EXPECTED_RESULT)
		}
	}
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_encode_jsonrpc_request
//--------------------------------------------------------------------------------

func TestRPC_encode_jsonrpc_request(t *testing.T) {
	//--------------------------------------------------
	ID := auto.ID()
	REQUEST_MAP := map[string]any{"id": ID, "method": "echo"}
	//----------
	EXPECTED_RESULT := fmt.Sprintf(`{"id":%d,"jsonrpc":"2.0","method":"echo"}`, ID)
	//----------------------------------------
	jsonrpc := RPC_jsonrpcStruct{RequestMap: REQUEST_MAP}
	//----------------------------------------
	result, err := jsonrpc.RPC_encode_jsonrpc_request()
	//----------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		if result != EXPECTED_RESULT {
			t.Errorf("result = %q but should = %q", result, EXPECTED_RESULT)
		}
	}
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_encode_jsonrpc_result_response
//--------------------------------------------------------------------------------

func TestRPC_encode_jsonrpc_result_response(t *testing.T) {
	//--------------------------------------------------
	ID := auto.ID()
	REQUEST_MAP := map[string]any{"id": ID}
	RESULT := true
	//----------
	EXPECTED_RESULT := fmt.Sprintf(`{"id":%d,"jsonrpc":"2.0","result":true}`, ID)
	//----------------------------------------
	jsonrpc := RPC_jsonrpcStruct{RequestMap: REQUEST_MAP}
	//----------------------------------------
	result, err := jsonrpc.RPC_encode_jsonrpc_result_response(RESULT)
	//----------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		if result != EXPECTED_RESULT {
			t.Errorf("result = %q but should = %q", result, EXPECTED_RESULT)
		}
	}
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_encode_jsonrpc_error_response
//--------------------------------------------------------------------------------

func TestRPC_encode_jsonrpc_error_response(t *testing.T) {
	//--------------------------------------------------
	ID := auto.ID()
	REQUEST_MAP := map[string]any{"id": ID}
	ERROR_MAP := map[string]any{"code": -32000, "message": "server error"}
	//----------
	EXPECTED_RESULT := fmt.Sprintf(`{"error":{"code":-32000,"message":"server error"},"id":%d,"jsonrpc":"2.0"}`, ID)
	//----------------------------------------
	jsonrpc := RPC_jsonrpcStruct{RequestMap: REQUEST_MAP}
	//----------------------------------------
	result, err := jsonrpc.RPC_encode_jsonrpc_error_response(ERROR_MAP)
	//----------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		if result != EXPECTED_RESULT {
			t.Errorf("result = %q but should = %q", result, EXPECTED_RESULT)
		}
	}
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_encode_jsonrpc_server_error_response
//--------------------------------------------------------------------------------

func TestRPC_encode_jsonrpc_server_error_response(t *testing.T) {
	//--------------------------------------------------
	ID := auto.ID()
	REQUEST_MAP := map[string]any{"id": ID}
	DATA := "some data"
	//----------
	EXPECTED_RESULT := fmt.Sprintf(`{"error":{"code":-32000,"data":"some data","message":"server error"},"id":%d,"jsonrpc":"2.0"}`, ID)
	//----------------------------------------
	jsonrpc := RPC_jsonrpcStruct{RequestMap: REQUEST_MAP}
	//----------------------------------------
	result, err := jsonrpc.RPC_encode_jsonrpc_server_error_response(DATA)
	//----------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		if result != EXPECTED_RESULT {
			t.Errorf("result = %q but should = %q", result, EXPECTED_RESULT)
		}
	}
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_encode_jsonrpc_invalid_request_response
//--------------------------------------------------------------------------------

func TestRPC_encode_jsonrpc_invalid_request_response(t *testing.T) {
	//--------------------------------------------------
	ID := auto.ID()
	REQUEST_MAP := map[string]any{"id": ID}
	DATA := "some data"
	//----------
	EXPECTED_RESULT := fmt.Sprintf(`{"error":{"code":-32600,"data":"some data","message":"invalid request"},"id":%d,"jsonrpc":"2.0"}`, ID)
	//----------------------------------------
	jsonrpc := RPC_jsonrpcStruct{RequestMap: REQUEST_MAP}
	//----------------------------------------
	result, err := jsonrpc.RPC_encode_jsonrpc_invalid_request_response(DATA)
	//----------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		if result != EXPECTED_RESULT {
			t.Errorf("result = %q but should = %q", result, EXPECTED_RESULT)
		}
	}
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_encode_jsonrpc_method_not_found_response
//--------------------------------------------------------------------------------

func TestRPC_encode_jsonrpc_method_not_found_response(t *testing.T) {
	//--------------------------------------------------
	ID := auto.ID()
	REQUEST_MAP := map[string]any{"id": ID}
	DATA := "some data"
	//----------
	EXPECTED_RESULT := fmt.Sprintf(`{"error":{"code":-32601,"data":"some data","message":"method not found"},"id":%d,"jsonrpc":"2.0"}`, ID)
	//----------------------------------------
	jsonrpc := RPC_jsonrpcStruct{RequestMap: REQUEST_MAP}
	//----------------------------------------
	result, err := jsonrpc.RPC_encode_jsonrpc_method_not_found_response(DATA)
	//----------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		if result != EXPECTED_RESULT {
			t.Errorf("result = %q but should = %q", result, EXPECTED_RESULT)
		}
	}
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_encode_jsonrpc_invalid_params_response
//--------------------------------------------------------------------------------

func TestRPC_encode_jsonrpc_invalid_params_response(t *testing.T) {
	//--------------------------------------------------
	ID := auto.ID()
	REQUEST_MAP := map[string]any{"id": ID}
	DATA := "some data"
	//----------
	EXPECTED_RESULT := fmt.Sprintf(`{"error":{"code":-32602,"data":"some data","message":"invalid params"},"id":%d,"jsonrpc":"2.0"}`, ID)
	//----------------------------------------
	jsonrpc := RPC_jsonrpcStruct{RequestMap: REQUEST_MAP}
	//----------------------------------------
	result, err := jsonrpc.RPC_encode_jsonrpc_invalid_params_response(DATA)
	//----------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		if result != EXPECTED_RESULT {
			t.Errorf("result = %q but should = %q", result, EXPECTED_RESULT)
		}
	}
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_encode_jsonrpc_internal_error_response
//--------------------------------------------------------------------------------

func TestRPC_encode_jsonrpc_internal_error_response(t *testing.T) {
	//--------------------------------------------------
	ID := auto.ID()
	REQUEST_MAP := map[string]any{"id": ID}
	DATA := "some data"
	//----------
	EXPECTED_RESULT := fmt.Sprintf(`{"error":{"code":-32603,"data":"some data","message":"internal error"},"id":%d,"jsonrpc":"2.0"}`, ID)
	//----------------------------------------
	jsonrpc := RPC_jsonrpcStruct{RequestMap: REQUEST_MAP}
	//----------------------------------------
	result, err := jsonrpc.RPC_encode_jsonrpc_internal_error_response(DATA)
	//----------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		if result != EXPECTED_RESULT {
			t.Errorf("result = %q but should = %q", result, EXPECTED_RESULT)
		}
	}
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_encode_jsonrpc_parse_error_response
//--------------------------------------------------------------------------------

func TestRPC_encode_jsonrpc_parse_error_response(t *testing.T) {
	//--------------------------------------------------
	ID := auto.ID()
	REQUEST_MAP := map[string]any{"id": ID}
	DATA := "some data"
	//----------
	EXPECTED_RESULT := fmt.Sprintf(`{"error":{"code":-32700,"data":"some data","message":"parse error"},"id":%d,"jsonrpc":"2.0"}`, ID)
	//----------------------------------------
	jsonrpc := RPC_jsonrpcStruct{RequestMap: REQUEST_MAP}
	//----------------------------------------
	result, err := jsonrpc.RPC_encode_jsonrpc_parse_error_response(DATA)
	//----------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		if result != EXPECTED_RESULT {
			t.Errorf("result = %q but should = %q", result, EXPECTED_RESULT)
		}
	}
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------

//--------------------------------------------------------------------------------
// RPC_decode_json
//--------------------------------------------------------------------------------

func TestRPC_decode_json(t *testing.T) {
	//--------------------------------------------------
	JSON_STRING := `{"result":true}`
	EXPECTED_RESULT := map[string]any{"result": true}
	//----------------------------------------
	result, err := RPC_decode_json(JSON_STRING)
	//----------------------------------------
	resultString := fmt.Sprint(result)
	exptectedResultString := fmt.Sprint(EXPECTED_RESULT)
	//--------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		if resultString != exptectedResultString {
			t.Errorf("resultString = %v but should = %v", result, exptectedResultString)
		}
	}
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_decode_jsonrpc_request
//--------------------------------------------------------------------------------

func TestRPC_decode_jsonrpc_request(t *testing.T) {
	//--------------------------------------------------
	ID := auto.ID()
	JSON_STRING := fmt.Sprintf(`{"method":"echo","id": %d}`, ID)
	EXPECTED_RESULT := map[string]any{"id": ID, "method": "echo"}
	//----------------------------------------
	result, err := RPC_decode_jsonrpc_request(JSON_STRING)
	//----------------------------------------
	resultString := fmt.Sprint(result)
	exptectedResultString := fmt.Sprint(EXPECTED_RESULT)
	//--------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		if resultString != exptectedResultString {
			t.Errorf("resultString = %v but should = %v", result, exptectedResultString)
		}
	}
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_decode_jsonrpc_response
//--------------------------------------------------------------------------------

func TestRPC_decode_jsonrpc_response(t *testing.T) {
	//--------------------------------------------------
	ID := auto.ID()
	JSON_STRING := fmt.Sprintf(`{"result":true,"id": %d}`, ID)
	EXPECTED_RESULT := map[string]any{"id": ID, "result": true}
	//----------------------------------------
	result, err := RPC_decode_jsonrpc_response(JSON_STRING)
	//----------------------------------------
	resultString := fmt.Sprint(result)
	exptectedResultString := fmt.Sprint(EXPECTED_RESULT)
	//--------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		if resultString != exptectedResultString {
			t.Errorf("resultString = %v but should = %v", result, exptectedResultString)
		}
	}
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------
