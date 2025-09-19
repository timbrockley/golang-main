//----------------------------------------------------------------------

package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/timbrockley/golang-main/conv"
)

//----------------------------------------------------------------------
// global variables / structs
//----------------------------------------------------------------------

var rpcInstance RPCStruct

//----------------------------------------------------------------------
//################################################################################
//----------------------------------------------------------------------

func Test_create_rpc_object1(t *testing.T) {
	//--------------------------------------------------
	// load test value to check later in another function
	rpcInstance = RPCStruct{ResponseURL: "new_url"}
	//--------------------------------------------------
}

//----------------------------------------------------------------------

func Test_create_rpc_object2(t *testing.T) {
	//--------------------------------------------------
	if rpcInstance.ResponseURL != "new_url" {
		t.Errorf("rpcInstance.ResponseURL = %q but should = \"new_url\"", rpcInstance.ResponseURL)
	}
	//--------------------------------------------------
}

//----------------------------------------------------------------------
// send request method
//----------------------------------------------------------------------

func TestRPC_send_request_method(t *testing.T) {
	//--------------------------------------------------
	server := httptest.NewServer(http.HandlerFunc(func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		var err error
		var requestBytes []byte
		var requestString string
		var responseString string
		contentType := httpRequest.Header.Get("Content-Type")
		requestBytes, err = io.ReadAll(httpRequest.Body)
		if err != nil {
			responseString = `{"error":"error occurred while trying to read request body"}`
		} else {
			//--------------------
			requestString = string(requestBytes)
			//--------------------
			if rpcInstance.Encoding == "base64" {
				requestString, err = conv.Base64_decode(requestString)
			} else if rpcInstance.Encoding == "base64url" {
				requestString, err = conv.Base64url_decode(requestString)
			}
			//--------------------
			if err != nil {
				responseString = fmt.Sprintf(`{"error":%q}`, err)
			} else {
				responseString = fmt.Sprintf(`{"content_type":%q,"request":%q}`, contentType, requestString)
			}
			//--------------------
			if rpcInstance.Encoding == "base64" {
				responseString = conv.Base64_encode(responseString)
			} else if rpcInstance.Encoding == "base64url" {
				responseString = conv.Base64url_encode(responseString)
			}
			responseWriter.Write([]byte(responseString))
			//--------------------
		}
	}))
	//--------------------
	defer server.Close()
	//--------------------------------------------------
	requestContentType := "text/plain; charset=UTF-8"
	requestHeadersMap := map[string]string{"Content-Type": requestContentType}
	requestString := "<REQUEST_DATA>"
	//--------------------
	EXPECTED_responseString := fmt.Sprintf(`{"content_type":%q,"request":%q}`, requestContentType, requestString)
	//--------------------------------------------------

	//--------------------------------------------------
	rpcInstance.ResponseHeadersMap = requestHeadersMap
	rpcInstance.ResponseURL = server.URL
	//--------------------------------------------------
	responseString, err := rpcInstance.RPC_send_request(requestString)
	//--------------------------------------------------

	//--------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		//--------------------
		if responseString != EXPECTED_responseString {
			t.Errorf("response = %q but should = %q", responseString, EXPECTED_responseString)
		}
		//--------------------
	}
	//--------------------------------------------------
}

//----------------------------------------------------------------------
// send request method - base64
//----------------------------------------------------------------------

func TestRPC_send_request_method_base64(t *testing.T) {
	//--------------------------------------------------
	rpcInstance.Encoding = "base64"
	//--------------------------------------------------
	server := httptest.NewServer(http.HandlerFunc(func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		var err error
		var requestBytes []byte
		var requestString string
		var responseString string
		contentType := httpRequest.Header.Get("Content-Type")
		requestBytes, err = io.ReadAll(httpRequest.Body)
		if err != nil {
			responseString = `{"error":"error occurred while trying to read request body"}`
		} else {
			//--------------------
			requestString = string(requestBytes)
			//--------------------
			if rpcInstance.Encoding == "base64" {
				requestString, err = conv.Base64_decode(requestString)
			} else if rpcInstance.Encoding == "base64url" {
				requestString, err = conv.Base64url_decode(requestString)
			}
			//--------------------
			if err != nil {
				responseString = fmt.Sprintf(`{"error":%q}`, err)
			} else {
				responseString = fmt.Sprintf(`{"content_type":%q,"request":%q}`, contentType, requestString)
			}
			//--------------------
			if rpcInstance.Encoding == "base64" {
				responseString = conv.Base64_encode(responseString)
			} else if rpcInstance.Encoding == "base64url" {
				responseString = conv.Base64url_encode(responseString)
			}
			responseWriter.Write([]byte(responseString))
			//--------------------
		}
	}))
	//--------------------
	defer server.Close()
	//--------------------------------------------------
	requestContentType := "text/plain; charset=UTF-8"
	requestHeadersMap := map[string]string{"Content-Type": requestContentType}
	requestString := "<REQUEST_DATA>"
	//--------------------
	EXPECTED_responseString := fmt.Sprintf(`{"content_type":%q,"request":%q}`, requestContentType, requestString)
	//--------------------------------------------------

	//--------------------------------------------------
	rpcInstance.ResponseHeadersMap = requestHeadersMap
	rpcInstance.ResponseURL = server.URL
	//--------------------------------------------------
	responseString, err := rpcInstance.RPC_send_request(requestString)
	//--------------------------------------------------

	//--------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		//--------------------
		if responseString != EXPECTED_responseString {
			t.Errorf("response = %q but should = %q", responseString, EXPECTED_responseString)
		}
		//--------------------
	}
	//--------------------------------------------------
	rpcInstance.Encoding = ""
	//--------------------------------------------------
}

//----------------------------------------------------------------------
// send request method - base64url
//----------------------------------------------------------------------

func TestRPC_send_request_method_base64url(t *testing.T) {
	//--------------------------------------------------
	rpcInstance.Encoding = "base64url"
	//--------------------------------------------------
	server := httptest.NewServer(http.HandlerFunc(func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		var err error
		var requestBytes []byte
		var requestString string
		var responseString string
		contentType := httpRequest.Header.Get("Content-Type")
		requestBytes, err = io.ReadAll(httpRequest.Body)
		if err != nil {
			responseString = `{"error":"error occurred while trying to read request body"}`
		} else {
			//--------------------
			requestString = string(requestBytes)
			//--------------------
			if rpcInstance.Encoding == "base64" {
				requestString, err = conv.Base64_decode(requestString)
			} else if rpcInstance.Encoding == "base64url" {
				requestString, err = conv.Base64url_decode(requestString)
			}
			//--------------------
			if err != nil {
				responseString = fmt.Sprintf(`{"error":%q}`, err)
			} else {
				responseString = fmt.Sprintf(`{"content_type":%q,"request":%q}`, contentType, requestString)
			}
			//--------------------
			if rpcInstance.Encoding == "base64" {
				responseString = conv.Base64_encode(responseString)
			} else if rpcInstance.Encoding == "base64url" {
				responseString = conv.Base64url_encode(responseString)
			}
			responseWriter.Write([]byte(responseString))
			//--------------------
		}
	}))
	//--------------------
	defer server.Close()
	//--------------------------------------------------
	requestContentType := "text/plain; charset=UTF-8"
	requestHeadersMap := map[string]string{"Content-Type": requestContentType}
	requestString := "<REQUEST_DATA>"
	//--------------------
	EXPECTED_responseString := fmt.Sprintf(`{"content_type":%q,"request":%q}`, requestContentType, requestString)
	//--------------------------------------------------

	//--------------------------------------------------
	rpcInstance.ResponseHeadersMap = requestHeadersMap
	rpcInstance.ResponseURL = server.URL
	//--------------------------------------------------
	responseString, err := rpcInstance.RPC_send_request(requestString)
	//--------------------------------------------------

	//--------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		//--------------------
		if responseString != EXPECTED_responseString {
			t.Errorf("response = %q but should = %q", responseString, EXPECTED_responseString)
		}
		//--------------------
	}
	//--------------------------------------------------
	rpcInstance.Encoding = ""
	//--------------------------------------------------
}

