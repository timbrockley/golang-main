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
	resultString = CR(map[string]any{"Writer": os.Stdout}) + LF(map[string]any{"Writer": os.Stdout}) + CRLF(map[string]any{"Writer": os.Stdout})
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
	expectedString = "\033[1A"    // CursorUp
	expectedString += "\033[1B"   // CursorDown
	expectedString += "\033[1C"   // CursorRight
	expectedString += "\033[1D"   // CursorLeft
	expectedString += "\033[H"    // CursorHome
	expectedString += "\033[1;1H" // CursorMove
	expectedString += "\033[s"    // CursorSave
	expectedString += "\033[u"    // CursorRestore
	expectedString += "\033[?25h" // CursorShow
	expectedString += "\033[?25l" // CursorHide
	//----------------------------------------
	resultString = CursorUp(1)
	resultString += CursorDown(1)
	resultString += CursorRight(1)
	resultString += CursorLeft(1)
	resultString += CursorHome()
	resultString += CursorMove(1, 1)
	resultString += CursorSave()
	resultString += CursorRestore()
	resultString += CursorShow()
	resultString += CursorHide()
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %v but got: %v", []byte(expectedString), []byte(resultString))
	}
	//----------------------------------------
	resultString = CursorUp(1, map[string]any{"Writer": os.Stdout})
	resultString += CursorDown(1, map[string]any{"Writer": os.Stdout})
	resultString += CursorRight(1, map[string]any{"Writer": os.Stdout})
	resultString += CursorLeft(1, map[string]any{"Writer": os.Stdout})
	resultString += CursorHome(map[string]any{"Writer": os.Stdout})
	resultString += CursorMove(1, 1, map[string]any{"Writer": os.Stdout})
	resultString += CursorSave(map[string]any{"Writer": os.Stdout})
	resultString += CursorRestore(map[string]any{"Writer": os.Stdout})
	resultString += CursorShow(map[string]any{"Writer": os.Stdout})
	resultString += CursorHide(map[string]any{"Writer": os.Stdout})
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
	expectedString = "\033[1S"  // ScrollUp
	expectedString += "\033[1T" // ScrollDown
	//----------------------------------------
	resultString = ScrollUp(1)
	resultString += ScrollDown(1)
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %v but got: %v", []byte(expectedString), []byte(resultString))
	}
	//----------------------------------------
	resultString = ScrollUp(1, map[string]any{"Writer": os.Stdout})
	resultString += ScrollDown(1, map[string]any{"Writer": os.Stdout})
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
	expectedString += "\033[2K"              // ClearLine
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
	resultString = ClearScrollbackBuffer(map[string]any{"Writer": os.Stdout})
	resultString += ClearWindow(map[string]any{"Writer": os.Stdout})
	resultString += ClearScreen(map[string]any{"Writer": os.Stdout})
	resultString += ClearLine(map[string]any{"Writer": os.Stdout})
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

func TestAlternativeScreenFunctions(t *testing.T) {
	//----------------------------------------
	var resultString, expectedString string
	//----------------------------------------
	oldStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	//----------------------------------------
	expectedString = "\033[?1049h"  // AlternativeScreenEnable
	expectedString += "\033[?1049l" // AlternativeScreenDisable
	//----------------------------------------
	resultString = AlternativeScreenEnable()
	resultString += AlternativeScreenDisable()
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %v but got: %v", []byte(expectedString), []byte(resultString))
	}
	//----------------------------------------
	resultString = AlternativeScreenEnable(map[string]any{"Writer": os.Stdout})
	resultString += AlternativeScreenDisable(map[string]any{"Writer": os.Stdout})
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

func TestColour(t *testing.T) {
	//----------------------------------------
	var resultString, expectedString string
	//----------------------------------------
	expectedString = "\033[30m"  // Colour(Black)
	expectedString += "\033[31m" // Colour(Red)
	expectedString += "\033[32m" // Colour(Green)
	expectedString += "\033[33m" // Colour(Yellow)
	expectedString += "\033[34m" // Colour(Blue)
	expectedString += "\033[35m" // Colour(Magenta)
	expectedString += "\033[36m" // Colour(Cyan)
	expectedString += "\033[37m" // Colour(White)
	expectedString += "\033[39m" // Colour(Default)
	//----------------------------------------
	expectedString += "\033[90m" // Colour(BrightBlack)
	expectedString += "\033[91m" // Colour(BrightRed)
	expectedString += "\033[92m" // Colour(BrightGreen)
	expectedString += "\033[93m" // Colour(BrightYellow)
	expectedString += "\033[94m" // Colour(BrightBlue)
	expectedString += "\033[95m" // Colour(BrightMagenta)
	expectedString += "\033[96m" // Colour(BrightCyan)
	expectedString += "\033[97m" // Colour(BrightWhite)
	expectedString += "\033[99m" // Colour(BrightDefault)
	//----------------------------------------
	expectedString += "\033[0m" // Reset()
	//----------------------------------------
	resultString = Colour(Black)
	resultString += Colour(Red)
	resultString += Colour(Green)
	resultString += Colour(Yellow)
	resultString += Colour(Blue)
	resultString += Colour(Magenta)
	resultString += Colour(Cyan)
	resultString += Colour(White)
	resultString += Colour(Default)
	//----------------------------------------
	resultString += Colour(BrightBlack)
	resultString += Colour(BrightRed)
	resultString += Colour(BrightGreen)
	resultString += Colour(BrightYellow)
	resultString += Colour(BrightBlue)
	resultString += Colour(BrightMagenta)
	resultString += Colour(BrightCyan)
	resultString += Colour(BrightWhite)
	resultString += Colour(BrightDefault)
	//----------------------------------------
	resultString += Reset()
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %v but got: %v", []byte(expectedString), []byte(resultString))
	}
	//----------------------------------------
}

