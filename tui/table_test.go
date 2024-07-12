package tui

import (
	"strings"
	"testing"
)

//------------------------------------------------------------

func TestRenderTable1(t *testing.T) {
	//----------------------------------------
	resultString := RenderTable([][]string{})
	//----------------------------------------
	expectedString := UnicodeTableStyle.TopLeft + UnicodeTableStyle.TopRight + "\n" + UnicodeTableStyle.BottomLeft + UnicodeTableStyle.BottomRight + "\n"
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %s but got: %s", expectedString, resultString)
	}
	//----------------------------------------
}

//------------------------------------------------------------

func TestRenderTable2(t *testing.T) {
	//----------------------------------------
	resultString := RenderTable([][]string{}, WithHeader)
	//----------------------------------------
	expectedString := UnicodeTableStyle.TopLeft + UnicodeTableStyle.TopRight + "\n" + UnicodeTableStyle.BottomLeft + UnicodeTableStyle.BottomRight + "\n"
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %s but got: %s", expectedString, resultString)
	}
	//----------------------------------------
}

//------------------------------------------------------------

func TestRenderTable3(t *testing.T) {
	//----------------------------------------
	resultString := RenderTable([][]string{
		{"John Doe", "30", "USA"},
		{"Jane Smith", "25", "Canada"},
	})
	//----------------------------------------
	expectedString := UnicodeTableStyle.TopLeft + strings.Repeat(UnicodeTableStyle.Horizontal, 12) + UnicodeTableStyle.TopMiddle + strings.Repeat(UnicodeTableStyle.Horizontal, 4) + UnicodeTableStyle.TopMiddle + strings.Repeat(UnicodeTableStyle.Horizontal, 8) + UnicodeTableStyle.TopRight + "\n"
	expectedString += UnicodeTableStyle.Vertical + " John Doe   " + UnicodeTableStyle.Vertical + " 30 " + UnicodeTableStyle.Vertical + " USA    " + UnicodeTableStyle.Vertical + "\n"
	expectedString += UnicodeTableStyle.Vertical + " Jane Smith " + UnicodeTableStyle.Vertical + " 25 " + UnicodeTableStyle.Vertical + " Canada " + UnicodeTableStyle.Vertical + "\n"
	expectedString += UnicodeTableStyle.BottomLeft + strings.Repeat(UnicodeTableStyle.Horizontal, 12) + UnicodeTableStyle.BottomMiddle + strings.Repeat(UnicodeTableStyle.Horizontal, 4) + UnicodeTableStyle.BottomMiddle + strings.Repeat(UnicodeTableStyle.Horizontal, 8) + UnicodeTableStyle.BottomRight + "\n"
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %s but got: %s", expectedString, resultString)
	}
	//----------------------------------------
}

//------------------------------------------------------------

func TestRenderTable4(t *testing.T) {
	//----------------------------------------
	resultString := RenderTable([][]string{
		{"Name", "Age", "Country"},
		{"John Doe", "30", "USA"},
		{"Jane Smith", "25", "Canada"},
	}, WithHeader)
	//----------------------------------------
	expectedString := UnicodeTableStyle.TopLeft + strings.Repeat(UnicodeTableStyle.Horizontal, 12) + UnicodeTableStyle.TopMiddle + strings.Repeat(UnicodeTableStyle.Horizontal, 5) + UnicodeTableStyle.TopMiddle + strings.Repeat(UnicodeTableStyle.Horizontal, 9) + UnicodeTableStyle.TopRight + "\n"
	expectedString += UnicodeTableStyle.Vertical + " Name       " + UnicodeTableStyle.Vertical + " Age " + UnicodeTableStyle.Vertical + " Country " + UnicodeTableStyle.Vertical + "\n"
	expectedString += UnicodeTableStyle.InnerLeft + strings.Repeat(UnicodeTableStyle.Horizontal, 12) + UnicodeTableStyle.InnerMiddle + strings.Repeat(UnicodeTableStyle.Horizontal, 5) + UnicodeTableStyle.InnerMiddle + strings.Repeat(UnicodeTableStyle.Horizontal, 9) + UnicodeTableStyle.InnerRight + "\n"
	expectedString += UnicodeTableStyle.Vertical + " John Doe   " + UnicodeTableStyle.Vertical + " 30  " + UnicodeTableStyle.Vertical + " USA     " + UnicodeTableStyle.Vertical + "\n"
	expectedString += UnicodeTableStyle.Vertical + " Jane Smith " + UnicodeTableStyle.Vertical + " 25  " + UnicodeTableStyle.Vertical + " Canada  " + UnicodeTableStyle.Vertical + "\n"
	expectedString += UnicodeTableStyle.BottomLeft + strings.Repeat(UnicodeTableStyle.Horizontal, 12) + UnicodeTableStyle.BottomMiddle + strings.Repeat(UnicodeTableStyle.Horizontal, 5) + UnicodeTableStyle.BottomMiddle + strings.Repeat(UnicodeTableStyle.Horizontal, 9) + UnicodeTableStyle.BottomRight + "\n"
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %s but got: %s", expectedString, resultString)
	}
	//----------------------------------------
}

