/*

Copyright 2023-2024, Tim Brockley. All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

*/

package rpc

import (
	"bytes"
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

func (rpcObject *RPCStruct) RPC_read_request() error {

	//--------------------------------------------------
	var err error
	var jsonInterface any
	//--------------------------------------------------
	var requestBytes []byte
	var requestString string
	//--------------------------------------------------
	rpcObject.RequestString = ""
	rpcObject.RequestMap = map[string]any{}
	//--------------------------------------------------
	rpcObject.RemoteIPAddr = GetRemoteIPAddr(rpcObject.HttpRequest)
	//--------------------------------------------------
	requestBytes, err = io.ReadAll(rpcObject.HttpRequest.Body)
	//--------------------------------------------------
	if err == nil {

		//--------------------------------------------------
		// replace the body with a new reader incase re-read by another function
		//
		rpcObject.HttpRequest.Body = io.NopCloser(bytes.NewBuffer(requestBytes))
		//
		//--------------------------------------------------
		rpcObject.RequestURL = rpcObject.HttpRequest.URL.String()
		//--------------------------------------------------
		requestString = string(requestBytes)
		//--------------------
		if rpcObject.Encoding == "base64" {
			requestString, err = conv.Base64_decode(requestString)
		} else if rpcObject.Encoding == "base64url" {
			requestString, err = conv.Base64url_decode(requestString)
		}
		//--------------------
		if err == nil {

			//--------------------
			rpcObject.RequestString = requestString
			//--------------------
			// if request looks like it might be single json try decode but don't report any errors
			match, _ := regexp.MatchString(`^\s*{.*}\s*$`, requestString)
			if match {

				//--------------------
				jsonInterface, err = RPC_decode_json(rpcObject.RequestString)
				//--------------------
				if err == nil && isObject(jsonInterface) {

					rpcObject.RequestMap = jsonInterface.(map[string]any)

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

func (rpcObject *RPCStruct) RPC_send_request(requestString string) (string, error) {

	//--------------------------------------------------
	var err error
	//--------------------------------------------------
	var httpRequest *http.Request
	var httpResponse *http.Response
	var responseBytes []byte
	//--------------------------------------------------
	responseString := ""
	//--------------------------------------------------
	if rpcObject.Encoding == "base64" {
		requestString = conv.Base64_encode(requestString)
	} else if rpcObject.Encoding == "base64url" {
		requestString = conv.Base64url_encode(requestString)
	}
	//--------------------
	httpRequest, err = http.NewRequest("POST", rpcObject.ResponseURL, bytes.NewBuffer([]byte(requestString)))
	//--------------------
	if err == nil {

		//--------------------
		if rpcObject.ResponseHeadersMap["Content-Type"] == "" {

			rpcObject.ResponseHeadersMap["Content-Type"] = ContentTypeText
		}
		//--------------------
		for headerKey, headerValue := range rpcObject.ResponseHeadersMap {

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
				if rpcObject.Encoding == "base64" {
					responseString, err = conv.Base64_decode(responseString)
				} else if rpcObject.Encoding == "base64url" {
					responseString, err = conv.Base64url_decode(responseString)
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

func (rpcObject *RPCStruct) RPC_send_json_request(requestMap map[string]any) (any, error) {

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
		if rpcObject.ResponseHeadersMap["Content-Type"] == "" {

			rpcObject.ResponseHeadersMap["Content-Type"] = ContentTypeJSON
		}
		//--------------------
		responseString, err = rpcObject.RPC_send_request(requestString)
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

func (rpcObject *RPCStruct) RPC_read_json_request() error {

	//--------------------------------------------------
	var err error
	var jsonInterface any
	//--------------------------------------------------
	err = rpcObject.RPC_read_request()
	//--------------------
	if err == nil {

		//--------------------
		// match what looks like a single json request
		match, _ := regexp.MatchString(`^\s*{.*}\s*$`, rpcObject.RequestString)
		if !match {

			err = errors.New("invalid request")

		} else {

			//--------------------
			jsonInterface, err = RPC_decode_json(rpcObject.RequestString)
			//--------------------
			if err == nil {

				if isObject(jsonInterface) {

					rpcObject.RequestMap = jsonInterface.(map[string]any)

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

func (rpcObject *RPCStruct) RPC_read_jsonrpc_request() error {

	//--------------------------------------------------
	var err error
	var jsonInterface any
	//--------------------------------------------------
	err = rpcObject.RPC_read_request()
	//--------------------
	if err == nil {

		//--------------------
		jsonInterface, err = RPC_decode_json(rpcObject.RequestString)
		//--------------------
		if err == nil {

			if isObject(jsonInterface) {

				rpcObject.RequestMap = jsonInterface.(map[string]any)

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

func (rpcObject *RPCStruct) RPC_send_jsonrpc_request(requestMap map[string]any) (any, error) {

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
		if rpcObject.ResponseHeadersMap["Content-Type"] == "" {

			rpcObject.ResponseHeadersMap["Content-Type"] = ContentTypeJSON
		}
		//--------------------
		responseString, err = rpcObject.RPC_send_request(requestString)
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

func (rpcObject *RPCStruct) RPC_send_response(responseString string) {

	//--------------------------------------------------
	if rpcObject.ResponseHeadersMap["Content-Type"] == "" {

		rpcObject.ResponseHeadersMap["Content-Type"] = ContentTypeText
	}
	//--------------------
	for headerKey, headerValue := range rpcObject.ResponseHeadersMap {

		rpcObject.ResponseWriter.Header().Set(headerKey, headerValue)
	}
	//--------------------
	if rpcObject.Encoding == "base64" {
		responseString = conv.Base64_encode(responseString)
	} else if rpcObject.Encoding == "base64url" {
		responseString = conv.Base64url_encode(responseString)
	}
	//--------------------
	fmt.Fprint(rpcObject.ResponseWriter, responseString)
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_send_json_response
//--------------------------------------------------------------------------------

func (rpcObject *RPCStruct) RPC_send_json_response(responseMap map[string]any) {

	//--------------------------------------------------
	var err error
	//--------------------------------------------------
	var responseString string
	//--------------------------------------------------
	if rpcObject.ResponseHeadersMap["Content-Type"] == "" {

		rpcObject.ResponseHeadersMap["Content-Type"] = ContentTypeJSON
	}
	//--------------------------------------------------
	responseString, err = RPC_encode_json(responseMap)
	if err != nil {

		//--------------------------------------------------
		rpcObject.RPC_send_response(fmt.Sprintf(`{"error":%q}`, err))
		//--------------------------------------------------

	} else {

		//--------------------------------------------------
		rpcObject.RPC_send_response(responseString)
		//--------------------------------------------------
	}
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_send_result_response
//--------------------------------------------------------------------------------

func (rpcObject *RPCStruct) RPC_send_result_response(result any) {

	//--------------------------------------------------
	responseMap := map[string]any{"result": result}
	//--------------------
	rpcObject.RPC_send_json_response(responseMap)
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_send_error_response
//--------------------------------------------------------------------------------

func (rpcObject *RPCStruct) RPC_send_error_response(errorString string) {

	//--------------------------------------------------
	responseMap := map[string]any{"error": errorString}
	//--------------------
	rpcObject.RPC_send_json_response(responseMap)
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_send_jsonrpc_result
//--------------------------------------------------------------------------------

func (rpcObject *RPCStruct) RPC_send_jsonrpc_result(result any) error {

	//--------------------------------------------------
	if !isObject(rpcObject.RequestMap) {
		rpcObject.RequestMap = map[string]any{}
	}
	//--------------------------------------------------
	responseMap := map[string]any{"jsonrpc": "2.0", "id": nil}
	//--------------------
	if isNumber(rpcObject.RequestMap["id"]) || isString(rpcObject.RequestMap["id"]) {
		responseMap["id"] = rpcObject.RequestMap["id"]
	}
	//--------------------
	responseMap["result"] = result
	//--------------------------------------------------
	responseString, err := RPC_encode_json(responseMap)
	//--------------------------------------------------
	if err == nil {
		rpcObject.RPC_send_response(responseString)
	}
	//--------------------------------------------------
	return err
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_send_jsonrpc_error
//--------------------------------------------------------------------------------

func (rpcObject *RPCStruct) RPC_send_jsonrpc_error(error any) error {

	//--------------------------------------------------
	if !isObject(rpcObject.RequestMap) {
		rpcObject.RequestMap = map[string]any{}
	}
	//--------------------------------------------------
	responseMap := map[string]any{"jsonrpc": "2.0", "id": nil}
	//--------------------
	if isNumber(rpcObject.RequestMap["id"]) || isString(rpcObject.RequestMap["id"]) {
		responseMap["id"] = rpcObject.RequestMap["id"]
	}
	//--------------------
	responseMap["error"] = error
	//--------------------------------------------------
	responseString, err := RPC_encode_json(responseMap)
	//--------------------------------------------------
	if err == nil {
		rpcObject.RPC_send_response(responseString)
	}
	//--------------------------------------------------
	return err
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_send_jsonrpc_server_error
//--------------------------------------------------------------------------------

func (rpcObject *RPCStruct) RPC_send_jsonrpc_server_error(Data ...any) error {

	//--------------------------------------------------
	errorMap := map[string]any{"code": -32000, "message": "server error"}
	//--------------------
	if Data != nil && Data[0] != nil {

		errorMap["data"] = Data[0]
	}
	//--------------------------------------------------
	return rpcObject.RPC_send_jsonrpc_error(errorMap)
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_send_jsonrpc_invalid_request
//--------------------------------------------------------------------------------

func (rpcObject *RPCStruct) RPC_send_jsonrpc_invalid_request(Data ...any) error {

	//--------------------------------------------------
	errorMap := map[string]any{"code": -32600, "message": "invalid request"}
	//--------------------
	if Data != nil && Data[0] != nil {

		errorMap["data"] = Data[0]
	}
	//--------------------------------------------------
	return rpcObject.RPC_send_jsonrpc_error(errorMap)
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_send_jsonrpc_method_not_found
//--------------------------------------------------------------------------------

func (rpcObject *RPCStruct) RPC_send_jsonrpc_method_not_found(Data ...any) error {

	//--------------------------------------------------
	errorMap := map[string]any{"code": -32601, "message": "method not found"}
	//--------------------
	if Data != nil && Data[0] != nil {

		errorMap["data"] = Data[0]
	}
	//--------------------------------------------------
	return rpcObject.RPC_send_jsonrpc_error(errorMap)
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_send_jsonrpc_invalid_params
//--------------------------------------------------------------------------------

func (rpcObject *RPCStruct) RPC_send_jsonrpc_invalid_params(Data ...any) error {

	//--------------------------------------------------
	errorMap := map[string]any{"code": -32602, "message": "invalid params"}
	//--------------------
	if Data != nil && Data[0] != nil {

		errorMap["data"] = Data[0]
	}
	//--------------------------------------------------
	return rpcObject.RPC_send_jsonrpc_error(errorMap)
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_send_jsonrpc_internal_error
//--------------------------------------------------------------------------------

func (rpcObject *RPCStruct) RPC_send_jsonrpc_internal_error(Data ...any) error {

	//--------------------------------------------------
	errorMap := map[string]any{"code": -32603, "message": "internal error"}
	//--------------------
	if Data != nil && Data[0] != nil {

		errorMap["data"] = Data[0]
	}
	//--------------------------------------------------
	return rpcObject.RPC_send_jsonrpc_error(errorMap)
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// RPC_send_jsonrpc_parse_error
//--------------------------------------------------------------------------------

func (rpcObject *RPCStruct) RPC_send_jsonrpc_parse_error(Data ...any) error {

	//--------------------------------------------------
	errorMap := map[string]any{"code": -32700, "message": "parse error"}
	//--------------------
	if Data != nil && Data[0] != nil {

		errorMap["data"] = Data[0]
	}
	//--------------------------------------------------
	return rpcObject.RPC_send_jsonrpc_error(errorMap)
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------

//--------------------------------------------------------------------------------
// RPC_echo
//--------------------------------------------------------------------------------

func (rpcObject *RPCStruct) RPC_echo() {

	//--------------------------------------------------
	_ = rpcObject.RPC_read_request() // read here anyway incase not already read
	//--------------------------------------------------
	contentType := rpcObject.HttpRequest.Header.Get("Content-Type")
	if contentType != "" {
		rpcObject.ResponseWriter.Header().Set("Content-Type", contentType)
	}
	//--------------------------------------------------
	REGEXP := regexp.MustCompile(`(?i)^X`)
	for name, values := range rpcObject.HttpRequest.Header {
		//--------------------
		if REGEXP.FindString(name) != "" {
			//--------------------
			rpcObject.ResponseWriter.Header().Set(name, strings.Join(values, ", "))
			//--------------------
			if strings.EqualFold(name, "X-Debug") && strings.EqualFold(values[0], "true") {
				//--------------------
				fmt.Println("\n" + strings.Repeat("-", 40))
				fmt.Println(rpcObject.HttpRequest.Method, rpcObject.HttpRequest.URL.String(), rpcObject.HttpRequest.Proto)
				fmt.Println(strings.Repeat("-", 40))
				fmt.Println(rpcObject.RequestString)
				fmt.Println(strings.Repeat("-", 40) + "\n")
				//--------------------
			}
			//--------------------
		}
		//--------------------
	}
	//--------------------------------------------------
	rpcObject.RPC_send_response(rpcObject.RequestString)
	//--------------------------------------------------
}

//--------------------------------------------------------------------------------
// JSONRPC_echo
//--------------------------------------------------------------------------------

func (rpcObject *RPCStruct) JSONRPC_echo() {

	//--------------------------------------------------
	_ = rpcObject.RPC_read_request() // read here anyway incase not already read
	//--------------------------------------------------
	contentType := rpcObject.HttpRequest.Header.Get("Content-Type")
	if contentType != "" {
		rpcObject.ResponseWriter.Header().Set("Content-Type", contentType)
	}
	//--------------------------------------------------
	REGEXP := regexp.MustCompile(`(?i)^X`)
	for name, values := range rpcObject.HttpRequest.Header {
		//--------------------
		if REGEXP.FindString(name) != "" {
			//--------------------
			rpcObject.ResponseWriter.Header().Set(name, strings.Join(values, ", "))
			//--------------------
			if strings.EqualFold(name, "X-Debug") && strings.EqualFold(values[0], "true") {
				//--------------------
				fmt.Println("\n" + strings.Repeat("-", 40))
				fmt.Println(rpcObject.HttpRequest.Method, rpcObject.HttpRequest.URL.String(), rpcObject.HttpRequest.Proto)
				fmt.Println(strings.Repeat("-", 40))
				fmt.Println(rpcObject.RequestString)
				fmt.Println(strings.Repeat("-", 40) + "\n")
				//--------------------
			}
			//--------------------
		}
		//--------------------
	}
	//--------------------------------------------------
	rpcObject.RPC_send_jsonrpc_result(rpcObject.RequestString)
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
	var rpcObject RPCStruct
	//--------------------
	rpcObject.ResponseWriter = responseWriter
	rpcObject.HttpRequest = httpRequest
	//--------------------
	rpcObject.ResponseURL = rpcObject.HttpRequest.URL.String()
	//--------------------
	rpcObject.ResponseHeadersMap = map[string]string{"Content-Type": ContentTypeJSON}
	//--------------------
	err = rpcObject.RPC_read_json_request()
	//--------------------------------------------------
	if err != nil {

		//--------------------
		rpcObject.RPC_send_error_response(fmt.Sprint(err))
		//--------------------

	} else {

		//--------------------------------------------------
		method, exists := rpcObject.RequestMap["method"]
		if !exists {

			//--------------------
			rpcObject.RPC_send_error_response("method is not defined")
			//--------------------

		} else {

			//--------------------------------------------------
			switch method {

			case "echo":
				rpcObject.RPC_echo()

			default:
				rpcObject.RPC_send_error_response("method not found")
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
	var rpcObject RPCStruct
	//--------------------
	rpcObject.ResponseWriter = responseWriter
	rpcObject.HttpRequest = httpRequest
	//--------------------
	rpcObject.ResponseURL = rpcObject.HttpRequest.URL.String()
	//--------------------
	contentType := httpRequest.Header.Get("Content-Type")
	if contentType == "" {
		contentType = ContentTypeJSON
	}
	//--------------------
	rpcObject.ResponseHeadersMap = map[string]string{"Content-Type": contentType}
	//--------------------
	err = rpcObject.RPC_read_jsonrpc_request()
	//--------------------------------------------------
	if err != nil {

		//--------------------
		rpcObject.RPC_send_jsonrpc_parse_error(fmt.Sprint(err))
		//--------------------

	} else {

		//--------------------------------------------------
		method, exists := rpcObject.RequestMap["method"]
		if !exists {

			//--------------------
			rpcObject.RPC_send_jsonrpc_method_not_found(rpcObject.RequestMap["method"])
			//--------------------

		} else if rpcObject.RequestMap["params"] != nil && !isArray(rpcObject.RequestMap["params"]) && !isObject(rpcObject.RequestMap["params"]) {

			//--------------------
			rpcObject.RPC_send_jsonrpc_invalid_params(rpcObject.RequestMap["params"])
			//--------------------

		} else {

			//--------------------------------------------------
			switch method {

			case "echo":
				rpcObject.JSONRPC_echo()

			default:
				rpcObject.RPC_send_jsonrpc_method_not_found(rpcObject.RequestMap["method"])
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
