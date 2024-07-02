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
	resultString = CR(WithStdout) + LF(WithStdout) + CRLF(WithStdout)
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
	resultString = CursorUp(0, WithStdout)
	resultString += CursorDown(0, WithStdout)
	resultString += CursorRight(0, WithStdout)
	resultString += CursorLeft(0, WithStdout)
	resultString += CursorHome(WithStdout)
	resultString += CursorMove(0, 0, WithStdout)
	resultString += CursorSave(WithStdout)
	resultString += CursorRestore(WithStdout)
	resultString += CursorShow(WithStdout)
	resultString += CursorHide(WithStdout)
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
	resultString = ScrollUp(0, WithStdout)
	resultString += ScrollDown(0, WithStdout)
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
	expectedString = "\033[3J"               // ClearScrollbackBuffer
	expectedString += "\033[2J\033[H"        // ClearWindow
	expectedString += "\033[2J\033[3J\033[H" // ClearScreen
	expectedString += "\033[2K\r"            // ClearLine
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
	resultString = ClearScrollbackBuffer(WithStdout)
	resultString += ClearWindow(WithStdout)
	resultString += ClearScreen(WithStdout)
	resultString += ClearLine(WithStdout)
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

func TestEffects(t *testing.T) {
	//----------------------------------------
	var resultString, expectedString string
	//----------------------------------------
	expectedString = "\033[30m"  // Effect(Black)
	expectedString += "\033[31m" // Effect(Red)
	expectedString += "\033[32m" // Effect(Green)
	expectedString += "\033[33m" // Effect(Yellow)
	expectedString += "\033[34m" // Effect(Blue)
	expectedString += "\033[35m" // Effect(Magenta)
	expectedString += "\033[36m" // Effect(Cyan)
	expectedString += "\033[37m" // Effect(White)
	expectedString += "\033[39m" // Effect(Default)
	//----------------------------------------
	expectedString += "\033[40m" // Effect(BlackBackground)
	expectedString += "\033[41m" // Effect(RedBackground)
	expectedString += "\033[42m" // Effect(GreenBackground)
	expectedString += "\033[43m" // Effect(YellowBackground)
	expectedString += "\033[44m" // Effect(BlueBackground)
	expectedString += "\033[45m" // Effect(MagentaBackground)
	expectedString += "\033[46m" // Effect(CyanBackground)
	expectedString += "\033[47m" // Effect(WhiteBackground)
	expectedString += "\033[49m" // Effect(DefaultBackground)
	//---------------------------------------
	expectedString += "\033[90m" // Effect(BrightBlack)
	expectedString += "\033[91m" // Effect(BrightRed)
	expectedString += "\033[92m" // Effect(BrightGreen)
	expectedString += "\033[93m" // Effect(BrightYellow)
	expectedString += "\033[94m" // Effect(BrightBlue)
	expectedString += "\033[95m" // Effect(BrightMagenta)
	expectedString += "\033[96m" // Effect(BrightCyan)
	expectedString += "\033[97m" // Effect(BrightWhite)
	expectedString += "\033[99m" // Effect(BrightDefault)
	//----------------------------------------
	expectedString += "\033[100m" // Effect(BrightBlackBackground)
	expectedString += "\033[101m" // Effect(BrightRedBackground)
	expectedString += "\033[102m" // Effect(BrightGreenBackground)
	expectedString += "\033[103m" // Effect(BrightYellowBackground)
	expectedString += "\033[104m" // Effect(BrightBlueBackground)
	expectedString += "\033[105m" // Effect(BrightMagentaBackground)
	expectedString += "\033[106m" // Effect(BrightCyanBackground)
	expectedString += "\033[107m" // Effect(BrightWhiteBackground)
	expectedString += "\033[109m" // Effect(BrightDefaultBackground)
	//----------------------------------------
	expectedString += "\033[1m" // Effect(Bold)
	expectedString += "\033[2m" // Effect(Dim)
	expectedString += "\033[4m" // Effect(Underline)
	expectedString += "\033[5m" // Effect(Blink)
	expectedString += "\033[7m" // Effect(Reverse)
	expectedString += "\033[8m" // Effect(Hide)
	//----------------------------------------
	expectedString += "\033[21m" // Effect(BoldOff)
	expectedString += "\033[22m" // Effect(DimOff)
	expectedString += "\033[24m" // Effect(UnderlineOff)
	expectedString += "\033[25m" // Effect(BlinkOff)
	expectedString += "\033[27m" // Effect(ReverseOff)
	expectedString += "\033[28m" // Effect(HideOff)
	//----------------------------------------
	expectedString += "\033[0m" // ResetEffect()
	//----------------------------------------
	resultString = Effect(Black)
	resultString += Effect(Red)
	resultString += Effect(Green)
	resultString += Effect(Yellow)
	resultString += Effect(Blue)
	resultString += Effect(Magenta)
	resultString += Effect(Cyan)
	resultString += Effect(White)
	resultString += Effect(Default)
	//----------------------------------------
	resultString += Effect(BlackBackground)
	resultString += Effect(RedBackground)
	resultString += Effect(GreenBackground)
	resultString += Effect(YellowBackground)
	resultString += Effect(BlueBackground)
	resultString += Effect(MagentaBackground)
	resultString += Effect(CyanBackground)
	resultString += Effect(WhiteBackground)
	resultString += Effect(DefaultBackground)
	//----------------------------------------
	resultString += Effect(BrightBlack)
	resultString += Effect(BrightRed)
	resultString += Effect(BrightGreen)
	resultString += Effect(BrightYellow)
	resultString += Effect(BrightBlue)
	resultString += Effect(BrightMagenta)
	resultString += Effect(BrightCyan)
	resultString += Effect(BrightWhite)
	resultString += Effect(BrightDefault)
	//----------------------------------------
	resultString += Effect(BrightBlackBackground)
	resultString += Effect(BrightRedBackground)
	resultString += Effect(BrightGreenBackground)
	resultString += Effect(BrightYellowBackground)
	resultString += Effect(BrightBlueBackground)
	resultString += Effect(BrightMagentaBackground)
	resultString += Effect(BrightCyanBackground)
	resultString += Effect(BrightWhiteBackground)
	resultString += Effect(BrightDefaultBackground)
	//----------------------------------------
	resultString += Effect(Bold)
	resultString += Effect(Dim)
	resultString += Effect(Underline)
	resultString += Effect(Blink)
	resultString += Effect(Reverse)
	resultString += Effect(Hide)
	//----------------------------------------
	resultString += Effect(BoldOff)
	resultString += Effect(DimOff)
	resultString += Effect(UnderlineOff)
	resultString += Effect(BlinkOff)
	resultString += Effect(ReverseOff)
	resultString += Effect(HideOff)
	//----------------------------------------
	resultString += ResetEffect()
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %v but got: %v", []byte(expectedString), []byte(resultString))
	}
	//----------------------------------------
}

//------------------------------------------------------------
