//------------------------------------------------------------

package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/timbrockley/golang-main/file"
	"github.com/timbrockley/golang-main/rpc"
	"golang.org/x/exp/slices"
)

//------------------------------------------------------------

const host = ":3000" // may be overriden in StartServer function

const pathRoot = "/www/golang/main/html/"

//------------------------------------------------------------

const TCPServerPort = 4000

const UDPServerPort = 4001
const UDPClientPort = 4002

const BufferSize = 256

//----------------------------------------

type NetworkStruct struct {
	//----------
	ServerAddr string
	//----------
	ClientAddr string
	//----------
	TCPListener net.Listener
	TCPConn     net.Conn
	//----------
	UDPConn net.PacketConn
	//----------
	RemoteAddr net.Addr
	//----------
	BytesRead    int
	BytesWritten int
	//----------
	KeepAlive bool
	//----------
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
	//----------
	for i := 1; i < len(os.Args); i++ {
		//----------
		if strings.HasPrefix(os.Args[i], "--debug") || strings.HasPrefix(os.Args[i], "debug") {
			//----------
			if slices.Contains(debugTrue, os.Args[i]) {
				flags["debug"] = true
			} else if slices.Contains(debugFalse, os.Args[i]) {
				flags["debug"] = false
			}
			//----------
		}
		//----------
		if os.Args[i] == "-test.v=true" {
			flags["test"] = true
		}
		//----------
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
	fmt.Println("starting server")
	//--------------------------------------------------
	executable, _ := os.Executable()
	//----------
	var HOST string
	//----------
	if file.Filename(executable) == "gin-bin" {
		HOST = ":3001"
	} else {
		HOST = host
	}
	//--------------------------------------------------
	log.Fatal(http.ListenAndServe(HOST, nil))
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
	//----------
	if file.FilePathExists(path) {

		//----------
		ext = file.FilenameExt(path)
		//----------
		if ext == "html" || ext == "css" || ext == "js" {
			//----------
			if ext == "html" {
				contentType = contentTypeTextHTML
			} else if ext == "css" {
				contentType = contentTypeTextCSS
			} else if ext == "js" {
				contentType = contentTypeJavascript
			}
			//----------
			responseWriter.Header().Set("Content-Type", contentType)
			//----------
		}
		//----------
		fmt.Println("path:", path)
		//----------
		http.ServeFile(responseWriter, httpRequest, path)
		//----------
	} else {
		//----------
		fmt.Fprint(responseWriter, `<html><body>404 Page not found</body></html>`)
		//----------
	}
	// --------------------------------------------------
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
		//----------
		if REGEXP.FindString(name) != "" {
			//----------
			responseWriter.Header().Set(name, strings.Join(values, ", "))
			//----------
			if strings.EqualFold(name, "X-Debug") && strings.EqualFold(values[0], "true") {
				//----------
				fmt.Println("\n" + strings.Repeat("-", 40))
				fmt.Println(httpRequest.Method, httpRequest.URL.String(), httpRequest.Proto)
				fmt.Println(strings.Repeat("-", 40))
				fmt.Println(string(requestBytes))
				fmt.Println(strings.Repeat("-", 40) + "\n")
				//----------
			}
			//----------
		}
		//----------
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
	//----------
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
		//----------
		addrSplit := strings.Split(addr, ":")
		//----------
		if fmt.Sprintf("%T", addrSplit) == "[]string" && fmt.Sprintf("%T", addrSplit[0]) == "string" {
			ipAddr = addrSplit[0]
		}
		//----------
		if fmt.Sprintf("%T", addrSplit) == "[]string" && fmt.Sprintf("%T", addrSplit[1]) == "string" {
			port, err = strconv.Atoi(addrSplit[1])
			if err != nil {
				port = 0
			}
		}
		//----------
	} else {
		//----------
		ipAddr = addr
		//----------
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
	//--------
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
			fmt.Printf("wating for connection (%s) ...\n", ipAddr)
		}
		//--------------------------------------------------
		defer networkObject.TCPListener.Close()
		//--------------------------------------------------
		for {
			//--------------------------------------------------
			networkObject.TCPConn, err = networkObject.TCPListener.Accept()
			//--------
			if err != nil {

				break

			} else {

				//--------------------------------------------------
				requestBytes = make([]byte, BufferSize)
				//--------------------------------------------------
				networkObject.BytesRead, err = networkObject.TCPConn.Read(requestBytes)
				//--------
				if flags["debug"] == true {
					fmt.Printf("(%s): %s\n", networkObject.TCPConn.RemoteAddr(), string(requestBytes[0:networkObject.BytesRead]))
				}
				//--------
				if err == nil && networkObject.BytesRead > 0 {
					//--------
					networkObject.BytesWritten, _ = networkObject.TCPConn.Write(requestBytes[0:networkObject.BytesRead])
					//--------
				}
				//--------------------------------------------------
				networkObject.TCPConn.Close()
				//--------------------------------------------------
				if !networkObject.KeepAlive {
					break
				}
				//--------------------------------------------------
			}
			//--------------------------------------------------
		}
		///--------------------------------------------------
	}
	//--------------------------------------------------
	return err
	//--------------------------------------------------
}

