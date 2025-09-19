/*

Copyright 2023-2024, Tim Brockley. All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

*/

package rpc

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/timbrockley/golang-main/conv"
)

//--------------------------------------------------------------------------------
// global variables / structs
//--------------------------------------------------------------------------------

const ContentTypeText string = "text/plain; charset=UTF-8"

const ContentTypeJSON string = "application/json; charset=UTF-8"

//--------------------------------------------------------------------------------

type RPCStruct struct {
	//--------------------
	ResponseWriter http.ResponseWriter
	HttpRequest    *http.Request
	//--------------------
	RequestURL    string
	RequestString string
	RequestMap    map[string]any
	//--------------------
	RemoteIPAddr string
	//--------------------
	ResponseURL        string
	ResponseHeadersMap map[string]string
	//--------------------
	Encoding string
	//--------------------
}

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------

//--------------------------------------------------------------------------------
// RPC_read_request
//--------------------------------------------------------------------------------

func (rpcInstance *RPCStruct) RPC_read_request() error {
	//--------------------------------------------------
	var err error
	var jsonInterface any
	//--------------------------------------------------
	var requestBytes []byte
	var requestString string
	//--------------------------------------------------
	rpcInstance.RequestString = ""
	rpcInstance.RequestMap = map[string]any{}
	//--------------------------------------------------
	rpcInstance.RemoteIPAddr = GetRemoteIPAddr(rpcInstance.HttpRequest)
	//--------------------------------------------------
	requestBytes, err = io.ReadAll(rpcInstance.HttpRequest.Body)
	//--------------------------------------------------
	if err == nil {
		//--------------------------------------------------
		// replace the body with a new reader incase re-read by another function
		//
		rpcInstance.HttpRequest.Body = io.NopCloser(bytes.NewBuffer(requestBytes))
		//
		//--------------------------------------------------
		rpcInstance.RequestURL = rpcInstance.HttpRequest.URL.String()
		//--------------------------------------------------
		requestString = string(requestBytes)
		//--------------------
		requestString, err = DecodeData(requestString, rpcInstance.Encoding)
		//--------------------
		if err == nil {
			//--------------------
			rpcInstance.RequestString = requestString
			//--------------------
			// if request looks like it might be single json try decode but don't report any errors
			match, _ := regexp.MatchString(`^\s*{.*}\s*$`, requestString)
			if match {
				//--------------------
				jsonInterface, err = RPC_decode_json(rpcInstance.RequestString)
				//--------------------
				if err == nil && isObject(jsonInterface) {
					rpcInstance.RequestMap = jsonInterface.(map[string]any)
				} else {
					err = nil // ignore error
				}
				//--------------------
			}
			//--------------------
		}
		//--------------------------------------------------
	}
	//--------------------------------------------------
	return err
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_send_request
//--------------------------------------------------------------------------------

func (rpcInstance *RPCStruct) RPC_send_request(requestString string) (string, error) {
	//--------------------------------------------------
	var err error
	//--------------------------------------------------
	var httpRequest *http.Request
	var httpResponse *http.Response
	var responseBytes []byte
	//--------------------------------------------------
	responseString := ""
	//--------------------------------------------------
	requestString, err = EncodeData(requestString, rpcInstance.Encoding)
	//--------------------
	if err == nil {
		//--------------------
		httpRequest, err = http.NewRequest("POST", rpcInstance.ResponseURL, bytes.NewBuffer([]byte(requestString)))
		//--------------------
		if err == nil {
			//--------------------
			if rpcInstance.ResponseHeadersMap["Content-Type"] == "" {
				rpcInstance.ResponseHeadersMap["Content-Type"] = ContentTypeText
			}
			//--------------------
			for headerKey, headerValue := range rpcInstance.ResponseHeadersMap {
				httpRequest.Header.Add(headerKey, headerValue)
			}
			//--------------------
			httpResponse, err = http.DefaultClient.Do(httpRequest)
			//--------------------
			if err == nil {
				//--------------------
				defer httpResponse.Body.Close()
				//--------------------
				responseBytes, err = io.ReadAll(httpResponse.Body)
				//--------------------
				if err == nil {
					//--------------------
					responseString = string(responseBytes)
					//--------------------
					responseString, err = DecodeData(responseString, rpcInstance.Encoding)
					//--------------------
				}
				//--------------------
			}
			//--------------------
		}
		//--------------------
	}
	//--------------------------------------------------
	return responseString, err
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_send_json_request
//--------------------------------------------------------------------------------

func (rpcInstance *RPCStruct) RPC_send_json_request(requestMap map[string]any) (any, error) {
	//--------------------------------------------------
	var err error
	//--------------------------------------------------
	var requestString string
	//--------------------------------------------------
	var responseString string
	var responseMap any
	//--------------------------------------------------
	requestString, err = RPC_encode_json(requestMap)
	//--------------------
	if err == nil {
		//--------------------
		if rpcInstance.ResponseHeadersMap["Content-Type"] == "" {
			rpcInstance.ResponseHeadersMap["Content-Type"] = ContentTypeJSON
		}
		//--------------------
		responseString, err = rpcInstance.RPC_send_request(requestString)
		//--------------------
		if err == nil {
			responseMap, err = RPC_decode_json(responseString)
		}
		//--------------------
	}
	//--------------------------------------------------
	return responseMap, err
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_read_json_request
//--------------------------------------------------------------------------------

func (rpcInstance *RPCStruct) RPC_read_json_request() error {
	//--------------------------------------------------
	var err error
	var jsonInterface any
	//--------------------------------------------------
	err = rpcInstance.RPC_read_request()
	//--------------------
	if err == nil {
		//--------------------
		// match what looks like a single json request
		match, _ := regexp.MatchString(`^\s*{.*}\s*$`, rpcInstance.RequestString)
		if !match {
			err = errors.New("invalid request")
		} else {
			//--------------------
			jsonInterface, err = RPC_decode_json(rpcInstance.RequestString)
			//--------------------
			if err == nil {
				if isObject(jsonInterface) {
					rpcInstance.RequestMap = jsonInterface.(map[string]any)
				} else {
					err = errors.New("parse error")
				}
			}
			//--------------------
		}
		//--------------------
	}
	//--------------------------------------------------
	return err
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_read_jsonrpc_request
//--------------------------------------------------------------------------------

func (rpcInstance *RPCStruct) RPC_read_jsonrpc_request() error {
	//--------------------------------------------------
	var err error
	var jsonInterface any
	//--------------------------------------------------
	err = rpcInstance.RPC_read_request()
	//--------------------
	if err == nil {
		//--------------------
		jsonInterface, err = RPC_decode_json(rpcInstance.RequestString)
		//--------------------
		if err == nil {
			if isObject(jsonInterface) {
				rpcInstance.RequestMap = jsonInterface.(map[string]any)
			} else {
				err = errors.New("parse error")
			}
		}
		//--------------------
	}
	//--------------------------------------------------
	return err
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_send_jsonrpc_request
//--------------------------------------------------------------------------------

func (rpcInstance *RPCStruct) RPC_send_jsonrpc_request(requestMap map[string]any) (any, error) {
	//--------------------------------------------------
	var err error
	//--------------------------------------------------
	var requestString string
	//--------------------------------------------------
	var responseString string
	var responseMap any
	//--------------------------------------------------
	if !isString(requestMap["method"]) {
		return nil, fmt.Errorf("method is not defined or is not a string")
	}
	//--------------------------------------------------
	if requestMap["params"] != nil && !isArray(requestMap["params"]) && !isObject(requestMap["params"]) {
		return nil, fmt.Errorf("params is not an array or an object")
	}
	//--------------------------------------------------
	if requestMap["id"] != nil && !isNumber(requestMap["id"]) && !isString(requestMap["id"]) {
		return nil, fmt.Errorf("id is not a number or a string")
	}
	//--------------------------------------------------
	requestMap["jsonrpc"] = "2.0"
	//--------------------
	if requestMap["id"] == nil {
		requestMap["id"] = auto.ID()
	}
	//--------------------------------------------------
	requestString, err = RPC_encode_json(requestMap)
	//--------------------
	if err == nil {
		//--------------------
		if rpcInstance.ResponseHeadersMap["Content-Type"] == "" {
			rpcInstance.ResponseHeadersMap["Content-Type"] = ContentTypeJSON
		}
		//--------------------
		responseString, err = rpcInstance.RPC_send_request(requestString)
		//--------------------
		if err == nil {
			responseMap, err = RPC_decode_json(responseString)
		}
		//--------------------
	}
	//--------------------------------------------------
	return responseMap, err
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_send_response
//--------------------------------------------------------------------------------

func (rpcInstance *RPCStruct) RPC_send_response(responseString string) {
	//--------------------------------------------------
	var err error
	//--------------------------------------------------
	if rpcInstance.ResponseHeadersMap["Content-Type"] == "" {
		rpcInstance.ResponseHeadersMap["Content-Type"] = ContentTypeText
	}
	//--------------------
	for headerKey, headerValue := range rpcInstance.ResponseHeadersMap {
		rpcInstance.ResponseWriter.Header().Set(headerKey, headerValue)
	}
	//--------------------
	responseString, err = EncodeData(responseString, rpcInstance.Encoding)
	//--------------------
	if err != nil {
		//--------------------------------------------------
		fmt.Fprintf(rpcInstance.ResponseWriter, `{"error":%q}`, err)
		//--------------------------------------------------
	} else {
		//--------------------------------------------------
		fmt.Fprint(rpcInstance.ResponseWriter, responseString)
		//--------------------------------------------------
	}
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_send_json_response
//--------------------------------------------------------------------------------

func (rpcInstance *RPCStruct) RPC_send_json_response(responseMap map[string]any) {
	//--------------------------------------------------
	var err error
	//--------------------------------------------------
	var responseString string
	//--------------------------------------------------
	if rpcInstance.ResponseHeadersMap["Content-Type"] == "" {
		rpcInstance.ResponseHeadersMap["Content-Type"] = ContentTypeJSON
	}
	//--------------------------------------------------
	responseString, err = RPC_encode_json(responseMap)
	//--------------------------------------------------
	if err != nil {
		//--------------------------------------------------
		rpcInstance.RPC_send_response(fmt.Sprintf(`{"error":%q}`, err))
		//--------------------------------------------------
	} else {
		//--------------------------------------------------
		rpcInstance.RPC_send_response(responseString)
		//--------------------------------------------------
	}
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_send_result_response
//--------------------------------------------------------------------------------

func (rpcInstance *RPCStruct) RPC_send_result_response(result any) {
	//--------------------------------------------------
	responseMap := map[string]any{"result": result}
	//--------------------
	rpcInstance.RPC_send_json_response(responseMap)
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_send_error_response
//--------------------------------------------------------------------------------

func (rpcInstance *RPCStruct) RPC_send_error_response(errorString string) {
	//--------------------------------------------------
	responseMap := map[string]any{"error": errorString}
	//--------------------
	rpcInstance.RPC_send_json_response(responseMap)
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_send_jsonrpc_result
//--------------------------------------------------------------------------------

func (rpcInstance *RPCStruct) RPC_send_jsonrpc_result(result any) error {
	//--------------------------------------------------
	if !isObject(rpcInstance.RequestMap) {
		rpcInstance.RequestMap = map[string]any{}
	}
	//--------------------------------------------------
	responseMap := map[string]any{"jsonrpc": "2.0", "id": nil}
	//--------------------
	if isNumber(rpcInstance.RequestMap["id"]) || isString(rpcInstance.RequestMap["id"]) {
		responseMap["id"] = rpcInstance.RequestMap["id"]
	}
	//--------------------
	responseMap["result"] = result
	//--------------------------------------------------
	responseString, err := RPC_encode_json(responseMap)
	//--------------------------------------------------
	if err == nil {
		rpcInstance.RPC_send_response(responseString)
	}
	//--------------------------------------------------
	return err
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_send_jsonrpc_error
//--------------------------------------------------------------------------------

func (rpcInstance *RPCStruct) RPC_send_jsonrpc_error(error any) error {
	//--------------------------------------------------
	if !isObject(rpcInstance.RequestMap) {
		rpcInstance.RequestMap = map[string]any{}
	}
	//--------------------------------------------------
	responseMap := map[string]any{"jsonrpc": "2.0", "id": nil}
	//--------------------
	if isNumber(rpcInstance.RequestMap["id"]) || isString(rpcInstance.RequestMap["id"]) {
		responseMap["id"] = rpcInstance.RequestMap["id"]
	}
	//--------------------
	responseMap["error"] = error
	//--------------------------------------------------
	responseString, err := RPC_encode_json(responseMap)
	//--------------------------------------------------
	if err == nil {
		rpcInstance.RPC_send_response(responseString)
	}
	//--------------------------------------------------
	return err
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_send_jsonrpc_server_error
//--------------------------------------------------------------------------------

func (rpcInstance *RPCStruct) RPC_send_jsonrpc_server_error(Data ...any) error {
	//--------------------------------------------------
	errorMap := map[string]any{"code": -32000, "message": "server error"}
	//--------------------
	if Data != nil && Data[0] != nil {
		errorMap["data"] = Data[0]
	}
	//--------------------------------------------------
	return rpcInstance.RPC_send_jsonrpc_error(errorMap)
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_send_jsonrpc_invalid_request
//--------------------------------------------------------------------------------

func (rpcInstance *RPCStruct) RPC_send_jsonrpc_invalid_request(Data ...any) error {
	//--------------------------------------------------
	errorMap := map[string]any{"code": -32600, "message": "invalid request"}
	//--------------------
	if Data != nil && Data[0] != nil {
		errorMap["data"] = Data[0]
	}
	//--------------------------------------------------
	return rpcInstance.RPC_send_jsonrpc_error(errorMap)
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_send_jsonrpc_method_not_found
//--------------------------------------------------------------------------------

func (rpcInstance *RPCStruct) RPC_send_jsonrpc_method_not_found(Data ...any) error {
	//--------------------------------------------------
	errorMap := map[string]any{"code": -32601, "message": "method not found"}
	//--------------------
	if Data != nil && Data[0] != nil {
		errorMap["data"] = Data[0]
	}
	//--------------------------------------------------
	return rpcInstance.RPC_send_jsonrpc_error(errorMap)
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_send_jsonrpc_invalid_params
//--------------------------------------------------------------------------------

func (rpcInstance *RPCStruct) RPC_send_jsonrpc_invalid_params(Data ...any) error {
	//--------------------------------------------------
	errorMap := map[string]any{"code": -32602, "message": "invalid params"}
	//--------------------
	if Data != nil && Data[0] != nil {
		errorMap["data"] = Data[0]
	}
	//--------------------------------------------------
	return rpcInstance.RPC_send_jsonrpc_error(errorMap)
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_send_jsonrpc_internal_error
//--------------------------------------------------------------------------------

func (rpcInstance *RPCStruct) RPC_send_jsonrpc_internal_error(Data ...any) error {
	//--------------------------------------------------
	errorMap := map[string]any{"code": -32603, "message": "internal error"}
	//--------------------
	if Data != nil && Data[0] != nil {
		errorMap["data"] = Data[0]
	}
	//--------------------------------------------------
	return rpcInstance.RPC_send_jsonrpc_error(errorMap)
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_send_jsonrpc_parse_error
//--------------------------------------------------------------------------------

func (rpcInstance *RPCStruct) RPC_send_jsonrpc_parse_error(Data ...any) error {
	//--------------------------------------------------
	errorMap := map[string]any{"code": -32700, "message": "parse error"}
	//--------------------
	if Data != nil && Data[0] != nil {
		errorMap["data"] = Data[0]
	}
	//--------------------------------------------------
	return rpcInstance.RPC_send_jsonrpc_error(errorMap)
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------

//--------------------------------------------------------------------------------
// RPC_echo
//--------------------------------------------------------------------------------

func (rpcInstance *RPCStruct) RPC_echo() {
	//--------------------------------------------------
	_ = rpcInstance.RPC_read_request() // read here anyway incase not already read
	//--------------------------------------------------
	contentType := rpcInstance.HttpRequest.Header.Get("Content-Type")
	if contentType != "" {
		rpcInstance.ResponseWriter.Header().Set("Content-Type", contentType)
	}
	//--------------------------------------------------
	REGEXP := regexp.MustCompile(`(?i)^X`)
	for name, values := range rpcInstance.HttpRequest.Header {
		//--------------------
		if REGEXP.FindString(name) != "" {
			//--------------------
			rpcInstance.ResponseWriter.Header().Set(name, strings.Join(values, ", "))
			//--------------------
			if strings.EqualFold(name, "X-Debug") && strings.EqualFold(values[0], "true") {
				//--------------------
				fmt.Println("\n" + strings.Repeat("-", 40))
				fmt.Println(rpcInstance.HttpRequest.Method, rpcInstance.HttpRequest.URL.String(), rpcInstance.HttpRequest.Proto)
				fmt.Println(strings.Repeat("-", 40))
				fmt.Println(rpcInstance.RequestString)
				fmt.Println(strings.Repeat("-", 40) + "\n")
				//--------------------
			}
			//--------------------
		}
		//--------------------
	}
	//--------------------------------------------------
	rpcInstance.RPC_send_response(rpcInstance.RequestString)
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// JSONRPC_echo
//--------------------------------------------------------------------------------

func (rpcInstance *RPCStruct) JSONRPC_echo() {
	//--------------------------------------------------
	_ = rpcInstance.RPC_read_request() // read here anyway incase not already read
	//--------------------------------------------------
	contentType := rpcInstance.HttpRequest.Header.Get("Content-Type")
	if contentType != "" {
		rpcInstance.ResponseWriter.Header().Set("Content-Type", contentType)
	}
	//--------------------------------------------------
	REGEXP := regexp.MustCompile(`(?i)^X`)
	for name, values := range rpcInstance.HttpRequest.Header {
		//--------------------
		if REGEXP.FindString(name) != "" {
			//--------------------
			rpcInstance.ResponseWriter.Header().Set(name, strings.Join(values, ", "))
			//--------------------
			if strings.EqualFold(name, "X-Debug") && strings.EqualFold(values[0], "true") {
				//--------------------
				fmt.Println("\n" + strings.Repeat("-", 40))
				fmt.Println(rpcInstance.HttpRequest.Method, rpcInstance.HttpRequest.URL.String(), rpcInstance.HttpRequest.Proto)
				fmt.Println(strings.Repeat("-", 40))
				fmt.Println(rpcInstance.RequestString)
				fmt.Println(strings.Repeat("-", 40) + "\n")
				//--------------------
			}
			//--------------------
		}
		//--------------------
	}
	//--------------------------------------------------
	rpcInstance.RPC_send_jsonrpc_result(rpcInstance.RequestString)
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------

//--------------------------------------------------------------------------------
// RPC_Handler
//--------------------------------------------------------------------------------

func RPC_Handler(responseWriter http.ResponseWriter, httpRequest *http.Request) {
	//--------------------------------------------------
	var err error
	//--------------------------------------------------
	var rpcInstance RPCStruct
	//--------------------
	rpcInstance.ResponseWriter = responseWriter
	rpcInstance.HttpRequest = httpRequest
	//--------------------
	rpcInstance.ResponseURL = rpcInstance.HttpRequest.URL.String()
	//--------------------
	rpcInstance.ResponseHeadersMap = map[string]string{"Content-Type": ContentTypeJSON}
	//--------------------
	err = rpcInstance.RPC_read_json_request()
	//--------------------------------------------------
	if err != nil {
		//--------------------
		rpcInstance.RPC_send_error_response(fmt.Sprint(err))
		//--------------------
	} else {
		//--------------------------------------------------
		method, exists := rpcInstance.RequestMap["method"]
		if !exists {
			//--------------------
			rpcInstance.RPC_send_error_response("method is not defined")
			//--------------------
		} else {
			//--------------------------------------------------
			switch method {

			case "echo":
				rpcInstance.RPC_echo()

			default:
				rpcInstance.RPC_send_error_response("method not found")
			}
			//--------------------------------------------------
		}
		//--------------------------------------------------
	}
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// JSONRPC_Handler
//--------------------------------------------------------------------------------

func JSONRPC_Handler(responseWriter http.ResponseWriter, httpRequest *http.Request) {
	//--------------------------------------------------
	var err error
	//--------------------------------------------------
	var rpcInstance RPCStruct
	//--------------------
	rpcInstance.ResponseWriter = responseWriter
	rpcInstance.HttpRequest = httpRequest
	//--------------------
	rpcInstance.ResponseURL = rpcInstance.HttpRequest.URL.String()
	//--------------------
	contentType := httpRequest.Header.Get("Content-Type")
	if contentType == "" {
		contentType = ContentTypeJSON
	}
	//--------------------
	rpcInstance.ResponseHeadersMap = map[string]string{"Content-Type": contentType}
	//--------------------
	err = rpcInstance.RPC_read_jsonrpc_request()
	//--------------------------------------------------
	if err != nil {
		//--------------------
		rpcInstance.RPC_send_jsonrpc_parse_error(fmt.Sprint(err))
		//--------------------
	} else {
		//--------------------------------------------------
		method, exists := rpcInstance.RequestMap["method"]
		if !exists {
			//--------------------
			rpcInstance.RPC_send_jsonrpc_method_not_found(rpcInstance.RequestMap["method"])
			//--------------------
		} else if rpcInstance.RequestMap["params"] != nil && !isArray(rpcInstance.RequestMap["params"]) && !isObject(rpcInstance.RequestMap["params"]) {
			//--------------------
			rpcInstance.RPC_send_jsonrpc_invalid_params(rpcInstance.RequestMap["params"])
			//--------------------
		} else {
			//--------------------------------------------------
			switch method {

			case "echo":
				rpcInstance.JSONRPC_echo()

			default:
				rpcInstance.RPC_send_jsonrpc_method_not_found(rpcInstance.RequestMap["method"])
			}
			//--------------------------------------------------
		}
		//--------------------------------------------------
	}
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------

//--------------------------------------------------------------------------------
// GetRemoteIPAddr
//--------------------------------------------------------------------------------

func GetRemoteIPAddr(httpRequest *http.Request) string {
	//--------------------------------------------------
	var ipAddr string
	var ips []string
	//--------------------------------------------------
	ipAddr = httpRequest.RemoteAddr
	//--------------------------------------------------
	bracketIndex := strings.Index(ipAddr, "[")
	if bracketIndex >= 0 {
		ipAddr = ""
	}
	//--------------------------------------------------
	if ipAddr == "" {
		ipAddr = httpRequest.Header.Get("X-Real-Ip")
	}
	//--------------------------------------------------
	if ipAddr == "" {
		ipAddr = httpRequest.Header.Get("X-Forwarded-For")
	}
	//--------------------------------------------------
	ips = strings.Split(ipAddr, ",")
	//--------------------
	if len(ips) >= 1 {
		//--------------------
		ipAddr = ips[0]
		//--------------------
		ipAddr = strings.TrimSpace(ipAddr)
		//--------------------
		colonIndex := strings.Index(ipAddr, ":")
		if colonIndex >= 0 {
			ipAddr = ipAddr[:colonIndex]
		}
		//--------------------
	} else {
		//--------------------
		ipAddr = ""
		//--------------------
	}
	//--------------------
	return ipAddr
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------

//--------------------------------------------------------------------------------
// EncodeData
//--------------------------------------------------------------------------------

func EncodeData(dataString string, encoding string) (string, error) {
	//--------------------------------------------------
	if dataString == "" || encoding == "" {
		return dataString, nil
	}
	//--------------------------------------------------
	switch encoding {
	case "obfuscate":
		return ObfuscateData(dataString, true, false)
	case "base64":
		dataString = conv.Base64_encode(dataString)
	case "base64url":
		dataString = conv.Base64url_encode(dataString)
	default:
		return "", errors.New("invalid encoding")
	}
	//--------------------------------------------------
	return dataString, nil
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// DecodeData
//--------------------------------------------------------------------------------

func DecodeData(dataString string, encoding string) (string, error) {
	//--------------------------------------------------
	if dataString == "" || encoding == "" {
		return dataString, nil
	}
	//--------------------------------------------------
	switch encoding {
	case "obfuscate":
		return ObfuscateData(dataString, false, true)
	case "base64":
		return conv.Base64_decode(dataString)
	case "base64url":
		return conv.Base64url_decode(dataString)
	default:
		return "", errors.New("invalid encoding")
	}
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// ObfuscateData
//--------------------------------------------------------------------------------

func ObfuscateData(dataString string, base64Encode bool, base64Decode bool, Value ...byte) (string, error) {
	//--------------------------------------------------
	if dataString == "" {
		return dataString, nil
	}
	//--------------------------------------------------
	var value byte
	//--------------------------------------------------
	if len(Value) > 0 {
		//--------------------------------------------------
		value = Value[0]
		//--------------------------------------------------
		if value == 0 {
			return "", errors.New("value should be an integer between 1 and 255")
		}
		//--------------------------------------------------
	} else {
		//--------------------------------------------------
		value = 0b10101010
		//--------------------------------------------------
	}
	//--------------------------------------------------
	var err error
	var dataBytes []byte
	//--------------------------------------------------
	if base64Decode {
		dataBytes, err = base64.StdEncoding.DecodeString(dataString)
		if err != nil {
			return "", err
		}
	} else {
		dataBytes = []byte(dataString)
	}
	//--------------------------------------------------
	for i := range dataBytes {
		dataBytes[i] ^= value
	}
	//--------------------------------------------------
	if base64Encode {
		return base64.StdEncoding.EncodeToString(dataBytes), nil
	} else {
		return string(dataBytes), nil
	}
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------
