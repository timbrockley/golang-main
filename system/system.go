//--------------------------------------------------------------------------------

package system

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/timbrockley/golang-main/file"
)

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------

//------------------------------------------------------------
// ConvertToString
//------------------------------------------------------------

func ConvertToString(value any) string {
	//------------------------------------------------------------
	switch typedValue := value.(type) {
	case string:
		return typedValue
	case []byte:
		return string(typedValue)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, complex64, complex128:
		return fmt.Sprint(typedValue)
	case bool:
		if typedValue {
			return "true"
		} else {
			return "false"
		}
	}
	return ""
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ConvertToBytes
//------------------------------------------------------------

func ConvertToBytes(value any) []byte {
	//------------------------------------------------------------
	switch typedValue := value.(type) {
	case string:
		return []byte(typedValue)
	case []byte:
		return typedValue
	}
	return []byte{}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ConvertToInt
//------------------------------------------------------------

func ConvertToInt(value any) int {
	//------------------------------------------------------------
	if value == nil {
		//----------
		return 0
		//----------
	} else {

		if fmt.Sprintf("%T", value) == "string" {
			//----------
			valueSplit := strings.Split(value.(string), ".")
			if len(valueSplit) > 0 {
				int64Val, _ := strconv.ParseInt(valueSplit[0], 10, 0)
				return int(int64Val)
			} else {
				return 0
			}
			//----------
		} else if fmt.Sprintf("%T", value) == "int" {
			//----------
			return value.(int)
			//----------
		} else if fmt.Sprintf("%T", value) == "int32" {
			//----------
			return value.(int)
			//----------
		} else if fmt.Sprintf("%T", value) == "int64" {
			//----------
			return value.(int)
			//----------
		} else if fmt.Sprintf("%T", value) == "float32" {
			//----------
			return int(value.(float32))
			//----------
		} else if fmt.Sprintf("%T", value) == "float64" {
			//----------
			return int(value.(float64))
			//----------
		} else if fmt.Sprintf("%T", value) == "bool" {
			//----------
			if value.(bool) {
				return 1
			} else {
				return 0
			}
			//----------
		} else {
			//----------
			return 0
			//----------
		}
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ConvertToInt32
//------------------------------------------------------------

func ConvertToInt32(value any) int32 {
	//------------------------------------------------------------
	if value == nil {
		//----------
		return 0
		//----------
	} else {

		if fmt.Sprintf("%T", value) == "string" {
			//----------
			valueSplit := strings.Split(value.(string), ".")
			if len(valueSplit) > 0 {
				int64Val, _ := strconv.ParseInt(valueSplit[0], 10, 0)
				return int32(int64Val)
			} else {
				return 0
			}
			//----------
		} else if fmt.Sprintf("%T", value) == "int" {
			//----------
			return value.(int32)
			//----------
		} else if fmt.Sprintf("%T", value) == "int32" {
			//----------
			return value.(int32)
			//----------
		} else if fmt.Sprintf("%T", value) == "int64" {
			//----------
			return value.(int32)
			//----------
		} else if fmt.Sprintf("%T", value) == "float32" {
			//----------
			return int32(value.(float32))
			//----------
		} else if fmt.Sprintf("%T", value) == "float64" {
			//----------
			return int32(value.(float64))
			//----------
		} else if fmt.Sprintf("%T", value) == "bool" {
			//----------
			if value.(bool) {
				return 1
			} else {
				return 0
			}
			//----------
		} else {
			//----------
			return 0
			//----------
		}
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ConvertToInt64
//------------------------------------------------------------

func ConvertToInt64(value any) int64 {
	//------------------------------------------------------------
	if value == nil {
		//----------
		return 0
		//----------
	} else {

		if fmt.Sprintf("%T", value) == "string" {
			//----------
			valueSplit := strings.Split(value.(string), ".")
			if len(valueSplit) > 0 {
				int64Val, _ := strconv.ParseInt(valueSplit[0], 10, 0)
				return int64Val
			} else {
				return 0
			}
			//----------
		} else if fmt.Sprintf("%T", value) == "int" {
			//----------
			return value.(int64)
			//----------
		} else if fmt.Sprintf("%T", value) == "int32" {
			//----------
			return value.(int64)
			//----------
		} else if fmt.Sprintf("%T", value) == "int64" {
			//----------
			return value.(int64)
			//----------
		} else if fmt.Sprintf("%T", value) == "float32" {
			//----------
			return int64(value.(float32))
			//----------
		} else if fmt.Sprintf("%T", value) == "float64" {
			//----------
			return int64(value.(float64))
			//----------
		} else if fmt.Sprintf("%T", value) == "bool" {
			//----------
			if value.(bool) {
				return 1
			} else {
				return 0
			}
			//----------
		} else {
			//----------
			return 0
			//----------
		}
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ConvertToFloat
//------------------------------------------------------------

func ConvertToFloat(value any) float64 {
	//------------------------------------------------------------
	return ConvertToFloat64(value)
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ConvertToFloat32
//------------------------------------------------------------

func ConvertToFloat32(value any) float32 {
	//------------------------------------------------------------
	return float32(ConvertToFloat64(value))
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ConvertToFloat64
//------------------------------------------------------------

func ConvertToFloat64(value any) float64 {
	//------------------------------------------------------------
	if value == nil {
		//----------
		return 0
		//----------
	} else {

		if fmt.Sprintf("%T", value) == "string" {
			//----------
			float64Val, _ := strconv.ParseFloat(value.(string), 64)
			return float64Val
			//----------
		} else if fmt.Sprintf("%T", value) == "int" {
			//----------
			return float64(value.(int))
			//----------
		} else if fmt.Sprintf("%T", value) == "int32" {
			//----------
			return float64(value.(int32))
			//----------
		} else if fmt.Sprintf("%T", value) == "int64" {
			//----------
			return float64(value.(int64))
			//----------
		} else if fmt.Sprintf("%T", value) == "float32" {
			//----------
			return float64(value.(float32))
			//----------
		} else if fmt.Sprintf("%T", value) == "float64" {
			//----------
			return value.(float64)
			//----------
		} else if fmt.Sprintf("%T", value) == "bool" {
			//----------
			if value.(bool) {
				return 1
			} else {
				return 0
			}
			//----------
		} else {
			//----------
			return 0
			//----------
		}
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ConvertToBool
//------------------------------------------------------------

func ConvertToBool(value any) bool {
	//------------------------------------------------------------
	switch typedValue := value.(type) {
	case string:
		return typedValue != "" && typedValue != "0" && !strings.EqualFold(typedValue, "false")
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, complex64, complex128:
		return typedValue != 0
	case bool:
		return typedValue
	}
	return false
	//------------------------------------------------------------
}

//--------------------------------------------------------------------------------
// GetENV
//--------------------------------------------------------------------------------

func GetENV(key string) string {
	//--------------------------------------------------------------------------------
	return os.Getenv(key)
	//--------------------------------------------------------------------------------
}

//--------------------------------------------------------------------------------
// GetENVs
//--------------------------------------------------------------------------------

func GetENVs() map[string]string {
	//--------------------------------------------------------------------------------
	envs := map[string]string{}
	//--------------------------------------------------------------------------------
	for _, e := range os.Environ() {
		//----------
		pair := strings.SplitN(e, "=", 2)
		//----------
		if pair[0] != "" && pair[1] != "" {
			//----------
			key := pair[0]
			//----------
			envs[key] = pair[1]
			//----------
		}
		//----------
	}
	//--------------------------------------------------------------------------------
	return envs
	//--------------------------------------------------------------------------------
}

//--------------------------------------------------------------------------------
// SetENV
//--------------------------------------------------------------------------------

func SetENV(key string, val string) error {
	//--------------------------------------------------------------------------------
	return os.Setenv(key, val)
	//--------------------------------------------------------------------------------
}

//--------------------------------------------------------------------------------
// SetENVs
//--------------------------------------------------------------------------------

func SetENVs(keyVals map[string]string) error {
	//--------------------------------------------------------------------------------
	var err error
	//--------------------------------------------------------------------------------
	for key, val := range keyVals {
		//----------
		err = os.Setenv(key, val)
		//----------
		if err != nil {
			break
		}
		//----------
	}
	//--------------------------------------------------------------------------------
	return err
	//--------------------------------------------------------------------------------
}

//--------------------------------------------------------------------------------
// LoadENVs
//--------------------------------------------------------------------------------

func LoadENVs(FilePath ...string) error {
	//--------------------------------------------------------------------------------
	var filePath, filename string
	//--------------------------------------------------------------------------------
	if FilePath != nil && FilePath[0] != "" {
		//----------
		filePath = FilePath[0]
		//----------
		if isDir, _ := file.IsDir(filePath); isDir {
			filename = FindENVFilename(filePath)
			filePath = file.FilePathJoin(filePath, filename)
		}
		//----------
	} else {
		//----------
		// runtime.Caller(0) => this script / runtime.Caller(1) => calling script
		_, filePath, _, _ = runtime.Caller(1)
		//----------
		filename = FindENVFilename(file.Path(filePath))
		filePath = file.FilePathJoin(file.Path(filePath), filename)
		//----------
		// if !file.FilePathExists(filePath) {
		// 	return nil
		// }
		//----------
	}
	//----------
	return godotenv.Load(filePath)
	//--------------------------------------------------------------------------------
}

//------------------------------------------------------------
// FindENVFilename
//------------------------------------------------------------

func FindENVFilename(path string) string {
	//------------------------------------------------------------
	HOSTNAME := strings.ToLower(GetHostname())
	//----------
	dockerYesNo := file.FilePathExists("/.dockerenv")
	//----------
	OS := strings.ToLower(GetOS())
	//------------------------------------------------------------
	FILENAME_HOSTNAME_DOCKER_EXT := "." + HOSTNAME + "_docker.env"
	FILENAME_HOSTNAME_OS_EXT := "." + HOSTNAME + "_" + OS + ".env"
	FILENAME_HOSTNAME_EXT := "." + HOSTNAME + ".env"
	FILENAME_DOCKER_EXT := ".docker.env"
	FILENAME_OS_EXT := "." + OS + ".env"
	FILENAME_EXT := ".env"
	//------------------------------------------------------------
	if dockerYesNo && file.FilePathExists(file.FilePathJoin(path, FILENAME_HOSTNAME_DOCKER_EXT)) {

		return FILENAME_HOSTNAME_DOCKER_EXT

	} else if file.FilePathExists(file.FilePathJoin(path, FILENAME_HOSTNAME_OS_EXT)) {

		return FILENAME_HOSTNAME_OS_EXT

	} else if file.FilePathExists(file.FilePathJoin(path, FILENAME_HOSTNAME_EXT)) {

		return FILENAME_HOSTNAME_EXT

	} else if dockerYesNo && file.FilePathExists(file.FilePathJoin(path, FILENAME_DOCKER_EXT)) {

		return FILENAME_DOCKER_EXT

	} else if file.FilePathExists(file.FilePathJoin(path, FILENAME_OS_EXT)) {

		return FILENAME_OS_EXT

	} else {

		return FILENAME_EXT
	}
	//------------------------------------------------------------
}

//--------------------------------------------------------------------------------
// GetLocalIPs
//--------------------------------------------------------------------------------

func GetLocalIPs() ([]net.IP, error) {
	//--------------------------------------------------------------------------------
	var ip_addrs []net.IP
	//--------------------------------------------------------------------------------
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	//--------------------------------------------------------------------------------
	for _, addr := range addresses {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip_addrs = append(ip_addrs, ipnet.IP)
			}
		}
	}
	//--------------------------------------------------------------------------------
	return ip_addrs, nil
	//--------------------------------------------------------------------------------
}

//--------------------------------------------------------------------------------
// GetLocalIP
//--------------------------------------------------------------------------------

func GetLocalIPAddr() string {
	//--------------------------------------------------------------------------------
	ip_addr := ""
	//--------------------------------------------------------------------------------
	ip_addrs, _ := GetLocalIPs()
	//----------
	if fmt.Sprintf("%T", ip_addrs) == "[]net.IP" && ip_addrs[0] != nil {
		ip_addr = fmt.Sprintf("%v", ip_addrs[0])
	} else {
		ip_addr = "127.0.0.1"
	}
	//--------------------------------------------------------------------------------
	return ip_addr
	//--------------------------------------------------------------------------------
}

//--------------------------------------------------------------------------------
// GetHostname
//--------------------------------------------------------------------------------

func GetHostname() string {
	//--------------------------------------------------------------------------------
	hostname, err := os.Hostname()
	if err != nil {
		hostname = ""
	}
	//--------------------------------------------------------------------------------
	return strings.ToLower(hostname)
	//--------------------------------------------------------------------------------
}

//--------------------------------------------------------------------------------
// GetOS
//--------------------------------------------------------------------------------

func GetOS() string {
	//--------------------------------------------------------------------------------
	OS := strings.ToLower(runtime.GOOS)
	//--------------------------------------------------------------------------------
	// if OS == "darwin" {
	// 	OS = "mac"
	// }
	//--------------------------------------------------------------------------------
	return OS
	//--------------------------------------------------------------------------------
}

//--------------------------------------------------------------------------------
// CLIParams
//--------------------------------------------------------------------------------

func CLIParams() []string {
	//--------------------------------------------------------------------------------
	return os.Args
	//--------------------------------------------------------------------------------
}

//--------------------------------------------------------------------------------
// CLIParam
//--------------------------------------------------------------------------------

func CLIParam(index int) string {
	//--------------------------------------------------------------------------------
	if fmt.Sprintf("%T", os.Args) != "[]string" || len(os.Args) < (index+1) {
		return ""
	}
	//--------------------------------------------------------------------------------
	return os.Args[index]
	//--------------------------------------------------------------------------------
}

//--------------------------------------------------------------------------------
// ################################################################################
//--------------------------------------------------------------------------------
