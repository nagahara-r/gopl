// Copyright © 2017 Yuki Nagahara

package params

import (
	"net/http"
	"reflect"
	"testing"
)

func TestPackAndUnpack(t *testing.T) {
	url := "http://localhost:12345/search"

	type S struct {
		ZipCode       int    `http:"zcode,zipcode"`
		MailAddress   string `http:"mail,mailaddress"`
		CreditCardNum string `http:"credit,creditcard"`
		Other         []string
	}
	tests := []struct {
		str S
		err error
	}{
		{
			S{2430465, "test@example.com", "1234567890123456", []string{"Other Message1", "Message2"}},
			nil,
		}, {
			S{1, "test@example.com", "1234567890123456", []string{""}},
			ErrZip,
		}, {
			S{1234567, "testexample", "1234567890123456", []string{""}},
			ErrMailAddress,
		}, {
			S{1234567, "test@example.com", "12345678901234567", []string{""}},
			ErrCreditCardNumber,
		}, {
			S{1234567, "", "1234567890123456", []string{""}},
			nil,
		}, {
			S{1234567, "test@example.com", "", []string{""}},
			nil,
		}, {
			S{2430465, "test@example.com", "1234567890123456", nil},
			nil,
		},
	}

	for _, test := range tests {
		structToURL, err := Pack(url, &test.str)

		// 期待結果がエラーの場合は、エラーが期待通りか確認します
		if err != nil {
			if err == test.err {
				continue
			}
			t.Errorf("%v", err)
		}

		// HTTPリクエストにします
		var result S
		req, err := http.NewRequest("GET", structToURL, nil)
		if err != nil {
			t.Errorf("%v", err)
		}

		// HTTPリクエストをアンパックします
		err = Unpack(req, &result)
		if err != nil {
			t.Errorf("%v", err)
		}

		// 元の構造体と同じか確認します
		if !reflect.DeepEqual(test.str, result) {
			t.Errorf("original(expected) = %v, result = %v", test.str, result)
		}
	}
}
