//------------------------------------------------------------

package file

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

//------------------------------------------------------------

var testTempPath, testTempGolangPath string
var testDataFilename, testLogFilename string

//------------------------------------------------------------
// initial setup
//------------------------------------------------------------

func TestMain(t *testing.T) {

	//------------------------------------------------------------
	testTempPath = os.TempDir()
	testTempGolangPath = testTempPath + "/golang"
	testLogFilename = testTempPath + "/golang.log"
	//------------------------------------------------------------
	if !FilePathExists(testTempGolangPath) {

		err := os.Mkdir(testTempGolangPath, 0700)
		if err != nil {

			t.Error(err)
		}
	}
	//------------------------------------------------------------
	// runtime.Caller(0) => this script / runtime.Caller(1) => calling script
	_, testDataFilename, _, _ = runtime.Caller(0)
	//------------------------------------------------------------
	testDataFilename = filepath.FromSlash(testDataFilename)
	testDataFilename = strings.Replace(testDataFilename, `/`, "-", -1)
	//--------------------------------------------------------------------------------
	testDataFilename = testDataFilename[0 : len(testDataFilename)-len(filepath.Ext(testDataFilename))]
	//--------------------------------------------------------------------------------
	testDataFilename = strings.TrimLeft(testDataFilename, "-")
	//------------------------------------------------------------
	testDataFilename = testTempGolangPath + `/` + testDataFilename + `.txt`
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Lock method
//------------------------------------------------------------

func TestLockUnlock(t *testing.T) {

	//------------------------------------------------------------
	var resultString, EXPECTED_string string
	var resultInt, EXPECTED_int int
	//------------------------------------------------------------
	resultString = fmt.Sprint(&FileMutex.SyncMutex)
	EXPECTED_string = "&{0 0}"
	//----------
	if resultString != EXPECTED_string {
		t.Errorf("expected result = %q but should = %q", resultString, EXPECTED_string)
	}
	//------------------------------------------------------------
	resultInt = FileMutex.LockValue
	EXPECTED_int = 0
	//----------
	if resultInt != EXPECTED_int {
		t.Errorf("expected result = %d but should = %d", resultInt, EXPECTED_int)
	}
	//------------------------------------------------------------
	FileMutex.Lock()
	FileMutex.lock()
	FileMutex.lock()
	//------------------------------------------------------------
	resultString = fmt.Sprint(&FileMutex.SyncMutex)
	EXPECTED_string = "&{1 0}"
	//----------
	if resultString != EXPECTED_string {
		t.Errorf("expected result = %q but should = %q", resultString, EXPECTED_string)
	}
	//------------------------------------------------------------
	resultInt = FileMutex.LockValue
	EXPECTED_int = 3
	//----------
	if resultInt != EXPECTED_int {
		t.Errorf("expected result = %d but should = %d", resultInt, EXPECTED_int)
	}
	//------------------------------------------------------------
	FileMutex.unlock()
	FileMutex.unlock()
	FileMutex.Unlock()
	//------------------------------------------------------------
	resultString = fmt.Sprint(&FileMutex.SyncMutex)
	EXPECTED_string = "&{0 0}"
	//----------
	if resultString != EXPECTED_string {
		t.Errorf("expected result = %q but should = %q", resultString, EXPECTED_string)
	}
	//------------------------------------------------------------
	resultInt = FileMutex.LockValue
	EXPECTED_int = 0
	//----------
	if resultInt != EXPECTED_int {
		t.Errorf("expected result = %d but should = %d", resultInt, EXPECTED_int)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// FilePathExists
//------------------------------------------------------------

func TestFilePathExists(t *testing.T) {

	//------------------------------------------------------------
	var result bool
	//------------------------------------------------------------
	result = FilePathExists(`MADEUP_PATH_fdsfhkdfghd7s8gds78f78`)
	//------------------------------------------------------------
	if result {

		t.Errorf("result = %v but should = %v", result, false)
	}
	//------------------------------------------------------------
	result = FilePathExists(`/`)
	//------------------------------------------------------------
	if !result {

		t.Errorf("result = %v but should = %v", result, true)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// IsDir
//------------------------------------------------------------

func TestIsDir(t *testing.T) {

	//------------------------------------------------------------
	var err error
	var result bool
	//------------------------------------------------------------
	result, err = IsDir(`MADEUP_PATH_fdsfhkdfghd7s8gds78f78`)
	//------------------------------------------------------------
	if err == nil || result {

		// t.Error(err)
		t.Errorf("result = %v but should = %v", result, false)
	}
	//------------------------------------------------------------
	result, err = IsDir(`/`)
	//------------------------------------------------------------
	if err != nil || !result {

		t.Errorf("result = %v but should = %v", result, true)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// IsFile
//------------------------------------------------------------

func TestIsFile(t *testing.T) {

	//------------------------------------------------------------
	var err error
	var result bool
	//------------------------------------------------------------
	result, err = IsFile(`MADEUP_PATH_fdsfhkdfghd7s8gds78f78`)
	//------------------------------------------------------------
	if err == nil || result {

		// t.Error(err)
		t.Errorf("result = %v but should = %v", result, false)
	}
	//------------------------------------------------------------
	result, err = IsFile(`/`)
	//------------------------------------------------------------
	if err != nil || result {

		t.Errorf("result = %v but should = %v", result, false)
	}
	//------------------------------------------------------------
	result, err = IsFile(`/etc/fstab`)
	//------------------------------------------------------------
	if err != nil || !result {

		t.Errorf("result = %v but should = %v", result, true)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Path
//------------------------------------------------------------

func TestPath(t *testing.T) {

	//------------------------------------------------------------
	var path string
	//------------------------------------------------------------
	path = Path()
	//------------------------------------------------------------
	if path == "" {

		t.Errorf(`path = %q`, path)
	}
	//------------------------------------------------------------
	path = Path(`/path1/path2/`)
	//------------------------------------------------------------
	if path != "/path1/path2/" {

		t.Errorf(`path = %q`, path)
	}
	//------------------------------------------------------------
	path = Path(`/path3/path4/`)
	//------------------------------------------------------------
	if path != "/path3/path4/" {

		t.Errorf(`path = %q`, path)
	}
	//------------------------------------------------------------
	path = Path(`/path5/path6/filename.txt`)
	//------------------------------------------------------------
	if path != "/path5/path6/" {

		t.Errorf(`path = %q`, path)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// PathJoin
//------------------------------------------------------------

func TestPathJoin(t *testing.T) {

	//------------------------------------------------------------
	var path, path1, path2 string
	//------------------------------------------------------------
	path1 = "/path1/path2/"
	path2 = "/path3/path4/"
	//------------------------------------------------------------
	path = PathJoin(path1, path2)
	//------------------------------------------------------------
	if path != "/path1/path2/path3/path4/" {

		t.Errorf(`filePath = %q`, path)
	}
	//------------------------------------------------------------
	path1 = "/path11/path12"
	path2 = "/path13/path14"
	//------------------------------------------------------------
	path = PathJoin(path1, path2)
	//------------------------------------------------------------
	if path != "/path11/path12/path13/path14/" {

		t.Errorf(`filePath = %q`, path)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// FilePathJoin
//------------------------------------------------------------

func TestFilePathJoin(t *testing.T) {

	//------------------------------------------------------------
	var filePath, path, filename string
	//------------------------------------------------------------
	path = "/path1/path2"
	filename = "filename.txt"
	//------------------------------------------------------------
	filePath = FilePathJoin(path, filename)
	//------------------------------------------------------------
	if filePath != "/path1/path2/filename.txt" {

		t.Errorf(`filePath = %q`, filePath)
	}
	//------------------------------------------------------------
	path = `/path3/path4/`
	filename = "filename.txt"
	//------------------------------------------------------------
	filePath = FilePathJoin(path, filename)
	//------------------------------------------------------------
	if filePath != "/path3/path4/filename.txt" {

		t.Errorf(`filePath = %q`, filePath)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// FilePathSplit
//------------------------------------------------------------

func TestFilePathSplit(t *testing.T) {

	//------------------------------------------------------------
	var path, filename string
	//------------------------------------------------------------
	path, filename = FilePathSplit()
	//------------------------------------------------------------
	if path == "" {

		t.Errorf(`path = %q`, path)
	}
	//----------
	if filename == "" {

		t.Errorf(`filename = %q`, filename)
	}
	//------------------------------------------------------------
	path, filename = FilePathSplit(`/path1/filename`)
	//------------------------------------------------------------
	if path != "/path1/" {

		t.Errorf(`path = %q`, path)
	}
	//----------
	if filename != "filename" {

		t.Errorf(`filename = %q`, filename)
	}
	//------------------------------------------------------------
	path, filename = FilePathSplit(`/path3/path4/`)
	//------------------------------------------------------------
	if path != "/path3/path4/" {

		t.Errorf(`path = %q`, path)
	}
	//----------
	if filename != "" {

		t.Errorf(`filename = %q`, filename)
	}
	//------------------------------------------------------------
	path, filename = FilePathSplit(`/path5/path6/filename.txt`)
	//------------------------------------------------------------
	if path != "/path5/path6/" {

		t.Errorf(`path = %q`, path)
	}
	//----------
	if filename != "filename.txt" {

		t.Errorf(`filename = %q`, filename)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// FilePathBaseToFilename
//------------------------------------------------------------

func TestFilePathBaseToFilename(t *testing.T) {

	//------------------------------------------------------------
	// runtime.Caller(0) => this script / runtime.Caller(1) => calling script
	_, filePath, _, _ := runtime.Caller(0)
	//------------------------------------------------------------
	filePath = FilePathBaseToFilename(filePath)
	//------------------------------------------------------------
	filePath = testTempGolangPath + `/` + filePath + `.txt`
	//------------------------------------------------------------
	if filePath != testDataFilename {

		t.Errorf("filename = %q but should = %q", filePath, testDataFilename)
	}
	//------------------------------------------------------------

	//------------------------------------------------------------
	filePath = FilePathBaseToFilename("/tmp/golang/test.txt")
	//------------------------------------------------------------
	if filePath != "tmp-golang-test" {

		t.Errorf("filename = %q but should = %q", filePath, "tmp-golang-test")
	}
	//------------------------------------------------------------

	//------------------------------------------------------------
	filePath = FilePathBaseToFilename("/tmp/golang/")
	//------------------------------------------------------------
	if filePath != "tmp-golang" {

		t.Errorf("filename = %q but should = %q", filePath, "tmp-golang")
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// FilePathBase
//------------------------------------------------------------

func TestFilePathBase(t *testing.T) {

	//------------------------------------------------------------
	var filePath string
	//------------------------------------------------------------

	//------------------------------------------------------------
	filePath = FilePathBase("/tmp/golang/test.txt")
	//------------------------------------------------------------
	if filePath != "/tmp/golang/test" {

		t.Errorf("filePath base = %q but should = %q", filePath, "tmp/golang/test")
	}
	//------------------------------------------------------------

	//------------------------------------------------------------
	filePath = FilePathBase("/tmp/golang/")
	//------------------------------------------------------------
	if filePath != "/tmp/golang" {

		t.Errorf("filePath base = %q but should = %q", filePath, "tmp/golang")
	}
	//------------------------------------------------------------

	//------------------------------------------------------------
	filePath = FilePathBase("/tmp/golang")
	//------------------------------------------------------------
	if filePath != "/tmp/golang" {

		t.Errorf("filePath base = %q but should = %q", filePath, "tmp/golang")
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// FilenameBase
//------------------------------------------------------------

func TestFilenameBase(t *testing.T) {

	//------------------------------------------------------------
	// runtime.Caller(0) => this script / runtime.Caller(1) => calling script
	_, filePath, _, _ := runtime.Caller(0)
	//------------------------------------------------------------
	filePath = FilenameBase(filePath)
	//------------------------------------------------------------
	if filePath != "file_test" {

		t.Errorf("filename base = %q but should = %q", filePath, "file_test")
	}
	//------------------------------------------------------------

	//------------------------------------------------------------
	filePath = FilenameBase("/tmp/golang/test")
	//------------------------------------------------------------
	if filePath != "test" {

		t.Errorf("filename base = %q but should = %q", filePath, "test")
	}
	//------------------------------------------------------------

	//------------------------------------------------------------
	filePath = FilenameBase("/tmp/golang/")
	//------------------------------------------------------------
	if filePath != "golang" {

		t.Errorf("filename base = %q but should = %q", filePath, "golang")
	}
	//------------------------------------------------------------

	//------------------------------------------------------------
	filePath = FilenameBase("/tmp/golang")
	//------------------------------------------------------------
	if filePath != "golang" {

		t.Errorf("filename base = %q but should = %q", filePath, "golang")
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Filename
//------------------------------------------------------------

func TestFilename(t *testing.T) {

	//------------------------------------------------------------
	// runtime.Caller(0) => this script / runtime.Caller(1) => calling script
	_, filePath, _, _ := runtime.Caller(0)
	//------------------------------------------------------------
	filePath = Filename(filePath)
	//------------------------------------------------------------
	if filePath != "file_test.go" {

		t.Errorf("filename = %q but should = %q", filePath, "file_test.go")
	}
	//------------------------------------------------------------

	//------------------------------------------------------------
	filePath = Filename("/tmp/golang/test.txt")
	//------------------------------------------------------------
	if filePath != "test.txt" {

		t.Errorf("filename = %q but should = %q", filePath, "test.txt")
	}
	//------------------------------------------------------------

	//------------------------------------------------------------
	filePath = Filename("/tmp/golang/")
	//------------------------------------------------------------
	if filePath != "" {

		t.Errorf("filename = %q but should = %q", filePath, "")
	}
	//------------------------------------------------------------

	//------------------------------------------------------------
	filePath = Filename("/tmp/golang")
	//------------------------------------------------------------
	if filePath != "golang" {

		t.Errorf("filename = %q but should = %q", filePath, "")
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// FilenameExt
//------------------------------------------------------------

func TestFilenameExt(t *testing.T) {

	//------------------------------------------------------------
	// runtime.Caller(0) => this script / runtime.Caller(1) => calling script
	_, filePath, _, _ := runtime.Caller(0)
	//------------------------------------------------------------
	filePath = FilenameExt(filePath)
	//------------------------------------------------------------
	if filePath != "go" {

		t.Errorf("filePath = %q but should = %q", filePath, "go")
	}
	//------------------------------------------------------------

	//------------------------------------------------------------
	filePath = FilenameExt("/tmp/golang/test.txt")
	//------------------------------------------------------------
	if filePath != "txt" {

		t.Errorf("extension = %q but should = %q", filePath, "txt")
	}
	//------------------------------------------------------------

	//------------------------------------------------------------
	filePath = FilenameExt("/tmp/golang/")
	//------------------------------------------------------------
	if filePath != "" {

		t.Errorf("extension = %q but should = %q", filePath, "")
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// FileSave
//------------------------------------------------------------

func TestFileSave(t *testing.T) {

	//------------------------------------------------------------
	err := FileSave(testDataFilename, "<TEST_DATA>")
	//----------
	if err != nil {

		t.Error(err)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// FileAppend
//------------------------------------------------------------

func TestFileAppend(t *testing.T) {

	//------------------------------------------------------------
	err := FileAppend(testDataFilename, "<TEST_DATA>")
	//------------------------------------------------------------
	if err != nil {
		t.Error(err)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// FileLoad
//------------------------------------------------------------

func TestFileLoad(t *testing.T) {

	//--------------------------------------------------------------------------------
	var err error
	//--------------------------------------------------------------------------------
	var dataString, restulString string
	//------------------------------------------------------------
	restulString = "<TEST_DATA><TEST_DATA>"
	//------------------------------------------------------------
	dataString, err = FileLoad(testDataFilename)
	//------------------------------------------------------------
	if err != nil {

		t.Error(err)

	} else {

		if dataString != restulString {

			t.Errorf(`dataString = %q but should = %q\n`, dataString, restulString)
		}
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// FileRemove
//------------------------------------------------------------

func TestFileRemove(t *testing.T) {

	//--------------------------------------------------------------------------------
	var err error
	//--------------------------------------------------------------------------------
	testTempFilename := testTempGolangPath + "/TEMP_FILE.txt"
	//--------------------------------------------------------------------------------
	err = FileSave(testTempFilename, "<TEST_DATA>")
	//----------
	if err != nil {

		t.Error(err)
	}
	//------------------------------------------------------------
	err = FileRemove(testTempFilename)
	//------------------------------------------------------------
	if err != nil {

		t.Error(err)

	} else {

		if FilePathExists(testTempFilename) {

			t.Error("error occurred while trying to remove file")
		}
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// TempPath
//------------------------------------------------------------

func TestTempPath(t *testing.T) {

	//------------------------------------------------------------
	tempPath, err := TempPath()
	//------------------------------------------------------------
	if err != nil {

		t.Error(err)
	}
	//----------
	if tempPath == "" {

		t.Errorf("tempPath should not = %q", "")
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// TempFilePath
//------------------------------------------------------------

func TestTempFilePath(t *testing.T) {

	//------------------------------------------------------------
	tempFilePath, err := TempFilePath()
	//------------------------------------------------------------
	if err != nil {

		t.Error(err)
	}
	//----------
	if tempFilePath == "" {

		t.Errorf("tempFilePath should not = %q", "")
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// LogFilePath
//------------------------------------------------------------

func TestLogFilePath(t *testing.T) {

	//------------------------------------------------------------
	logFilePath := LogFilePath()
	//------------------------------------------------------------
	if logFilePath != testLogFilename {

		t.Errorf("logFilePath = %q but should = %q", logFilePath, testLogFilename)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Log
//------------------------------------------------------------

func TestLog(t *testing.T) {

	//------------------------------------------------------------
	dataBytes1, _ := os.ReadFile(testLogFilename)
	//----------
	len1 := len(dataBytes1)
	//------------------------------------------------------------
	err := Log("<TEST_DATA>")
	//------------------------------------------------------------
	if err != nil {

		t.Error(err)

	} else {

		//----------
		dataBytes2, _ := os.ReadFile(testLogFilename)
		//----------
		len2 := len(dataBytes2)
		//----------
		if len1 >= len2 || len2 == 0 {

			t.Error("error occurred while trying to write to log file")
		}
		//----------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