//------------------------------------------------------------

func (networkObject *NetworkStruct) TCPListen() error {
	//--------------------------------------------------
	var err error
	//--------------------------------------------------
	serverIPAddr, serverPort := SplitAddrPort(networkObject.ServerAddr)
	//--------
	if serverPort == 0 {
		serverPort = TCPServerPort
	}
	//--------------------------------------------------
	networkObject.TCPListener, err = net.Listen("tcp4", fmt.Sprintf("%s:%d", serverIPAddr, serverPort))
	//--------------------------------------------------
	return err
	//--------------------------------------------------
}

//------------------------------------------------------------

func (networkObject *NetworkStruct) TCPListenConn() error {
	//--------------------------------------------------
	var err error
	//--------------------------------------------------
	serverIPAddr, serverPort := SplitAddrPort(networkObject.ServerAddr)
	//--------
	if serverPort == 0 {
		serverPort = TCPServerPort
	}
	//--------------------------------------------------
	networkObject.TCPListener, err = net.Listen("tcp4", fmt.Sprintf("%s:%d", serverIPAddr, serverPort))
	//--------------------------------------------------
	if err == nil {
		//--------
		defer networkObject.TCPListener.Close()
		//--------
		networkObject.TCPConn, err = networkObject.TCPListener.Accept()
		//--------
	}
	//--------------------------------------------------
	return err
	//--------------------------------------------------
}

//------------------------------------------------------------

func (networkObject *NetworkStruct) TCPReadBytes() ([]byte, error) {
	//--------------------------------------------------
	var err error
	//--------------------------------------------------
	bytes := make([]byte, BufferSize)
	//--------
	networkObject.BytesRead, err = networkObject.TCPConn.Read(bytes)
	//--------
	if err == nil && networkObject.BytesRead > 0 {
		bytes = bytes[0:networkObject.BytesRead]
	}
	//--------------------------------------------------
	return bytes, err
	//--------------------------------------------------
}

//------------------------------------------------------------