//----------------------------------------------------------------------
// send json request method
//----------------------------------------------------------------------

func TestRPC_send_json_request_method(t *testing.T) {
	//--------------------------------------------------
	server := httptest.NewServer(http.HandlerFunc(func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		requestContentType := httpRequest.Header.Get("Content-Type")
		requestBytes, err := io.ReadAll(httpRequest.Body)
		if err != nil {
			responseWriter.Write([]byte(`{"error":"error occurred while trying to read request"}`))
		} else {
			requestMap := map[string]any{}
			err = json.Unmarshal(requestBytes, &requestMap)
			if err != nil {
				responseWriter.Write([]byte(`{"error":"error occurred while trying to json decode request"}`))
			} else {
				responseMap := map[string]any{"content_type": requestContentType, "request": requestMap}
				responseBytes, err := conv.JSON_Marshal(responseMap)
				if err != nil {
					responseWriter.Write([]byte(`{"error":"error occurred while trying to json encode response"}`))
				} else {
					responseWriter.Write(responseBytes)
				}
			}
		}
	}))
	//--------------------
	defer server.Close()
	//--------------------------------------------------
	requestContentType := "application/json; charset=UTF-8"
	requestHeadersMap := map[string]string{"Content-Type": requestContentType}
	requestMap := map[string]any{"test_key": "test_value"}
	//--------------------
	EXPECTED_responseMap := map[string]any{"content_type": requestContentType, "request": requestMap}
	//--------------------------------------------------

	//--------------------------------------------------
	rpcInstance.ResponseURL = server.URL
	rpcInstance.ResponseHeadersMap = requestHeadersMap
	//--------------------------------------------------
	responseMap, err := rpcInstance.RPC_send_json_request(requestMap)
	//--------------------------------------------------

	//--------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		//--------------------
		if fmt.Sprint(responseMap) != fmt.Sprint(EXPECTED_responseMap) {
			t.Errorf("response = %q but should = %q", responseMap, EXPECTED_responseMap)
		}
		//--------------------
	}
	//--------------------------------------------------
}

//----------------------------------------------------------------------
// read request method
//----------------------------------------------------------------------

func TestRPC_read_request_method(t *testing.T) {
	//--------------------------------------------------
	type testDataStruct struct {
		requestString string
		requestMap    map[string]any
		errorString   string
	}
	//--------------------------------------------------
	testData := []testDataStruct{
		{"", map[string]any{}, "unexpected end of JSON input"},
		{`{"invalid_json_data"}`, map[string]any{}, "invalid character '}' after object key"},
		{`{"test_key":"test_value"}`, map[string]any{"test_key": "test_value"}, ""},
	}
	//--------------------------------------------------
	for index, test := range testData {

		//--------------------------------------------------
		rpcInstance.HttpRequest = httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(test.requestString)))
		//--------------------------------------------------
		err := rpcInstance.RPC_read_request()
		//--------------------
		errorString := fmt.Sprint(err)
		//--------------------
		if err != nil && errorString != test.errorString {
			t.Errorf("index %d: result error = %q but expected error = %q", index, errorString, test.errorString)
		} else {

			//--------------------
			if rpcInstance.RequestString != test.requestString {
				t.Errorf("index %d: resultString = %q but should = %q", index, rpcInstance.RequestString, test.requestString)
			}
			//--------------------
			if fmt.Sprint(rpcInstance.RequestMap) != fmt.Sprint(test.requestMap) {
				t.Errorf("index %d: requestMap = %q but should = %q", index, fmt.Sprint(rpcInstance.RequestMap), fmt.Sprint(test.requestMap))
			}
			//--------------------
		}
		//--------------------------------------------------
	}
	//--------------------------------------------------
}

func TestRPC_read_request_method_base64(t *testing.T) {
	//--------------------------------------------------
	type testDataStruct struct {
		requestString string
		errorString   string
	}
	//--------------------------------------------------
	testData := []testDataStruct{
		{"", "unexpected end of JSON input"},
		{`{"invalid_json_data"}`, "invalid character '}' after object key"},
		{`{"test_key":"test_value"}`, ""},
	}
	//--------------------------------------------------
	for index, test := range testData {

		//--------------------------------------------------
		rpcInstance.HttpRequest = httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(test.requestString)))
		//--------------------------------------------------
		err := rpcInstance.RPC_read_request()
		//--------------------
		errorString := fmt.Sprint(err)
		//--------------------
		if err != nil && errorString != test.errorString {
			t.Errorf("index %d: result error = %q but expected error = %q", index, errorString, test.errorString)
		} else {
			//--------------------
			if rpcInstance.RequestString != test.requestString {
				t.Errorf("index %d: resultString = %q but should = %q", index, rpcInstance.RequestString, test.requestString)
			}
			//--------------------
		}
		//--------------------------------------------------
	}
	//--------------------------------------------------
}

//----------------------------------------------------------------------
// read json request method
//----------------------------------------------------------------------

func TestRPC_read_json_request_method(t *testing.T) {
	//--------------------------------------------------
	type testDataStruct struct {
		requestString string
		requestMap    any
		errorString   string
	}
	//--------------------------------------------------
	testData := []testDataStruct{
		{`["INVALID_REQUEST"]`, map[string]any{}, "invalid request"},
		{`{"invalid_json_data"}`, map[string]any{}, "invalid character '}' after object key"},
		{`{"invalid_json_data"}`, map[string]any{}, "invalid character '}' after object key"},
		{`[{"test_key":"test_value"}]`, map[string]any{}, "invalid request"},
		{`{"test_key":"test_value"}`, map[string]any{"test_key": "test_value"}, ""},
	}
	//--------------------------------------------------
	for index, test := range testData {

		//--------------------------------------------------
		rpcInstance.HttpRequest = httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(test.requestString)))
		//--------------------------------------------------
		err := rpcInstance.RPC_read_json_request()
		//--------------------
		errorString := fmt.Sprint(err)
		//--------------------
		if err != nil && errorString != test.errorString {
			t.Errorf("index %d: result error = %q but expected error = %q", index, errorString, test.errorString)
		} else {

			//--------------------
			if rpcInstance.RequestString != test.requestString {
				t.Errorf("index %d: resultString = %q but should = %q", index, rpcInstance.RequestString, test.requestString)
			}
			//--------------------
			if fmt.Sprint(rpcInstance.RequestMap) != fmt.Sprint(test.requestMap) {
				t.Errorf("index %d: requestMap = %q but should = %q", index, fmt.Sprint(rpcInstance.RequestMap), fmt.Sprint(test.requestMap))
			}
			//--------------------
		}
		//--------------------------------------------------
	}
	//--------------------------------------------------
}

//----------------------------------------------------------------------
// read json request method
//----------------------------------------------------------------------