//------------------------------------------------------------

func TestBackgroundColour(t *testing.T) {
	//----------------------------------------
	var resultString, expectedString string
	//----------------------------------------
	expectedString = "\033[40m"  // Colour(Black, true)
	expectedString += "\033[41m" // Colour(Red, true)
	expectedString += "\033[42m" // Colour(Green, true)
	expectedString += "\033[43m" // Colour(Yellow, true)
	expectedString += "\033[44m" // Colour(Blue, true)
	expectedString += "\033[45m" // Colour(Magenta, true)
	expectedString += "\033[46m" // Colour(Cyan, true)
	expectedString += "\033[47m" // Colour(White, true)
	expectedString += "\033[49m" // Colour(Default, true)
	//----------------------------------------
	expectedString += "\033[100m" // Colour(BrightBlack, true)
	expectedString += "\033[101m" // Colour(BrightRed, true)
	expectedString += "\033[102m" // Colour(BrightGreen, true)
	expectedString += "\033[103m" // Colour(BrightYellow, true)
	expectedString += "\033[104m" // Colour(BrightBlue, true)
	expectedString += "\033[105m" // Colour(BrightMagenta, true)
	expectedString += "\033[106m" // Colour(BrightCyan, true)
	expectedString += "\033[107m" // Colour(BrightWhite, true)
	expectedString += "\033[109m" // Colour(BrightDefault, true)
	//----------------------------------------
	expectedString += "\033[0m" // Reset()
	//----------------------------------------
	resultString = Colour(Black, true)
	resultString += Colour(Red, true)
	resultString += Colour(Green, true)
	resultString += Colour(Yellow, true)
	resultString += Colour(Blue, true)
	resultString += Colour(Magenta, true)
	resultString += Colour(Cyan, true)
	resultString += Colour(White, true)
	resultString += Colour(Default, true)
	//----------------------------------------
	resultString += Colour(BrightBlack, true)
	resultString += Colour(BrightRed, true)
	resultString += Colour(BrightGreen, true)
	resultString += Colour(BrightYellow, true)
	resultString += Colour(BrightBlue, true)
	resultString += Colour(BrightMagenta, true)
	resultString += Colour(BrightCyan, true)
	resultString += Colour(BrightWhite, true)
	resultString += Colour(BrightDefault, true)
	//----------------------------------------
	resultString += Reset()
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %v but got: %v", []byte(expectedString), []byte(resultString))
	}
	//----------------------------------------
}

//------------------------------------------------------------

func TestColour256(t *testing.T) {
	//----------------------------------------
	var resultString, expectedString string
	//----------------------------------------
	expectedString = "\033[38;5;0m" // Colour256(0)
	//----------------------------------------
	expectedString += "\033[48;5;0m" // Colour256(0, true)
	//----------------------------------------
	expectedString += "\033[0m" // Reset()
	//----------------------------------------
	resultString = Colour256(0)
	//----------------------------------------
	resultString += Colour256(0, true)
	//----------------------------------------
	resultString += Reset()
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %v but got: %v", []byte(expectedString), []byte(resultString))
	}
	//----------------------------------------
}

//------------------------------------------------------------

func TestColourRGB(t *testing.T) {
	//----------------------------------------
	var resultString, expectedString string
	//----------------------------------------
	expectedString = "\033[38;2;255;255;255m" // ColourRGB(0xFF, 0xFF, 0xFF)
	//----------------------------------------
	expectedString += "\033[48;2;255;255;255m" // ColourRGB(0xFF, 0xFF, 0xFF, true)
	//----------------------------------------
	expectedString += "\033[0m" // Reset()
	//----------------------------------------
	resultString = ColourRGB(0xFF, 0xFF, 0xFF)
	//----------------------------------------
	resultString += ColourRGB(0xFF, 0xFF, 0xFF, true)
	//----------------------------------------
	resultString += Reset()
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %v but got: %v", []byte(expectedString), []byte(resultString))
	}
	//----------------------------------------
}

//------------------------------------------------------------

func TestEffect(t *testing.T) {
	//----------------------------------------
	var resultString, expectedString string
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
	expectedString += "\033[0m" // Reset()
	//----------------------------------------
	resultString = Effect(Bold)
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
	resultString += Reset()
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %v but got: %v", []byte(expectedString), []byte(resultString))
	}
	//----------------------------------------
}

//------------------------------------------------------------
