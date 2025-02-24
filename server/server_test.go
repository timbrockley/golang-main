//------------------------------------------------------------

package server

import (
	"bytes"
	"encoding/binary"
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
	//--------------------
	EXPECTED_contentType := "text/plain; charset=UTF-8"
	EXPECTED_responseString := "<ECHO_TEST>"
	//--------------------
	recorder := httptest.NewRecorder()
	//--------------------
	httpRequest := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(requestString)))
	//--------------------
	httpRequest.Header.Set("Content-Type", requestContentType)
	//--------------------

	//--------------------------------------------------
	Echo(recorder, httpRequest)
	//--------------------------------------------------

	//--------------------
	httpResponse := recorder.Result()
	//--------------------
	defer httpResponse.Body.Close()
	//--------------------
	responseBytes, err := io.ReadAll(httpResponse.Body)
	//--------------------
	if err != nil {
		t.Error(err)
	} else {
		//--------------------
		if string(responseBytes) != EXPECTED_responseString {
			t.Errorf("response %q expected %q", string(responseBytes), EXPECTED_responseString)
		}
		//--------------------
	}
	//--------------------
	responseContentType := httpRequest.Header.Get("Content-Type")
	//--------------------
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
	//--------------------
	defer httpResponse.Body.Close()
	//--------------------
	responseBytes, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		t.Error(err)
	}
	//--------------------------------------------------
	match, _ := regexp.MatchString("User-Agent", string(responseBytes))
	if !match {
		//--------------------
		t.Errorf("could not match test headers")
		//--------------------
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
		//--------------------
		t.Errorf("ipAddr should = %q but = %q", "127.0.0.1", ipAddr)
		//--------------------
	}
	//--------------------
	if port != 80 {
		//--------------------
		t.Errorf("port should = %d but = %d", 80, port)
		//--------------------
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
	networkInstance := NetworkStruct{ServerAddr: serverIPAddr}
	//--------
	serverIPAddr, serverPort := SplitAddrPort(networkInstance.ServerAddr)
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
		err := networkInstance.TCPServerEcho()
		//--------
		if err != nil {
			t.Error(err)
		}
		//--------------------------------------------------
	}()
	//--------------------------------------------------
	wg.Wait()
	//--------------------
	time.Sleep(10 * time.Millisecond)
	//--------------------------------------------------
	requestString := "TCPServerEcho test"
	//--------------------
	TCPConn, err := net.Dial("tcp4", fmt.Sprintf("%s:%d", serverIPAddr, serverPort))
	//--------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		//--------------------
		defer TCPConn.Close()
		//--------------------
		_, err := TCPConn.Write([]byte(requestString))
		//--------------------
		_ = TCPConn.(*net.TCPConn).CloseWrite()
		//--------------------
		if err != nil {
			t.Error(err)
		} else {
			//--------------------
			responseBytes, _ := io.ReadAll(TCPConn)
			//--------------------
			responseString := string(responseBytes)
			//--------------------
			if responseString != requestString {
				//--------------------
				t.Errorf("response = %q but should = %q", responseString, requestString)
				//--------------------
			}
			//--------------------
		}
		//--------------------
	}
	//--------------------------------------------------
}

//------------------------------------------------------------