//------------------------------------------------------------

func TestRenderTable5(t *testing.T) {
	//----------------------------------------
	resultString := RenderTable([][]string{
		{"Name", "Age", "Country"},
		{"John Doe", "30", "USA"},
		{"Jane Smith", "25", "Canada"},
	}, WithHeader, WithMaxTableWidth(10))
	//----------------------------------------
	expectedString := UnicodeTableStyle.TopLeft + strings.Repeat(UnicodeTableStyle.Horizontal, 9) + "\n"
	expectedString += UnicodeTableStyle.Vertical + " Name    " + "\n"
	expectedString += UnicodeTableStyle.InnerLeft + strings.Repeat(UnicodeTableStyle.Horizontal, 9) + "\n"
	expectedString += UnicodeTableStyle.Vertical + " John Doe" + "\n"
	expectedString += UnicodeTableStyle.Vertical + " Jane Smi" + "\n"
	expectedString += UnicodeTableStyle.BottomLeft + strings.Repeat(UnicodeTableStyle.Horizontal, 9) + "\n"
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %s but got: %s", expectedString, resultString)
	}
	//----------------------------------------
}

//------------------------------------------------------------

func TestRenderTable6(t *testing.T) {
	//----------------------------------------
	resultString := RenderTable([][]string{
		{"Name", "Age", "Country"},
		{"John Doe", "30", "USA"},
		{"Jane Smith", "25", "Canada"},
	}, WithHeader, WithMaxColumnWidth(1), WithPadding(0))
	//----------------------------------------
	expectedString := UnicodeTableStyle.TopLeft + strings.Repeat(UnicodeTableStyle.Horizontal, 1) + UnicodeTableStyle.TopMiddle + strings.Repeat(UnicodeTableStyle.Horizontal, 1) + UnicodeTableStyle.TopMiddle + strings.Repeat(UnicodeTableStyle.Horizontal, 1) + UnicodeTableStyle.TopRight + "\n"
	expectedString += UnicodeTableStyle.Vertical + "N" + UnicodeTableStyle.Vertical + "A" + UnicodeTableStyle.Vertical + "C" + UnicodeTableStyle.Vertical + "\n"
	expectedString += UnicodeTableStyle.InnerLeft + strings.Repeat(UnicodeTableStyle.Horizontal, 1) + UnicodeTableStyle.InnerMiddle + strings.Repeat(UnicodeTableStyle.Horizontal, 1) + UnicodeTableStyle.InnerMiddle + strings.Repeat(UnicodeTableStyle.Horizontal, 1) + UnicodeTableStyle.InnerRight + "\n"
	expectedString += UnicodeTableStyle.Vertical + "J" + UnicodeTableStyle.Vertical + "3" + UnicodeTableStyle.Vertical + "U" + UnicodeTableStyle.Vertical + "\n"
	expectedString += UnicodeTableStyle.Vertical + "J" + UnicodeTableStyle.Vertical + "2" + UnicodeTableStyle.Vertical + "C" + UnicodeTableStyle.Vertical + "\n"
	expectedString += UnicodeTableStyle.BottomLeft + strings.Repeat(UnicodeTableStyle.Horizontal, 1) + UnicodeTableStyle.BottomMiddle + strings.Repeat(UnicodeTableStyle.Horizontal, 1) + UnicodeTableStyle.BottomMiddle + strings.Repeat(UnicodeTableStyle.Horizontal, 1) + UnicodeTableStyle.BottomRight + "\n"
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %s but got: %s", expectedString, resultString)
	}
	//----------------------------------------
}

//------------------------------------------------------------

