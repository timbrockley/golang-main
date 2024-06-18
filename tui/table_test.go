package table

import (
	"strings"
	"testing"
)

//------------------------------------------------------------

func TestRenderTable1(t *testing.T) {
	//----------------------------------------
	resultString := RenderTable([][]string{}, false)
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
	resultString := RenderTable([][]string{}, true)
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
	}, false)
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
	}, true)
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