func TestTCPClient(t *testing.T) {
	//--------------------------------------------------
	// TCPClient
	//--------------------------------------------------
	serverIPAddr := "127.0.0.1"
	//--------
	networkInstance := NetworkStruct{ServerAddr: serverIPAddr}
	//--------
	serverIPAddr, serverPort := SplitAddrPort(networkInstance.ServerAddr)
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
		requestBytes, err := io.ReadAll(TCPConn)
		//--------
		if err == nil {
			TCPConn.Write(requestBytes)
		}
		//--------------------------------------------------
	}()
	//--------------------------------------------------
	wg.Wait()
	//--------------------
	time.Sleep(10 * time.Millisecond)
	//--------------------------------------------------
	requestString := "TCPClient test"
	//--------------------
	responseBytes, err := networkInstance.TCPClient([]byte(requestString))
	//--------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		//--------------------
		responseString := string(responseBytes)
		//--------------------
		if responseString != requestString {
			//--------------------
			t.Errorf("responseString should = %q but = %q", requestString, responseString)
			//--------------------
		}
		//--------------------
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
	networkInstance := NetworkStruct{ServerAddr: fmt.Sprintf("%s:%d", serverIPAddr, serverPort)}
	//--------------------------------------------------
	wg := sync.WaitGroup{}
	//--------
	wg.Add(1)
	//--------------------------------------------------
	go func() {
		//--------
		wg.Done()
		//--------
		err := networkInstance.UDPServerEcho()
		//--------
		if err != nil {
			t.Error(err)
		}
		//--------
	}()
	//--------------------------------------------------
	wg.Wait()
	//--------------------
	time.Sleep(10 * time.Millisecond)
	//--------------------------------------------------
	var err error
	var UDPConn net.PacketConn
	var n int
	//--------------------------------------------------
	requestString := "UDPServerEcho test"
	//--------------------
	UDPConn, err = net.ListenPacket("udp4", fmt.Sprintf("%s:%d", clientIPAddr, clientPort))
	//--------------------
	if err != nil {
		t.Error(err)
	} else {
		//--------------------
		remoteAddr, _ := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%d", serverIPAddr, serverPort))
		//--------------------
		_, err = UDPConn.WriteTo([]byte(requestString), remoteAddr)
		//--------------------------------------------------
		if err != nil {
			t.Error(err)
		} else {
			//--------------------
			defer UDPConn.Close()
			//--------------------
			responseBytes := make([]byte, BufferSize)
			//--------
			n, _, err = UDPConn.ReadFrom(responseBytes)
			//--------------------
			if err != nil {
				t.Error(err)
			} else {
				//--------------------
				responseString := string(responseBytes[0:n])
				//--------------------
				if responseString != requestString {
					//--------------------
					t.Errorf("response = %q but should = %q", responseString, requestString)
					//--------------------
				}
				//--------------------
			}
			//--------------------
		}
		//--------------------
	}
	//--------------------------------------------------
}

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
	networkInstance := NetworkStruct{ServerAddr: fmt.Sprintf("%s:%d", serverIPAddr, serverPort), ClientAddr: fmt.Sprintf("%s:%d", clientIPAddr, clientPort)}
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
		//--------------------
		if err == nil {
			//--------------------
			defer UDPConn.Close()
			//--------------------
			requestBytes := make([]byte, BufferSize)
			//--------
			n, remoteAddr, err := UDPConn.ReadFrom(requestBytes)
			//--------------------------------------------------
			if err == nil && n > 0 {
				//--------------------
				_, err = UDPConn.WriteTo(requestBytes[0:n], remoteAddr)
				//--------------------
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
	//--------------------
	time.Sleep(10 * time.Millisecond)
	//--------------------------------------------------
	requestString := "UDPClient test"
	//--------------------
	responseBytes, err := networkInstance.UDPClient([]byte(requestString))
	//--------------------
	if err != nil {
		t.Error(err)
	} else {
		//--------------------
		responseString := string(responseBytes)
		//--------------------
		if responseString != requestString {
			//--------------------
			t.Errorf("response = %q but should = %q", responseString, requestString)
			//--------------------
		}
		//--------------------
	}
	//--------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

func TestSocketServerEcho(t *testing.T) {
	//--------------------------------------------------
	// SocketServerEcho
	//--------------------------------------------------
	socketAddr := "golang-socket-test.sock"
	//--------
	socketInstance := SocketStruct{Addr: socketAddr}
	//--------------------------------------------------
	wg := sync.WaitGroup{}
	//--------
	wg.Add(1)
	//--------------------------------------------------
	go func() {
		//--------------------------------------------------
		wg.Done()
		//--------
		err := socketInstance.SocketServerEcho()
		//--------
		if err != nil {
			t.Error(err)
		}
		//--------------------------------------------------
	}()
	//--------------------------------------------------
	wg.Wait()
	//--------------------
	time.Sleep(100 * time.Millisecond)
	//--------------------------------------------------
	requestString := "SocketServerEcho test"
	requestBytes := []byte(requestString)
	//--------------------
	conn, err := net.Dial("unix", socketAddr)
	//--------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		//--------------------
		defer conn.Close()
		//--------------------
		headerLength := uint8(5) // (base header length + optional extended header length)
		//--------------------
		combinedRequestBytes := make([]byte, int(headerLength)+len(requestBytes))
		combinedRequestBytes[0] = headerLength
		binary.BigEndian.PutUint32(combinedRequestBytes[1:5], uint32(len(requestBytes)))
		//--------------------
		/*
			create optional extended header here
		*/
		//--------------------
		copy(combinedRequestBytes[headerLength:], requestBytes)
		//--------------------
		_, err = conn.Write(combinedRequestBytes)
		//--------------------
		responseBytes, _ := io.ReadAll(conn)
		//--------------------
		if err != nil {
			t.Error(err)
		} else {
			//--------------------
			responseString := string(responseBytes)
			//--------------------
			if responseString != requestString {
				//--------------------
				t.Errorf("response = %q but should = %q", responseString, requestString)
				//--------------------
			}
			//--------------------
		}
		//--------------------
	}
	//--------------------------------------------------
}

//------------------------------------------------------------

func TestSocketClient(t *testing.T) {
	//--------------------------------------------------
	// SocketClient
	//--------------------------------------------------
	socketAddr := "golang-socket-test.sock"
	//--------
	socketInstance := SocketStruct{Addr: socketAddr}
	//--------------------------------------------------
	wg := sync.WaitGroup{}
	//--------
	wg.Add(1)
	//--------------------------------------------------
	go func() {
		//--------------------------------------------------
		wg.Done()
		//--------
		socketListener, _ := net.Listen("unix", socketAddr)
		//--------
		conn, _ := socketListener.Accept()
		//--------
		defer socketListener.Close()
		defer conn.Close()
		//--------
		requestBytes := make([]byte, 4096)
		//--------
		n, err := conn.Read(requestBytes)
		//--------
		if err == nil && n > 0 {
			conn.Write(requestBytes[5:n]) // ignore header
		}
		//--------------------------------------------------
	}()
	//--------------------------------------------------
	wg.Wait()
	//--------------------
	time.Sleep(10 * time.Millisecond)
	//--------------------------------------------------
	requestString := "SocketClient test"
	//--------------------
	responseBytes, err := socketInstance.SocketClient([]byte(requestString))
	//--------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		//--------------------
		responseString := string(responseBytes)
		//--------------------
		if responseString != requestString {
			//--------------------
			t.Errorf("responseString should = %q but = %q", requestString, responseString)
			//--------------------
		}
		//--------------------
	}
	//--------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------
