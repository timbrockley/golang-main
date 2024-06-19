package table

import (
	"strings"
	"testing"
)

//------------------------------------------------------------

func TestRenderTable1(t *testing.T) {
	//----------------------------------------
	resultString := RenderTable([][]string{}, Options{Header: false})
	//----------------------------------------
	expectedString := topLeft + topRight + "\n" + bottomLeft + bottomRight + "\n"
	//----------------------------------------
	if expectedString != resultString {
		t.Errorf("expected: %s but got: %s", expectedString, resultString)
	}
	//----------------------------------------
}

//------------------------------------------------------------

func TestRenderTable2(t *testing.T) {
	//----------------------------------------
	resultString := RenderTable([][]string{}, Options{Header: true})
	//----------------------------------------
	expectedString := topLeft + topRight + "\n" + bottomLeft + bottomRight + "\n"
	//----------------------------------------
	if expectedString != resultString {
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
	}, Options{Header: false})
	//----------------------------------------
	expectedString := topLeft + strings.Repeat(horizontal, 10) + topMiddle + strings.Repeat(horizontal, 2) + topMiddle + strings.Repeat(horizontal, 6) + topRight + "\n"
	expectedString += vertical + "John Doe  " + vertical + "30" + vertical + "USA   " + vertical + "\n"
	expectedString += vertical + "Jane Smith" + vertical + "25" + vertical + "Canada" + vertical + "\n"
	expectedString += bottomLeft + strings.Repeat(horizontal, 10) + bottomMiddle + strings.Repeat(horizontal, 2) + bottomMiddle + strings.Repeat(horizontal, 6) + bottomRight + "\n"
	//----------------------------------------
	if expectedString != resultString {
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
	}, Options{Header: true})
	//----------------------------------------
	expectedString := topLeft + strings.Repeat(horizontal, 10) + topMiddle + strings.Repeat(horizontal, 3) + topMiddle + strings.Repeat(horizontal, 7) + topRight + "\n"
	expectedString += vertical + "Name      " + vertical + "Age" + vertical + "Country" + vertical + "\n"
	expectedString += innerLeft + strings.Repeat(horizontal, 10) + innerMiddle + strings.Repeat(horizontal, 3) + innerMiddle + strings.Repeat(horizontal, 7) + innerRight + "\n"
	expectedString += vertical + "John Doe  " + vertical + "30 " + vertical + "USA    " + vertical + "\n"
	expectedString += vertical + "Jane Smith" + vertical + "25 " + vertical + "Canada " + vertical + "\n"
	expectedString += bottomLeft + strings.Repeat(horizontal, 10) + bottomMiddle + strings.Repeat(horizontal, 3) + bottomMiddle + strings.Repeat(horizontal, 7) + bottomRight + "\n"
	//----------------------------------------
	if expectedString != resultString {
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
	}, Options{Header: true, MaxWidth: 10})
	//----------------------------------------
	expectedString := topLeft + strings.Repeat(horizontal, 9) + "\n"
	expectedString += vertical + "Name     " + "\n"
	expectedString += innerLeft + strings.Repeat(horizontal, 9) + "\n"
	expectedString += vertical + "John Doe " + "\n"
	expectedString += vertical + "Jane Smit" + "\n"
	expectedString += bottomLeft + strings.Repeat(horizontal, 9) + "\n"
	//----------------------------------------
	if expectedString != resultString {
		t.Errorf("expected: %s but got: %s", expectedString, resultString)
	}
	//----------------------------------------
}

//------------------------------------------------------------

func TestTabwriterTable1(t *testing.T) {
	//----------------------------------------
	resultString := TabwriterTable([][]string{}, Options{})
	//----------------------------------------
	expectedString := ""
	//----------------------------------------
	if expectedString != resultString {
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
	}, Options{})
	//----------------------------------------
	expectedString := "John Doe    30  USA\n"
	expectedString += "Jane Smith  25  Canada\n"
	//----------------------------------------
	if expectedString != resultString {
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
	}, Options{Header: true})
	//----------------------------------------
	expectedString := "Name        Age  Country\n"
	expectedString += "----        ---  -------\n"
	expectedString += "John Doe    30   USA\n"
	expectedString += "Jane Smith  25   Canada\n"
	//----------------------------------------
	if expectedString != resultString {
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
	}, Options{Header: true, MaxWidth: 10})
	//----------------------------------------
	expectedString := "Name      \n"
	expectedString += "----      \n"
	expectedString += "John Doe  \n"
	expectedString += "Jane Smith\n"
	//----------------------------------------
	if expectedString != resultString {
		t.Errorf("expected: %s but got: %s", expectedString, resultString)
	}
	//----------------------------------------
}

//------------------------------------------------------------