func TestRPC_read_jsonrpc_request_method(t *testing.T) {
	//--------------------------------------------------
	type testDataStruct struct {
		requestString string
		requestMap    any
		errorString   string
	}
	//--------------------------------------------------
	testData := []testDataStruct{
		{`["INVALID_REQUEST"]`, map[string]any{}, "parse error"},
		{`{"invalid_json_data"}`, map[string]any{}, "invalid character '}' after object key"},
		{`{"invalid_json_data"}`, map[string]any{}, "invalid character '}' after object key"},
		{`[{"test_key":"test_value"}]`, map[string]any{}, "parse error"},
		{`{"test_key":"test_value"}`, map[string]any{"test_key": "test_value"}, ""},
	}
	//--------------------------------------------------
	for index, test := range testData {

		//--------------------------------------------------
		rpcInstance.HttpRequest = httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(test.requestString)))
		//--------------------------------------------------
		err := rpcInstance.RPC_read_jsonrpc_request()
		//--------------------
		errorString := fmt.Sprint(err)
		//--------------------
		if err != nil && errorString != test.errorString {
			t.Errorf("index %d: result error = %q but expected error = %q", index, errorString, test.errorString)
		} else {

			//--------------------
			if rpcInstance.RequestString != test.requestString {
				t.Errorf("index %d: resultString = %q but should = %q", index, rpcInstance.RequestString, test.requestString)
			}
			//--------------------
			if fmt.Sprint(rpcInstance.RequestMap) != fmt.Sprint(test.requestMap) {
				t.Errorf("index %d: requestMap = %q but should = %q", index, fmt.Sprint(rpcInstance.RequestMap), fmt.Sprint(test.requestMap))
			}
			//--------------------
		}
		//--------------------------------------------------
	}
	//--------------------------------------------------
}

//----------------------------------------------------------------------
// send response method
//----------------------------------------------------------------------

func TestRPC_send_response_method(t *testing.T) {
	//--------------------------------------------------
	requestString := "<TEST_DATA>"
	requestContentType := "text/plain; charset=UTF-8"
	//--------------------
	EXPECTED_responseContentType := "text/plain; charset=UTF-8"
	EXPECTED_responseString := "<TEST_DATA>"
	//--------------------------------------------------
	httptestRecorder := httptest.NewRecorder()
	httpRequest := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(requestString)))
	//--------------------
	rpcInstance.ResponseWriter = httptestRecorder
	rpcInstance.HttpRequest = httpRequest
	rpcInstance.ResponseHeadersMap = map[string]string{"Content-Type": requestContentType}
	//--------------------------------------------------

	//--------------------
	rpcInstance.RPC_send_response(requestString)
	//--------------------

	//--------------------------------------------------
	httpResponse := httptestRecorder.Result()
	//--------------------
	defer httpResponse.Body.Close()
	//--------------------
	responseBytes, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		t.Error(err)
	} else {

		//--------------------
		responseContentType := httpResponse.Header.Get("Content-Type")
		responseResponseString := string(responseBytes)
		//--------------------
		if responseContentType != EXPECTED_responseContentType {
			t.Errorf("Content-Type = %q but should = %q", responseContentType, EXPECTED_responseContentType)
		}
		//--------------------
		if responseResponseString != EXPECTED_responseString {
			t.Errorf("response = %q but should = %q", responseResponseString, EXPECTED_responseString)
		}
		//--------------------
	}
	//--------------------------------------------------
}

//----------------------------------------------------------------------
// send response method - base64
//----------------------------------------------------------------------

func TestRPC_send_response_method_base64(t *testing.T) {
	//--------------------------------------------------
	rpcInstance.Encoding = "base64"
	//--------------------------------------------------
	requestContentType := "text/plain; charset=UTF-8"
	requestString := "<TEST_DATA>"
	//--------------------
	EXPECTED_responseContentType := "text/plain; charset=UTF-8"
	EXPECTED_responseString := "<TEST_DATA>"
	//--------------------------------------------------
	httptestRecorder := httptest.NewRecorder()
	httpRequest := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(requestString)))
	//--------------------
	rpcInstance.ResponseWriter = httptestRecorder
	rpcInstance.HttpRequest = httpRequest
	rpcInstance.ResponseHeadersMap = map[string]string{"Content-Type": requestContentType}
	//--------------------------------------------------

	//--------------------
	rpcInstance.RPC_send_response(requestString)
	//--------------------

	//--------------------------------------------------
	httpResponse := httptestRecorder.Result()
	//--------------------
	defer httpResponse.Body.Close()
	//--------------------
	responseBytes, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		t.Error(err)
	} else {

		//--------------------
		responseContentType := httpResponse.Header.Get("Content-Type")
		responseString := string(responseBytes)
		//--------------------
		if rpcInstance.Encoding == "base64" {
			responseString, _ = conv.Base64_decode(responseString)
		} else if rpcInstance.Encoding == "base64url" {
			responseString, _ = conv.Base64url_decode(responseString)
		}
		//--------------------
		if responseContentType != EXPECTED_responseContentType {
			t.Errorf("Content-Type = %q but should = %q", responseContentType, EXPECTED_responseContentType)
		}
		//--------------------
		if responseString != EXPECTED_responseString {
			t.Errorf("response = %q but should = %q", responseString, EXPECTED_responseString)
		}
		//--------------------
	}
	//--------------------------------------------------
	rpcInstance.Encoding = ""
	//--------------------------------------------------
}

//----------------------------------------------------------------------
// send json response method
//----------------------------------------------------------------------

func TestRPC_send_json_response_method(t *testing.T) {
	//--------------------------------------------------
	requestContentType := "application/json; charset=UTF-8"
	requestString := ""
	requestMap := map[string]any{"test_key": "test_value"}
	//--------------------
	EXPECTED_responseContentType := "application/json; charset=UTF-8"
	EXPECTED_responseString := `{"test_key":"test_value"}`
	//--------------------------------------------------
	httptestRecorder := httptest.NewRecorder()
	httpRequest := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(requestString)))
	//--------------------
	rpcInstance.ResponseWriter = httptestRecorder
	rpcInstance.HttpRequest = httpRequest
	rpcInstance.ResponseHeadersMap = map[string]string{"Content-Type": requestContentType}
	//--------------------------------------------------

	//--------------------
	rpcInstance.RPC_send_json_response(requestMap)
	//--------------------

	//--------------------------------------------------
	httpResponse := httptestRecorder.Result()
	//--------------------
	defer httpResponse.Body.Close()
	//--------------------
	responseBytes, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		t.Error(err)
	} else {

		//--------------------
		responseContentType := httpResponse.Header.Get("Content-Type")
		responseString := string(responseBytes)
		//--------------------
		if responseContentType != EXPECTED_responseContentType {
			t.Errorf("Content-Type = %q but should = %q", responseContentType, EXPECTED_responseContentType)
		}
		//--------------------
		if responseString != EXPECTED_responseString {
			t.Errorf("response = %q but should = %q", responseString, EXPECTED_responseString)
		}
		//--------------------
	}
	//--------------------------------------------------
}

//----------------------------------------------------------------------
// send result response method
//----------------------------------------------------------------------

