/*

Copyright 2023-2024, Tim Brockley. All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

*/

package file

import (
	"errors"
	"fmt"
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/juju/fslock"
)

//------------------------------------------------------------

type FileMutexStruct struct {
	StateMutex          sync.Mutex
	IsLocked            bool
	LockLevelStateMutex sync.Mutex
	LockLevel           int
	SyncMutex           sync.Mutex
}

var FileMutex FileMutexStruct

//------------------------------------------------------------
//################################################################################
//------------------------------------------------------------

//------------------------------------------------------------
// mutex methods !!! ONLY WORK IN CURRENT RUNNING PROCESS !!!
//------------------------------------------------------------

//------------------------------------------------------------
// internal mutex (onlys performs lock once to prevent internal "lockup")
//------------------------------------------------------------

func (fm *FileMutexStruct) lock() {
	//--------------------
	fm.LockLevelStateMutex.Lock()
	fm.LockLevel++
	//--------------------
	if fm.LockLevel == 1 {
		fm.LockLevelStateMutex.Unlock()
		fm.Lock()
		return
	}
	//--------------------
	fm.LockLevelStateMutex.Unlock()
	//--------------------
}

func (fm *FileMutexStruct) unlock() error {
	//--------------------
	fm.LockLevelStateMutex.Lock()
	fm.LockLevel--
	//--------------------
	if fm.LockLevel == 0 {
		fm.LockLevelStateMutex.Unlock()
		return fm.Unlock()
	}
	//---------------------
	fm.LockLevelStateMutex.Unlock()
	return nil
	//---------------------
}

//------------------------------------------------------------
// main mutex (locks every time)
//------------------------------------------------------------

func (fm *FileMutexStruct) Lock() {
	//---------------------
	fm.StateMutex.Lock()
	defer fm.StateMutex.Unlock()
	//---------------------
	fm.LockLevelStateMutex.Lock()
	fm.LockLevel++
	fm.LockLevelStateMutex.Unlock()
	//---------------------
	fm.SyncMutex.Lock()
	fm.IsLocked = true
	//---------------------
}

func (fm *FileMutexStruct) Unlock() error {
	//---------------------
	fm.StateMutex.Lock()
	defer fm.StateMutex.Unlock()
	//---------------------
	fm.LockLevelStateMutex.Lock()
	fm.LockLevel--
	fm.LockLevelStateMutex.Unlock()
	//---------------------
	if !fm.IsLocked {
		return fmt.Errorf("already unlocked")
	} else {
		fm.SyncMutex.Unlock()
		fm.IsLocked = false
		return nil
	}
	//---------------------
}

//------------------------------------------------------------
// FilePathExists
//------------------------------------------------------------

func FilePathExists(filePath string) bool {
	//------------------------------------------------------------
	filePath = filepath.FromSlash(filePath)
	//------------------------------------------------------------
	_, err := os.Stat(filePath)
	//------------------------------------------------------------
	return err == nil
	//------------------------------------------------------------
}

//------------------------------------------------------------
// IsDirectory
//------------------------------------------------------------

func IsDirectory(filePath string) (bool, error) {
	//------------------------------------------------------------
	var err error
	var fileInfo fs.FileInfo
	//------------------------------------------------------------
	filePath = filepath.FromSlash(filePath)
	//------------------------------------------------------------
	fileInfo, err = os.Stat(filePath)
	//------------------------------------------------------------
	return err == nil && fileInfo.IsDir(), err
	//------------------------------------------------------------
}

//------------------------------------------------------------
// IsFile
//------------------------------------------------------------

func IsFile(filePath string) (bool, error) {
	//------------------------------------------------------------
	var err error
	var fileInfo fs.FileInfo
	//------------------------------------------------------------
	filePath = filepath.FromSlash(filePath)
	//------------------------------------------------------------
	fileInfo, err = os.Stat(filePath)
	//------------------------------------------------------------
	return err == nil && !fileInfo.IsDir(), err
	//------------------------------------------------------------
}

//------------------------------------------------------------
// FileLoad
//------------------------------------------------------------

