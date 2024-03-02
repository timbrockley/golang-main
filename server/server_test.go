//------------------------------------------------------------

package server

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"regexp"
	"sync"
	"testing"
	"time"
)

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

func TestEcho(t *testing.T) {

	//--------------------------------------------------
	// Echo
	//--------------------------------------------------
	requestContentType := "text/plain; charset=UTF-8"
	requestString := "<ECHO_TEST>"
	//----------
	EXPECTED_contentType := "text/plain; charset=UTF-8"
	EXPECTED_responseString := "<ECHO_TEST>"
	//----------
	recorder := httptest.NewRecorder()
	//----------
	httpRequest := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(requestString)))
	//----------
	httpRequest.Header.Set("Content-Type", requestContentType)
	//----------

	//--------------------------------------------------
	Echo(recorder, httpRequest)
	//--------------------------------------------------

	//----------
	httpResponse := recorder.Result()
	//----------
	defer httpResponse.Body.Close()
	//----------
	responseBytes, err := io.ReadAll(httpResponse.Body)
	//----------
	if err != nil {

		t.Error(err)

	} else {

		//----------
		if string(responseBytes) != EXPECTED_responseString {
			t.Errorf("response %q expected %q", string(responseBytes), EXPECTED_responseString)
		}
		//----------
	}
	//----------
	responseContentType := httpRequest.Header.Get("Content-Type")
	//----------
	if responseContentType != EXPECTED_contentType {

		t.Errorf("content type = %q but should be %q", responseContentType, EXPECTED_contentType)
	}
	//--------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

func TestHeaders(t *testing.T) {

	//--------------------------------------------------
	// Headers
	//--------------------------------------------------
	s := httptest.NewServer(http.HandlerFunc(Headers))
	//--------------------------------------------------
	httpResponse, err := http.Get(s.URL)
	if err != nil {
		t.Error(err)
	}
	//----------
	defer httpResponse.Body.Close()
	//----------
	responseBytes, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		t.Error(err)
	}
	//--------------------------------------------------
	match, _ := regexp.MatchString("User-Agent", string(responseBytes))
	if !match {
		//----------
		t.Errorf("could not match test headers")
		//----------
	}
	//--------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

func TestSplitAddrPort(t *testing.T) {

	//--------------------------------------------------
	// SplitAddrPort
	//--------------------------------------------------
	addr := "127.0.0.1:80"
	//--------------------------------------------------
	ipAddr, port := SplitAddrPort(addr)
	//--------------------------------------------------
	if ipAddr != "127.0.0.1" {
		//----------
		t.Errorf("ipAddr should = %q but = %q", "127.0.0.1", ipAddr)
		//----------
	}
	//----------
	if port != 80 {
		//----------
		t.Errorf("port should = %d but = %d", 80, port)
		//----------
	}
	//--------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

func TestTCPServerEcho(t *testing.T) {

	//--------------------------------------------------
	// TCPServerEcho
	//--------------------------------------------------
	serverIPAddr := "127.0.0.1"
	//--------
	networkObject := NetworkStruct{ServerAddr: serverIPAddr}
	//--------
	serverIPAddr, serverPort := SplitAddrPort(networkObject.ServerAddr)
	//--------
	if serverPort == 0 {
		serverPort = TCPServerPort
	}
	//--------------------------------------------------
	wg := sync.WaitGroup{}
	//--------
	wg.Add(1)
	//--------------------------------------------------
	go func() {
		//--------------------------------------------------
		wg.Done()
		//--------
		err := networkObject.TCPServerEcho()
		//--------
		if err != nil {
			t.Error(err)
		}
		//--------------------------------------------------
	}()
	//--------------------------------------------------
	wg.Wait()
	//----------
	time.Sleep(10 * time.Millisecond)
	//--------------------------------------------------
	requestString := "TCPServerEcho test"
	//----------
	TCPConn, err := net.Dial("tcp4", fmt.Sprintf("%s:%d", serverIPAddr, serverPort))
	//--------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		//----------
		defer TCPConn.Close()
		//----------
		_, err := TCPConn.Write([]byte(requestString))
		//----------
		responseBytes, _ := io.ReadAll(TCPConn)
		//----------
		if err != nil {
			t.Error(err)
		} else {
			//----------
			responseString := string(responseBytes)
			//----------
			if responseString != requestString {
				//----------
				t.Errorf("response = %q but should = %q", responseString, requestString)
				//----------
			}
			//----------
		}
		//----------
	}
	//--------------------------------------------------
}

