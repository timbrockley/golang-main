//------------------------------------------------------------

package conv

import (
	"fmt"
	"testing"
)

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

//------------------------------------------------------------
// Base64_encode
//------------------------------------------------------------

func TestBase64_encode(t *testing.T) {

	//------------------------------------------------------------
	var dataString, base64String, resultString string
	//------------------------------------------------------------
	dataString = ""
	//------------------------------------------------------------
	resultString = Base64_encode(dataString)
	//--------------------------------------------------
	if resultString != dataString {

		t.Errorf("resultString = %q but should = %q", resultString, dataString)
	}
	//--------------------------------------------------
	dataString = "ABC <> &quot; \u00A3 \u65E5\u672C\u8A9E\U0001f427"
	base64String = "QUJDIDw+ICZxdW90OyDCoyDml6XmnKzoqp7wn5Cn"
	//------------------------------------------------------------
	resultString = Base64_encode(dataString)
	//--------------------------------------------------
	if resultString != base64String {

		t.Errorf("resultString = %q but should = %q", resultString, base64String)
	}
	//--------------------------------------------------
}

//------------------------------------------------------------
// Base64_decode
//------------------------------------------------------------

func TestBase64_decode(t *testing.T) {

	//------------------------------------------------------------
	var err error
	var dataString, base64String, resultString string
	//------------------------------------------------------------
	dataString = ""
	//------------------------------------------------------------
	resultString, err = Base64_decode(dataString)
	//--------------------------------------------------
	if err != nil {

		t.Error(err)

	} else {

		//--------------------
		if resultString != dataString {

			t.Errorf("resultString = %q but should = %q", resultString, dataString)
		}
		//--------------------
	}
	//------------------------------------------------------------
	dataString = "ABC <> &quot; \u00A3 \u65E5\u672C\u8A9E\U0001f427"
	base64String = "QUJDIDw+ICZxdW90OyDCoyDml6XmnKzoqp7wn5Cn"
	//------------------------------------------------------------
	resultString, err = Base64_decode(base64String)
	//--------------------------------------------------
	if err != nil {

		t.Error(err)

	} else {

		//--------------------
		if resultString != dataString {

			t.Errorf("resultString = %q but should = %q", resultString, dataString)
		}
		//--------------------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Base64url_encode
//------------------------------------------------------------

func TestBase64url_encode(t *testing.T) {

	//------------------------------------------------------------
	var dataString, base64urlString, resultString string
	//------------------------------------------------------------
	dataString = ""
	//------------------------------------------------------------
	resultString = Base64url_encode(dataString)
	//--------------------------------------------------
	if resultString != dataString {

		t.Errorf("resultString = %q but should = %q", resultString, dataString)
	}
	//--------------------------------------------------
	dataString = "\u65E5\u672C\u8A9E\U0001f427"
	base64urlString = "5pel5pys6Kqe8J-Qpw"
	//------------------------------------------------------------
	resultString = Base64url_encode(dataString)
	//--------------------------------------------------
	if resultString != base64urlString {

		t.Errorf("resultString = %q but should = %q", resultString, base64urlString)
	}
	//--------------------------------------------------
}

//------------------------------------------------------------
// Base64url_decode
//------------------------------------------------------------

func TestBase64url_decode(t *testing.T) {

	//------------------------------------------------------------
	var err error
	var dataString, base64urlString, resultString string
	//------------------------------------------------------------
	dataString = ""
	//------------------------------------------------------------
	resultString, err = Base64url_decode(dataString)
	//--------------------------------------------------
	if err != nil {

		t.Error(err)

	} else {

		//--------------------
		if resultString != dataString {

			t.Errorf("resultString = %q but should = %q", resultString, dataString)
		}
		//--------------------
	}
	//------------------------------------------------------------
	dataString = "\u65E5\u672C\u8A9E\U0001f427"
	base64urlString = "5pel5pys6Kqe8J-Qpw"
	//------------------------------------------------------------
	resultString, err = Base64url_decode(base64urlString)
	//--------------------------------------------------
	if err != nil {

		t.Error(err)

	} else {

		//--------------------
		if resultString != dataString {

			t.Errorf("resultString = %q but should = %q", resultString, dataString)
		}
		//--------------------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Base64_Base64url
//------------------------------------------------------------

func TestBase64_Base64url(t *testing.T) {

	//------------------------------------------------------------
	var dataString, base64String, base64urlString, resultString string
	//------------------------------------------------------------
	dataString = ""
	//------------------------------------------------------------
	resultString = Base64_Base64url(dataString)
	//--------------------------------------------------
	if resultString != dataString {

		t.Errorf("resultString = %q but should = %q", resultString, dataString)
	}
	//------------------------------------------------------------
	base64String = "5pel5pys6Kqe8J+Qpw=="
	base64urlString = "5pel5pys6Kqe8J-Qpw"
	//------------------------------------------------------------
	resultString = Base64_Base64url(base64String)
	//--------------------------------------------------
	if resultString != base64urlString {

		t.Errorf("resultString = %q but should = %q", resultString, base64urlString)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Base64url_Base64
//------------------------------------------------------------

func TestBase64url_Base64(t *testing.T) {

	//------------------------------------------------------------
	var dataString, base64String, base64urlString, resultString string
	//------------------------------------------------------------
	dataString = ""
	//------------------------------------------------------------
	resultString = Base64url_Base64(dataString)
	//--------------------------------------------------
	if resultString != dataString {

		t.Errorf("resultString = %q but should = %q", resultString, dataString)
	}
	//------------------------------------------------------------
	base64String = "5pel5pys6Kqe8J+Qpw=="
	base64urlString = "5pel5pys6Kqe8J-Qpw"
	//------------------------------------------------------------
	resultString = Base64url_Base64(base64urlString)
	//--------------------------------------------------
	if resultString != base64String {

		t.Errorf("resultString = %q but should = %q", resultString, base64String)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

//------------------------------------------------------------
// Base91_encode
//------------------------------------------------------------

func TestBase91_encode(t *testing.T) {

	//------------------------------------------------------------
	var dataString, base91String, resultString string
	//------------------------------------------------------------
	dataString = ""
	//------------------------------------------------------------
	resultString = Base91_encode(dataString, false)
	//--------------------------------------------------
	if resultString != dataString {

		t.Errorf("resultString = %q but should = %q", resultString, dataString)
	}
	//--------------------------------------------------
	dataString = "May your trails be crooked, winding, lonesome, dangerous, leading to the most amazing view. May your mountains rise into and above the clouds."
	base91String = "8D9KR`0eLUd/ZQFl62>vb,1RL%%&~8bju\"sQ;mmaU=UfU)1T70<^rm?i;Ct)/p;R(&^m5PKimf2+H[QSd/[E<oTPgZh>DZ%y;#,aIl]U>vP:3pPIqSwPmLwre3:W.{6U)/wP;mYBxgP[UCsS)/[EOiqMgZR*Sk<Rd/=8jL=ibg7+b[C"
	//------------------------------------------------------------
	resultString = Base91_encode(dataString, false)
	//--------------------------------------------------
	if resultString != base91String {

		t.Errorf("resultString = %q but should = %q", resultString, base91String)
	}
	//--------------------------------------------------
	dataString = "May your trails be crooked, winding, lonesome, dangerous, leading to the most amazing view. May your mountains rise into and above the clouds."
	base91String = "8D9KR-g0eLUd/ZQFl62>vb,1RL%%&~8bju-qsQ;mmaU=UfU)1T70<^rm?i;Ct)/p;R(&^m5PKimf2+H[QSd/[E<oTPgZh>DZ%y;#,aIl]U>vP:3pPIqSwPmLwre3:W.{6U)/wP;mYBxgP[UCsS)/[EOiqMgZR*Sk<Rd/=8jL=ibg7+b[C"
	//------------------------------------------------------------
	resultString = Base91_encode(dataString, true)
	//--------------------------------------------------
	if resultString != base91String {

		t.Errorf("resultString = %q but should = %q", resultString, base91String)
	}
	//--------------------------------------------------
}

//------------------------------------------------------------
// Base91_decode
//------------------------------------------------------------

func TestBase91_decode(t *testing.T) {

	//------------------------------------------------------------
	var err error
	var dataString, base91String, resultString string
	//------------------------------------------------------------
	dataString = ""
	//------------------------------------------------------------
	resultString, err = Base91_decode(dataString, false)
	//--------------------------------------------------
	if err != nil {

		t.Error(err)

	} else {

		//--------------------
		if resultString != dataString {

			t.Errorf("resultString = %q but should = %q", resultString, dataString)
		}
		//--------------------
	}
	//------------------------------------------------------------
	dataString = "May your trails be crooked, winding, lonesome, dangerous, leading to the most amazing view. May your mountains rise into and above the clouds."
	base91String = "8D9KR`0eLUd/ZQFl62>vb,1RL%%&~8bju\"sQ;mmaU=UfU)1T70<^rm?i;Ct)/p;R(&^m5PKimf2+H[QSd/[E<oTPgZh>DZ%y;#,aIl]U>vP:3pPIqSwPmLwre3:W.{6U)/wP;mYBxgP[UCsS)/[EOiqMgZR*Sk<Rd/=8jL=ibg7+b[C"
	//------------------------------------------------------------
	resultString, err = Base91_decode(base91String, false)
	//--------------------------------------------------
	if err != nil {

		t.Error(err)

	} else {

		//--------------------
		if resultString != dataString {

			t.Errorf("resultString = %q but should = %q", resultString, dataString)
		}
		//--------------------
	}
	//------------------------------------------------------------
	dataString = "May your trails be crooked, winding, lonesome, dangerous, leading to the most amazing view. May your mountains rise into and above the clouds."
	base91String = "8D9KR-g0eLUd/ZQFl62>vb,1RL%%&~8bju-qsQ;mmaU=UfU)1T70<^rm?i;Ct)/p;R(&^m5PKimf2+H[QSd/[E<oTPgZh>DZ%y;#,aIl]U>vP:3pPIqSwPmLwre3:W.{6U)/wP;mYBxgP[UCsS)/[EOiqMgZR*Sk<Rd/=8jL=ibg7+b[C"
	//------------------------------------------------------------
	resultString, err = Base91_decode(base91String, true)
	//--------------------------------------------------
	if err != nil {

		t.Error(err)

	} else {

		//--------------------
		if resultString != dataString {

			t.Errorf("resultString = %q but should = %q", resultString, dataString)
		}
		//--------------------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ############################################################
//------------------------------------------------------------

//------------------------------------------------------------
// JSON_MarshalIndent
//------------------------------------------------------------

func TestJSON_MarshalIndent(t *testing.T) {

	//------------------------------------------------------------
	jsonMap := []interface{}{map[string]interface{}{"test_key1": "test_value1_<=>"}, map[string]interface{}{"test_key2": "test_value2"}}
	//--------------------
	jsonString := "[\n\t{\n\t\t\"test_key1\": \"test_value1_<=>\"\n\t},\n\t{\n\t\t\"test_key2\": \"test_value2\"\n\t}\n]"
	//------------------------------------------------------------

	//--------------------------------------------------
	resultBytes, err := JSON_MarshalIndent(jsonMap, "", "\t")
	//--------------------------------------------------

	//--------------------
	if err != nil {

		t.Error(err)

	} else {

		//--------------------
		if string(resultBytes) != jsonString {

			t.Errorf("string(resultBytes) = %q but should = %q", string(resultBytes), jsonString)
		}
		//--------------------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// JSON_marshal
//------------------------------------------------------------

func TestJSON_marshal(t *testing.T) {

	//------------------------------------------------------------
	jsonMap := []interface{}{map[string]interface{}{"test_key1": "test_value1_<=>"}, map[string]interface{}{"test_key2": "test_value2"}}
	//--------------------
	jsonString := "[{\"test_key1\":\"test_value1_<=>\"},{\"test_key2\":\"test_value2\"}]"
	//------------------------------------------------------------

	//--------------------------------------------------
	resultBytes, err := JSON_Marshal(jsonMap)
	//--------------------------------------------------

	//--------------------
	if err != nil {

		t.Error(err)

	} else {

		//--------------------
		if string(resultBytes) != jsonString {

			t.Errorf("string(resultBytes) = %q but should = %q", string(resultBytes), jsonString)
		}
		//--------------------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// JSON_encode
//------------------------------------------------------------

func TestJSON_encode(t *testing.T) {

	//------------------------------------------------------------
	jsonMap := []interface{}{map[string]interface{}{"test_key1": "test_value1"}, map[string]interface{}{"test_key2": "test_value2"}}
	//--------------------
	jsonString := "[{\"test_key1\":\"test_value1\"},{\"test_key2\":\"test_value2\"}]"
	//------------------------------------------------------------

	//--------------------------------------------------
	resultString, err := JSON_encode(jsonMap)
	//--------------------------------------------------

	//--------------------
	if err != nil {

		t.Error(err)

	} else {

		//--------------------
		if resultString != jsonString {

			t.Errorf("resultString = %q but should = %q", resultString, jsonString)
		}
		//--------------------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// JSON_decode
//------------------------------------------------------------

func TestJSON_decode(t *testing.T) {

	//------------------------------------------------------------
	jsonString := "[{\"test_key1\":\"test_value1\"},{\"test_key2\":\"test_value2\"}]"
	//--------------------
	jsonMap := []interface{}{map[string]interface{}{"test_key1": "test_value1"}, map[string]interface{}{"test_key2": "test_value2"}}
	//------------------------------------------------------------

	//--------------------------------------------------
	resultMap, err := JSON_decode(jsonString)
	//--------------------------------------------------

	//--------------------
	if err != nil {

		t.Error(err)

	} else {

		//--------------------
		if fmt.Sprint(resultMap) != fmt.Sprint(jsonMap) {

			t.Errorf("resultMap = %#v but should = %#v", resultMap, jsonMap)
		}
		//--------------------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ############################################################
//------------------------------------------------------------