func FileLoad(filePath string) (string, error) {
	//------------------------------------------------------------
	filePath = filepath.FromSlash(filePath)
	//------------------------------------------------------------
	var err error
	var dataBytes []byte
	var dataString string
	//------------------------------------------------------------
	if !FilePathExists(filePath) {
		return dataString, errors.New("file does not exist")
	}
	//------------------------------------------------------------
	fslock := fslock.New(filePath)
	fslock.Lock()
	defer fslock.Unlock()
	//------------------------------------------------------------
	dataBytes, err = os.ReadFile(filePath)
	//--------------------
	if err == nil {
		//--------------------
		dataString = string(dataBytes)
		//--------------------
	}
	//------------------------------------------------------------
	return dataString, err
	//------------------------------------------------------------
}

//------------------------------------------------------------
// FileSave
//------------------------------------------------------------

func FileSave(filePath string, data string) error {
	//------------------------------------------------------------
	filePath = filepath.FromSlash(filePath)
	//------------------------------------------------------------
	fslock := fslock.New(filePath)
	fslock.Lock()
	defer fslock.Unlock()
	//------------------------------------------------------------
	file, err := os.Create(filePath)
	//--------------------
	if err == nil {
		//--------------------
		defer file.Close()
		//--------------------
		_, err = file.WriteString(data)
		//--------------------
	}
	//------------------------------------------------------------
	return err
	//------------------------------------------------------------
}

//------------------------------------------------------------
// FileAppend
//------------------------------------------------------------

func FileAppend(filePath string, data string) error {
	//------------------------------------------------------------
	filePath = filepath.FromSlash(filePath)
	//------------------------------------------------------------
	fslock := fslock.New(filePath)
	fslock.Lock()
	defer fslock.Unlock()
	//------------------------------------------------------------
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	//--------------------
	if err == nil && data != "" {
		//--------------------
		defer file.Close()
		//--------------------
		_, err = file.WriteString(data)
		//--------------------
	}
	//------------------------------------------------------------
	return err
	//------------------------------------------------------------
}

//------------------------------------------------------------
// FileRemove
//------------------------------------------------------------