//------------------------------------------------------------

func TestTCPListen(t *testing.T) {

	//--------------------------------------------------
	// TCPListen
	//--------------------------------------------------
	var err error
	//--------------------------------------------------
	serverIPAddr := "127.0.0.1"
	//--------------------------------------------------
	networkObject := NetworkStruct{ServerAddr: serverIPAddr}
	//--------------------------------------------------
	err = networkObject.TCPListen()
	//----------
	if err != nil {
		t.Error(err)
	} else {
		//----------
		defer networkObject.TCPListener.Close()
		//----------
		listenerType := fmt.Sprintf("%T", networkObject.TCPListener)
		//----------
		if listenerType != "*net.TCPListener" {
			//----------
			t.Errorf("TCPListen should create listener type %q but it created type = %q", "*net.TCPListener", listenerType)
			//----------
		}
		//----------
	}
	//--------------------------------------------------
}

func TestTCPListenConn(t *testing.T) {

	//--------------------------------------------------
	// TCPListenConn
	//--------------------------------------------------
	serverIPAddr := "127.0.0.1"
	//--------------------------------------------------
	networkObject := NetworkStruct{ServerAddr: serverIPAddr}
	//--------
	serverIPAddr, serverPort := SplitAddrPort(networkObject.ServerAddr)
	//--------
	if serverPort == 0 {
		serverPort = TCPServerPort
	}
	//--------------------------------------------------
	wg := sync.WaitGroup{}
	//--------
	wg.Add(1)
	//--------------------------------------------------
	go func() {
		//--------
		wg.Done()
		//--------
		err := networkObject.TCPListenConn()
		//--------
		if err != nil {
			t.Error(err)
		} else {
			//--------
			defer networkObject.TCPConn.Close()
			//--------
			requestBytes := make([]byte, BufferSize)
			//--------
			n, err := networkObject.TCPConn.Read(requestBytes)
			//--------
			if err == nil && n > 0 {
				_, _ = networkObject.TCPConn.Write(requestBytes[0:n])
			}
			//--------
		}
		//--------
	}()
	//--------------------------------------------------
	wg.Wait()
	//----------
	time.Sleep(10 * time.Millisecond)
	//--------------------------------------------------
	requestString := "TCPListenConn test"
	//----------
	responseString := ""
	//----------
	TCPConn, err := net.Dial("tcp4", fmt.Sprintf("%s:%d", serverIPAddr, serverPort))
	//--------------------------------------------------
	if err == nil {
		//----------
		defer TCPConn.Close()
		//----------
		_, err = TCPConn.Write([]byte(requestString))
		//----------
		if err == nil {
			responseBytes, err := io.ReadAll(TCPConn)
			if err == nil {
				responseString = string(responseBytes)
			}
		}
		//----------
	}
	//--------------------------------------------------
	if err != nil {
		t.Error("error connecting to server:", err)
	} else {
		//----------
		if responseString != requestString {
			//----------
			t.Errorf("responseString should = %q but = %q", requestString, responseString)
			//----------
		}
		//----------
	}
	//--------------------------------------------------
}