func (networkObject *NetworkStruct) TCPWriteBytes(bytes []byte) error {
	//--------------------------------------------------
	var err error
	//--------------------------------------------------
	networkObject.BytesWritten, err = networkObject.TCPConn.Write(bytes)
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
	//--------
	if serverPort == 0 {
		serverPort = TCPServerPort
	}
	//--------------------------------------------------
	networkObject.TCPConn, err = net.Dial("tcp4", fmt.Sprintf("%s:%d", serverIPAddr, serverPort))
	//--------------------------------------------------
	if err == nil {
		//----------
		defer networkObject.TCPConn.Close()
		//----------
		networkObject.BytesWritten, err = networkObject.TCPConn.Write(requestBytes)
		//----------
		if err == nil {
			responseBytes, _ = io.ReadAll(networkObject.TCPConn)
		}
		//----------
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
	var err error
	//--------------------------------------------------
	serverIPAddr, serverPort := SplitAddrPort(networkObject.ServerAddr)
	//----------
	if serverPort == 0 {
		serverPort = UDPServerPort
	}
	//--------------------------------------------------
	ipAddr := fmt.Sprintf("%s:%d", serverIPAddr, serverPort)
	//--------------------------------------------------
	if flags["debug"] == true {
		fmt.Printf("wating for connection (%s) ...\n", ipAddr)
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
			requestBytes := make([]byte, BufferSize)
			//--------
			networkObject.BytesRead, networkObject.RemoteAddr, err = networkObject.UDPConn.ReadFrom(requestBytes)
			//--------
			if flags["debug"] == true {
				fmt.Printf("(%s): %s\n", networkObject.RemoteAddr, string(requestBytes[0:networkObject.BytesRead]))
			}
			//--------------------------------------------------
			if err == nil && networkObject.BytesRead > 0 {
				//--------
				networkObject.BytesWritten, err = networkObject.UDPConn.WriteTo(requestBytes[0:networkObject.BytesRead], networkObject.RemoteAddr)
				//--------
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

func (networkObject *NetworkStruct) UDPListen() error {
	//--------------------------------------------------
	var err error
	//--------------------------------------------------
	serverIPAddr, serverPort := SplitAddrPort(networkObject.ServerAddr)
	//----------
	if serverPort == 0 {
		serverPort = UDPServerPort
	}
	//--------------------------------------------------
	networkObject.UDPConn, err = net.ListenPacket("udp4", fmt.Sprintf("%s:%d", serverIPAddr, serverPort))
	//--------------------------------------------------
	return err
	//--------------------------------------------------
}

//------------------------------------------------------------

func (networkObject *NetworkStruct) UDPReadBytes() ([]byte, error) {
	//--------------------------------------------------
	var err error
	//--------------------------------------------------
	bytes := make([]byte, BufferSize)
	//--------
	networkObject.BytesRead, networkObject.RemoteAddr, err = networkObject.UDPConn.ReadFrom(bytes)
	//--------------------------------------------------
	if err == nil {
		bytes = bytes[0:networkObject.BytesRead]
	}
	//--------------------------------------------------
	return bytes, err
	//--------------------------------------------------
}

//------------------------------------------------------------

func (networkObject *NetworkStruct) UDPWriteBytes(bytes []byte) error {
	//--------------------------------------------------
	var err error
	//--------------------------------------------------
	networkObject.BytesWritten, err = networkObject.UDPConn.WriteTo(bytes, networkObject.RemoteAddr)
	//--------------------------------------------------
	return err
	//--------------------------------------------------
}

//------------------------------------------------------------

func (networkObject *NetworkStruct) UDPClient(requestBytes []byte) ([]byte, error) {
	//--------------------------------------------------
	var err error
	var responseBytes []byte
	var n int
	//--------------------------------------------------
	serverIPAddr, serverPort := SplitAddrPort(networkObject.ServerAddr)
	//----------
	if serverPort == 0 {
		serverPort = UDPServerPort
	}
	//--------------------------------------------------
	clientIPAddr, clientPort := SplitAddrPort(networkObject.ClientAddr)
	//----------
	if clientPort == 0 {
		clientPort = UDPClientPort
	}
	//--------------------------------------------------
	networkObject.UDPConn, err = net.ListenPacket("udp4", fmt.Sprintf("%s:%d", clientIPAddr, clientPort))
	//----------
	if err == nil {
		//----------
		remoteAddr, _ := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%d", serverIPAddr, serverPort))
		//----------
		_, err = networkObject.UDPConn.WriteTo(requestBytes, remoteAddr)
		//--------------------------------------------------
		if err == nil {
			//----------
			defer networkObject.UDPConn.Close()
			//----------
			responseBytes = make([]byte, BufferSize)
			//--------
			n, _, err = networkObject.UDPConn.ReadFrom(responseBytes)
			//----------
			if err == nil && n > 0 {
				responseBytes = responseBytes[0:n]
			}
			//----------
		}
		//----------
	}
	//--------------------------------------------------
	return responseBytes, err
	//--------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------
