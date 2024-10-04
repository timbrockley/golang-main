/*

Copyright 2023-2024, Tim Brockley. All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

*/

package server

import (
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"syscall"

	"github.com/timbrockley/golang-main/file"
	"github.com/timbrockley/golang-main/rpc"
	"golang.org/x/exp/slices"
)

//------------------------------------------------------------

const host = ":3000"

const pathRoot = "/www/golang/main/html/"

//------------------------------------------------------------

var httpServer *http.Server

var stopChan chan os.Signal

//------------------------------------------------------------

const TCPServerPort = 4000

const UDPServerPort = 4001
const UDPClientPort = 4002

const BufferSize = 1024

//----------------------------------------

type NetworkStruct struct {
	//--------------------
	ServerAddr string
	//--------------------
	ClientAddr string
	//--------------------
	TCPListener net.Listener
	TCPConn     net.Conn
	//--------------------
	UDPConn net.PacketConn
	//--------------------
	RemoteAddr net.Addr
	//--------------------
	BufferSize int
	//--------------------
	KeepAlive bool
	//--------------------
}

//----------------------------------------

type SocketStruct struct {
	//--------------------
	Addr string
	//--------------------
	Listener net.Listener
	Conn     net.Conn
	//--------------------
	KeepAlive bool
	//--------------------
}

//------------------------------------------------------------

const contentTypeText string = "text/plain; charset=UTF-8"
const contentTypeTextHTML string = "text/html; charset=UTF-8"
const contentTypeTextCSS string = "text/css; charset=UTF-8"
const contentTypeJavascript string = "application/javascript; charset=UTF-8"

//------------------------------------------------------------

var flags = map[string]any{}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

func init() {

	//--------------------------------------------------
	// set debug / test flags
	//--------------------------------------------------
	debugTrue := []string{"--debug", "--debug=true", "debug", "debug=true"}
	debugFalse := []string{"--debug=false", "debug=false"}
	//--------------------
	for i := 1; i < len(os.Args); i++ {
		//--------------------
		if strings.HasPrefix(os.Args[i], "--debug") || strings.HasPrefix(os.Args[i], "debug") {
			//--------------------
			if slices.Contains(debugTrue, os.Args[i]) {
				flags["debug"] = true
			} else if slices.Contains(debugFalse, os.Args[i]) {
				flags["debug"] = false
			}
			//--------------------
		}
		//--------------------
		if os.Args[i] == "-test.v=true" {
			flags["test"] = true
		}
		//--------------------
	}
	//--------------------------------------------------
	// check if environment variable set
	if flags["debug"] == nil && strings.EqualFold(os.Getenv("GOLANG_DEBUG"), "true") {
		flags["debug"] = true
	}
	//--------------------------------------------------
	// default debug = true if not testing or already set true or false
	if flags["test"] != true && flags["debug"] == nil {
		flags["debug"] = true
	}
	//--------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

func StartServer() {

	//--------------------------------------------------
	http.HandleFunc("/", servePath)
	//----------------------------------------
	http.HandleFunc("/echo", Echo)
	http.HandleFunc("/headers", Headers)
	//--------------------------------------------------
	http.HandleFunc("/rpc", rpc.RPC_Handler)
	//--------------------------------------------------
	http.HandleFunc("/jsonrpc", rpc.JSONRPC_Handler)
	//--------------------------------------------------
	log.Println("starting server")
	//--------------------------------------------------
	httpServer = &http.Server{Addr: host}
	//--------------------------------------------------
	stopChan = make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
	//--------------------------------------------------
	go func() {
		//----------------------------------------
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %s", err)
		}
		//----------------------------------------
	}()
	//--------------------------------------------------
	// Wait until we get a stop signal or channel closes
	<-stopChan
	//--------------------------------------------------
	if err := httpServer.Shutdown(context.Background()); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}
	//--------------------------------------------------
	log.Println("Shutting down server")
	//--------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