func TestRenderTable7(t *testing.T) {
	//----------------------------------------
	resultString := RenderTable([][]string{
		{"Name", "Age", "Country"},
		{"John Doe", "30", "USA"},
		{"Jane Smith", "25", "Canada"},
	}, WithHeader, WithMaxColumnWidths([]int{1, 2, 3}), WithPadding(0))
	//----------------------------------------
	expectedString := UnicodeTableStyle.TopLeft + strings.Repeat(UnicodeTableStyle.Horizontal, 1) + UnicodeTableStyle.TopMiddle + strings.Repeat(UnicodeTableStyle.Horizontal, 2) + UnicodeTableStyle.TopMiddle + strings.Repeat(UnicodeTableStyle.Horizontal, 3) + UnicodeTableStyle.TopRight + "\n"
	expectedString += UnicodeTableStyle.Vertical + "N" + UnicodeTableStyle.Vertical + "Ag" + UnicodeTableStyle.Vertical + "Cou" + UnicodeTableStyle.Vertical + "\n"
	expectedString += UnicodeTableStyle.InnerLeft + strings.Repeat(UnicodeTableStyle.Horizontal, 1) + UnicodeTableStyle.InnerMiddle + strings.Repeat(UnicodeTableStyle.Horizontal, 2) + UnicodeTableStyle.InnerMiddle + strings.Repeat(UnicodeTableStyle.Horizontal, 3) + UnicodeTableStyle.InnerRight + "\n"
	expectedString += UnicodeTableStyle.Vertical + "J" + UnicodeTableStyle.Vertical + "30" + UnicodeTableStyle.Vertical + "USA" + UnicodeTableStyle.Vertical + "\n"
	expectedString += UnicodeTableStyle.Vertical + "J" + UnicodeTableStyle.Vertical + "25" + UnicodeTableStyle.Vertical + "Can" + UnicodeTableStyle.Vertical + "\n"
	expectedString += UnicodeTableStyle.BottomLeft + strings.Repeat(UnicodeTableStyle.Horizontal, 1) + UnicodeTableStyle.BottomMiddle + strings.Repeat(UnicodeTableStyle.Horizontal, 2) + UnicodeTableStyle.BottomMiddle + strings.Repeat(UnicodeTableStyle.Horizontal, 3) + UnicodeTableStyle.BottomRight + "\n"
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %s but got: %s", expectedString, resultString)
	}
	//----------------------------------------
}

//------------------------------------------------------------

func TestTabwriterTable1(t *testing.T) {
	//----------------------------------------
	resultString := TabwriterTable([][]string{})
	//----------------------------------------
	expectedString := ""
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %s but got: %s", expectedString, resultString)
	}
	//----------------------------------------
}

//------------------------------------------------------------

func TestTabwriterTable2(t *testing.T) {
	//----------------------------------------
	resultString := TabwriterTable([][]string{
		{"John Doe", "30", "USA"},
		{"Jane Smith", "25", "Canada"},
	})
	//----------------------------------------
	expectedString := "John Doe    30  USA\n"
	expectedString += "Jane Smith  25  Canada\n"
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %s but got: %s", expectedString, resultString)
	}
	//----------------------------------------
}

//------------------------------------------------------------

func TestTabwriterTable3(t *testing.T) {
	//----------------------------------------
	resultString := TabwriterTable([][]string{
		{"Name", "Age", "Country"},
		{"John Doe", "30", "USA"},
		{"Jane Smith", "25", "Canada"},
	}, WithHeader)
	//----------------------------------------
	expectedString := "Name        Age  Country\n"
	expectedString += "----        ---  -------\n"
	expectedString += "John Doe    30   USA\n"
	expectedString += "Jane Smith  25   Canada\n"
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %s but got: %s", expectedString, resultString)
	}
	//----------------------------------------
}

//------------------------------------------------------------

func TestTabwriterTable4(t *testing.T) {
	//----------------------------------------
	resultString := TabwriterTable([][]string{
		{"Name", "Age", "Country"},
		{"John Doe", "30", "USA"},
		{"Jane Smith", "25", "Canada"},
	}, WithHeader, WithMaxTableWidth(10))
	//----------------------------------------
	expectedString := "Name      \n"
	expectedString += "----      \n"
	expectedString += "John Doe  \n"
	expectedString += "Jane Smith\n"
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %s but got: %s", expectedString, resultString)
	}
	//----------------------------------------
}

//------------------------------------------------------------

func TestTabwriterTable5(t *testing.T) {
	//----------------------------------------
	resultString := TabwriterTable([][]string{
		{"Name", "Age", "Country"},
		{"John Doe", "30", "USA"},
		{"Jane Smith", "25", "Canada"},
	}, WithHeader, WithMaxColumnWidth(1))
	//----------------------------------------
	expectedString := "N  A  C\n"
	expectedString += "-  -  -\n"
	expectedString += "J  3  U\n"
	expectedString += "J  2  C\n"
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %s but got: %s", expectedString, resultString)
	}
	//----------------------------------------
}

//------------------------------------------------------------

func TestTabwriterTable6(t *testing.T) {
	//----------------------------------------
	resultString := TabwriterTable([][]string{
		{"Name", "Age", "Country"},
		{"John Doe", "30", "USA"},
		{"Jane Smith", "25", "Canada"},
	}, WithHeader, WithMaxColumnWidths([]int{1, 2, 3}))
	//----------------------------------------
	expectedString := "N  Ag  Cou\n"
	expectedString += "-  --  ---\n"
	expectedString += "J  30  USA\n"
	expectedString += "J  25  Can\n"
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %s but got: %s", expectedString, resultString)
	}
	//----------------------------------------
}

//------------------------------------------------------------
