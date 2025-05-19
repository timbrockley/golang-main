//------------------------------------------------------------

package crypto

import (
	"testing"
)

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

//------------------------------------------------------------
// ObfuscateV4
//------------------------------------------------------------

func TestObfuscateV4(t *testing.T) {
	//------------------------------------------------------------
	var dataString, result, EXPECTED_result string
	//------------------------------------------------------------
	dataString = "test BBB>>>www|||qqq 123"
	EXPECTED_result = "*\x27+\x22~-\x5C-\x60m\x60k\x279\x22*\x22\x5C-\x5C~\x60l\x27"
	//------------------------------------------------------------
	result = ObfuscateV4(dataString)
	//------------------------------------------------------------
	if result != EXPECTED_result {
		t.Errorf("result = %q but should = %q", result, EXPECTED_result)
	}
	//------------------------------------------------------------
	dataString = "*\x27+\x22~-\x5C-\x60m\x60k\x279\x22*\x22\x5C-\x5C~\x60l\x27"
	EXPECTED_result = "test BBB>>>www|||qqq 123"
	//------------------------------------------------------------
	result = ObfuscateV4(dataString)
	//------------------------------------------------------------
	if result != EXPECTED_result {
		t.Errorf("result = %q but should = %q", result, EXPECTED_result)
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
	type testPair struct {
		charByte   byte
		resultByte byte
	}
	//------------------------------------------------------------
	testData := []testPair{
		{0, 0}, {31, 31}, {32, 126}, {126, 32}, {127, 127}, {255, 255},
	}
	//--------------------------------------------------
	for index, test := range testData {

		//------------------------------------------------------------
		result = SlideByteV4(test.charByte)
		//------------------------------------------------------------
		if result != test.resultByte {
			t.Errorf("index %v: result = %v but should = %v", index, result, test.resultByte)
		}
		//------------------------------------------------------------
	}
}

//------------------------------------------------------------
// ObfuscateV4_encode
//------------------------------------------------------------

func TestObfuscateV4_encode(t *testing.T) {
	//------------------------------------------------------------
	var dataString, result string
	var resultBytes, EXPECTED_resultBytes []byte
	//------------------------------------------------------------
	dataString = "\x00 ABC \n \r \x22 \x7C \x27 \x77 \x60 \x3E \u65e5\u672c\u8a9e\U0001F427"
	EXPECTED_resultBytes = []byte{0, 126, 93, 126, 91, 126, 92, 110, 151, 92, 114, 230, 124, 172, 92, 113, 170, 119, 240, 92, 97, 126, 62, 92, 92, 92, 103, 126, 230, 126, 165, 126, 156, 126, 232, 126, 158, 126, 159, 144, 167}
	//------------------------------------------------------------
	result = ObfuscateV4_encode(dataString)
	//------------------------------------------------------------
	resultBytes = []byte(result)
	//------------------------------------------------------------
	if string(resultBytes) != string(EXPECTED_resultBytes) {
		t.Errorf("result = %v but should = %v", resultBytes, EXPECTED_resultBytes)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateV4_decode
//------------------------------------------------------------

func TestObfuscateV4_decode(t *testing.T) {
	//------------------------------------------------------------
	var dataString, result, EXPECTED_result string
	//------------------------------------------------------------
	dataString = string([]byte{0, 126, 93, 126, 91, 126, 92, 110, 151, 92, 114, 230, 124, 172, 92, 113, 170, 119, 240, 92, 97, 126, 62, 92, 92, 92, 103, 126, 230, 126, 165, 126, 156, 126, 232, 126, 158, 126, 159, 144, 167})
	EXPECTED_result = "\x00 ABC \n \r \x22 \x7C \x27 \x77 \x60 \x3E \u65e5\u672c\u8a9e\U0001F427"
	//------------------------------------------------------------
	result = ObfuscateV4_decode(dataString)
	//------------------------------------------------------------
	if result != EXPECTED_result {
		t.Errorf("result = %q but should = %q", result, EXPECTED_result)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateV4_base_encode
//------------------------------------------------------------

func TestObfuscateV4_base_encode(t *testing.T) {
	//------------------------------------------------------------
	var dataString, result, EXPECTED_result string
	//------------------------------------------------------------
	dataString = "test BBB>>>www|||qqq 123 ABC @ XYZ"
	EXPECTED_result = "1S62)LYnA-BxU2H0[Lzi.zv8,LX&atLXKG1LREUl:::"
	//------------------------------------------------------------
	result = ObfuscateV4_base_encode(dataString)
	//------------------------------------------------------------
	if result != EXPECTED_result {
		t.Errorf("result = %v but should = %v", result, EXPECTED_result)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateV4_base_decode
//------------------------------------------------------------

func TestObfuscateV4_base_decode(t *testing.T) {
	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	var dataString, result, EXPECTED_result string
	//------------------------------------------------------------
	dataString = "1S62)LYnA-BxU2H0[Lzi.zv8,LX&atLXKG1LREUl:::"
	EXPECTED_result = "test BBB>>>www|||qqq 123 ABC @ XYZ"
	//------------------------------------------------------------
	result, err = ObfuscateV4_base_decode(dataString)
	//------------------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		if result != EXPECTED_result {
			t.Errorf("result = %q but should = %q", result, EXPECTED_result)
		}
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateV4_base64_encode
//------------------------------------------------------------

func TestObfuscateV4_base64_encode(t *testing.T) {
	//------------------------------------------------------------
	var dataString, result, EXPECTED_result string
	//------------------------------------------------------------
	dataString = "test BBB>>>www|||qqq 123"
	EXPECTED_result = "KicrIn4tXC1gbWBrJzkiKiJcLVx+YGwn"
	//------------------------------------------------------------
	result = ObfuscateV4_base64_encode(dataString)
	//------------------------------------------------------------
	if result != EXPECTED_result {
		t.Errorf("result = %v but should = %v", result, EXPECTED_result)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateV4_base64_decode
//------------------------------------------------------------

func TestObfuscateV4_base64_decode(t *testing.T) {
	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	var dataString, result, EXPECTED_result string
	//------------------------------------------------------------
	dataString = "KicrIn4tXC1gbWBrJzkiKiJcLVx+YGwn"
	EXPECTED_result = "test BBB>>>www|||qqq 123"
	//------------------------------------------------------------
	result, err = ObfuscateV4_base64_decode(dataString)
	//------------------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		if result != EXPECTED_result {
			t.Errorf("result = %q but should = %q", result, EXPECTED_result)
		}
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateV4_base64url_encode
//------------------------------------------------------------

func TestObfuscateV4_base64url_encode(t *testing.T) {
	//------------------------------------------------------------
	var dataString, result, EXPECTED_result string
	//------------------------------------------------------------
	dataString = "test BBB>>>www|||qqq 123"
	EXPECTED_result = "KicrIn4tXC1gbWBrJzkiKiJcLVx-YGwn"
	//------------------------------------------------------------
	result = ObfuscateV4_base64url_encode(dataString)
	//------------------------------------------------------------
	if result != EXPECTED_result {
		t.Errorf("result = %v but should = %v", result, EXPECTED_result)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateV4_base64url_decode
//------------------------------------------------------------

func TestObfuscateV4_base64url_decode(t *testing.T) {
	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	var dataString, result, EXPECTED_result string
	//------------------------------------------------------------
	dataString = "KicrIn4tXC1gbWBrJzkiKiJcLVx-YGwn"
	EXPECTED_result = "test BBB>>>www|||qqq 123"
	//------------------------------------------------------------
	result, err = ObfuscateV4_base64url_decode(dataString)
	//------------------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		if result != EXPECTED_result {
			t.Errorf("result = %q but should = %q", result, EXPECTED_result)
		}
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateV4_base91_encode
//------------------------------------------------------------

func TestObfuscateV4_base91_encode(t *testing.T) {
	//------------------------------------------------------------
	var dataString, result, EXPECTED_result string
	//------------------------------------------------------------
	dataString = "test BBB>>>www|||qqq 123"
	EXPECTED_result = "OU/w-d}u)}H;#-ql>NXG%w.-du)TWm!&B"
	//------------------------------------------------------------
	result = ObfuscateV4_base91_encode(dataString)
	//------------------------------------------------------------
	if result != EXPECTED_result {
		t.Errorf("result = %v but should = %v", result, EXPECTED_result)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateV4_base91_decode
//------------------------------------------------------------

func TestObfuscateV4_base91_decode(t *testing.T) {
	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	var dataString, result, EXPECTED_result string
	//------------------------------------------------------------
	dataString = "OU/w-d}u)}H;#-ql>NXG%w.-du)TWm!&B"
	EXPECTED_result = "test BBB>>>www|||qqq 123"
	//------------------------------------------------------------
	result, err = ObfuscateV4_base91_decode(dataString)
	//------------------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		if result != EXPECTED_result {
			t.Errorf("result = %q but should = %q", result, EXPECTED_result)
		}
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

//------------------------------------------------------------
// ObfuscateV5
//------------------------------------------------------------

func TestObfuscateV5(t *testing.T) {
	//------------------------------------------------------------
	var dataString, result, EXPECTED_result string
	var resultBytes, EXPECTED_resultBytes []byte
	//------------------------------------------------------------
	dataString = "test BBB>>>www|||qqqzzz 123 \x00\x09\x0A ~~~"
	EXPECTED_resultBytes = []byte{42, 45, 43, 45, 126, 36, 92, 126, 96, 108, 96, 126, 39, 22, 34, 126, 34, 57, 45, 42, 36, 92, 36, 92, 109, 96, 107, 39, 31, 39, 21, 34, 32, 32, 32}
	//------------------------------------------------------------
	result = ObfuscateV5(dataString)
	//------------------------------------------------------------
	resultBytes = []byte(result)
	//------------------------------------------------------------
	if string(resultBytes) != string(EXPECTED_resultBytes) {
		t.Errorf("result = %v but should = %v", resultBytes, EXPECTED_resultBytes)
	}
	//------------------------------------------------------------
	dataString = string([]byte{42, 45, 43, 45, 126, 36, 92, 126, 96, 108, 96, 126, 39, 22, 34, 126, 34, 57, 45, 42, 36, 92, 36, 92, 109, 96, 107, 39, 31, 39, 21, 34, 32, 32, 32})
	EXPECTED_result = "test BBB>>>www|||qqqzzz 123 \x00\x09\x0A ~~~"
	//------------------------------------------------------------
	result = ObfuscateV5(dataString)
	//------------------------------------------------------------
	if result != EXPECTED_result {
		t.Errorf("result bytes = %v but should = %v", []byte(result), []byte(EXPECTED_result))
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
	type testPair struct {
		charByte   byte
		resultByte byte
	}
	//------------------------------------------------------------
	testData := []testPair{
		{0, 31}, {31, 0}, {32, 126}, {126, 32}, {127, 127}, {128, 255}, {255, 128},
	}
	//--------------------------------------------------
	for index, test := range testData {

		//------------------------------------------------------------
		result = SlideByteV5(test.charByte)
		//------------------------------------------------------------
		if result != test.resultByte {
			t.Errorf("index %v: result = %v but should = %v", index, result, test.resultByte)
		}
		//------------------------------------------------------------
	}
}

//------------------------------------------------------------
// ObfuscateV5_encode
//------------------------------------------------------------

func TestObfuscateV5_encode(t *testing.T) {
	//------------------------------------------------------------
	var dataString, result, EXPECTED_result string
	//------------------------------------------------------------
	dataString = "test BBB>>>www|||qqqzzz 123 ~~~"
	EXPECTED_result = "*-q+--~---b-d-g~-gl-a~-q9-q*---b-d-b-d-gm-ak-a-s-s-s"
	//------------------------------------------------------------
	result = ObfuscateV5_encode(dataString)
	//------------------------------------------------------------
	if result != EXPECTED_result {
		t.Errorf("result = %q but should = %q", result, EXPECTED_result)
	}
	//------------------------------------------------------------
	dataString = "ABC \u00a9 \u65e5\u672c\u8a9e\U0001f427"
	EXPECTED_result = string([]byte{93, 227, 91, 151, 189, 225, 126, 224, 232, 216, 153, 45, 98, 211, 126, 213, 214, 143, 153, 239, 218})
	//------------------------------------------------------------
	result = ObfuscateV5_encode(dataString)
	//------------------------------------------------------------
	if result != EXPECTED_result {
		t.Errorf("result = %q but should = %q", result, EXPECTED_result)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateV5_decode
//------------------------------------------------------------

func TestObfuscateV5_decode(t *testing.T) {
	//------------------------------------------------------------
	var dataString, result, EXPECTED_result string
	//------------------------------------------------------------
	dataString = "*-q+--~---b-d-g~-gl-a~-q9-q*---b-d-b-d-gm-ak-a-s-s-s"
	EXPECTED_result = "test BBB>>>www|||qqqzzz 123 ~~~"
	//------------------------------------------------------------
	result = ObfuscateV5_decode(dataString)
	//------------------------------------------------------------
	if result != EXPECTED_result {
		t.Errorf("result = %q but should = %q", result, EXPECTED_result)
	}
	//------------------------------------------------------------
	dataString = string([]byte{93, 227, 91, 151, 189, 225, 126, 224, 232, 216, 153, 45, 98, 211, 126, 213, 214, 143, 153, 239, 218})
	EXPECTED_result = "ABC \u00a9 \u65e5\u672c\u8a9e\U0001f427"
	//------------------------------------------------------------
	result = ObfuscateV5_decode(dataString)
	//------------------------------------------------------------
	if result != EXPECTED_result {
		t.Errorf("result = %q but should = %q", result, EXPECTED_result)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateV5_base_encode
//------------------------------------------------------------

func TestObfuscateV5_base_encode(t *testing.T) {
	//------------------------------------------------------------
	var dataString, result, EXPECTED_result string
	//------------------------------------------------------------
	dataString = "test BBB>>>www|||qqq 123 ABC @ XYZ"
	EXPECTED_result = "1S62)LYnA-BxU2H0[Lzi.zv8,LX&atLXKG1LREUl:::"
	//------------------------------------------------------------
	result = ObfuscateV5_base_encode(dataString)
	//------------------------------------------------------------
	if result != EXPECTED_result {
		t.Errorf("result = %v but should = %v", result, EXPECTED_result)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateV5_base_decode
//------------------------------------------------------------

func TestObfuscateV5_base_decode(t *testing.T) {
	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	var dataString, result, EXPECTED_result string
	//------------------------------------------------------------
	dataString = "1S62)LYnA-BxU2H0[Lzi.zv8,LX&atLXKG1LREUl:::"
	EXPECTED_result = "test BBB>>>www|||qqq 123 ABC @ XYZ"
	//------------------------------------------------------------
	result, err = ObfuscateV5_base_decode(dataString)
	//------------------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		if result != EXPECTED_result {
			t.Errorf("result = %q but should = %q", result, EXPECTED_result)
		}
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateV5_base64_encode
//------------------------------------------------------------

func TestObfuscateV5_base64_encode(t *testing.T) {
	//------------------------------------------------------------
	var dataString, result, EXPECTED_result string
	//------------------------------------------------------------
	dataString = "test BBB>>>www|||qqq 123 ABC @ XYZ"
	EXPECTED_result = "Ki0rLX5tXGtgXWBbJ14iRiI5LSp+XGxcfmBcJ34nfiJFRA=="
	//------------------------------------------------------------
	result = ObfuscateV5_base64_encode(dataString)
	//------------------------------------------------------------
	if result != EXPECTED_result {
		t.Errorf("result = %v but should = %v", result, EXPECTED_result)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateV5_base64_decode
//------------------------------------------------------------

func TestObfuscateV5_base64_decode(t *testing.T) {
	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	var dataString, result, EXPECTED_result string
	//------------------------------------------------------------
	dataString = "Ki0rLX5tXGtgXWBbJ14iRiI5LSp+XGxcfmBcJ34nfiJFRA=="
	EXPECTED_result = "test BBB>>>www|||qqq 123 ABC @ XYZ"
	//------------------------------------------------------------
	result, err = ObfuscateV5_base64_decode(dataString)
	//------------------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		if result != EXPECTED_result {
			t.Errorf("result = %q but should = %q", result, EXPECTED_result)
		}
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateV5_base64url_encode
//------------------------------------------------------------

func TestObfuscateV5_base64url_encode(t *testing.T) {
	//------------------------------------------------------------
	var dataString, result, EXPECTED_result string
	//------------------------------------------------------------
	dataString = "test BBB>>>www|||qqq 123 ABC @ XYZ"
	EXPECTED_result = "Ki0rLX5tXGtgXWBbJ14iRiI5LSp-XGxcfmBcJ34nfiJFRA"
	//------------------------------------------------------------
	result = ObfuscateV5_base64url_encode(dataString)
	//------------------------------------------------------------
	if result != EXPECTED_result {
		t.Errorf("result = %v but should = %v", result, EXPECTED_result)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateV5_base64url_decode
//------------------------------------------------------------

func TestObfuscateV5_base64url_decode(t *testing.T) {
	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	var dataString, result, EXPECTED_result string
	//------------------------------------------------------------
	dataString = "Ki0rLX5tXGtgXWBbJ14iRiI5LSp-XGxcfmBcJ34nfiJFRA"
	EXPECTED_result = "test BBB>>>www|||qqq 123 ABC @ XYZ"
	//------------------------------------------------------------
	result, err = ObfuscateV5_base64url_decode(dataString)
	//------------------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		if result != EXPECTED_result {
			t.Errorf("result = %q but should = %q", result, EXPECTED_result)
		}
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateV5_base91_encode
//------------------------------------------------------------

func TestObfuscateV5_base91_encode(t *testing.T) {
	//------------------------------------------------------------
	var dataString, result, EXPECTED_result string
	//------------------------------------------------------------
	dataString = "test BBB>>>www|||qqq 123 ABC @ XYZ"
	EXPECTED_result = "Dlba(}^*?Sdp-ql<N8GQzQoX5QWk!0,_h-dU3>%}@cAM"
	//------------------------------------------------------------
	result = ObfuscateV5_base91_encode(dataString)
	//------------------------------------------------------------
	if result != EXPECTED_result {
		t.Errorf("result = %v but should = %v", result, EXPECTED_result)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ObfuscateV5_base91_decode
//------------------------------------------------------------

func TestObfuscateV5_base91_decode(t *testing.T) {
	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	var dataString, result, EXPECTED_result string
	//------------------------------------------------------------
	dataString = "Dlba(}^*?Sdp-ql<N8GQzQoX5QWk!0,_h-dU3>%}@cAM"
	EXPECTED_result = "test BBB>>>www|||qqq 123 ABC @ XYZ"
	//------------------------------------------------------------
	result, err = ObfuscateV5_base91_decode(dataString)
	//------------------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		if result != EXPECTED_result {
			t.Errorf("result = %q but should = %q", result, EXPECTED_result)
		}
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

//------------------------------------------------------------
// SwapStringV0
//------------------------------------------------------------

func TestSwapStringV0(t *testing.T) {
	//------------------------------------------------------------
	var dataString, result, EXPECTED_result string
	var resultBytes, EXPECTED_resultBytes []byte
	//------------------------------------------------------------
	dataString = "test BBB>>>www|||qqq 123"
	EXPECTED_resultBytes = []byte{42, 57, 43, 42, 126, 92, 92, 92, 96, 96, 96, 39, 39, 39, 34, 34, 34, 45, 45, 45, 126, 109, 108, 107}
	//------------------------------------------------------------
	result = SwapStringV0(dataString)
	//------------------------------------------------------------
	resultBytes = []byte(result)
	//------------------------------------------------------------
	if string(resultBytes) != string(EXPECTED_resultBytes) {
		t.Errorf("result = %v but should = %v", resultBytes, EXPECTED_resultBytes)
	}
	//------------------------------------------------------------
	dataString = string([]byte{42, 57, 43, 42, 126, 92, 92, 92, 96, 96, 96, 39, 39, 39, 34, 34, 34, 45, 45, 45, 126, 109, 108, 107})
	EXPECTED_result = "test BBB>>>www|||qqq 123"
	//------------------------------------------------------------
	result = SwapStringV0(dataString)
	//------------------------------------------------------------
	if result != EXPECTED_result {
		t.Errorf("result bytes = %v but should = %v", []byte(result), []byte(EXPECTED_result))
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// SwapStringV0_encode
//------------------------------------------------------------

func TestSwapStringV0_encode(t *testing.T) {
	//------------------------------------------------------------
	var dataString, result string
	var resultBytes, EXPECTED_resultBytes []byte
	//------------------------------------------------------------
	dataString = "\x00 ABC \n \r \x22 \x7C \x27 \x77 \x60 \x3E \u65e5\u672c\u8a9e\U0001f427qqqqAqKI\x5CqKI\x5C"
	EXPECTED_resultBytes = []byte{31, 126, 93, 45, 98, 91, 126, 21, 126, 18, 126, 124, 126, 45, 113, 126, 119, 126, 45, 97, 126, 62, 126, 45, 103, 126, 153, 232, 218, 153, 227, 211, 151, 213, 225, 143, 224, 239, 216, 45, 45, 45, 45, 45, 45, 45, 45, 93, 45, 45, 83, 85, 66, 45, 45, 83, 85, 66}
	//------------------------------------------------------------
	result = SwapStringV0_encode(dataString)
	//------------------------------------------------------------
	resultBytes = []byte(result)
	//------------------------------------------------------------
	if string(resultBytes) != string(EXPECTED_resultBytes) {
		t.Errorf("result = %v but should = %v", resultBytes, EXPECTED_resultBytes)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// SwapStringV0_decode
//------------------------------------------------------------

func TestSwapStringV0_decode(t *testing.T) {
	//------------------------------------------------------------
	dataBytes := []byte{31, 126, 93, 45, 98, 91, 126, 21, 126, 18, 126, 124, 126, 45, 113, 126, 119, 126, 45, 97, 126, 62, 126, 45, 103, 126, 153, 232, 218, 153, 227, 211, 151, 213, 225, 143, 224, 239, 216, 45, 45, 45, 45, 45, 45, 45, 45, 93, 45, 45, 83, 85, 66, 45, 45, 83, 85, 66}
	EXPECTED_result := "\x00 ABC \n \r \x22 \x7C \x27 \x77 \x60 \x3E \u65e5\u672c\u8a9e\U0001f427qqqqAqKI\x5CqKI\x5C"
	//------------------------------------------------------------
	result := SwapStringV0_decode(string(dataBytes))
	//------------------------------------------------------------
	if result != EXPECTED_result {
		t.Errorf("result = %q but should = %q", result, EXPECTED_result)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// SwapStringV0_base_encode
//------------------------------------------------------------

func TestSwapStringV0_base_encode(t *testing.T) {
	//------------------------------------------------------------
	var dataString, result, EXPECTED_result string
	//------------------------------------------------------------
	dataString = "A>|\U0001f427"
	EXPECTED_result = "B!sp>n=%="
	//------------------------------------------------------------
	result = SwapStringV0_base_encode(dataString)
	//------------------------------------------------------------
	if result != EXPECTED_result {
		t.Errorf("result = %v but should = %v", result, EXPECTED_result)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// SwapStringV0_base_decode
//------------------------------------------------------------

func TestSwapStringV0_base_decode(t *testing.T) {
	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	var dataString, result, EXPECTED_result string
	//------------------------------------------------------------
	dataString = "B!sp>n=%="
	EXPECTED_result = "A>|\U0001f427"
	//------------------------------------------------------------
	result, err = SwapStringV0_base_decode(dataString)
	//------------------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		if result != EXPECTED_result {
			t.Errorf("result = %q but should = %q", result, EXPECTED_result)
		}
	}
	//------------------------------------------------------------
}

//--
//------------------------------------------------------------
// SwapStringV0_base64_encode
//------------------------------------------------------------

func TestSwapStringV0_base64_encode(t *testing.T) {
	//------------------------------------------------------------
	var dataString, result, EXPECTED_result string
	//------------------------------------------------------------
	dataString = "A>|\U0001f427"
	EXPECTED_result = "XWAij+Dv2A=="
	//------------------------------------------------------------
	result = SwapStringV0_base64_encode(dataString)
	//------------------------------------------------------------
	if result != EXPECTED_result {
		t.Errorf("result = %v but should = %v", result, EXPECTED_result)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// SwapStringV0_base64_decode
//------------------------------------------------------------

func TestSwapStringV0_base64_decode(t *testing.T) {
	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	var dataString, result, EXPECTED_result string
	//------------------------------------------------------------
	dataString = "XWAij+Dv2A=="
	EXPECTED_result = "A>|\U0001f427"
	//------------------------------------------------------------
	result, err = SwapStringV0_base64_decode(dataString)
	//------------------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		if result != EXPECTED_result {
			t.Errorf("result = %q but should = %q", result, EXPECTED_result)
		}
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// SwapStringV0_base64url_encode
//------------------------------------------------------------

func TestSwapStringV0_base64url_encode(t *testing.T) {
	//------------------------------------------------------------
	var dataString, result, EXPECTED_result string
	//------------------------------------------------------------
	dataString = "A>|\U0001f427"
	EXPECTED_result = "XWAij-Dv2A"
	//------------------------------------------------------------
	result = SwapStringV0_base64url_encode(dataString)
	//------------------------------------------------------------
	if result != EXPECTED_result {
		t.Errorf("result = %v but should = %v", result, EXPECTED_result)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// SwapStringV0_base64url_decode
//------------------------------------------------------------

func TestSwapStringV0_base64url_decode(t *testing.T) {
	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	var dataString, result, EXPECTED_result string
	//------------------------------------------------------------
	dataString = "XWAij-Dv2A"
	EXPECTED_result = "A>|\U0001f427"
	//------------------------------------------------------------
	result, err = SwapStringV0_base64url_decode(dataString)
	//------------------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		if result != EXPECTED_result {
			t.Errorf("result = %q but should = %q", result, EXPECTED_result)
		}
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// SwapStringV0_base91_encode
//------------------------------------------------------------

func TestSwapStringV0_base91_encode(t *testing.T) {
	//------------------------------------------------------------
	var dataString, result, EXPECTED_result string
	//------------------------------------------------------------
	dataString = "A>|\U0001f427"
	EXPECTED_result = "CBx+](ZyN"
	//------------------------------------------------------------
	result = SwapStringV0_base91_encode(dataString)
	//------------------------------------------------------------
	if result != EXPECTED_result {
		t.Errorf("result = %v but should = %v", result, EXPECTED_result)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// SwapStringV0_base91_decode
//------------------------------------------------------------

func TestSwapStringV0_base91_decode(t *testing.T) {
	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	var dataString, result, EXPECTED_result string
	//------------------------------------------------------------
	dataString = "CBx+](ZyN"
	EXPECTED_result = "A>|\U0001f427"
	//------------------------------------------------------------
	result, err = SwapStringV0_base91_decode(dataString)
	//------------------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		if result != EXPECTED_result {
			t.Errorf("result = %q but should = %q", result, EXPECTED_result)
		}
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------
