package tui

import (
	"io"
	"os"
	"testing"
)

//------------------------------------------------------------

func TestReturnOutput(t *testing.T) {
	//----------------------------------------
	var resultString, expectedString string
	//----------------------------------------
	oldStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	//----------------------------------------
	expectedString = "XXXXXXXXXX"
	//----------------------------------------
	resultString = ReturnOutput("XXXXXXXXXX")
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %v but got: %v", []byte(expectedString), []byte(resultString))
	}
	//----------------------------------------
	expectedString = "XXXXX"
	//----------------------------------------
	resultString = ReturnOutput("XXXXXXXXXX", Stdout, MaxWidth(5))
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