func TestRPC_send_result_response_method(t *testing.T) {
	//--------------------------------------------------
	requestContentType := "application/json; charset=UTF-8"
	requestString := ""
	requestMap := map[string]any{"test_key": "test_value"}
	//--------------------
	EXPECTED_responseContentType := "application/json; charset=UTF-8"
	EXPECTED_responseString := `{"result":{"test_key":"test_value"}}`
	//--------------------------------------------------
	httptestRecorder := httptest.NewRecorder()
	httpRequest := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(requestString)))
	//--------------------
	rpcInstance.ResponseWriter = httptestRecorder
	rpcInstance.HttpRequest = httpRequest
	rpcInstance.ResponseHeadersMap = map[string]string{"Content-Type": requestContentType}
	//--------------------------------------------------

	//--------------------
	rpcInstance.RPC_send_result_response(requestMap)
	//--------------------

	//--------------------------------------------------
	httpResponse := httptestRecorder.Result()
	//--------------------
	defer httpResponse.Body.Close()
	//--------------------
	responseBytes, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		t.Error(err)
	} else {

		//--------------------
		responseContentType := httpResponse.Header.Get("Content-Type")
		responseString := string(responseBytes)
		//--------------------
		if responseContentType != EXPECTED_responseContentType {
			t.Errorf("Content-Type = %q but should = %q", responseContentType, EXPECTED_responseContentType)
		}
		//--------------------
		if responseString != EXPECTED_responseString {
			t.Errorf("response = %q but should = %q", responseString, EXPECTED_responseString)
		}
		//--------------------
	}
	//--------------------------------------------------
}

//----------------------------------------------------------------------
// send error response method
//----------------------------------------------------------------------

func TestRPC_send_error_response_method(t *testing.T) {
	//--------------------------------------------------
	requestContentType := "application/json; charset=UTF-8"
	requestString := ""
	requestErrorString := "this is an error"
	//--------------------
	EXPECTED_responseContentType := "application/json; charset=UTF-8"
	EXPECTED_responseString := `{"error":"this is an error"}`
	//--------------------------------------------------
	httptestRecorder := httptest.NewRecorder()
	httpRequest := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(requestString)))
	//--------------------
	rpcInstance.ResponseWriter = httptestRecorder
	rpcInstance.HttpRequest = httpRequest
	rpcInstance.ResponseHeadersMap = map[string]string{"Content-Type": requestContentType}
	//--------------------------------------------------

	//--------------------
	rpcInstance.RPC_send_error_response(requestErrorString)
	//--------------------

	//--------------------------------------------------
	httpResponse := httptestRecorder.Result()
	//--------------------
	defer httpResponse.Body.Close()
	//--------------------
	responseBytes, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		t.Error(err)
	} else {

		//--------------------
		responseContentType := httpResponse.Header.Get("Content-Type")
		responseString := string(responseBytes)
		//--------------------
		if responseContentType != EXPECTED_responseContentType {
			t.Errorf("Content-Type = %q but should = %q", responseContentType, EXPECTED_responseContentType)
		}
		//--------------------
		if responseString != EXPECTED_responseString {
			t.Errorf("response = %q but should = %q", responseString, EXPECTED_responseString)
		}
		//--------------------
	}
	//--------------------------------------------------
}

//----------------------------------------------------------------------
// send JSON-RPC request method
//----------------------------------------------------------------------

func TestRPC_send_jsonrpc_request_method(t *testing.T) {
	//--------------------------------------------------
	server := httptest.NewServer(http.HandlerFunc(func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		ServerRequestContentType := httpRequest.Header.Get("Content-Type")
		ServerRequestBytes, err := io.ReadAll(httpRequest.Body)
		if err != nil {
			responseWriter.Write([]byte(`{"error":"error occurred while trying to read request"}`))
		} else {
			ServerRequestMap := map[string]any{}
			err = json.Unmarshal(ServerRequestBytes, &ServerRequestMap)
			if err != nil {
				responseWriter.Write([]byte(`{"error":"error occurred while trying to json decode request"}`))
			} else {
				serverResponseMap := map[string]any{"content_type": ServerRequestContentType, "request": ServerRequestMap}
				responseBytes, err := conv.JSON_Marshal(serverResponseMap)
				if err != nil {
					responseWriter.Write([]byte(`{"error":"error occurred while trying to json encode response"}`))
				} else {
					responseWriter.Write(responseBytes)
				}
			}
		}
	}))
	//--------------------
	defer server.Close()
	//--------------------------------------------------
	requestContentType := "application/json; charset=UTF-8"
	requestHeadersMap := map[string]string{"Content-Type": requestContentType}
	requestMap := map[string]any{"test_key": "test_value"}
	//--------------------
	EXPECTED_responseMap := map[string]any{"content_type": requestContentType, "request": requestMap}
	//--------------------------------------------------

	//--------------------------------------------------
	var err error
	var responseMap any
	//--------------------------------------------------
	rpcInstance.ResponseURL = server.URL
	rpcInstance.ResponseHeadersMap = requestHeadersMap
	//--------------------------------------------------

	//--------------------------------------------------
	_, err = rpcInstance.RPC_send_jsonrpc_request(requestMap)
	//--------------------
	if err == nil {
		t.Error("undefined method not being reported")
	}
	//--------------------
	if fmt.Sprint(err) != "method is not defined or is not a string" {
		t.Error(err)
	}
	//--------------------------------------------------

	//--------------------------------------------------
	requestMap["method"] = 0
	//--------------------
	_, err = rpcInstance.RPC_send_jsonrpc_request(requestMap)
	//--------------------
	if err == nil {
		t.Error("undefined method not being reported")
	}
	//--------------------
	if fmt.Sprint(err) != "method is not defined or is not a string" {
		t.Error(err)
	}
	//--------------------------------------------------

	//--------------------------------------------------
	requestMap["method"] = "echo"
	requestMap["params"] = 0
	//--------------------
	_, err = rpcInstance.RPC_send_jsonrpc_request(requestMap)
	//--------------------
	if err == nil {
		t.Error("invalid params not being reported")
	}
	//--------------------
	if fmt.Sprint(err) != "params is not an array or an object" {
		t.Error(err)
	}
	//--------------------------------------------------

	//--------------------------------------------------
	requestMap["params"] = []any{}
	//--------------------
	_, err = rpcInstance.RPC_send_jsonrpc_request(requestMap)
	//--------------------
	if err != nil {
		t.Error(err)
	}
	//--------------------------------------------------

	//--------------------------------------------------
	requestMap["params"] = map[string]any{}
	//--------------------
	_, err = rpcInstance.RPC_send_jsonrpc_request(requestMap)
	//--------------------
	if err != nil {
		t.Error(err)
	}
	//--------------------------------------------------

	//--------------------------------------------------
	requestMap["params"] = []any{}
	//--------------------
	responseMap, err = rpcInstance.RPC_send_jsonrpc_request(requestMap)
	//--------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		//--------------------
		if fmt.Sprint(responseMap) != fmt.Sprint(EXPECTED_responseMap) {
			t.Errorf("response = %q but should = %q", responseMap, EXPECTED_responseMap)
		}
		//--------------------
	}
	//--------------------------------------------------
}

//----------------------------------------------------------------------
// send JSON-RPC result response method
//----------------------------------------------------------------------

