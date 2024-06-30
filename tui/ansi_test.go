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
	expectedString = "0x1B[0A"
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
	expectedString = "0x1B[0B"
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

func TestCursorRight(t *testing.T) {
	//----------------------------------------
	var resultString, expectedString string
	//----------------------------------------
	oldStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	//----------------------------------------
	expectedString = "0x1B[0C"
	//----------------------------------------
	resultString = CursorRight(0)
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %v but got: %v", []byte(expectedString), []byte(resultString))
	}
	//----------------------------------------
	resultString = CursorRight(0, Stdout)
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

func TestCursorLeft(t *testing.T) {
	//----------------------------------------
	var resultString, expectedString string
	//----------------------------------------
	oldStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	//----------------------------------------
	expectedString = "0x1B[0D"
	//----------------------------------------
	resultString = CursorLeft(0)
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %v but got: %v", []byte(expectedString), []byte(resultString))
	}
	//----------------------------------------
	resultString = CursorLeft(0, Stdout)
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

func TestCursorHome(t *testing.T) {
	//----------------------------------------
	var resultString, expectedString string
	//----------------------------------------
	oldStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	//----------------------------------------
	expectedString = "0x1B[H"
	//----------------------------------------
	resultString = CursorHome()
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %v but got: %v", []byte(expectedString), []byte(resultString))
	}
	//----------------------------------------
	resultString = CursorHome(Stdout)
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

func TestCursorMove(t *testing.T) {
	//----------------------------------------
	var resultString, expectedString string
	//----------------------------------------
	oldStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	//----------------------------------------
	expectedString = "0x1B[0;0H"
	//----------------------------------------
	resultString = CursorMove(0, 0)
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %v but got: %v", []byte(expectedString), []byte(resultString))
	}
	//----------------------------------------
	resultString = CursorMove(0, 0, Stdout)
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

func TestScrollUp(t *testing.T) {
	//----------------------------------------
	var resultString, expectedString string
	//----------------------------------------
	oldStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	//----------------------------------------
	expectedString = "0x1B[0S"
	//----------------------------------------
	resultString = ScrollUp(0)
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %v but got: %v", []byte(expectedString), []byte(resultString))
	}
	//----------------------------------------
	resultString = ScrollUp(0, Stdout)
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

func TestScrollDown(t *testing.T) {
	//----------------------------------------
	var resultString, expectedString string
	//----------------------------------------
	oldStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	//----------------------------------------
	expectedString = "0x1B[0T"
	//----------------------------------------
	resultString = ScrollDown(0)
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %v but got: %v", []byte(expectedString), []byte(resultString))
	}
	//----------------------------------------
	resultString = ScrollDown(0, Stdout)
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

func TestClearScreen(t *testing.T) {
	//----------------------------------------
	var resultString, expectedString string
	//----------------------------------------
	oldStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	//----------------------------------------
	expectedString = "0x1B[2J"
	//----------------------------------------
	resultString = ClearScreen(false)
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %s but got: %s", expectedString, resultString)
	}
	//----------------------------------------
	expectedString = "0x1B[3J"
	//----------------------------------------
	resultString = ClearScreen(true, Stdout)
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

func TestClearLine(t *testing.T) {
	//----------------------------------------
	var resultString, expectedString string
	//----------------------------------------
	oldStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	//----------------------------------------
	expectedString = "0x1B[2K"
	//----------------------------------------
	resultString = ClearLine()
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %s but got: %s", expectedString, resultString)
	}
	//----------------------------------------
	resultString = ClearLine(Stdout)
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

func TestCursorShow(t *testing.T) {
	//----------------------------------------
	var resultString, expectedString string
	//----------------------------------------
	oldStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	//----------------------------------------
	expectedString = "0x1B[?25h"
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

func TestCursorHide(t *testing.T) {
	//----------------------------------------
	var resultString, expectedString string
	//----------------------------------------
	oldStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	//----------------------------------------
	expectedString = "0x1B[?25l"
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
