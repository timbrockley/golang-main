package tui

import (
	"io"
	"os"
	"testing"
)

//------------------------------------------------------------

func TestRender(t *testing.T) {
	//----------------------------------------
	var resultString1, expectedString1 string
	//----------------------------------------
	oldStdout1 := os.Stdout
	reader1, writer1, _ := os.Pipe()
	os.Stdout = writer1
	//----------------------------------------
	expectedString1 = "XXXXXXXXXX"
	//----------------------------------------
	resultString1 = Render("XXXXXXXXXX")
	//----------------------------------------
	if resultString1 != expectedString1 {
		t.Errorf("expected: %v but got: %v", []byte(expectedString1), []byte(resultString1))
	}
	//----------------------------------------
	expectedString1 = "XXXXX"
	//----------------------------------------
	resultString1 = Render("XXXXXXXXXX", map[string]any{"Writer": os.Stdout, "MaxWidth": 5})
	//----------------------------------------
	if resultString1 != expectedString1 {
		t.Errorf("expected: %v but got: %v", []byte(expectedString1), []byte(resultString1))
	}
	//----------------------------------------
	writer1.Close()
	capturedStdout1, _ := io.ReadAll(reader1)
	os.Stdout = oldStdout1
	//----------------------------------------
	if string(capturedStdout1) != expectedString1 {
		t.Errorf("expected: %v but got: %v", []byte(expectedString1), capturedStdout1)
	}
	//----------------------------------------
	var resultString2, expectedString2 string
	//----------------------------------------
	reader2, writer2, _ := os.Pipe()
	//----------------------------------------
	expectedString2 = "XXXXXXXXXX"
	//----------------------------------------
	resultString2 = Render("XXXXXXXXXX", map[string]any{"Writer": writer2})
	//----------------------------------------
	if resultString2 != expectedString2 {
		t.Errorf("expected: %v but got: %v", []byte(expectedString2), []byte(resultString2))
	}
	//----------------------------------------
	writer2.Close()
	capturedStdout2, _ := io.ReadAll(reader2)
	//----------------------------------------
	if string(capturedStdout2) != expectedString2 {
		t.Errorf("expected: %v but got: %v", []byte(expectedString2), capturedStdout2)
	}
	//----------------------------------------
}

//------------------------------------------------------------