func TestTCPReadBytesTCPWriteBytes(t *testing.T) {

	//--------------------------------------------------
	// TCPReadBytes / TCPWriteBytes
	//--------------------------------------------------
	serverIPAddr := "127.0.0.1"
	//--------
	networkObject := NetworkStruct{ServerAddr: serverIPAddr}
	//--------
	serverIPAddr, serverPort := SplitAddrPort(networkObject.ServerAddr)
	//--------
	if serverPort == 0 {
		serverPort = TCPServerPort
	}
	//--------------------------------------------------
	wg := sync.WaitGroup{}
	//--------
	wg.Add(1)
	//--------------------------------------------------
	go func() {
		//--------------------------------------------------
		wg.Done()
		//--------------------------------------------------
		var err error
		var requestBytes []byte
		//--------------------------------------------------
		networkObject.TCPListener, err = net.Listen("tcp4", fmt.Sprintf("%s:%d", serverIPAddr, serverPort))
		//--------
		if err != nil {
			t.Error(err)
		} else {
			//--------------------------------------------------
			defer networkObject.TCPListener.Close()
			//--------------------------------------------------
			networkObject.TCPConn, err = networkObject.TCPListener.Accept()
			//--------
			if err != nil {
				t.Error(err)
			} else {
				//--------
				defer networkObject.TCPConn.Close()
				//--------
				requestBytes, err = networkObject.TCPReadBytes()
				//--------
				if err == nil {
					_ = networkObject.TCPWriteBytes(requestBytes)
				}
				//--------
			}
			//--------------------------------------------------
		}
		//--------------------------------------------------
	}()
	//--------------------------------------------------
	wg.Wait()
	//----------
	time.Sleep(10 * time.Millisecond)
	//--------------------------------------------------
	requestString := "TCPReadBytes / TCPWriteBytes test"
	//----------
	responseString := ""
	//----------
	TCPConn, err := net.Dial("tcp4", fmt.Sprintf("%s:%d", serverIPAddr, serverPort))
	//--------------------------------------------------
	if err == nil {
		//----------
		defer TCPConn.Close()
		//----------
		_, err = TCPConn.Write([]byte(requestString))
		//----------
		if err == nil {
			responseBytes, err := io.ReadAll(TCPConn)
			if err == nil {
				responseString = string(responseBytes)
			}
		}
		//----------
		if responseString != requestString {
			//----------
			t.Errorf("responseString should = %q but = %q", requestString, responseString)
			//----------
		}
		//----------
	}
	//--------------------------------------------------
}

//------------------------------------------------------------

func TestTCPCient(t *testing.T) {

	//--------------------------------------------------
	// TCPCient
	//--------------------------------------------------
	serverIPAddr := "127.0.0.1"
	//--------
	networkObject := NetworkStruct{ServerAddr: serverIPAddr}
	//--------
	serverIPAddr, serverPort := SplitAddrPort(networkObject.ServerAddr)
	//--------
	if serverPort == 0 {
		serverPort = TCPServerPort
	}
	//--------------------------------------------------
	wg := sync.WaitGroup{}
	//--------
	wg.Add(1)
	//--------------------------------------------------
	go func() {
		//--------------------------------------------------
		wg.Done()
		//--------
		TCPListener, _ := net.Listen("tcp4", fmt.Sprintf("%s:%d", serverIPAddr, serverPort))
		//--------
		TCPConn, _ := TCPListener.Accept()
		//--------
		defer TCPListener.Close()
		defer TCPConn.Close()
		//--------
		requestBytes := make([]byte, BufferSize)
		//--------
		n, err := TCPConn.Read(requestBytes)
		//--------
		if err == nil && n > 0 {
			_, _ = TCPConn.Write(requestBytes[0:n])
		}
		//--------------------------------------------------
	}()
	//--------------------------------------------------
	wg.Wait()
	//----------
	time.Sleep(10 * time.Millisecond)
	//--------------------------------------------------
	requestString := "TCPClient test"
	//----------
	responseBytes, err := networkObject.TCPClient([]byte(requestString))
	//--------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		//----------
		responseString := string(responseBytes)
		//----------
		if responseString != requestString {
			//----------
			t.Errorf("responseString should = %q but = %q", requestString, responseString)
			//----------
		}
		//----------
	}
	//--------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