func TestRPC_send_jsonrpc_result_method(t *testing.T) {
	//--------------------------------------------------
	requestContentType := "application/json; charset=UTF-8"
	requestString := ""
	requestResult := map[string]any{"test_key": "test_value"}
	//--------------------
	EXPECTED_responseContentType := "application/json; charset=UTF-8"
	EXPECTED_responseString := `{"id":null,"jsonrpc":"2.0","result":{"test_key":"test_value"}}`
	//--------------------------------------------------
	httptestRecorder := httptest.NewRecorder()
	httpRequest := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(requestString)))
	//--------------------
	rpcInstance.ResponseWriter = httptestRecorder
	rpcInstance.HttpRequest = httpRequest
	rpcInstance.ResponseHeadersMap = map[string]string{"Content-Type": requestContentType}
	//--------------------------------------------------

	//--------------------
	rpcInstance.RPC_send_jsonrpc_result(requestResult)
	//--------------------

	//--------------------------------------------------
	httpResponse := httptestRecorder.Result()
	//--------------------
	defer httpResponse.Body.Close()
	//--------------------
	responseBytes, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		t.Error(err)
	} else {

		//--------------------
		responseContentType := httpResponse.Header.Get("Content-Type")
		responseString := string(responseBytes)
		//--------------------
		if responseContentType != EXPECTED_responseContentType {
			t.Errorf("Content-Type = %q but should = %q", responseContentType, EXPECTED_responseContentType)
		}
		//--------------------
		if responseString != EXPECTED_responseString {
			t.Errorf("response = %q but should = %q", responseString, EXPECTED_responseString)
		}
		//--------------------
	}
	//--------------------------------------------------
}

//----------------------------------------------------------------------
// send JSON-RPC error response method
//----------------------------------------------------------------------

func TestRPC_send_jsonrpc_error_method(t *testing.T) {
	//--------------------------------------------------
	requestContentType := "application/json; charset=UTF-8"
	requestString := ""
	requestError := map[string]any{"code": -32700, "message": "parse error"}
	//--------------------
	EXPECTED_responseContentType := "application/json; charset=UTF-8"
	EXPECTED_responseString := `{"error":{"code":-32700,"message":"parse error"},"id":null,"jsonrpc":"2.0"}`
	//--------------------------------------------------
	httptestRecorder := httptest.NewRecorder()
	httpRequest := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(requestString)))
	//--------------------
	rpcInstance.ResponseWriter = httptestRecorder
	rpcInstance.HttpRequest = httpRequest
	rpcInstance.ResponseHeadersMap = map[string]string{"Content-Type": requestContentType}
	//--------------------------------------------------

	//--------------------
	rpcInstance.RPC_send_jsonrpc_error(requestError)
	//--------------------

	//--------------------------------------------------
	httpResponse := httptestRecorder.Result()
	//--------------------
	defer httpResponse.Body.Close()
	//--------------------
	responseBytes, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		t.Error(err)
	} else {

		//--------------------
		responseContentType := httpResponse.Header.Get("Content-Type")
		responseString := string(responseBytes)
		//--------------------
		if responseContentType != EXPECTED_responseContentType {
			t.Errorf("Content-Type = %q but should = %q", responseContentType, EXPECTED_responseContentType)
		}
		//--------------------
		if responseString != EXPECTED_responseString {
			t.Errorf("response = %q but should = %q", responseString, EXPECTED_responseString)
		}
		//--------------------
	}
	//--------------------------------------------------
}

//----------------------------------------------------------------------
// send JSON-RPC internal error response method
//----------------------------------------------------------------------

// no data param passed

func TestRPC_send_jsonrpc_internal_error1_method(t *testing.T) {
	//--------------------------------------------------
	requestContentType := "application/json; charset=UTF-8"
	requestString := ""
	//--------------------
	EXPECTED_responseContentType := "application/json; charset=UTF-8"
	EXPECTED_responseString := `{"error":{"code":-32603,"message":"internal error"},"id":null,"jsonrpc":"2.0"}`
	//--------------------------------------------------
	httptestRecorder := httptest.NewRecorder()
	httpRequest := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(requestString)))
	//--------------------
	rpcInstance.ResponseWriter = httptestRecorder
	rpcInstance.HttpRequest = httpRequest
	rpcInstance.ResponseHeadersMap = map[string]string{"Content-Type": requestContentType}
	//--------------------------------------------------

	//--------------------
	rpcInstance.RPC_send_jsonrpc_internal_error()
	//--------------------

	//--------------------------------------------------
	httpResponse := httptestRecorder.Result()
	//--------------------
	defer httpResponse.Body.Close()
	//--------------------
	responseBytes, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		t.Error(err)
	} else {

		//--------------------
		responseContentType := httpResponse.Header.Get("Content-Type")
		responseString := string(responseBytes)
		//--------------------
		if responseContentType != EXPECTED_responseContentType {
			t.Errorf("Content-Type = %q but should = %q", responseContentType, EXPECTED_responseContentType)
		}
		//--------------------
		if responseString != EXPECTED_responseString {
			t.Errorf("response = %q but should = %q", responseString, EXPECTED_responseString)
		}
		//--------------------
	}
	//--------------------------------------------------
}

// data param passed

func TestRPC_send_jsonrpc_internal_error2_method(t *testing.T) {
	//--------------------------------------------------
	requestContentType := "application/json; charset=UTF-8"
	requestString := ""
	//--------------------
	EXPECTED_responseContentType := "application/json; charset=UTF-8"
	EXPECTED_responseString := `{"error":{"code":-32603,"data":"test data","message":"internal error"},"id":null,"jsonrpc":"2.0"}`
	//--------------------------------------------------
	httptestRecorder := httptest.NewRecorder()
	httpRequest := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(requestString)))
	//--------------------
	rpcInstance.ResponseWriter = httptestRecorder
	rpcInstance.HttpRequest = httpRequest
	rpcInstance.ResponseHeadersMap = map[string]string{"Content-Type": requestContentType}
	//--------------------------------------------------

	//--------------------
	rpcInstance.RPC_send_jsonrpc_internal_error("test data")
	//--------------------

	//--------------------------------------------------
	httpResponse := httptestRecorder.Result()
	//--------------------
	defer httpResponse.Body.Close()
	//--------------------
	responseBytes, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		t.Error(err)
	} else {

		//--------------------
		responseContentType := httpResponse.Header.Get("Content-Type")
		responseString := string(responseBytes)
		//--------------------
		if responseContentType != EXPECTED_responseContentType {
			t.Errorf("Content-Type = %q but should = %q", responseContentType, EXPECTED_responseContentType)
		}
		//--------------------
		if responseString != EXPECTED_responseString {
			t.Errorf("response = %q but should = %q", responseString, EXPECTED_responseString)
		}
		//--------------------
	}
	//--------------------------------------------------
}

//----------------------------------------------------------------------
// send JSON-RPC invalid params response method
//----------------------------------------------------------------------

func TestRPC_send_jsonrpc_invalid_params_method(t *testing.T) {
	//--------------------------------------------------
	requestContentType := "application/json; charset=UTF-8"
	requestString := ""
	//--------------------
	EXPECTED_responseContentType := "application/json; charset=UTF-8"
	EXPECTED_responseString := `{"error":{"code":-32602,"message":"invalid params"},"id":null,"jsonrpc":"2.0"}`
	//--------------------------------------------------
	httptestRecorder := httptest.NewRecorder()
	httpRequest := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(requestString)))
	//--------------------
	rpcInstance.ResponseWriter = httptestRecorder
	rpcInstance.HttpRequest = httpRequest
	rpcInstance.ResponseHeadersMap = map[string]string{"Content-Type": requestContentType}
	//--------------------------------------------------

	//--------------------
	rpcInstance.RPC_send_jsonrpc_invalid_params()
	//--------------------

	//--------------------------------------------------
	httpResponse := httptestRecorder.Result()
	//--------------------
	defer httpResponse.Body.Close()
	//--------------------
	responseBytes, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		t.Error(err)
	} else {

		//--------------------
		responseContentType := httpResponse.Header.Get("Content-Type")
		responseString := string(responseBytes)
		//--------------------
		if responseContentType != EXPECTED_responseContentType {
			t.Errorf("Content-Type = %q but should = %q", responseContentType, EXPECTED_responseContentType)
		}
		//--------------------
		if responseString != EXPECTED_responseString {
			t.Errorf("response = %q but should = %q", responseString, EXPECTED_responseString)
		}
		//--------------------
	}
	//--------------------------------------------------
}

