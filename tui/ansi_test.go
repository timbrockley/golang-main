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
	expectedString = "\x1B[0A"    // CursorUp
	expectedString += "\x1B[0B"   // CursorDown
	expectedString += "\x1B[0C"   // CursorRight
	expectedString += "\x1B[0D"   // CursorLeft
	expectedString += "\x1B[H"    // CursorHome
	expectedString += "\x1B[0;0H" // CursorMove
	expectedString += "\x1B[s"    // CursorSave
	expectedString += "\x1B[u"    // CursorRestore
	expectedString += "\x1B[?25h" // CursorShow
	expectedString += "\x1B[?25l" // CursorHide
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

/*

	2024-07-01: worked ok when run in non-testing environment but failed in testing

	func TestCursorRowCol(t *testing.T) {
		//----------------------------------------
		_,_, err := CursorRowCol()
		if err != nil {
			t.Error(err)
		}
		//----------------------------------------
	}

*/

//------------------------------------------------------------

func TestScrollbackFunctions(t *testing.T) {
	//----------------------------------------
	var resultString, expectedString string
	//----------------------------------------
	oldStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	//----------------------------------------
	expectedString = "\x1B[0S"  // ScrollUp
	expectedString += "\x1B[0T" // ScrollDown
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
	expectedString = "\x1B[3J"               // ClearScrollbackBuffer
	expectedString += "\x1B[2J\x1B[H"        // ClearWindow
	expectedString += "\x1B[2J\x1B[3J\x1B[H" // ClearScreen
	expectedString += "\x1B[2K\r"            // ClearLine
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