//--------------------------------------------------------------------------------
// servePath
//--------------------------------------------------------------------------------

func servePath(responseWriter http.ResponseWriter, httpRequest *http.Request) {

	//--------------------------------------------------
	var path, ext, contentType string
	//--------------------------------------------------
	path = strings.TrimRight(file.PathJoin(pathRoot, httpRequest.URL.String()), "/")
	//--------------------
	if file.FilePathExists(path) {

		//--------------------
		ext = file.FilenameExt(path)
		//--------------------
		if ext == "html" || ext == "css" || ext == "js" {
			//--------------------
			if ext == "html" {
				contentType = contentTypeTextHTML
			} else if ext == "css" {
				contentType = contentTypeTextCSS
			} else if ext == "js" {
				contentType = contentTypeJavascript
			}
			//--------------------
			responseWriter.Header().Set("Content-Type", contentType)
			//--------------------
		}
		//--------------------
		// fmt.Println("path:", path)
		//--------------------
		http.ServeFile(responseWriter, httpRequest, path)
		//--------------------
	} else {
		//--------------------
		fmt.Fprint(responseWriter, `<html><body>404 Page not found</body></html>`)
		//--------------------
	}
	//--------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

func Echo(responseWriter http.ResponseWriter, httpRequest *http.Request) {

	//--------------------------------------------------
	requestBytes, err := io.ReadAll(httpRequest.Body)
	if err != nil {
		log.Println(err)
	}
	//--------------------------------------------------
	contentType := httpRequest.Header.Get("Content-Type")
	if contentType != "" {
		responseWriter.Header().Set("Content-Type", contentType)
	}
	//--------------------------------------------------
	REGEXP := regexp.MustCompile(`(?i)^X`)
	for name, values := range httpRequest.Header {
		//--------------------
		if REGEXP.FindString(name) != "" {
			//--------------------
			responseWriter.Header().Set(name, strings.Join(values, ", "))
			//--------------------
			if strings.EqualFold(name, "X-Debug") && strings.EqualFold(values[0], "true") {
				//--------------------
				fmt.Println("\n" + strings.Repeat("-", 40))
				fmt.Println(httpRequest.Method, httpRequest.URL.String(), httpRequest.Proto)
				fmt.Println(strings.Repeat("-", 40))
				fmt.Println(string(requestBytes))
				fmt.Println(strings.Repeat("-", 40) + "\n")
				//--------------------
			}
			//--------------------
		}
		//--------------------
	}
	//--------------------------------------------------
	fmt.Fprintf(responseWriter, "%s", string(requestBytes))
	//--------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

func Headers(responseWriter http.ResponseWriter, httpRequest *http.Request) {

	//--------------------------------------------------
	var headers []string
	//--------------------------------------------------
	for name, values := range httpRequest.Header {
		for _, value := range values {
			headers = append(headers, fmt.Sprintf("%v: %v", name, value))
		}
	}
	//--------------------------------------------------
	sort.Strings(headers)
	//--------------------------------------------------
	responseWriter.Header().Set("Content-Type", contentTypeText)
	//--------------------------------------------------
	ipAddr := rpc.GetRemoteIPAddr(httpRequest)
	//--------------------
	fmt.Fprintf(responseWriter, "Remote IP Address: %v\n\n", ipAddr)
	//--------------------------------------------------
	for _, header := range headers {
		fmt.Fprintf(responseWriter, "%s\n", header)
	}
	//--------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

func SplitAddrPort(addr string) (string, int) {
	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	var ipAddr string = ""
	var port int = 0
	//------------------------------------------------------------
	colonIndex := strings.Index(addr, ":")
	if colonIndex >= 0 {
		//--------------------
		addrSplit := strings.Split(addr, ":")
		//--------------------
		if fmt.Sprintf("%T", addrSplit) == "[]string" && fmt.Sprintf("%T", addrSplit[0]) == "string" {
			ipAddr = addrSplit[0]
		}
		//--------------------
		if fmt.Sprintf("%T", addrSplit) == "[]string" && fmt.Sprintf("%T", addrSplit[1]) == "string" {
			port, err = strconv.Atoi(addrSplit[1])
			if err != nil {
				port = 0
			}
		}
		//--------------------
	} else {
		//--------------------
		ipAddr = addr
		//--------------------
	}
	//------------------------------------------------------------
	return ipAddr, port
	//------------------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

func (networkObject *NetworkStruct) TCPServerEcho() error {
	//--------------------------------------------------
	var err error
	var requestBytes []byte
	//--------------------------------------------------
	serverIPAddr, serverPort := SplitAddrPort(networkObject.ServerAddr)
	//--------------------
	if serverPort == 0 {
		serverPort = TCPServerPort
	}
	//--------------------------------------------------
	ipAddr := fmt.Sprintf("%s:%d", serverIPAddr, serverPort)
	//--------------------------------------------------
	networkObject.TCPListener, err = net.Listen("tcp4", ipAddr)
	//--------------------------------------------------
	if err == nil {
		//--------------------------------------------------
		if flags["debug"] == true {
			fmt.Printf("starting server: (%s) ...\n", ipAddr)
		}
		//--------------------------------------------------
		defer networkObject.TCPListener.Close()
		//--------------------------------------------------
		for {
			//--------------------------------------------------
			networkObject.TCPConn, err = networkObject.TCPListener.Accept()
			//--------------------
			if err != nil {

				break

			} else {

				//--------------------------------------------------
				go func(conn net.Conn) {
					//--------------------------------------------------
					defer conn.Close()
					//--------------------------------------------------
					requestBytes, err = io.ReadAll(conn)
					//--------------------------------------------------
					if flags["debug"] == true {
						fmt.Printf("(%s): %s\n", conn.RemoteAddr(), string(requestBytes))
					}
					//--------------------------------------------------
					conn.Write(requestBytes)
					//--------------------------------------------------
				}(networkObject.TCPConn)
				//--------------------------------------------------
				if !networkObject.KeepAlive {
					break
				}
				//--------------------------------------------------
			}
			//--------------------------------------------------
		}
		//--------------------------------------------------
	}
	//--------------------------------------------------
	return err
	//--------------------------------------------------
}

//------------------------------------------------------------

func (networkObject *NetworkStruct) TCPClient(requestBytes []byte) ([]byte, error) {
	//--------------------------------------------------
	var err error
	var responseBytes []byte
	//--------------------------------------------------
	serverIPAddr, serverPort := SplitAddrPort(networkObject.ServerAddr)
	//--------------------
	if serverPort == 0 {
		serverPort = TCPServerPort
	}
	//--------------------------------------------------
	networkObject.TCPConn, err = net.Dial("tcp4", fmt.Sprintf("%s:%d", serverIPAddr, serverPort))
	//--------------------------------------------------
	if err == nil {
		//--------------------
		defer networkObject.TCPConn.Close()
		//--------------------
		_, err = networkObject.TCPConn.Write(requestBytes)
		//--------------------
		if err == nil {
			//--------------------
			err = networkObject.TCPConn.(*net.TCPConn).CloseWrite()
			//--------------------
			if err == nil {
				responseBytes, _ = io.ReadAll(networkObject.TCPConn)
			}
			//--------------------
		}
		//--------------------
	}
	//--------------------------------------------------
	return responseBytes, err
	//--------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

func (networkObject *NetworkStruct) UDPServerEcho() error {
	//--------------------------------------------------
	/*
		does not use a go routine as this causes issues when trying to process connections
		all testing without fixed buffer lengths failed with deadlock

		TCP Sockets and Unix Domain Sockets both work with go routines
		and don't required a fixed buffer length
	*/
	//--------------------------------------------------
	var err error
	var bytesRead int
	//--------------------------------------------------
	serverIPAddr, serverPort := SplitAddrPort(networkObject.ServerAddr)
	//--------------------
	if serverPort == 0 {
		serverPort = UDPServerPort
	}
	//--------------------------------------------------
	ipAddr := fmt.Sprintf("%s:%d", serverIPAddr, serverPort)
	//--------------------------------------------------
	if flags["debug"] == true {
		fmt.Printf("starting server: (%s) ...\n", ipAddr)
	}
	//--------------------------------------------------
	for {
		//--------------------------------------------------
		networkObject.UDPConn, err = net.ListenPacket("udp4", fmt.Sprintf("%s:%d", serverIPAddr, serverPort))
		//--------------------------------------------------
		if err != nil {

			break

		} else {

			//--------------------------------------------------
			// fixed buffer size used as can only read 1 udp packet per connection
			//--------------------------------------------------
			bufferSize := BufferSize
			//--------------------
			if networkObject.BufferSize > 0 {
				bufferSize = networkObject.BufferSize
			}
			//--------------------------------------------------
			requestBytes := make([]byte, bufferSize)
			//--------------------
			bytesRead, networkObject.RemoteAddr, err = networkObject.UDPConn.ReadFrom(requestBytes)
			//--------------------
			if flags["debug"] == true {
				fmt.Printf("(%s): %s\n", networkObject.RemoteAddr, string(requestBytes[0:bytesRead]))
			}
			//--------------------------------------------------
			if bytesRead > 0 {
				//--------------------
				_, err = networkObject.UDPConn.WriteTo(requestBytes[0:bytesRead], networkObject.RemoteAddr)
				//--------------------
			}
			//--------------------------------------------------
			networkObject.UDPConn.Close()
			//--------------------------------------------------
			if !networkObject.KeepAlive {
				break
			}
			//--------------------------------------------------
		}
		//--------------------------------------------------
	}
	//--------------------------------------------------
	return err
	//--------------------------------------------------
}

//------------------------------------------------------------

func (networkObject *NetworkStruct) UDPClient(requestBytes []byte) ([]byte, error) {
	//--------------------------------------------------
	var err error
	var responseBytes []byte
	var bytesRead int
	//--------------------------------------------------
	serverIPAddr, serverPort := SplitAddrPort(networkObject.ServerAddr)
	//--------------------
	if serverPort == 0 {
		serverPort = UDPServerPort
	}
	//--------------------------------------------------
	clientIPAddr, clientPort := SplitAddrPort(networkObject.ClientAddr)
	//--------------------
	if clientPort == 0 {
		clientPort = UDPClientPort
	}
	//--------------------------------------------------
	networkObject.UDPConn, err = net.ListenPacket("udp4", fmt.Sprintf("%s:%d", clientIPAddr, clientPort))
	//--------------------
	if err == nil {
		//--------------------------------------------------
		remoteAddr, _ := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%d", serverIPAddr, serverPort))
		//--------------------------------------------------
		_, err = networkObject.UDPConn.WriteTo(requestBytes, remoteAddr)
		//--------------------------------------------------
		if err == nil {
			//--------------------
			defer networkObject.UDPConn.Close()
			//--------------------------------------------------
			// fixed buffer size used as can only read 1 udp packet per connection
			//--------------------------------------------------
			bufferSize := BufferSize
			//--------------------------------------------------
			if networkObject.BufferSize > 0 {
				bufferSize = networkObject.BufferSize
			}
			//--------------------------------------------------
			responseBytes = make([]byte, bufferSize)
			//--------------------------------------------------
			bytesRead, _, err = networkObject.UDPConn.ReadFrom(responseBytes)
			//--------------------------------------------------
			if bytesRead > 0 {
				responseBytes = responseBytes[0:bytesRead]
			}
			//--------------------------------------------------
		}
		//--------------------------------------------------
	}
	//--------------------------------------------------
	return responseBytes, err
	//--------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

func (socketObject *SocketStruct) SocketServerEcho() error {
	//--------------------------------------------------
	var err error
	var baseHeaderBytes, extendedbaseHeaderBytes, requestBytes []byte
	var baseHeaderLength, extendedHeaderLength, headerLength uint8
	var bodyLength uint32
	var bytesRead int
	//--------------------------------------------------
	_ = os.Remove(socketObject.Addr)
	//--------------------------------------------------
	socketObject.Listener, err = net.Listen("unix", socketObject.Addr)
	//--------------------------------------------------
	if err == nil {
		//--------------------------------------------------
		if flags["debug"] == true {
			fmt.Printf("starting server: (%s) ...\n", socketObject.Addr)
		}
		//--------------------------------------------------
		defer socketObject.Listener.Close()
		//--------------------------------------------------
		for {
			//--------------------------------------------------
			socketObject.Conn, err = socketObject.Listener.Accept()
			//--------------------
			if err != nil {

				break

			} else {

				//--------------------------------------------------
				go func(conn net.Conn) {
					//--------------------------------------------------
					defer conn.Close()
					//--------------------------------------------------
					// baseHeaderLength = header length (1 byte) + body length (4 bytes)
					// extendedHeaderLength = optional extended header (headerLength - baseHeaderLength)
					baseHeaderLength = 5
					//--------------------
					baseHeaderBytes = make([]byte, baseHeaderLength)
					_, err = conn.Read(baseHeaderBytes)
					if err == nil {
						//--------------------
						headerLength = uint8(baseHeaderBytes[0])
						bodyLength = binary.BigEndian.Uint32(baseHeaderBytes[1:5])
						//--------------------
						extendedHeaderLength = headerLength - baseHeaderLength
						if extendedHeaderLength > 0 {
							extendedbaseHeaderBytes = make([]byte, extendedHeaderLength)
							_, err = conn.Read(extendedbaseHeaderBytes)
						}
						//--------------------
						if err == nil {
							//--------------------
							requestBytes = make([]byte, bodyLength)
							bytesRead, err = conn.Read(requestBytes)
							//--------------------
							if bytesRead > 0 {
								//--------------------
								if flags["debug"] == true {
									fmt.Printf("client request: %s\n", string(requestBytes))
								}
								//--------------------
								conn.Write(requestBytes)
								//--------------------
							}
							//--------------------
						}
						//--------------------
					}
					//--------------------------------------------------
				}(socketObject.Conn)
				//--------------------------------------------------
				if !socketObject.KeepAlive {
					break
				}
				//--------------------------------------------------
			}
			//--------------------------------------------------
		}
		//--------------------------------------------------
	}
	//--------------------------------------------------
	return err
	//--------------------------------------------------
}

//------------------------------------------------------------

func (socketObject *SocketStruct) SocketClient(requestBytes []byte) ([]byte, error) {
	//--------------------------------------------------
	var err error
	var combinedRequestBytes, responseBytes []byte
	var headerLength uint8
	//--------------------------------------------------
	socketObject.Conn, err = net.Dial("unix", socketObject.Addr)
	//--------------------------------------------------
	if err == nil {
		//--------------------
		defer socketObject.Conn.Close()
		//--------------------
		headerLength = 5 // (base header length + optional extended header length)
		//--------------------
		combinedRequestBytes = make([]byte, int(headerLength)+len(requestBytes))
		combinedRequestBytes[0] = headerLength
		binary.BigEndian.PutUint32(combinedRequestBytes[1:5], uint32(len(requestBytes)))
		//--------------------
		/*
			create optional extended header here
		*/
		//--------------------
		copy(combinedRequestBytes[headerLength:], requestBytes)
		//--------------------
		_, err = socketObject.Conn.Write(combinedRequestBytes)
		//--------------------
		if err == nil {
			responseBytes, _ = io.ReadAll(socketObject.Conn)
		}
		//--------------------
	}
	//--------------------------------------------------
	return responseBytes, err
	//--------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------
