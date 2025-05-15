/*

Copyright 2023-2024, Tim Brockley. All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

*/

package system

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/timbrockley/golang-main/file"
)

//------------------------------------------------------------
//###########################################################
//------------------------------------------------------------

//------------------------------------------------------------
// ToString
//------------------------------------------------------------

func ToString(value any) string {
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
// ToBytes
//------------------------------------------------------------

func ToBytes(value any) []byte {
	//------------------------------------------------------------
	switch typedValue := value.(type) {
	case []byte:
		return typedValue
	case string:
		return []byte(typedValue)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, complex64, complex128:
		return []byte(fmt.Sprint(typedValue))
	default:
		return []byte{}
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ToInt
//------------------------------------------------------------

func ToInt(value any) int {
	//------------------------------------------------------------
	if value == nil {
		return 0
	}
	//------------------------------------------------------------
	switch typedValue := value.(type) {
	case string:
		valueSplit := strings.Split(typedValue, ".")
		if len(valueSplit) > 0 {
			int64Val, _ := strconv.ParseInt(valueSplit[0], 10, 0)
			return int(int64Val)
		} else {
			return 0
		}
	case int:
		return int(typedValue)
	case int32:
		return int(typedValue)
	case int64:
		return int(typedValue)
	case float32:
		return int(typedValue)
	case float64:
		return int(typedValue)
	case bool:
		if typedValue {
			return 1
		} else {
			return 0
		}
	default:
		return 0
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ToInt32
//------------------------------------------------------------

func ToInt32(value any) int32 {
	//------------------------------------------------------------
	if value == nil {
		return 0
	}
	//------------------------------------------------------------
	switch typedValue := value.(type) {
	case string:
		valueSplit := strings.Split(typedValue, ".")
		if len(valueSplit) > 0 {
			int64Val, _ := strconv.ParseInt(valueSplit[0], 10, 0)
			return int32(int64Val)
		} else {
			return 0
		}
	case int:
		return int32(typedValue)
	case int32:
		return int32(typedValue)
	case int64:
		return int32(typedValue)
	case float32:
		return int32(typedValue)
	case float64:
		return int32(typedValue)
	case bool:
		if typedValue {
			return 1
		} else {
			return 0
		}
	default:
		return 0
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ToInt64
//------------------------------------------------------------

func ToInt64(value any) int64 {
	//------------------------------------------------------------
	if value == nil {
		return 0
	}
	//------------------------------------------------------------
	switch typedValue := value.(type) {
	case string:
		valueSplit := strings.Split(typedValue, ".")
		if len(valueSplit) > 0 {
			int64Val, _ := strconv.ParseInt(valueSplit[0], 10, 0)
			return int64Val
		} else {
			return 0
		}
	case int:
		return int64(typedValue)
	case int32:
		return int64(typedValue)
	case int64:
		return int64(typedValue)
	case float32:
		return int64(typedValue)
	case float64:
		return int64(typedValue)
	case bool:
		if typedValue {
			return 1
		} else {
			return 0
		}
	default:
		return 0
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ToFloat
//------------------------------------------------------------

func ToFloat(value any) float64 {
	//------------------------------------------------------------
	return ToFloat64(value)
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ToFloat32
//------------------------------------------------------------

func ToFloat32(value any) float32 {
	//------------------------------------------------------------
	return float32(ToFloat64(value))
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ToFloat64
//------------------------------------------------------------

func ToFloat64(value any) float64 {
	//------------------------------------------------------------
	if value == nil {
		return 0
	}
	//------------------------------------------------------------
	switch typedValue := value.(type) {
	case string:
		float64Val, _ := strconv.ParseFloat(value.(string), 64)
		return float64Val
	case int:
		return float64(typedValue)
	case int32:
		return float64(typedValue)
	case int64:
		return float64(typedValue)
	case float32:
		return float64(typedValue)
	case float64:
		return float64(typedValue)
	case bool:
		if typedValue {
			return 1
		} else {
			return 0
		}
	default:
		return 0
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ToBool
//------------------------------------------------------------

func ToBool(value any) bool {
	//------------------------------------------------------------
	switch typedValue := value.(type) {
	case string:
		return typedValue != "" && typedValue != "0" && !strings.EqualFold(typedValue, "false")
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, complex64, complex128:
		return typedValue != 0
	case bool:
		return typedValue
	default:
		return false
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ToStringSlice
//------------------------------------------------------------

func ToStringSlice(value any) []string {
	//------------------------------------------------------------
	switch typedValue := value.(type) {
	case []string:
		return typedValue
	default:
		return []string{}
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ToIntSlice
//------------------------------------------------------------

func ToIntSlice(value any) []int {
	//------------------------------------------------------------
	switch typedValue := value.(type) {
	case []int:
		return typedValue
	default:
		return []int{}
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ToWriter
//------------------------------------------------------------

func ToWriter(value any) io.Writer {
	//------------------------------------------------------------
	switch typedValue := value.(type) {
	case io.Writer:
		return typedValue
	default:
		return io.Discard
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// CopyMap
//------------------------------------------------------------

func CopyMap(m map[string]interface{}) map[string]interface{} {
	newMap := make(map[string]interface{})
	for k, v := range m {
		vm, ok := v.(map[string]interface{})
		if ok {
			newMap[k] = CopyMap(vm)
		} else {
			newMap[k] = v
		}
	}
	return newMap
}

//------------------------------------------------------------
// GetENV
//------------------------------------------------------------

func GetENV(key string) string {
	//------------------------------------------------------------
	return os.Getenv(key)
	//------------------------------------------------------------
}

//------------------------------------------------------------
// GetENVs
//------------------------------------------------------------

func GetENVs() map[string]string {
	//------------------------------------------------------------
	envs := map[string]string{}
	//------------------------------------------------------------
	for _, e := range os.Environ() {
		//--------------------
		pair := strings.SplitN(e, "=", 2)
		//--------------------
		if pair[0] != "" && pair[1] != "" {
			//--------------------
			key := pair[0]
			//--------------------
			envs[key] = pair[1]
			//--------------------
		}
		//--------------------
	}
	//------------------------------------------------------------
	return envs
	//------------------------------------------------------------
}

//------------------------------------------------------------
// SetENV
//------------------------------------------------------------

func SetENV(key string, val string) error {
	//------------------------------------------------------------
	return os.Setenv(key, val)
	//------------------------------------------------------------
}

//------------------------------------------------------------
// SetENVs
//------------------------------------------------------------

func SetENVs(keyVals map[string]string) error {
	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	for key, val := range keyVals {
		//--------------------
		err = os.Setenv(key, val)
		//--------------------
		if err != nil {
			break
		}
		//--------------------
	}
	//------------------------------------------------------------
	return err
	//------------------------------------------------------------
}

//------------------------------------------------------------
// LoadENVs
//------------------------------------------------------------

func LoadENVs(FilePath ...string) error {
	//------------------------------------------------------------
	var filePath, filename string
	//------------------------------------------------------------
	if FilePath != nil && FilePath[0] != "" {
		//--------------------
		filePath = FilePath[0]
		//--------------------
		if IsDirectory, _ := file.IsDirectory(filePath); IsDirectory {
			filename = FindENVFilename(filePath)
			filePath = file.FilePathJoin(filePath, filename)
		}
		//--------------------
	} else {
		//--------------------
		// runtime.Caller(0) => this script / runtime.Caller(1) => calling script
		_, filePath, _, _ = runtime.Caller(1)
		//--------------------
		filename = FindENVFilename(file.Path(filePath))
		filePath = file.FilePathJoin(file.Path(filePath), filename)
		//--------------------
		// if !file.FilePathExists(filePath) {
		// 	return nil
		// }
		//--------------------
	}
	//--------------------
	return godotenv.Load(filePath)
	//------------------------------------------------------------
}

//------------------------------------------------------------
// FindENVFilename
//------------------------------------------------------------

func FindENVFilename(path string) string {
	//------------------------------------------------------------
	HOSTNAME := strings.ToLower(GetHostname())
	//--------------------
	dockerYesNo := file.FilePathExists("/.dockerenv")
	//--------------------
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

//------------------------------------------------------------
// GetLocalIPs
//------------------------------------------------------------

func GetLocalIPs() ([]net.IP, error) {
	//------------------------------------------------------------
	var ip_addrs []net.IP
	//------------------------------------------------------------
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	//------------------------------------------------------------
	for _, addr := range addresses {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip_addrs = append(ip_addrs, ipnet.IP)
			}
		}
	}
	//------------------------------------------------------------
	return ip_addrs, nil
	//------------------------------------------------------------
}

//------------------------------------------------------------
// GetLocalIP
//------------------------------------------------------------

func GetLocalIPAddr() string {
	//------------------------------------------------------------
	ip_addr := ""
	//------------------------------------------------------------
	ip_addrs, _ := GetLocalIPs()
	//--------------------
	if fmt.Sprintf("%T", ip_addrs) == "[]net.IP" && ip_addrs[0] != nil {
		ip_addr = fmt.Sprintf("%v", ip_addrs[0])
	} else {
		ip_addr = "127.0.0.1"
	}
	//------------------------------------------------------------
	return ip_addr
	//------------------------------------------------------------
}

//------------------------------------------------------------
// GetHostname
//------------------------------------------------------------

func GetHostname() string {
	//------------------------------------------------------------
	hostname, err := os.Hostname()
	if err != nil {
		hostname = ""
	}
	//------------------------------------------------------------
	return strings.ToLower(hostname)
	//------------------------------------------------------------
}

//------------------------------------------------------------
// GetOS
//------------------------------------------------------------

func GetOS() string {
	//------------------------------------------------------------
	OS := strings.ToLower(runtime.GOOS)
	//------------------------------------------------------------
	// if OS == "darwin" {
	// 	OS = "mac"
	// }
	//------------------------------------------------------------
	return OS
	//------------------------------------------------------------
}

//------------------------------------------------------------
// CLIParams
//------------------------------------------------------------

func CLIParams() []string {
	//------------------------------------------------------------
	return os.Args
	//------------------------------------------------------------
}

//------------------------------------------------------------
// CLIParam
//------------------------------------------------------------

func CLIParam(index int) string {
	//------------------------------------------------------------
	if fmt.Sprintf("%T", os.Args) != "[]string" || len(os.Args) < (index+1) {
		return ""
	}
	//------------------------------------------------------------
	return os.Args[index]
	//------------------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

//------------------------------------------------------------
// GenerateOTP
//------------------------------------------------------------

func GenerateOTP(timestamp int, secret string) (string, error) {
	//------------------------------------------------------------
	if timestamp <= 0 {
		return "", errors.New("invalid timestamp")
	}
	if secret == "" {
		return "", errors.New("invalid secret")
	}
	//------------------------------------------------------------
	message := make([]byte, 8)
	for i := 7; i >= 0; i-- {
		message[i] = byte(timestamp / 30)
		timestamp >>= 8
	}
	//------------------------------------------------------------
	key, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(strings.ToUpper(secret))
	if err != nil {
		return "", errors.New("secret contains invalid base32 characters")
	}
	//------------------------------------------------------------
	hash := hmac.New(sha1.New, key)
	hash.Write(message)
	hmacResult := hash.Sum(nil)
	//------------------------------------------------------------
	offset := hmacResult[len(hmacResult)-1] & 0xF
	code := (int(hmacResult[offset+0]&0x7F) << 24) |
		(int(hmacResult[offset+1]&0xFF) << 16) |
		(int(hmacResult[offset+2]&0xFF) << 8) |
		(int(hmacResult[offset+3] & 0xFF))
	//------------------------------------------------------------
	otp := code % 1000000
	return fmt.Sprintf("%06d", otp), nil
	// ------------------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------