//----------------------------------------------------------------------
// send JSON-RPC invalid request response method
//----------------------------------------------------------------------

func TestRPC_send_jsonrpc_invalid_request_method(t *testing.T) {
	//--------------------------------------------------
	requestContentType := "application/json; charset=UTF-8"
	requestString := ""
	//--------------------
	EXPECTED_responseContentType := "application/json; charset=UTF-8"
	EXPECTED_responseString := `{"error":{"code":-32600,"message":"invalid request"},"id":null,"jsonrpc":"2.0"}`
	//--------------------------------------------------
	httptestRecorder := httptest.NewRecorder()
	httpRequest := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(requestString)))
	//--------------------
	rpcInstance.ResponseWriter = httptestRecorder
	rpcInstance.HttpRequest = httpRequest
	rpcInstance.ResponseHeadersMap = map[string]string{"Content-Type": requestContentType}
	//--------------------------------------------------

	//--------------------
	rpcInstance.RPC_send_jsonrpc_invalid_request()
	//--------------------

	//--------------------------------------------------
	httpResponse := httptestRecorder.Result()
	//--------------------
	defer httpResponse.Body.Close()
	//--------------------
	responseBytes, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		t.Error(err)
	} else {

		//--------------------
		responseContentType := httpResponse.Header.Get("Content-Type")
		responseString := string(responseBytes)
		//--------------------
		if responseContentType != EXPECTED_responseContentType {
			t.Errorf("Content-Type = %q but should = %q", responseContentType, EXPECTED_responseContentType)
		}
		//--------------------
		if responseString != EXPECTED_responseString {
			t.Errorf("response = %q but should = %q", responseString, EXPECTED_responseString)
		}
		//--------------------
	}
	//--------------------------------------------------
}

//----------------------------------------------------------------------
// send JSON-RPC method not found response method
//----------------------------------------------------------------------

func TestRPC_send_jsonrpc_method_not_found_method(t *testing.T) {
	//--------------------------------------------------
	requestContentType := "application/json; charset=UTF-8"
	requestString := ""
	//--------------------
	EXPECTED_responseContentType := "application/json; charset=UTF-8"
	EXPECTED_responseString := `{"error":{"code":-32601,"message":"method not found"},"id":null,"jsonrpc":"2.0"}`
	//--------------------------------------------------
	httptestRecorder := httptest.NewRecorder()
	httpRequest := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(requestString)))
	//--------------------
	rpcInstance.ResponseWriter = httptestRecorder
	rpcInstance.HttpRequest = httpRequest
	rpcInstance.ResponseHeadersMap = map[string]string{"Content-Type": requestContentType}
	//--------------------------------------------------

	//--------------------
	rpcInstance.RPC_send_jsonrpc_method_not_found()
	//--------------------

	//--------------------------------------------------
	httpResponse := httptestRecorder.Result()
	//--------------------
	defer httpResponse.Body.Close()
	//--------------------
	responseBytes, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		t.Error(err)
	} else {

		//--------------------
		responseContentType := httpResponse.Header.Get("Content-Type")
		responseString := string(responseBytes)
		//--------------------
		if responseContentType != EXPECTED_responseContentType {
			t.Errorf("Content-Type = %q but should = %q", responseContentType, EXPECTED_responseContentType)
		}
		//--------------------
		if responseString != EXPECTED_responseString {
			t.Errorf("response = %q but should = %q", responseString, EXPECTED_responseString)
		}
		//--------------------
	}
	//--------------------------------------------------
}

//----------------------------------------------------------------------
// send JSON-RPC parse error response method
//----------------------------------------------------------------------

func TestRPC_send_jsonrpc_parse_error(t *testing.T) {
	//--------------------------------------------------
	requestContentType := "application/json; charset=UTF-8"
	requestString := ""
	//--------------------
	EXPECTED_responseContentType := "application/json; charset=UTF-8"
	EXPECTED_responseString := `{"error":{"code":-32700,"message":"parse error"},"id":null,"jsonrpc":"2.0"}`
	//--------------------------------------------------
	httptestRecorder := httptest.NewRecorder()
	httpRequest := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(requestString)))
	//--------------------
	rpcInstance.ResponseWriter = httptestRecorder
	rpcInstance.HttpRequest = httpRequest
	rpcInstance.ResponseHeadersMap = map[string]string{"Content-Type": requestContentType}
	//--------------------------------------------------

	//--------------------
	rpcInstance.RPC_send_jsonrpc_parse_error()
	//--------------------

	//--------------------------------------------------
	httpResponse := httptestRecorder.Result()
	//--------------------
	defer httpResponse.Body.Close()
	//--------------------
	responseBytes, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		t.Error(err)
	} else {

		//--------------------
		responseContentType := httpResponse.Header.Get("Content-Type")
		responseString := string(responseBytes)
		//--------------------
		if responseContentType != EXPECTED_responseContentType {
			t.Errorf("Content-Type = %q but should = %q", responseContentType, EXPECTED_responseContentType)
		}
		//--------------------
		if responseString != EXPECTED_responseString {
			t.Errorf("response = %q but should = %q", responseString, EXPECTED_responseString)
		}
		//--------------------
	}
	//--------------------------------------------------
}

//----------------------------------------------------------------------
// send JSON-RPC server error response method
//----------------------------------------------------------------------

func TestRPC_send_jsonrpc_server_error(t *testing.T) {
	//--------------------------------------------------
	requestContentType := "application/json; charset=UTF-8"
	requestString := ""
	//--------------------
	EXPECTED_responseContentType := "application/json; charset=UTF-8"
	EXPECTED_responseString := `{"error":{"code":-32000,"message":"server error"},"id":null,"jsonrpc":"2.0"}`
	//--------------------------------------------------
	httptestRecorder := httptest.NewRecorder()
	httpRequest := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(requestString)))
	//--------------------
	rpcInstance.ResponseWriter = httptestRecorder
	rpcInstance.HttpRequest = httpRequest
	rpcInstance.ResponseHeadersMap = map[string]string{"Content-Type": requestContentType}
	//--------------------------------------------------

	//--------------------
	rpcInstance.RPC_send_jsonrpc_server_error()
	//--------------------

	//--------------------------------------------------
	httpResponse := httptestRecorder.Result()
	//--------------------
	defer httpResponse.Body.Close()
	//--------------------
	responseBytes, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		t.Error(err)
	} else {

		//--------------------
		responseContentType := httpResponse.Header.Get("Content-Type")
		responseString := string(responseBytes)
		//--------------------
		if responseContentType != EXPECTED_responseContentType {
			t.Errorf("Content-Type = %q but should = %q", responseContentType, EXPECTED_responseContentType)
		}
		//--------------------
		if responseString != EXPECTED_responseString {
			t.Errorf("response = %q but should = %q", responseString, EXPECTED_responseString)
		}
		//--------------------
	}
	//--------------------------------------------------
}

