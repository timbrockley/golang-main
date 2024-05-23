/*

Copyright (c) 2023-2024, Tim Brockley. All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

*/

package rpc

import (
	"errors"
	"fmt"
	"sync"

	"github.com/timbrockley/golang-main/conv"
)

//--------------------------------------------------------------------------------

type RPC_jsonrpcStruct struct {
	RequestMap  map[string]any
	ResponseMap map[string]any
}

//--------------------------------------------------------------------------------
// auto increment ID
//--------------------------------------------------------------------------------

type autoInc struct {
	sync.Mutex
	id int
}

func (a *autoInc) ID() int {
	a.Lock()
	defer a.Unlock()
	a.id++
	return a.id
}

var auto autoInc

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------

//--------------------------------------------------------------------------------
// RPC_encode_json
//--------------------------------------------------------------------------------

func RPC_encode_json(jsonMap map[string]any) (string, error) {
	//--------------------------------------------------
	return conv.JSON_encode(jsonMap)
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------

//--------------------------------------------------------------------------------
// RPC_encode_json_request
//--------------------------------------------------------------------------------

func RPC_encode_json_request(requestMap map[string]any) (string, error) {
	//--------------------------------------------------
	return RPC_encode_json(requestMap)
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_encode_json_response
//--------------------------------------------------------------------------------

func RPC_encode_json_response(responseMap map[string]any) (string, error) {
	//--------------------------------------------------
	return RPC_encode_json(responseMap)
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_encode_json_result_response
//--------------------------------------------------------------------------------

func RPC_encode_json_result_response(result any) (string, error) {
	//--------------------------------------------------
	responseMap := map[string]any{"result": result}
	//--------------------------------------------------
	return RPC_encode_json_response(responseMap)
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_encode_json_error_response
//--------------------------------------------------------------------------------

func RPC_encode_json_error_response(error string, Data ...any) (string, error) {
	//--------------------------------------------------
	responseMap := map[string]any{"error": error}
	//--------------------------------------------------
	if Data != nil && Data[0] != nil {
		responseMap["data"] = Data[0]
	}
	//--------------------------------------------------
	return RPC_encode_json_response(responseMap)
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------

//--------------------------------------------------------------------------------
// RPC_encode_jsonrpc_request
//--------------------------------------------------------------------------------

func (jsonrpc *RPC_jsonrpcStruct) RPC_encode_jsonrpc_request() (string, error) {
	//--------------------------------------------------
	if !isString(jsonrpc.RequestMap["method"]) {
		return "", fmt.Errorf("method is not defined or is not a string")
	}
	//--------------------------------------------------
	if jsonrpc.RequestMap["params"] != nil && !isArray(jsonrpc.RequestMap["params"]) && !isObject(jsonrpc.RequestMap["params"]) {
		return "", fmt.Errorf("params is not an array or an object")
	}
	//--------------------------------------------------
	if jsonrpc.RequestMap["id"] != nil && !isNumber(jsonrpc.RequestMap["id"]) && !isString(jsonrpc.RequestMap["id"]) {
		return "", fmt.Errorf("id is not a number or a string")
	}
	//--------------------------------------------------
	jsonrpc.RequestMap["jsonrpc"] = "2.0"
	//----------
	if jsonrpc.RequestMap["id"] == nil {
		jsonrpc.RequestMap["id"] = auto.ID()
	}
	//--------------------------------------------------
	return RPC_encode_json_response(jsonrpc.RequestMap)
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_encode_jsonrpc_response
//--------------------------------------------------------------------------------

func (jsonrpc *RPC_jsonrpcStruct) RPC_encode_jsonrpc_response() (string, error) {
	//--------------------------------------------------
	if jsonrpc.ResponseMap == nil {
		jsonrpc.ResponseMap = map[string]any{}
	}
	//--------------------------------------------------
	jsonrpc.ResponseMap["jsonrpc"] = "2.0"
	//--------------------------------------------------
	if jsonrpc.ResponseMap["id"] == nil {

		if isNumber(jsonrpc.RequestMap["id"]) || isString(jsonrpc.RequestMap["id"]) {

			jsonrpc.ResponseMap["id"] = jsonrpc.RequestMap["id"]

		} else {

			jsonrpc.ResponseMap["id"] = nil // ensure a map entry for id is present
		}
	}
	//--------------------------------------------------
	return RPC_encode_json_response(jsonrpc.ResponseMap)
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_encode_jsonrpc_result_response
//--------------------------------------------------------------------------------

func (jsonrpc *RPC_jsonrpcStruct) RPC_encode_jsonrpc_result_response(result any) (string, error) {
	//--------------------------------------------------
	if jsonrpc.ResponseMap == nil {
		jsonrpc.ResponseMap = map[string]any{}
	}
	//--------------------------------------------------
	jsonrpc.ResponseMap["result"] = result
	//--------------------------------------------------
	return jsonrpc.RPC_encode_jsonrpc_response()
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_encode_jsonrpc_error_response
//--------------------------------------------------------------------------------

func (jsonrpc *RPC_jsonrpcStruct) RPC_encode_jsonrpc_error_response(errorMap map[string]any) (string, error) {
	//--------------------------------------------------
	if jsonrpc.ResponseMap == nil {
		jsonrpc.ResponseMap = map[string]any{}
	}
	//--------------------------------------------------
	jsonrpc.ResponseMap["error"] = errorMap
	//--------------------------------------------------
	return jsonrpc.RPC_encode_jsonrpc_response()
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_encode_jsonrpc_server_error_response
//--------------------------------------------------------------------------------

func (jsonrpc *RPC_jsonrpcStruct) RPC_encode_jsonrpc_server_error_response(Data ...any) (string, error) {
	//--------------------------------------------------
	errorMap := map[string]any{"code": -32000, "message": "server error"}
	//--------------------------------------------------
	if Data != nil && Data[0] != nil {

		errorMap["data"] = Data[0]
	}
	//--------------------------------------------------
	return jsonrpc.RPC_encode_jsonrpc_error_response(errorMap)
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_encode_jsonrpc_invalid_request_response
//--------------------------------------------------------------------------------

func (jsonrpc *RPC_jsonrpcStruct) RPC_encode_jsonrpc_invalid_request_response(Data ...any) (string, error) {
	//--------------------------------------------------
	errorMap := map[string]any{"code": -32600, "message": "invalid request"}
	//--------------------------------------------------
	if Data != nil && Data[0] != nil {

		errorMap["data"] = Data[0]
	}
	//--------------------------------------------------
	return jsonrpc.RPC_encode_jsonrpc_error_response(errorMap)
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_encode_jsonrpc_method_not_found_response
//--------------------------------------------------------------------------------

func (jsonrpc *RPC_jsonrpcStruct) RPC_encode_jsonrpc_method_not_found_response(Data ...any) (string, error) {
	//--------------------------------------------------
	errorMap := map[string]any{"code": -32601, "message": "method not found"}
	//--------------------------------------------------
	if Data != nil && Data[0] != nil {

		errorMap["data"] = Data[0]
	}
	//--------------------------------------------------
	return jsonrpc.RPC_encode_jsonrpc_error_response(errorMap)
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_encode_jsonrpc_invalid_params_response
//--------------------------------------------------------------------------------

func (jsonrpc *RPC_jsonrpcStruct) RPC_encode_jsonrpc_invalid_params_response(Data ...any) (string, error) {
	//--------------------------------------------------
	errorMap := map[string]any{"code": -32602, "message": "invalid params"}
	//--------------------------------------------------
	if Data != nil && Data[0] != nil {

		errorMap["data"] = Data[0]
	}
	//--------------------------------------------------
	return jsonrpc.RPC_encode_jsonrpc_error_response(errorMap)
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_encode_jsonrpc_internal_error_response
//--------------------------------------------------------------------------------

func (jsonrpc *RPC_jsonrpcStruct) RPC_encode_jsonrpc_internal_error_response(Data ...any) (string, error) {
	//--------------------------------------------------
	errorMap := map[string]any{"code": -32603, "message": "internal error"}
	//--------------------------------------------------
	if Data != nil && Data[0] != nil {

		errorMap["data"] = Data[0]
	}
	//--------------------------------------------------
	return jsonrpc.RPC_encode_jsonrpc_error_response(errorMap)
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_encode_jsonrpc_parse_error_response
//--------------------------------------------------------------------------------

func (jsonrpc *RPC_jsonrpcStruct) RPC_encode_jsonrpc_parse_error_response(Data ...any) (string, error) {
	//--------------------------------------------------
	errorMap := map[string]any{"code": -32700, "message": "parse error"}
	//--------------------------------------------------
	if Data != nil && Data[0] != nil {

		errorMap["data"] = Data[0]
	}
	//--------------------------------------------------
	return jsonrpc.RPC_encode_jsonrpc_error_response(errorMap)
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------

//--------------------------------------------------------------------------------
// RPC_decode_json
//--------------------------------------------------------------------------------

func RPC_decode_json(jsonString string) (map[string]any, error) {
	//--------------------------------------------------
	var err error
	var jsonInterface any
	//--------------------------------------------------
	jsonMap := map[string]any{}
	//--------------------------------------------------
	jsonInterface, err = conv.JSON_decode(jsonString)
	//--------------------------------------------------
	if err == nil {

		if isObject(jsonInterface) {

			jsonMap = jsonInterface.(map[string]any)

		} else {

			err = errors.New("parse error")
		}
	}
	//--------------------------------------------------
	return jsonMap, err
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------

//--------------------------------------------------------------------------------
// RPC_decode_jsonrpc_request
//--------------------------------------------------------------------------------

func RPC_decode_jsonrpc_request(requestString string) (map[string]any, error) {
	//--------------------------------------------------
	var err error
	var jsonInterface any
	//--------------------------------------------------
	requestMap := map[string]any{}
	//----------
	jsonInterface, err = RPC_decode_json(requestString)
	//----------
	if err == nil {

		if isObject(jsonInterface) {

			requestMap = jsonInterface.(map[string]any)

		} else {

			err = errors.New("parse error")
		}
	}
	//--------------------------------------------------
	return requestMap, err
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------

//--------------------------------------------------------------------------------
// RPC_decode_jsonrpc_response
//--------------------------------------------------------------------------------

func RPC_decode_jsonrpc_response(responseString string) (map[string]any, error) {
	//--------------------------------------------------
	var err error
	var jsonInterface any
	//--------------------------------------------------
	requestMap := map[string]any{}
	//--------------------------------------------------
	jsonInterface, err = RPC_decode_json(responseString)
	//--------------------------------------------------
	if err == nil {

		if isObject(jsonInterface) {

			requestMap = jsonInterface.(map[string]any)

		} else {

			err = errors.New("parse error")
		}
	}
	//--------------------------------------------------
	return requestMap, err
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------

//--------------------------------------------------------------------------------
// isString
//--------------------------------------------------------------------------------

func isString(value interface{}) bool {
	return fmt.Sprintf("%T", value) == "string"
}

//--------------------------------------------------------------------------------
// isNumber
//--------------------------------------------------------------------------------

func isNumber(value interface{}) bool {
	switch value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true
	case float32, float64, complex64, complex128:
		return true
	}
	return false
}

//--------------------------------------------------------------------------------
// isArray
//--------------------------------------------------------------------------------

func isArray(value interface{}) bool {
	return fmt.Sprintf("%T", value) == "[]interface {}"
}

//--------------------------------------------------------------------------------
// isObject
//--------------------------------------------------------------------------------

func isObject(value interface{}) bool {
	return fmt.Sprintf("%T", value) == "map[string]interface {}"
}

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------
