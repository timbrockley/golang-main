package tui

import (
	"io"
	"os"
	"testing"
)

//------------------------------------------------------------

func TestCR(t *testing.T) {
	//----------------------------------------
	var resultString, expectedString string
	//----------------------------------------
	oldStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	//----------------------------------------
	expectedString = "\r"
	//----------------------------------------
	resultString = CR()
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %s but got: %s", expectedString, resultString)
	}
	//----------------------------------------
	resultString = CR(Stdout)
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %s but got: %s", expectedString, resultString)
	}
	//----------------------------------------
	writer.Close()
	capturedStdout, _ := io.ReadAll(reader)
	resultString = string(capturedStdout)
	os.Stdout = oldStdout
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %s but got: %s", expectedString, resultString)
	}
	//----------------------------------------
}

//------------------------------------------------------------

func TestLF(t *testing.T) {
	//----------------------------------------
	var resultString, expectedString string
	//----------------------------------------
	oldStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	//----------------------------------------
	expectedString = "\n"
	//----------------------------------------
	resultString = LF()
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %s but got: %s", expectedString, resultString)
	}
	//----------------------------------------
	resultString = LF(Stdout)
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %s but got: %s", expectedString, resultString)
	}
	//----------------------------------------
	writer.Close()
	capturedStdout, _ := io.ReadAll(reader)
	resultString = string(capturedStdout)
	os.Stdout = oldStdout
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %s but got: %s", expectedString, resultString)
	}
	//----------------------------------------
}

//------------------------------------------------------------

func TestCRLF(t *testing.T) {
	//----------------------------------------
	var resultString, expectedString string
	//----------------------------------------
	oldStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	//----------------------------------------
	expectedString = "\r\n"
	//----------------------------------------
	resultString = CRLF()
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %s but got: %s", expectedString, resultString)
	}
	//----------------------------------------
	resultString = CRLF(Stdout)
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %s but got: %s", expectedString, resultString)
	}
	//----------------------------------------
	writer.Close()
	capturedStdout, _ := io.ReadAll(reader)
	resultString = string(capturedStdout)
	os.Stdout = oldStdout
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %s but got: %s", expectedString, resultString)
	}
	//----------------------------------------
}

//------------------------------------------------------------

func TestCursorUp(t *testing.T) {
	//----------------------------------------
	var resultString, expectedString string
	//----------------------------------------
	oldStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	//----------------------------------------
	expectedString = "\033[0A"
	//----------------------------------------
	resultString = CursorUp(0)
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %v but got: %v", []byte(expectedString), []byte(resultString))
	}
	//----------------------------------------
	resultString = CursorUp(0, Stdout)
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %v but got: %v", []byte(expectedString), []byte(resultString))
	}
	//----------------------------------------
	writer.Close()
	capturedStdout, _ := io.ReadAll(reader)
	os.Stdout = oldStdout
	//----------------------------------------
	if string(capturedStdout) != expectedString {
		t.Errorf("expected: %v but got: %v", []byte(expectedString), capturedStdout)
	}
	//----------------------------------------
}

//------------------------------------------------------------

func TestCursorDown(t *testing.T) {
	//----------------------------------------
	var resultString, expectedString string
	//----------------------------------------
	oldStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	//----------------------------------------
	expectedString = "\033[0B"
	//----------------------------------------
	resultString = CursorDown(0)
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %v but got: %v", []byte(expectedString), []byte(resultString))
	}
	//----------------------------------------
	resultString = CursorDown(0, Stdout)
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %v but got: %v", []byte(expectedString), []byte(resultString))
	}
	//----------------------------------------
	writer.Close()
	capturedStdout, _ := io.ReadAll(reader)
	os.Stdout = oldStdout
	//----------------------------------------
	if string(capturedStdout) != expectedString {
		t.Errorf("expected: %v but got: %v", []byte(expectedString), capturedStdout)
	}
	//----------------------------------------
}

//------------------------------------------------------------

func TestCursorHide(t *testing.T) {
	//----------------------------------------
	var resultString, expectedString string
	//----------------------------------------
	oldStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	//----------------------------------------
	expectedString = "\033[?25l"
	//----------------------------------------
	resultString = CursorHide()
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %v but got: %v", []byte(expectedString), []byte(resultString))
	}
	//----------------------------------------
	resultString = CursorHide(Stdout)
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %v but got: %v", []byte(expectedString), []byte(resultString))
	}
	//----------------------------------------
	writer.Close()
	capturedStdout, _ := io.ReadAll(reader)
	os.Stdout = oldStdout
	//----------------------------------------
	if string(capturedStdout) != expectedString {
		t.Errorf("expected: %v but got: %v", []byte(expectedString), capturedStdout)
	}
	//----------------------------------------
}

//------------------------------------------------------------

func TestCursorShow(t *testing.T) {
	//----------------------------------------
	var resultString, expectedString string
	//----------------------------------------
	oldStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	//----------------------------------------
	expectedString = "\033[?25h"
	//----------------------------------------
	resultString = CursorShow()
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %v but got: %v", []byte(expectedString), []byte(resultString))
	}
	//----------------------------------------
	resultString = CursorShow(Stdout)
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %v but got: %v", []byte(expectedString), []byte(resultString))
	}
	//----------------------------------------
	writer.Close()
	capturedStdout, _ := io.ReadAll(reader)
	os.Stdout = oldStdout
	//----------------------------------------
	if string(capturedStdout) != expectedString {
		t.Errorf("expected: %v but got: %v", []byte(expectedString), capturedStdout)
	}
	//----------------------------------------
}

//------------------------------------------------------------