//----------------------------------------------------------------------
//################################################################################
//----------------------------------------------------------------------

//----------------------------------------------------------------------
// echo method
//----------------------------------------------------------------------

func TestRPC_echo_method(t *testing.T) {
	//--------------------------------------------------
	requestContentType := "application/json; charset=UTF-8"
	requestString := "{\"method\":\"echo\",\"data\":\"ECHO_TEST ABC <> &quot; \u00A3 \u65E5\u672C\u8A9E\U0001f427\"}"
	//--------------------
	EXPECTED_responseContentType := "application/json; charset=UTF-8"
	EXPECTED_responseString := `{"method":"echo","data":"ECHO_TEST ABC <> &quot; ` + "\u00A3 \u65E5\u672C\u8A9E\U0001f427" + `"}`
	//--------------------------------------------------
	httptestRecorder := httptest.NewRecorder()
	httpRequest := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(requestString)))
	//--------------------
	rpcInstance.ResponseWriter = httptestRecorder
	rpcInstance.HttpRequest = httpRequest
	rpcInstance.ResponseHeadersMap = map[string]string{"Content-Type": requestContentType}
	//--------------------------------------------------

	//--------------------
	rpcInstance.RPC_echo()
	//--------------------

	//--------------------------------------------------
	httpResponse := httptestRecorder.Result()
	//--------------------
	defer httpResponse.Body.Close()
	//--------------------
	responseBytes, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		t.Error(err)
	} else {

		//--------------------
		responseContentType := httpResponse.Header.Get("Content-Type")
		responseString := string(responseBytes)
		//--------------------
		if responseContentType != EXPECTED_responseContentType {
			t.Errorf("Content-Type = %q but should = %q", responseContentType, EXPECTED_responseContentType)
		}
		//--------------------
		if responseString != EXPECTED_responseString {
			t.Errorf("response = %q but should = %q", responseString, EXPECTED_responseString)
		}
		//--------------------
	}
	//--------------------------------------------------
}

//----------------------------------------------------------------------
// echo method
//----------------------------------------------------------------------

func TestJSONRPC_echo_method(t *testing.T) {
	//--------------------------------------------------
	requestContentType := "application/json; charset=UTF-8"
	requestString := "{\"id\":101,\"method\":\"echo\",\"params\":[\"ECHO_TEST ABC <> &quot; \u00A3 \u65E5\u672C\u8A9E\U0001f427\"]}"
	//--------------------
	EXPECTED_responseContentType := "application/json; charset=UTF-8"
	EXPECTED_responseString := `{"id":101,"jsonrpc":"2.0","result":"{\"id\":101,\"method\":\"echo\",\"params\":[\"ECHO_TEST ABC <> &quot; ` + "\u00A3 \u65E5\u672C\u8A9E\U0001f427" + `\"]}"}`
	//--------------------------------------------------
	httptestRecorder := httptest.NewRecorder()
	httpRequest := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(requestString)))
	//--------------------
	rpcInstance.ResponseWriter = httptestRecorder
	rpcInstance.HttpRequest = httpRequest
	rpcInstance.ResponseHeadersMap = map[string]string{"Content-Type": requestContentType}
	//--------------------------------------------------

	//--------------------
	rpcInstance.JSONRPC_echo()
	//--------------------

	//--------------------------------------------------
	httpResponse := httptestRecorder.Result()
	//--------------------
	defer httpResponse.Body.Close()
	//--------------------
	responseBytes, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		t.Error(err)
	} else {

		//--------------------
		responseContentType := httpResponse.Header.Get("Content-Type")
		responseString := string(responseBytes)
		//--------------------
		if responseContentType != EXPECTED_responseContentType {
			t.Errorf("Content-Type = %q but should = %q", responseContentType, EXPECTED_responseContentType)
		}
		//--------------------
		if responseString != EXPECTED_responseString {
			t.Errorf("response = %q but should = %q", responseString, EXPECTED_responseString)
		}
		//--------------------
	}
	//--------------------------------------------------
}

//----------------------------------------------------------------------
//################################################################################
//----------------------------------------------------------------------
/*

	Stadard Functions

*/
//----------------------------------------------------------------------
//################################################################################

//----------------------------------------------------------------------
// RPC handler
//----------------------------------------------------------------------

func TestRPC_Handler(t *testing.T) {
	//--------------------------------------------------
	EXPECTED_responseContentType := "application/json; charset=UTF-8"
	//--------------------------------------------------
	type testDataStruct struct {
		requestString  string
		responseString string
	}
	//--------------------------------------------------
	testData := []testDataStruct{
		{"INVALID_REQUEST", `{"error":"invalid request"}`},
		{`{"invalid_json_data"}`, `{"error":"invalid character '}' after object key"}`},
		{`{"no_method_key":"test_value"}`, `{"error":"method is not defined"}`},
		{`{"method":"NO_SUCH_METHOD"}`, `{"error":"method not found"}`},
		{`{"method":"echo","data":"ECHO_TEST ABC <> &quot; ` + "\u00A3 \u65E5\u672C\u8A9E\U0001f427\"}", `{"method":"echo","data":"ECHO_TEST ABC <> &quot; ` + "\u00A3 \u65E5\u672C\u8A9E\U0001f427" + `"}`},
	}
	//--------------------------------------------------
	for index, test := range testData {

		//--------------------------------------------------
		httptestRecorder := httptest.NewRecorder()
		httpRequest := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(test.requestString)))
		//--------------------------------------------------

		//--------------------------------------------------
		RPC_Handler(httptestRecorder, httpRequest)
		//--------------------------------------------------

		//--------------------------------------------------
		httpResponse := httptestRecorder.Result()
		//--------------------
		defer httpResponse.Body.Close()
		//--------------------
		responseBytes, err := io.ReadAll(httpResponse.Body)
		if err != nil {
			t.Error(err)
		} else {

			//--------------------
			responseContentType := httpResponse.Header.Get("Content-Type")
			responseString := string(responseBytes)
			//--------------------
			if responseContentType != EXPECTED_responseContentType {
				t.Errorf("index %d: Content-Type = %q but should = %q", index, responseContentType, EXPECTED_responseContentType)
			}
			//--------------------
			if responseString != test.responseString {
				t.Errorf("index %d: response = %q but should = %q", index, responseString, test.responseString)
			}
			//--------------------
		}
		//--------------------------------------------------
	}
	//--------------------------------------------------
}

//----------------------------------------------------------------------
// RPC handler
//----------------------------------------------------------------------

