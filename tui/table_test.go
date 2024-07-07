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
	expectedString := topLeft + topRight + "\n" + bottomLeft + bottomRight + "\n"
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
	expectedString := topLeft + topRight + "\n" + bottomLeft + bottomRight + "\n"
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
	expectedString := topLeft + strings.Repeat(horizontal, 10) + topMiddle + strings.Repeat(horizontal, 2) + topMiddle + strings.Repeat(horizontal, 6) + topRight + "\n"
	expectedString += vertical + "John Doe  " + vertical + "30" + vertical + "USA   " + vertical + "\n"
	expectedString += vertical + "Jane Smith" + vertical + "25" + vertical + "Canada" + vertical + "\n"
	expectedString += bottomLeft + strings.Repeat(horizontal, 10) + bottomMiddle + strings.Repeat(horizontal, 2) + bottomMiddle + strings.Repeat(horizontal, 6) + bottomRight + "\n"
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
	expectedString := topLeft + strings.Repeat(horizontal, 10) + topMiddle + strings.Repeat(horizontal, 3) + topMiddle + strings.Repeat(horizontal, 7) + topRight + "\n"
	expectedString += vertical + "Name      " + vertical + "Age" + vertical + "Country" + vertical + "\n"
	expectedString += innerLeft + strings.Repeat(horizontal, 10) + innerMiddle + strings.Repeat(horizontal, 3) + innerMiddle + strings.Repeat(horizontal, 7) + innerRight + "\n"
	expectedString += vertical + "John Doe  " + vertical + "30 " + vertical + "USA    " + vertical + "\n"
	expectedString += vertical + "Jane Smith" + vertical + "25 " + vertical + "Canada " + vertical + "\n"
	expectedString += bottomLeft + strings.Repeat(horizontal, 10) + bottomMiddle + strings.Repeat(horizontal, 3) + bottomMiddle + strings.Repeat(horizontal, 7) + bottomRight + "\n"
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
	expectedString := topLeft + strings.Repeat(horizontal, 9) + "\n"
	expectedString += vertical + "Name     " + "\n"
	expectedString += innerLeft + strings.Repeat(horizontal, 9) + "\n"
	expectedString += vertical + "John Doe " + "\n"
	expectedString += vertical + "Jane Smit" + "\n"
	expectedString += bottomLeft + strings.Repeat(horizontal, 9) + "\n"
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
	}, WithHeader, WithMaxColumnWidth(1))
	//----------------------------------------
	expectedString := topLeft + strings.Repeat(horizontal, 1) + topMiddle + strings.Repeat(horizontal, 1) + topMiddle + strings.Repeat(horizontal, 1) + topRight + "\n"
	expectedString += vertical + "N" + vertical + "A" + vertical + "C" + vertical + "\n"
	expectedString += innerLeft + strings.Repeat(horizontal, 1) + innerMiddle + strings.Repeat(horizontal, 1) + innerMiddle + strings.Repeat(horizontal, 1) + innerRight + "\n"
	expectedString += vertical + "J" + vertical + "3" + vertical + "U" + vertical + "\n"
	expectedString += vertical + "J" + vertical + "2" + vertical + "C" + vertical + "\n"
	expectedString += bottomLeft + strings.Repeat(horizontal, 1) + bottomMiddle + strings.Repeat(horizontal, 1) + bottomMiddle + strings.Repeat(horizontal, 1) + bottomRight + "\n"
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
	}, WithHeader, WithMaxColumnWidths([]int{1, 2, 3}))
	//----------------------------------------
	expectedString := topLeft + strings.Repeat(horizontal, 1) + topMiddle + strings.Repeat(horizontal, 2) + topMiddle + strings.Repeat(horizontal, 3) + topRight + "\n"
	expectedString += vertical + "N" + vertical + "Ag" + vertical + "Cou" + vertical + "\n"
	expectedString += innerLeft + strings.Repeat(horizontal, 1) + innerMiddle + strings.Repeat(horizontal, 2) + innerMiddle + strings.Repeat(horizontal, 3) + innerRight + "\n"
	expectedString += vertical + "J" + vertical + "30" + vertical + "USA" + vertical + "\n"
	expectedString += vertical + "J" + vertical + "25" + vertical + "Can" + vertical + "\n"
	expectedString += bottomLeft + strings.Repeat(horizontal, 1) + bottomMiddle + strings.Repeat(horizontal, 2) + bottomMiddle + strings.Repeat(horizontal, 3) + bottomRight + "\n"
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
