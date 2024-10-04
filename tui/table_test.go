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
	expectedString := UnicodeBorderStyle.TopLeft + UnicodeBorderStyle.TopRight + "\n" + UnicodeBorderStyle.BottomLeft + UnicodeBorderStyle.BottomRight + "\n"
	//----------------------------------------
	if resultString != expectedString {
		t.Errorf("expected: %s but got: %s", expectedString, resultString)
	}
	//----------------------------------------
}

//------------------------------------------------------------

func TestRenderTable2(t *testing.T) {
	//----------------------------------------
	resultString := RenderTable([][]string{}, map[string]any{"Header": true})
	//----------------------------------------
	expectedString := UnicodeBorderStyle.TopLeft + UnicodeBorderStyle.TopRight + "\n" + UnicodeBorderStyle.BottomLeft + UnicodeBorderStyle.BottomRight + "\n"
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
	expectedString := UnicodeBorderStyle.TopLeft + strings.Repeat(UnicodeBorderStyle.Horizontal, 10) + UnicodeBorderStyle.TopMiddle + strings.Repeat(UnicodeBorderStyle.Horizontal, 2) + UnicodeBorderStyle.TopMiddle + strings.Repeat(UnicodeBorderStyle.Horizontal, 6) + UnicodeBorderStyle.TopRight + "\n"
	expectedString += UnicodeBorderStyle.Vertical + "John Doe  " + UnicodeBorderStyle.Vertical + "30" + UnicodeBorderStyle.Vertical + "USA   " + UnicodeBorderStyle.Vertical + "\n"
	expectedString += UnicodeBorderStyle.Vertical + "Jane Smith" + UnicodeBorderStyle.Vertical + "25" + UnicodeBorderStyle.Vertical + "Canada" + UnicodeBorderStyle.Vertical + "\n"
	expectedString += UnicodeBorderStyle.BottomLeft + strings.Repeat(UnicodeBorderStyle.Horizontal, 10) + UnicodeBorderStyle.BottomMiddle + strings.Repeat(UnicodeBorderStyle.Horizontal, 2) + UnicodeBorderStyle.BottomMiddle + strings.Repeat(UnicodeBorderStyle.Horizontal, 6) + UnicodeBorderStyle.BottomRight + "\n"
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
	}, map[string]any{"Header": true, "Padding": 1})
	//----------------------------------------
	expectedString := UnicodeBorderStyle.TopLeft + strings.Repeat(UnicodeBorderStyle.Horizontal, 12) + UnicodeBorderStyle.TopMiddle + strings.Repeat(UnicodeBorderStyle.Horizontal, 5) + UnicodeBorderStyle.TopMiddle + strings.Repeat(UnicodeBorderStyle.Horizontal, 9) + UnicodeBorderStyle.TopRight + "\n"
	expectedString += UnicodeBorderStyle.Vertical + " Name       " + UnicodeBorderStyle.Vertical + " Age " + UnicodeBorderStyle.Vertical + " Country " + UnicodeBorderStyle.Vertical + "\n"
	expectedString += UnicodeBorderStyle.InnerLeft + strings.Repeat(UnicodeBorderStyle.Horizontal, 12) + UnicodeBorderStyle.InnerMiddle + strings.Repeat(UnicodeBorderStyle.Horizontal, 5) + UnicodeBorderStyle.InnerMiddle + strings.Repeat(UnicodeBorderStyle.Horizontal, 9) + UnicodeBorderStyle.InnerRight + "\n"
	expectedString += UnicodeBorderStyle.Vertical + " John Doe   " + UnicodeBorderStyle.Vertical + " 30  " + UnicodeBorderStyle.Vertical + " USA     " + UnicodeBorderStyle.Vertical + "\n"
	expectedString += UnicodeBorderStyle.Vertical + " Jane Smith " + UnicodeBorderStyle.Vertical + " 25  " + UnicodeBorderStyle.Vertical + " Canada  " + UnicodeBorderStyle.Vertical + "\n"
	expectedString += UnicodeBorderStyle.BottomLeft + strings.Repeat(UnicodeBorderStyle.Horizontal, 12) + UnicodeBorderStyle.BottomMiddle + strings.Repeat(UnicodeBorderStyle.Horizontal, 5) + UnicodeBorderStyle.BottomMiddle + strings.Repeat(UnicodeBorderStyle.Horizontal, 9) + UnicodeBorderStyle.BottomRight + "\n"
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
	}, map[string]any{"Header": true, "Padding": 1, "MaxWidth": 10})
	//----------------------------------------
	expectedString := UnicodeBorderStyle.TopLeft + strings.Repeat(UnicodeBorderStyle.Horizontal, 8) + UnicodeBorderStyle.TopRight + "\n"
	expectedString += UnicodeBorderStyle.Vertical + " Name   " + UnicodeBorderStyle.Vertical + "\n"
	expectedString += UnicodeBorderStyle.InnerLeft + strings.Repeat(UnicodeBorderStyle.Horizontal, 8) + UnicodeBorderStyle.InnerRight + "\n"
	expectedString += UnicodeBorderStyle.Vertical + " John D " + UnicodeBorderStyle.Vertical + "\n"
	expectedString += UnicodeBorderStyle.Vertical + " Jane S " + UnicodeBorderStyle.Vertical + "\n"
	expectedString += UnicodeBorderStyle.BottomLeft + strings.Repeat(UnicodeBorderStyle.Horizontal, 8) + UnicodeBorderStyle.BottomRight + "\n"
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
	}, map[string]any{"Header": true, "MaxColumnWidth": 1, "Padding": 0})
	//----------------------------------------
	expectedString := UnicodeBorderStyle.TopLeft + strings.Repeat(UnicodeBorderStyle.Horizontal, 1) + UnicodeBorderStyle.TopMiddle + strings.Repeat(UnicodeBorderStyle.Horizontal, 1) + UnicodeBorderStyle.TopMiddle + strings.Repeat(UnicodeBorderStyle.Horizontal, 1) + UnicodeBorderStyle.TopRight + "\n"
	expectedString += UnicodeBorderStyle.Vertical + "N" + UnicodeBorderStyle.Vertical + "A" + UnicodeBorderStyle.Vertical + "C" + UnicodeBorderStyle.Vertical + "\n"
	expectedString += UnicodeBorderStyle.InnerLeft + strings.Repeat(UnicodeBorderStyle.Horizontal, 1) + UnicodeBorderStyle.InnerMiddle + strings.Repeat(UnicodeBorderStyle.Horizontal, 1) + UnicodeBorderStyle.InnerMiddle + strings.Repeat(UnicodeBorderStyle.Horizontal, 1) + UnicodeBorderStyle.InnerRight + "\n"
	expectedString += UnicodeBorderStyle.Vertical + "J" + UnicodeBorderStyle.Vertical + "3" + UnicodeBorderStyle.Vertical + "U" + UnicodeBorderStyle.Vertical + "\n"
	expectedString += UnicodeBorderStyle.Vertical + "J" + UnicodeBorderStyle.Vertical + "2" + UnicodeBorderStyle.Vertical + "C" + UnicodeBorderStyle.Vertical + "\n"
	expectedString += UnicodeBorderStyle.BottomLeft + strings.Repeat(UnicodeBorderStyle.Horizontal, 1) + UnicodeBorderStyle.BottomMiddle + strings.Repeat(UnicodeBorderStyle.Horizontal, 1) + UnicodeBorderStyle.BottomMiddle + strings.Repeat(UnicodeBorderStyle.Horizontal, 1) + UnicodeBorderStyle.BottomRight + "\n"
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
	}, map[string]any{"Header": true, "MaxColumnWidths": []int{1, 2, 3}, "Padding": 0})
	//----------------------------------------
	expectedString := UnicodeBorderStyle.TopLeft + strings.Repeat(UnicodeBorderStyle.Horizontal, 1) + UnicodeBorderStyle.TopMiddle + strings.Repeat(UnicodeBorderStyle.Horizontal, 2) + UnicodeBorderStyle.TopMiddle + strings.Repeat(UnicodeBorderStyle.Horizontal, 3) + UnicodeBorderStyle.TopRight + "\n"
	expectedString += UnicodeBorderStyle.Vertical + "N" + UnicodeBorderStyle.Vertical + "Ag" + UnicodeBorderStyle.Vertical + "Cou" + UnicodeBorderStyle.Vertical + "\n"
	expectedString += UnicodeBorderStyle.InnerLeft + strings.Repeat(UnicodeBorderStyle.Horizontal, 1) + UnicodeBorderStyle.InnerMiddle + strings.Repeat(UnicodeBorderStyle.Horizontal, 2) + UnicodeBorderStyle.InnerMiddle + strings.Repeat(UnicodeBorderStyle.Horizontal, 3) + UnicodeBorderStyle.InnerRight + "\n"
	expectedString += UnicodeBorderStyle.Vertical + "J" + UnicodeBorderStyle.Vertical + "30" + UnicodeBorderStyle.Vertical + "USA" + UnicodeBorderStyle.Vertical + "\n"
	expectedString += UnicodeBorderStyle.Vertical + "J" + UnicodeBorderStyle.Vertical + "25" + UnicodeBorderStyle.Vertical + "Can" + UnicodeBorderStyle.Vertical + "\n"
	expectedString += UnicodeBorderStyle.BottomLeft + strings.Repeat(UnicodeBorderStyle.Horizontal, 1) + UnicodeBorderStyle.BottomMiddle + strings.Repeat(UnicodeBorderStyle.Horizontal, 2) + UnicodeBorderStyle.BottomMiddle + strings.Repeat(UnicodeBorderStyle.Horizontal, 3) + UnicodeBorderStyle.BottomRight + "\n"
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
	}, map[string]any{"Header": true})
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
	}, map[string]any{"Header": true, "MaxWidth": 10})
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
	}, map[string]any{"Header": true, "MaxColumnWidth": 1})
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
	}, map[string]any{"Header": true, "MaxColumnWidths": []int{1, 2, 3}})
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
