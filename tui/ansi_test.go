package tui

import (
	"io"
	"os"
	"testing"
)

//------------------------------------------------------------

func TestCR_LF_CRLF(t *testing.T) {
	//----------------------------------------
	var resultString, expectedString string
	//----------------------------------------
	oldStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	//----------------------------------------
	expectedString = "\r\n\r\n"
	//----------------------------------------
	resultString = CR() + LF() + CRLF()
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %v but got: %v", []byte(expectedString), []byte(resultString))
	}
	//----------------------------------------
	resultString = CR(Stdout) + LF(Stdout) + CRLF(Stdout)
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

func TestCursorFunctions(t *testing.T) {
	//----------------------------------------
	var resultString, expectedString string
	//----------------------------------------
	oldStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	//----------------------------------------
	expectedString = "\033[0A"    // CursorUp
	expectedString += "\033[0B"   // CursorDown
	expectedString += "\033[0C"   // CursorRight
	expectedString += "\033[0D"   // CursorLeft
	expectedString += "\033[H"    // CursorHome
	expectedString += "\033[0;0H" // CursorMove
	expectedString += "\033[s"    // CursorSave
	expectedString += "\033[u"    // CursorRestore
	expectedString += "\033[?25h" // CursorShow
	expectedString += "\033[?25l" // CursorHide
	//----------------------------------------
	resultString = CursorUp(0)
	resultString += CursorDown(0)
	resultString += CursorRight(0)
	resultString += CursorLeft(0)
	resultString += CursorHome()
	resultString += CursorMove(0, 0)
	resultString += CursorSave()
	resultString += CursorRestore()
	resultString += CursorShow()
	resultString += CursorHide()
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %v but got: %v", []byte(expectedString), []byte(resultString))
	}
	//----------------------------------------
	resultString = CursorUp(0, Stdout)
	resultString += CursorDown(0, Stdout)
	resultString += CursorRight(0, Stdout)
	resultString += CursorLeft(0, Stdout)
	resultString += CursorHome(Stdout)
	resultString += CursorMove(0, 0, Stdout)
	resultString += CursorSave(Stdout)
	resultString += CursorRestore(Stdout)
	resultString += CursorShow(Stdout)
	resultString += CursorHide(Stdout)
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

func TestScrollbackFunctions(t *testing.T) {
	//----------------------------------------
	var resultString, expectedString string
	//----------------------------------------
	oldStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	//----------------------------------------
	expectedString = "\033[0S"  // ScrollUp
	expectedString += "\033[0T" // ScrollDown
	//----------------------------------------
	resultString = ScrollUp(0)
	resultString += ScrollDown(0)
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %v but got: %v", []byte(expectedString), []byte(resultString))
	}
	//----------------------------------------
	resultString = ScrollUp(0, Stdout)
	resultString += ScrollDown(0, Stdout)
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

func TestClearFunctions(t *testing.T) {
	//----------------------------------------
	var resultString, expectedString string
	//----------------------------------------
	oldStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	//----------------------------------------
	expectedString = "\033[3J"         // ClearScrollbackBuffer
	expectedString += "\033[2J"        // ClearWindow
	expectedString += "\033[2J\033[3J" // ClearScreen
	expectedString += "\033[2K"        // ClearLine
	//----------------------------------------
	resultString = ClearScrollbackBuffer()
	resultString += ClearWindow()
	resultString += ClearScreen()
	resultString += ClearLine()
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %v but got: %v", []byte(expectedString), []byte(resultString))
	}
	//----------------------------------------
	resultString = ClearScrollbackBuffer(Stdout)
	resultString += ClearWindow(Stdout)
	resultString += ClearScreen(Stdout)
	resultString += ClearLine(Stdout)
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