func TestUDPServerEcho(t *testing.T) {

	//--------------------------------------------------
	// UDPServerEcho
	//--------------------------------------------------
	serverIPAddr := "127.0.0.1"
	serverPort := UDPServerPort
	//--------
	clientIPAddr := "127.0.0.1"
	clientPort := UDPClientPort
	//--------------------------------------------------
	networkObject := NetworkStruct{ServerAddr: fmt.Sprintf("%s:%d", serverIPAddr, serverPort)}
	//--------------------------------------------------
	wg := sync.WaitGroup{}
	//--------
	wg.Add(1)
	//--------------------------------------------------
	go func() {
		//--------
		wg.Done()
		//--------
		err := networkObject.UDPServerEcho()
		//--------
		if err != nil {
			t.Error(err)
		}
		//--------
	}()
	//--------------------------------------------------
	wg.Wait()
	//----------
	time.Sleep(10 * time.Millisecond)
	//--------------------------------------------------
	var err error
	var UDPConn net.PacketConn
	var n int
	//--------------------------------------------------
	requestString := "UDPServerEcho test"
	//----------
	UDPConn, err = net.ListenPacket("udp4", fmt.Sprintf("%s:%d", clientIPAddr, clientPort))
	//----------
	if err != nil {
		t.Error(err)
	} else {
		//----------
		remoteAddr, _ := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%d", serverIPAddr, serverPort))
		//----------
		_, err = UDPConn.WriteTo([]byte(requestString), remoteAddr)
		//--------------------------------------------------
		if err != nil {
			t.Error(err)
		} else {
			//----------
			defer UDPConn.Close()
			//----------
			responseBytes := make([]byte, BufferSize)
			//--------
			n, _, err = UDPConn.ReadFrom(responseBytes)
			//----------
			if err != nil {
				t.Error(err)
			} else {
				//----------
				responseString := string(responseBytes[0:n])
				//----------
				if responseString != requestString {
					//----------
					t.Errorf("response = %q but should = %q", responseString, requestString)
					//----------
				}
				//----------
			}
			//----------
		}
		//----------
	}
	//--------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

func TestUDPListen(t *testing.T) {

	//--------------------------------------------------
	// UDPListen
	//--------------------------------------------------
	var err error
	//--------------------------------------------------
	serverIPAddr := "127.0.0.1"
	//--------------------------------------------------
	networkObject := NetworkStruct{ServerAddr: serverIPAddr}
	//--------------------------------------------------
	err = networkObject.UDPListen()
	//--------------------------------------------------
	if err != nil {
		t.Error("error creating packet connection:", err)
	} else {
		//----------
		defer networkObject.UDPConn.Close()
		//----------
		connectionType := fmt.Sprintf("%T", networkObject.UDPConn)
		//----------
		if connectionType != "*net.UDPConn" {
			//----------
			t.Errorf("UDPListen should create connection type %q but it created type = %q", "*net.UDPConn", connectionType)
			//----------
		}
		//----------
	}
	//--------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

func TestUDPReadBytes(t *testing.T) {

	//--------------------------------------------------
	// UDPReadBytes
	//--------------------------------------------------
	var err error
	var UDPConn net.PacketConn
	var n int
	var bytes []byte
	//--------------------------------------------------
	serverAddr := fmt.Sprintf("127.0.0.1:%d", UDPServerPort)
	message := "test message to read"
	messageBytes := []byte(message)
	//--------------------------------------------------
	UDPConn, err = net.ListenPacket("udp4", serverAddr)
	//--------------------------------------------------
	if err != nil {
		t.Error("error creating packet connection:", err)
	} else {

		//--------------------------------------------------
		defer UDPConn.Close()
		//--------------------------------------------------
		addr, _ := net.ResolveUDPAddr("udp4", serverAddr)
		//--------------------------------------------------
		n, err = UDPConn.WriteTo(messageBytes, addr)
		//--------------------------------------------------
		if err != nil {
			t.Error("error writing bytes for test:", err)
		} else {

			//--------------------------------------------------
			networkObject := NetworkStruct{ServerAddr: serverAddr, UDPConn: UDPConn}
			//--------------------------------------------------
			bytes, err = networkObject.UDPReadBytes()
			//--------------------------------------------------
			if err != nil {
				t.Error("error reading bytes:", err)
			} else {

				//----------
				resultString := string(bytes[0:n])
				length := len(resultString)
				//----------
				if length != n {
					//----------
					t.Errorf("returned bytes length = %d but should = %d", length, n)
					//----------
				}
				//----------
				if fmt.Sprint(bytes) != fmt.Sprint(messageBytes) {
					//----------
					t.Errorf("returned bytes = %v but should = %v", bytes, messageBytes)
					//----------
				}
				//----------
			}
		}
	}
	//--------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

func TestUDPWriteBytes(t *testing.T) {

	//--------------------------------------------------
	// UDPWriteBytes
	//--------------------------------------------------
	var err error
	var UDPConn net.PacketConn
	var n int
	var remoteAddr net.Addr
	var bytes []byte
	//--------------------------------------------------
	serverAddr := fmt.Sprintf("127.0.0.1:%d", UDPServerPort)
	message := "test message to write"
	messageBytes := []byte(message)
	bytesLength := len(messageBytes)
	//--------------------------------------------------
	UDPConn, err = net.ListenPacket("udp4", serverAddr)
	//--------------------------------------------------
	if err != nil {
		t.Error("error creating packet connection:", err)
	} else {

		//--------------------------------------------------
		defer UDPConn.Close()
		//--------------------------------------------------
		addr, _ := net.ResolveUDPAddr("udp4", serverAddr)
		//--------------------------------------------------
		networkObject := NetworkStruct{ServerAddr: serverAddr, UDPConn: UDPConn, RemoteAddr: addr}
		//--------------------------------------------------
		err = networkObject.UDPWriteBytes(messageBytes)
		//--------------------------------------------------
		if err != nil {
			t.Error("error writing bytes:", err)
		} else {

			//--------------------------------------------------
			bytes = make([]byte, BufferSize)
			//----------
			n, remoteAddr, err = UDPConn.ReadFrom(bytes)
			//--------------------------------------------------
			if err != nil {
				t.Error("error reading written bytes:", err)
			} else {

				//----------
				bytes = bytes[0:n]
				//----------
				if n != bytesLength {
					//----------
					t.Errorf("returned bytes length = %d but should = %d", n, bytesLength)
					//----------
				}
				//----------
				if fmt.Sprint(bytes) != fmt.Sprint(messageBytes) {
					//----------
					t.Errorf("returned bytes = %v but should = %v", bytes, messageBytes)
					//----------
				}
				//----------
				if fmt.Sprint(remoteAddr) != fmt.Sprint(addr) {
					//----------
					t.Errorf("returned remote address = %q but should = %q", remoteAddr, addr)
					//----------
				}
				//----------
			}
		}
	}
	//--------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

func TestUDPClient(t *testing.T) {

	//--------------------------------------------------
	// UDPClient
	//--------------------------------------------------
	serverIPAddr := "127.0.0.1"
	serverPort := UDPServerPort
	//--------
	clientIPAddr := "127.0.0.1"
	clientPort := UDPClientPort
	//--------------------------------------------------
	networkObject := NetworkStruct{ServerAddr: fmt.Sprintf("%s:%d", serverIPAddr, serverPort), ClientAddr: fmt.Sprintf("%s:%d", clientIPAddr, clientPort)}
	//--------------------------------------------------
	wg := sync.WaitGroup{}
	//--------
	wg.Add(1)
	//--------------------------------------------------
	go func() {
		//--------
		wg.Done()
		//--------
		UDPConn, err := net.ListenPacket("udp4", fmt.Sprintf("%s:%d", serverIPAddr, serverPort))
		//----------
		if err == nil {
			//----------
			defer UDPConn.Close()
			//----------
			requestBytes := make([]byte, BufferSize)
			//--------
			n, remoteAddr, err := UDPConn.ReadFrom(requestBytes)
			//--------------------------------------------------
			if err == nil && n > 0 {
				//----------
				_, err = UDPConn.WriteTo(requestBytes[0:n], remoteAddr)
				//----------
			}
			//--------
			if err != nil {
				t.Error(err)
			}
			//--------
		}
		//--------
	}()
	//--------------------------------------------------
	wg.Wait()
	//----------
	time.Sleep(10 * time.Millisecond)
	//--------------------------------------------------
	requestString := "UDPClient test"
	//----------
	responseBytes, err := networkObject.UDPClient([]byte(requestString))
	//----------
	if err != nil {
		t.Error(err)
	} else {
		//----------
		responseString := string(responseBytes)
		//----------
		if responseString != requestString {
			//----------
			t.Errorf("response = %q but should = %q", responseString, requestString)
			//----------
		}
		//----------
	}
	//--------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------
