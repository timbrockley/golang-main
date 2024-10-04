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
	expectedString = string([]byte{0x1B, 0x5B, 0x31, 0x3B, 0x31, 0x48, 0x1B, 0x5B, 0x30, 0x6D, 0xE2, 0x94, 0x8C, 0xE2, 0x94, 0x90, 0x1B, 0x5B, 0x32, 0x3B, 0x31, 0x48, 0x1B, 0x5B, 0x30, 0x6D, 0xE2, 0x94, 0x94, 0xE2, 0x94, 0x98})
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
	expectedString = string([]byte{0x1B, 0x5B, 0x31, 0x3B, 0x31, 0x48, 0x1B, 0x5B, 0x30, 0x6D, 0xE2, 0x94, 0x8C, 0xE2, 0x94, 0x90, 0x1B, 0x5B, 0x32, 0x3B, 0x31, 0x48, 0x1B, 0x5B, 0x30, 0x6D, 0xE2, 0x94, 0x94, 0xE2, 0x94, 0x98})

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
