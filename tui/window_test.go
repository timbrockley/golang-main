package tui

import (
	"io"
	"os"
	"testing"
)

//------------------------------------------------------------

func TestRenderWindow1(t *testing.T) {
	//----------------------------------------
	var resultString, expectedString string
	//----------------------------------------
	oldStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	//----------------------------------------
	expectedString = "\033[1;1H"  // CursorMove(1,1)
	expectedString += "\u250C"    // TopLeft
	expectedString += "\u2510"    // TopRight
	expectedString += "\033[2;1H" // CursorMove(2,1)
	expectedString += "\u2514"    // BottomLeft
	expectedString += "\u2518"    // BottomRight
	//----------------------------------------
	resultString = RenderWindow([]string{}, map[string]any{"Writer": os.Stdout})
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

func TestRenderWindow2(t *testing.T) {
	//----------------------------------------
	var resultString, expectedString string
	//----------------------------------------
	oldStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	//----------------------------------------
	// Horizontal:   "\u2500",
	// Vertical:     "\u2502",

	expectedString = "\033[1;1H"  // CursorMove(1,1)
	expectedString += "\u250C"    // TopLeft
	expectedString += "\u2510"    // TopRight
	expectedString += "\033[2;1H" // CursorMove(2,1)
	expectedString += "\u2514"    // BottomLeft
	expectedString += "\u2518"    // BottomRight
	//----------------------------------------
	resultString = RenderWindow([]string{}, map[string]any{"MaxWidth": 2, "MaxHeight": 2, "Writer": os.Stdout})
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