func TestJSONRPC_Handler(t *testing.T) {
	//--------------------------------------------------
	EXPECTED_responseContentType := "application/json; charset=UTF-8"
	//--------------------------------------------------
	type testDataStruct struct {
		requestString  string
		responseString string
	}
	//--------------------------------------------------
	testData := []testDataStruct{
		{"INVALID_JSON", `{"error":{"code":-32700,"data":"invalid character 'I' looking for beginning of value","message":"parse error"},"id":null,"jsonrpc":"2.0"}`},
		{`{"invalid_json"}`, `{"error":{"code":-32700,"data":"invalid character '}' after object key","message":"parse error"},"id":null,"jsonrpc":"2.0"}`},
		{`{"id":102,"no_method_key":"test_value"}`, `{"error":{"code":-32601,"message":"method not found"},"id":102,"jsonrpc":"2.0"}`},
		{`{"id":102,"method":"NO_SUCH_METHOD"}`, `{"error":{"code":-32601,"data":"NO_SUCH_METHOD","message":"method not found"},"id":102,"jsonrpc":"2.0"}`},
		{`{"id":103,"method":"echo","params":0}`, `{"error":{"code":-32602,"data":0,"message":"invalid params"},"id":103,"jsonrpc":"2.0"}`},
		{`{"id":104,"method":"echo","params":["ECHO_TEST ABC <> &quot; ` + "\u00A3 \u65E5\u672C\u8A9E\U0001f427\"]}", `{"id":104,"jsonrpc":"2.0","result":"{\"id\":104,\"method\":\"echo\",\"params\":[\"ECHO_TEST ABC <> &quot; ` + "\u00A3 \u65E5\u672C\u8A9E\U0001f427" + `\"]}"}`},
	}
	//--------------------------------------------------
	for index, test := range testData {

		//--------------------------------------------------
		httptestRecorder := httptest.NewRecorder()
		httpRequest := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(test.requestString)))
		//--------------------------------------------------

		//--------------------------------------------------
		JSONRPC_Handler(httptestRecorder, httpRequest)
		//--------------------------------------------------

		//--------------------------------------------------
		httpResponse := httptestRecorder.Result()
		//--------------------
		defer httpResponse.Body.Close()
		//--------------------
		responseBytes, err := io.ReadAll(httpResponse.Body)
		if err != nil {
			t.Error(err)
		} else {

			//--------------------
			responseContentType := httpResponse.Header.Get("Content-Type")
			responseString := string(responseBytes)
			//--------------------
			if responseContentType != EXPECTED_responseContentType {
				t.Errorf("index %d: Content-Type = %q but should = %q", index, responseContentType, EXPECTED_responseContentType)
			}
			//--------------------
			if responseString != test.responseString {
				t.Errorf("index %d: response = %q but should = %q", index, responseString, test.responseString)
			}
			//--------------------
		}
		//--------------------------------------------------
	}
	//--------------------------------------------------
}

//----------------------------------------------------------------------
//################################################################################
//----------------------------------------------------------------------

//----------------------------------------------------------------------
// EncodeData
//----------------------------------------------------------------------

func TestEncodeData(t *testing.T) {
	//--------------------------------------------------
	type testDataStruct struct {
		inputString    string
		encoding       string
		expectedString string
		expectedError  string
	}
	//--------------------------------------------------
	testData := []testDataStruct{
		{"ABCD", "INVALID_ENCODING", "", "invalid encoding"},
		{"", "", "", ""},
		{"ABCD", "obfuscate", "6+jp7g==", ""},
		{"ABCD", "base64", "QUJDRA==", ""},
		{"ABCD", "base64url", "QUJDRA", ""},
	}
	//--------------------------------------------------
	for _, test := range testData {
		//--------------------------------------------------
		resultString, err := EncodeData(test.inputString, test.encoding)
		//--------------------------------------------------
		if resultString != test.expectedString {
			//--------------------
			t.Errorf("returned result = %q but should = %q", resultString, test.expectedString)
			//--------------------
		}
		//--------------------------------------------------
		if err != nil {
			//--------------------
			resultError := err.Error()
			//--------------------
			if resultError != test.expectedError {
				t.Errorf("returned error = %q but should = %q", resultError, test.expectedError)
			}
			//--------------------
		} else {
			//--------------------
			if test.expectedError != "" {
				t.Errorf("returned error is nil but should = %q", test.expectedError)
			}
			//--------------------
		}
		//--------------------------------------------------
	}
	//--------------------------------------------------
}

//----------------------------------------------------------------------
// DecodeData
//----------------------------------------------------------------------

func TestDecodeData(t *testing.T) {
	//--------------------------------------------------
	type testDataStruct struct {
		inputString    string
		encoding       string
		expectedString string
		expectedError  string
	}
	//--------------------------------------------------
	testData := []testDataStruct{
		{"ABCD", "INVALID_ENCODING", "", "invalid encoding"},
		{"", "", "", ""},
		{"6+jp7g==", "obfuscate", "ABCD", ""},
		{"QUJDRA==", "base64", "ABCD", ""},
		{"QUJDRA", "base64url", "ABCD", ""},
	}
	//--------------------------------------------------
	for _, test := range testData {
		//--------------------------------------------------
		resultString, err := DecodeData(test.inputString, test.encoding)
		//--------------------------------------------------
		if resultString != test.expectedString {
			//--------------------
			t.Errorf("returned result = %q but should = %q", resultString, test.expectedString)
			//--------------------
		}
		//--------------------------------------------------
		if err != nil {
			//--------------------
			resultError := err.Error()
			//--------------------
			if resultError != test.expectedError {
				t.Errorf("returned error = %q but should = %q", resultError, test.expectedError)
			}
			//--------------------
		} else {
			//--------------------
			if test.expectedError != "" {
				t.Errorf("returned error is nil but should = %q", test.expectedError)
			}
			//--------------------
		}
		//--------------------------------------------------
	}
	//--------------------------------------------------
}

//----------------------------------------------------------------------
//################################################################################
//----------------------------------------------------------------------

//------------------------------------------------------------
// ObfuscateData
//------------------------------------------------------------

func TestObfuscateData(t *testing.T) {
	//--------------------------------------------------
	testCases := []struct {
		dataString     string
		base64Encode   bool
		base64Decode   bool
		value          byte
		useValue       bool
		expectedString string
		expectedError  string
	}{
		{"ABCD", false, false, 0, true, "", "value should be an integer between 1 and 255"},
		{"", false, false, 0, false, "", ""},
		{"\xEB\xE8\xE9\xEE", false, false, 0, false, "ABCD", ""},
		{"ABCD", true, false, 0, false, "6+jp7g==", ""},
		{"ABCD", true, false, 170, true, "6+jp7g==", ""},
		{"6+jp7g==", false, true, 0, false, "ABCD", ""},
		{"6+jp7g==", false, true, 170, true, "ABCD", ""},
	}
	//--------------------------------------------------
	for _, testCase := range testCases {
		//--------------------------------------------------
		var resultString string
		var err error
		//--------------------------------------------------
		if testCase.useValue {
			resultString, err = ObfuscateData(testCase.dataString, testCase.base64Encode, testCase.base64Decode, testCase.value)
		} else {
			resultString, err = ObfuscateData(testCase.dataString, testCase.base64Encode, testCase.base64Decode)
		}
		//--------------------------------------------------
		if resultString != testCase.expectedString {
			//--------------------------------------------------
			t.Errorf("(%q) resultString = %q %v but should = %q %v", testCase.dataString, resultString, []byte(resultString), testCase.expectedString, []byte(testCase.expectedString))
			//--------------------------------------------------
		}
		//--------------------------------------------------
		if err != nil {
			//--------------------
			resultError := err.Error()
			//--------------------
			if resultError != testCase.expectedError {
				t.Errorf("returned error = %q but should = %q", resultError, testCase.expectedError)
			}
			//--------------------
		} else {
			//--------------------
			if testCase.expectedError != "" {
				t.Errorf("returned error is nil but should = %q", testCase.expectedError)
			}
			//--------------------
		}
		//--------------------------------------------------
	}
	//--------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------