func FileRemove(filePath string) error {
	//------------------------------------------------------------
	filePath = filepath.FromSlash(filePath)
	//------------------------------------------------------------
	if !FilePathExists(filePath) {
		return errors.New("file does not exist")
	}
	//------------------------------------------------------------
	fslock := fslock.New(filePath)
	fslock.Lock()
	defer fslock.Unlock()
	//------------------------------------------------------------
	return os.Remove(filePath)
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Path
//------------------------------------------------------------

func Path(FilePath ...string) string {
	//------------------------------------------------------------
	var filePath string
	//------------------------------------------------------------
	if FilePath != nil && FilePath[0] != "" {
		filePath = FilePath[0]
	} else {
		// runtime.Caller(0) => this script / runtime.Caller(1) => calling script
		_, filePath, _, _ = runtime.Caller(1)
	}
	//------------------------------------------------------------
	if filepath.Ext(filePath) != "" {
		return strings.TrimRight(filepath.Dir(filePath), "/") + "/"
	} else {
		return strings.TrimRight(filePath, "/") + "/"
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// PathJoin
//------------------------------------------------------------

func PathJoin(path1 string, path2 string) string {
	//------------------------------------------------------------
	return strings.TrimRight(path1, "/") + "/" + strings.Trim(path2, "/") + "/"
	//------------------------------------------------------------
}

//------------------------------------------------------------
// FilePathJoin
//------------------------------------------------------------

func FilePathJoin(path string, filename string) string {
	//------------------------------------------------------------
	return strings.TrimRight(path, "/") + "/" + filename
	//------------------------------------------------------------
}

//------------------------------------------------------------
// FilePathSplit
//------------------------------------------------------------

func FilePathSplit(FilePath ...string) (string, string) {
	//------------------------------------------------------------
	var filePath string
	//------------------------------------------------------------
	if FilePath != nil && FilePath[0] != "" {
		filePath = FilePath[0]
	} else {
		// runtime.Caller(0) => this script / runtime.Caller(1) => calling script
		_, filePath, _, _ = runtime.Caller(1)
	}
	//------------------------------------------------------------
	return filepath.Split(filePath)
	//------------------------------------------------------------
}

//------------------------------------------------------------
// FilePathBaseToFilename
//------------------------------------------------------------

func FilePathBaseToFilename(FilePath ...string) string {
	//------------------------------------------------------------
	var filePath string
	//------------------------------------------------------------
	if FilePath != nil && FilePath[0] != "" {
		filePath = FilePath[0]
	} else {
		// runtime.Caller(0) => this script / runtime.Caller(1) => calling script
		_, filePath, _, _ = runtime.Caller(1)
	}
	//------------------------------------------------------------
	filePath = filepath.FromSlash(filePath)
	filePath = strings.Replace(filePath, "/", "-", -1)
	//------------------------------------------------------------
	filePath = filePath[0 : len(filePath)-len(filepath.Ext(filePath))]
	//------------------------------------------------------------
	m1 := regexp.MustCompile(`^-|-$`)
	//------------------------------------------------------------
	return m1.ReplaceAllString(filePath, "")
	//------------------------------------------------------------
}

//------------------------------------------------------------
// FilePathBase
//------------------------------------------------------------

func FilePathBase(FilePath ...string) string {
	//------------------------------------------------------------
	var filePath string
	//------------------------------------------------------------
	if FilePath != nil && FilePath[0] != "" {
		filePath = FilePath[0]
	} else {
		// runtime.Caller(0) => this script / runtime.Caller(1) => calling script
		_, filePath, _, _ = runtime.Caller(1)
	}
	//------------------------------------------------------------
	filePath = filepath.FromSlash(filePath)
	//------------------------------------------------------------
	filePath = filePath[0 : len(filePath)-len(filepath.Ext(filePath))]
	//------------------------------------------------------------
	return strings.TrimRight(filePath, "/")
	//------------------------------------------------------------
}

//------------------------------------------------------------
// FilenameBase
//------------------------------------------------------------

func FilenameBase(FilePath ...string) string {
	//------------------------------------------------------------
	var filePath string
	//------------------------------------------------------------
	if FilePath != nil && FilePath[0] != "" {
		filePath = FilePath[0]
	} else {
		// runtime.Caller(0) => this script / runtime.Caller(1) => calling script
		_, filePath, _, _ = runtime.Caller(1)
	}
	//------------------------------------------------------------
	filePath = filepath.FromSlash(filePath)
	//------------------------------------------------------------
	filePath = filepath.Base(filePath)
	//------------------------------------------------------------
	return filePath[0 : len(filePath)-len(filepath.Ext(filePath))]
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Filename
//------------------------------------------------------------

func Filename(FilePath ...string) string {
	//------------------------------------------------------------
	var filePath string
	//------------------------------------------------------------
	if FilePath != nil && FilePath[0] != "" {
		filePath = FilePath[0]
	} else {
		// runtime.Caller(0) => this script / runtime.Caller(1) => calling script
		_, filePath, _, _ = runtime.Caller(1)
	}
	//------------------------------------------------------------
	filePath = filepath.FromSlash(filePath)
	//------------------------------------------------------------
	if filePath[len(filePath)-1:] == "/" {
		return ""
	}
	//------------------------------------------------------------
	return filepath.Base(filePath)
	//------------------------------------------------------------
}

//------------------------------------------------------------
// FilenameExt
//------------------------------------------------------------

func FilenameExt(FilePath ...string) string {
	//------------------------------------------------------------
	var filePath string
	//------------------------------------------------------------
	if FilePath != nil && FilePath[0] != "" {
		filePath = FilePath[0]
	} else {
		// runtime.Caller(0) => this script / runtime.Caller(1) => calling script
		_, filePath, _, _ = runtime.Caller(1)
	}
	//------------------------------------------------------------
	filePath = filepath.FromSlash(filePath)
	//------------------------------------------------------------
	return strings.TrimLeft(filepath.Ext(filePath), ".")
	//------------------------------------------------------------
}

//------------------------------------------------------------
// TempPath
//------------------------------------------------------------

func TempPath() (string, error) {
	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	tempPath := strings.TrimRight(os.TempDir(), "/") + "/golang/"
	//------------------------------------------------------------
	if !FilePathExists(tempPath) {
		err = os.MkdirAll(tempPath, 0o700)
	}
	//------------------------------------------------------------
	return filepath.FromSlash(tempPath), err
	//------------------------------------------------------------
}

//------------------------------------------------------------
// TempFilePath
//------------------------------------------------------------

func TempFilePath() (string, error) {
	//------------------------------------------------------------
	tempFilePath, err := TempPath()
	if err == nil {
		tempFilePath += fmt.Sprintf("tmp_%d_%d.tmp", time.Now().UnixMicro(), rand.Int())
	}
	//------------------------------------------------------------
	return tempFilePath, err
	//------------------------------------------------------------
}

//------------------------------------------------------------
// LogFilePath
//------------------------------------------------------------

func LogFilePath() string {
	//------------------------------------------------------------
	return filepath.FromSlash(os.TempDir() + "/golang.log")
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Log
//------------------------------------------------------------

func Log(messageString string, FilePath ...string) error {
	//------------------------------------------------------------
	var logFilePath, logLineString string
	//------------------------------------------------------------
	var callingFilePath, callingPath, callingFilename string
	var callingLineNumber int
	var pc uintptr
	var ok bool
	//------------------------------------------------------------
	if FilePath != nil && FilePath[0] != "" {
		logFilePath = filepath.FromSlash(FilePath[0])
	} else {
		logFilePath = LogFilePath()
	}
	//------------------------------------------------------------
	// check if file exists before fslock used because fslock creates a file
	if !FilePathExists(logFilePath) {
		logLineString = "utm\tcymd\thms\tpath\tfilename\tline\terror\n"
	}
	//------------------------------------------------------------
	fslock := fslock.New(logFilePath)
	fslock.Lock()
	defer fslock.Unlock()
	//------------------------------------------------------------
	timeNow := time.Now()
	utm := timeNow.UnixMicro()
	t := timeNow.UTC()
	//--------------------
	// runtime.Caller(0) => this script / runtime.Caller(1) => calling script
	pc, callingFilePath, callingLineNumber, ok = runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		callingPath, callingFilename = filepath.Split(callingFilePath)
		callingPath = strings.TrimRight(callingPath, "/")
	}
	//------------------------------------------------------------
	replacer := strings.NewReplacer(
		"\x5C", "\\\\", // \x5C = backslash
		"\x09", "\\t", // \x09 = tab
		"\x0A", "\\n", // \x0A = newline
		"\x0D", "\\r", // \x0D = carriage return
	// 	"\x22", "\\q", // \x22 = double quotes
	// 	"\x27", "\\a", // \x27 = apostrophe
	// 	"\x60", "\\g", // \x60 = grave accent
	)
	messageString = replacer.Replace(messageString)
	//------------------------------------------------------------
	escapedMessageString := ""
	//--------------------
	for i := 0; i < len(messageString); i++ {
		charByte := messageString[i]
		if charByte >= 0x20 && charByte < 0x7F {
			escapedMessageString += string(charByte)
		} else {
			escapedMessageString += fmt.Sprintf("\\x%02X", charByte)
		}
	}
	//------------------------------------------------------------
	logLineString += fmt.Sprintf(
		"%d\t%d%02d%02d\t%02d%02d%02d\t%v\t%v\t%v\t%v\n",
		utm,
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second(),
		callingPath,
		callingFilename,
		callingLineNumber,
		escapedMessageString)
	//------------------------------------------------------------
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	//--------------------
	if err == nil {
		//--------------------
		defer file.Close()
		//--------------------
		_, err = file.WriteString(logLineString)
		//--------------------
	}
	//------------------------------------------------------------
	return err
	//------------------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------
