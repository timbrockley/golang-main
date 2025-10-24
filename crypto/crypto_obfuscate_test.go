//----------------------------------------------------------------------

package crypto

import (
	"testing"
)

//----------------------------------------------------------------------
//######################################################################
//----------------------------------------------------------------------

//------------------------------------------------------------
// ObfuscateV0
//------------------------------------------------------------

func TestObfuscateV0(t *testing.T) {
	//------------------------------------------------------------
	type testRecord struct {
		input  string
		output string
	}
	//------------------------------------------------------------
	testData := []testRecord{
		{"", ""},
		{"abc", "=<;"},
		{"hello", "6922/"},
		{"test BBB>>>www|||qqq 123", string([]byte{42, 57, 43, 42, 126, 92, 92, 92, 96, 96, 96, 39, 39, 39, 34, 34, 34, 45, 45, 45, 126, 109, 108, 107})},
		{string([]byte{42, 57, 43, 42, 126, 92, 92, 92, 96, 96, 96, 39, 39, 39, 34, 34, 34, 45, 45, 45, 126, 109, 108, 107}), "test BBB>>>www|||qqq 123"},
	}
	//--------------------------------------------------
	for index, test := range testData {
		//------------------------------------------------------------
		result := ObfuscateV0(test.input)
		//------------------------------------------------------------
		if result != test.output {
			t.Errorf("index %v: result = %v but should = %v", index, []byte(result), []byte(test.output))
		}
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// SlideByteV0
//------------------------------------------------------------

func TestSlideByteV0(t *testing.T) {
	//------------------------------------------------------------
	var result byte
	//------------------------------------------------------------
	type testRecord struct {
		input  byte
		output byte
	}
	//------------------------------------------------------------
	testData := []testRecord{
		{0, 31}, {31, 0}, {32, 126}, {126, 32}, {127, 127}, {128, 255}, {255, 128},
	}
	//--------------------------------------------------
	for index, test := range testData {
		//------------------------------------------------------------
		result = SlideByteV0(test.input)
		//------------------------------------------------------------
		if result != test.output {
			t.Errorf("index %v: result = %v but should = %v", index, result, test.output)
		}
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateV0Encode
//------------------------------------------------------------

func TestObfuscateV0Encode(t *testing.T) {
	//------------------------------------------------------------
	type testRecord struct {
		input    string
		output   string
		encoding string
	}
	//------------------------------------------------------------
	testData := []testRecord{
		{"", "", ""},
		{"\x00 ABC \n \r \x22 \x7C \x27 \x77 \x60 \x3E \u65e5\u672c\u8a9e\U0001f427qqqqAqKI\x5CqKI\x5C", string([]byte{31, 126, 93, 45, 98, 91, 126, 21, 126, 18, 126, 124, 126, 45, 113, 126, 119, 126, 45, 97, 126, 62, 126, 45, 103, 126, 153, 232, 218, 153, 227, 211, 151, 213, 225, 143, 224, 239, 216, 45, 45, 45, 45, 45, 45, 45, 45, 93, 45, 45, 83, 85, 66, 45, 45, 83, 85, 66}), ""},
		{"A>|\U0001f427", "B!sp>n=%=", "base"},
		{"A>|\U0001f427", "XWAij+Dv2A==", "base64"},
		{"A>|\U0001f427", "XWAij-Dv2A", "base64url"},
		{"A>|\U0001f427", "CBx+](ZyN", "base91"},
		{"A>|\U0001f427", "5D60228FE0EFD8", "hex"},
	}
	//--------------------------------------------------
	for index, test := range testData {
		//------------------------------------------------------------
		result := ObfuscateV0Encode(test.input, test.encoding)
		//------------------------------------------------------------
		if result != test.output {
			t.Errorf("index %v: result = %v but should = %v", index, []byte(result), []byte(test.output))
		}
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateV0Decode
//------------------------------------------------------------

func TestObfuscateV0Decode(t *testing.T) {
	//------------------------------------------------------------
	type testRecord struct {
		input    string
		output   string
		encoding string
	}
	//------------------------------------------------------------
	testData := []testRecord{
		{"", "", ""},
		{string([]byte{31, 126, 93, 45, 98, 91, 126, 21, 126, 18, 126, 124, 126, 45, 113, 126, 119, 126, 45, 97, 126, 62, 126, 45, 103, 126, 153, 232, 218, 153, 227, 211, 151, 213, 225, 143, 224, 239, 216, 45, 45, 45, 45, 45, 45, 45, 45, 93, 45, 45, 83, 85, 66, 45, 45, 83, 85, 66}), "\x00 ABC \n \r \x22 \x7C \x27 \x77 \x60 \x3E \u65e5\u672c\u8a9e\U0001f427qqqqAqKI\x5CqKI\x5C", ""},
		{"B!sp>n=%=", "A>|\U0001f427", "base"},
		{"XWAij+Dv2A==", "A>|\U0001f427", "base64"},
		{"XWAij-Dv2A", "A>|\U0001f427", "base64url"},
		{"CBx+](ZyN", "A>|\U0001f427", "base91"},
		{"5D60228FE0EFD8", "A>|\U0001f427", "hex"},
	}
	//--------------------------------------------------
	for index, test := range testData {
		//------------------------------------------------------------
		result, err := ObfuscateV0Decode(test.input, test.encoding)
		//------------------------------------------------------------
		if err != nil {
			t.Error(err)
		} else if result != test.output {
			t.Errorf("index %v: result = %v but should = %v", index, []byte(result), []byte(test.output))
		}
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
}

//----------------------------------------------------------------------
//######################################################################
//----------------------------------------------------------------------

//------------------------------------------------------------
// ObfuscateV4
//------------------------------------------------------------

func TestObfuscateV4(t *testing.T) {
	//------------------------------------------------------------
	type testRecord struct {
		input    string
		output   string
		mixChars bool
	}
	//------------------------------------------------------------
	testData := []testRecord{
		{"", "", false},
		{"hello", "6922/", false},
		{"hello", "6229/", true},
		{"test BBB>>>www|||qqq 123", "*\x27+\x22~-\x5C-\x60m\x60k\x279\x22*\x22\x5C-\x5C~\x60l\x27", true},
		{"*\x27+\x22~-\x5C-\x60m\x60k\x279\x22*\x22\x5C-\x5C~\x60l\x27", "test BBB>>>www|||qqq 123", true},
	}
	//--------------------------------------------------
	for index, test := range testData {
		//------------------------------------------------------------
		result := ObfuscateV4(test.input, WithMixChars(test.mixChars))
		//------------------------------------------------------------
		if result != test.output {
			t.Errorf("index %v: result = %v but should = %v", index, []byte(result), []byte(test.output))
		}
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// SlideByteV4
//------------------------------------------------------------

func TestSlideByteV4(t *testing.T) {
	//------------------------------------------------------------
	var result byte
	//------------------------------------------------------------
	type testRecord struct {
		input  byte
		output byte
	}
	//------------------------------------------------------------
	testData := []testRecord{
		{0, 0}, {31, 31}, {32, 126}, {126, 32}, {127, 127}, {255, 255},
	}
	//--------------------------------------------------
	for index, test := range testData {
		//------------------------------------------------------------
		result = SlideByteV4(test.input)
		//------------------------------------------------------------
		if result != test.output {
			t.Errorf("index %v: result = %v but should = %v", index, result, test.output)
		}
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateV4Encode
//------------------------------------------------------------

func TestObfuscateV4Encode(t *testing.T) {
	//------------------------------------------------------------
	type testRecord struct {
		input    string
		output   string
		encoding string
		MixChars bool
	}
	//------------------------------------------------------------
	testData := []testRecord{
		{"", "", "", false},
		{"hello", "6922/", "", false},
		{"hello", "6229/", "", true},
		{"\x00 ABC \n \r \x22 \x7C \x27 \x77 \x60 \x3E \u65e5\u672c\u8a9e\U0001F427", string([]byte{0, 126, 93, 126, 91, 126, 92, 110, 151, 92, 114, 230, 124, 172, 92, 113, 170, 119, 240, 92, 97, 126, 62, 92, 92, 92, 103, 126, 230, 126, 165, 126, 156, 126, 232, 126, 158, 126, 159, 144, 167}), "", true},
		{"test BBB>>>www|||qqq 123 ABC @ XYZ", "1S62)LYnA-BxU2H0[Lzi.zv8,LX&atLXKG1LREUl:::", "base", true},
		{"test BBB>>>www|||qqq 123", "KicrIn4tXC1gbWBrJzkiKiJcLVx+YGwn", "base64", true},
		{"test BBB>>>www|||qqq 123", "KicrIn4tXC1gbWBrJzkiKiJcLVx-YGwn", "base64url", true},
		{"test BBB>>>www|||qqq 123", "OU/w-d}u)}H;#-ql>NXG%w.-du)TWm!&B", "base91", true},
		{"test BBB>>>www|||qqq 123", "2A272B227E2D5C2D606D606B2739222A225C2D5C7E606C27", "hex", true},
	}
	//--------------------------------------------------
	for index, test := range testData {
		//------------------------------------------------------------
		result := ObfuscateV4Encode(test.input, WithEncoding(test.encoding), WithMixChars(test.MixChars))
		//------------------------------------------------------------
		if result != test.output {
			t.Errorf("index %v: result = %v but should = %v", index, []byte(result), []byte(test.output))
		}
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateV4Decode
//------------------------------------------------------------

func TestObfuscateV4Decode(t *testing.T) {
	//------------------------------------------------------------
	type testRecord struct {
		input    string
		output   string
		encoding string
		mixChars bool
	}
	//------------------------------------------------------------
	testData := []testRecord{
		{"", "", "", false},
		{"hello", "6922/", "", false},
		{"hello", "6229/", "", true},
		{string([]byte{0, 126, 93, 126, 91, 126, 92, 110, 151, 92, 114, 230, 124, 172, 92, 113, 170, 119, 240, 92, 97, 126, 62, 92, 92, 92, 103, 126, 230, 126, 165, 126, 156, 126, 232, 126, 158, 126, 159, 144, 167}), "\x00 ABC \n \r \x22 \x7C \x27 \x77 \x60 \x3E \u65e5\u672c\u8a9e\U0001F427", "", true},
		{"1S62)LYnA-BxU2H0[Lzi.zv8,LX&atLXKG1LREUl:::", "test BBB>>>www|||qqq 123 ABC @ XYZ", "base", true},
		{"KicrIn4tXC1gbWBrJzkiKiJcLVx+YGwn", "test BBB>>>www|||qqq 123", "base64", true},
		{"KicrIn4tXC1gbWBrJzkiKiJcLVx-YGwn", "test BBB>>>www|||qqq 123", "base64url", true},
		{"OU/w-d}u)}H;#-ql>NXG%w.-du)TWm!&B", "test BBB>>>www|||qqq 123", "base91", true},
		{"2A272B227E2D5C2D606D606B2739222A225C2D5C7E606C27", "test BBB>>>www|||qqq 123", "hex", true},
	}
	//--------------------------------------------------
	for index, test := range testData {
		//------------------------------------------------------------
		result, err := ObfuscateV4Decode(test.input, WithEncoding(test.encoding), WithMixChars(test.mixChars))
		//------------------------------------------------------------
		if err != nil {
			t.Error(err)
		} else if result != test.output {
			t.Errorf("index %v: result = %v but should = %v", index, []byte(result), []byte(test.output))
		}
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
}

//----------------------------------------------------------------------
//######################################################################
//----------------------------------------------------------------------

//------------------------------------------------------------
// ObfuscateV5
//------------------------------------------------------------

func TestObfuscateV5(t *testing.T) {
	//------------------------------------------------------------
	type testRecord struct {
		input  string
		output string
	}
	//------------------------------------------------------------
	testData := []testRecord{
		{"", ""},
		{"abc", "=<;"},
		{"hello", "6229/"},
		{"test BBB>>>www|||qqqzzz 123 \x00\x09\x0A ~~~", string([]byte{42, 45, 43, 45, 126, 36, 92, 126, 96, 108, 96, 126, 39, 22, 34, 126, 34, 57, 45, 42, 36, 92, 36, 92, 109, 96, 107, 39, 31, 39, 21, 34, 32, 32, 32})},
		{string([]byte{42, 45, 43, 45, 126, 36, 92, 126, 96, 108, 96, 126, 39, 22, 34, 126, 34, 57, 45, 42, 36, 92, 36, 92, 109, 96, 107, 39, 31, 39, 21, 34, 32, 32, 32}), "test BBB>>>www|||qqqzzz 123 \x00\x09\x0A ~~~"},
	}
	//--------------------------------------------------
	for index, test := range testData {
		//------------------------------------------------------------
		result := ObfuscateV5(test.input)
		//------------------------------------------------------------
		if result != test.output {
			t.Errorf("index %v: result = %v but should = %v", index, []byte(result), []byte(test.output))
		}
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// SlideByteV5
//------------------------------------------------------------

func TestSlideByteV5(t *testing.T) {
	//------------------------------------------------------------
	var result byte
	//------------------------------------------------------------
	type testRecord struct {
		input  byte
		output byte
	}
	//------------------------------------------------------------
	testData := []testRecord{
		{0, 31}, {31, 0}, {32, 126}, {126, 32}, {127, 127}, {128, 255}, {255, 128},
	}
	//--------------------------------------------------
	for index, test := range testData {
		//------------------------------------------------------------
		result = SlideByteV5(test.input)
		//------------------------------------------------------------
		if result != test.output {
			t.Errorf("index %v: result = %v but should = %v", index, result, test.output)
		}
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateV5Encode
//------------------------------------------------------------

func TestObfuscateV5Encode(t *testing.T) {
	//------------------------------------------------------------
	type testRecord struct {
		input    string
		output   string
		encoding string
	}
	//------------------------------------------------------------
	testData := []testRecord{
		{"", "", ""},
		{"test BBB>>>www|||qqqzzz 123 ~~~", "*-q+--~---b-d-g~-gl-a~-q9-q*---b-d-b-d-gm-ak-a-s-s-s", ""},
		{"ABC \u00a9 \u65e5\u672c\u8a9e\U0001f427", string([]byte{93, 227, 91, 151, 189, 225, 126, 224, 232, 216, 153, 45, 98, 211, 126, 213, 214, 143, 153, 239, 218}), ""},
		{"test BBB>>>www|||qqq 123 ABC @ XYZ", "1S62)LYnA-BxU2H0[Lzi.zv8,LX&atLXKG1LREUl:::", "base"},
		{"test BBB>>>www|||qqq 123 ABC @ XYZ", "Ki0rLX5tXGtgXWBbJ14iRiI5LSp+XGxcfmBcJ34nfiJFRA==", "base64"},
		{"test BBB>>>www|||qqq 123 ABC @ XYZ", "Ki0rLX5tXGtgXWBbJ14iRiI5LSp-XGxcfmBcJ34nfiJFRA", "base64url"},
		{"test BBB>>>www|||qqq 123 ABC @ XYZ", "Dlba(}^*?Sdp-ql<N8GQzQoX5QWk!0,_h-dU3>%}@cAM", "base91"},
		{"test BBB>>>www|||qqq 123", "2A272B227E2D5C2D606D606B2739222A225C2D5C7E606C27", "hex"},
	}
	//--------------------------------------------------
	for index, test := range testData {
		//------------------------------------------------------------
		result := ObfuscateV5Encode(test.input, test.encoding)
		//------------------------------------------------------------
		if result != test.output {
			t.Errorf("index %v: result = %v but should = %v", index, []byte(result), []byte(test.output))
		}
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateV5Decode
//------------------------------------------------------------

func TestObfuscateV5Decode(t *testing.T) {
	//------------------------------------------------------------
	type testRecord struct {
		input    string
		output   string
		encoding string
	}
	//------------------------------------------------------------
	testData := []testRecord{
		{"", "", ""},
		{"*-q+--~---b-d-g~-gl-a~-q9-q*---b-d-b-d-gm-ak-a-s-s-s", "test BBB>>>www|||qqqzzz 123 ~~~", ""},
		{string([]byte{93, 227, 91, 151, 189, 225, 126, 224, 232, 216, 153, 45, 98, 211, 126, 213, 214, 143, 153, 239, 218}), "ABC \u00a9 \u65e5\u672c\u8a9e\U0001f427", ""},
		{"1S62)LYnA-BxU2H0[Lzi.zv8,LX&atLXKG1LREUl:::", "test BBB>>>www|||qqq 123 ABC @ XYZ", "base"},
		{"Ki0rLX5tXGtgXWBbJ14iRiI5LSp+XGxcfmBcJ34nfiJFRA==", "test BBB>>>www|||qqq 123 ABC @ XYZ", "base64"},
		{"Ki0rLX5tXGtgXWBbJ14iRiI5LSp-XGxcfmBcJ34nfiJFRA", "test BBB>>>www|||qqq 123 ABC @ XYZ", "base64url"},
		{"Dlba(}^*?Sdp-ql<N8GQzQoX5QWk!0,_h-dU3>%}@cAM", "test BBB>>>www|||qqq 123 ABC @ XYZ", "base91"},
		{"2A272B227E2D5C2D606D606B2739222A225C2D5C7E606C27", "test BBB>>>www|||qqq 123", "hex"},
	}
	//--------------------------------------------------
	for index, test := range testData {
		//------------------------------------------------------------
		result, err := ObfuscateV5Decode(test.input, test.encoding)
		//------------------------------------------------------------
		if err != nil {
			t.Error(err)
		} else if result != test.output {
			t.Errorf("index %v: result = %v but should = %v", index, []byte(result), []byte(test.output))
		}
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
}

//----------------------------------------------------------------------
//######################################################################
//----------------------------------------------------------------------

//------------------------------------------------------------
// ObfuscateXOR
//------------------------------------------------------------

func TestObfuscateXOR(t *testing.T) {
	//------------------------------------------------------------
	type testRecord struct {
		input  string
		output string
		value  byte
	}
	//------------------------------------------------------------
	testData := []testRecord{
		{"", "", 0},
		{"", "", 32},
		{"ABCD", "abcd", 32},
		{"abcd", "ABCD", 32},
	}
	//--------------------------------------------------
	for index, test := range testData {
		//------------------------------------------------------------
		result, err := ObfuscateXOR(test.input, test.value)
		//------------------------------------------------------------
		if err != nil && test.value == 0 && err.Error() != "value should be an integer between 1 and 255" {
			t.Error("unexpected error: ", err.Error())
		} else if err != nil && test.value != 0 {
			t.Error(err)
		} else if result != test.output {
			t.Errorf("index %v: result = %v but should = %v", index, []byte(result), []byte(test.output))
		}
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateXOREncode
//------------------------------------------------------------

func TestObfuscateXOREncode(t *testing.T) {
	//------------------------------------------------------------
	type testRecord struct {
		input    string
		output   string
		value    byte
		encoding string
	}
	//------------------------------------------------------------
	testData := []testRecord{
		{"", "", 0, ""},
		{"", "", 32, ""},
		{"ABCD", "C=HdZ", 32, "base"},
		{"ABCD", "YWJjZA==", 32, "base64"},
		{"ABCD", "YWJjZA", 32, "base64url"},
		{"ABCD", "#G(IZ", 32, "base91"},
		{"ABCD", "61626364", 32, "hex"},
	}
	//--------------------------------------------------
	for index, test := range testData {
		//------------------------------------------------------------
		result, err := ObfuscateXOREncode(test.input, test.value, test.encoding)
		//------------------------------------------------------------
		if err != nil && test.value == 0 && err.Error() != "value should be an integer between 1 and 255" {
			t.Error("unexpected error: ", err.Error())
		} else if err != nil && test.value != 0 {
			t.Error(err)
		} else if result != test.output {
			t.Errorf("index %v: result = %v but should = %v", index, []byte(result), []byte(test.output))
		}
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateXORDecode
//------------------------------------------------------------

func TestObfuscateXORDecode(t *testing.T) {
	//------------------------------------------------------------
	type testRecord struct {
		input    string
		output   string
		value    byte
		encoding string
	}
	//------------------------------------------------------------
	testData := []testRecord{
		{"", "", 0, ""},
		{"", "", 32, ""},
		{"C=HdZ", "ABCD", 32, "base"},
		{"YWJjZA==", "ABCD", 32, "base64"},
		{"YWJjZA", "ABCD", 32, "base64url"},
		{"#G(IZ", "ABCD", 32, "base91"},
		{"61626364", "ABCD", 32, "hex"},
	}
	//--------------------------------------------------
	for index, test := range testData {
		//------------------------------------------------------------
		result, err := ObfuscateXORDecode(test.input, test.value, test.encoding)
		//------------------------------------------------------------
		if err != nil && test.value == 0 && err.Error() != "value should be an integer between 1 and 255" {
			t.Error("unexpected error: ", err.Error())
		} else if err != nil && test.value != 0 {
			t.Error(err)
		} else if result != test.output {
			t.Errorf("index %v: result = %v but should = %v", index, []byte(result), []byte(test.output))
		}
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
}

//----------------------------------------------------------------------
//######################################################################
//----------------------------------------------------------------------
