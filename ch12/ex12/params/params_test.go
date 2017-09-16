// Copyright © 2017 Yuki Nagahara

package params

import (
	"log"
	"strings"
	"testing"
)

func TestPack(t *testing.T) {
	url := "http://localhost:12345/search"

	type S struct {
		ZipCode       int    `http:"zcode,zipcode"`
		MailAddress   string `http:"mail,mailaddress"`
		CreditCardNum string `http:"credit,creditcard"`
	}
	tests := []struct {
		str      S
		expected []string
		err      error
	}{
		{
			S{2430465, "test@example.com", "1234567890123456"},
			[]string{"mail=test@example.com", "zcode=2430465", "credit=1234567890123456"},
			nil,
		}, {
			S{1, "test@example.com", "1234567890123456"},
			nil,
			ErrZip,
		}, {
			S{1234567, "testexample", "1234567890123456"},
			nil,
			ErrMailAddress,
		}, {
			S{1234567, "test@example.com", "12345678901234567"},
			nil,
			ErrCreditCardNumber,
		}, {
			S{1234567, "", "1234567890123456"},
			[]string{"zcode=1234567", "credit=1234567890123456"},
			nil,
		}, {
			S{1234567, "test@example.com", ""},
			[]string{"zcode=1234567", "mail=test@example.com"},
			nil,
		},
	}

	for _, test := range tests {
		result, err := Pack(url, &test.str)

		if err != nil {
			if err == test.err {
				continue
			}
			t.Errorf("%v", err)
		}

		if !isExpected(result, url, test.expected) {
			t.Errorf("want = %v, result = %v", test.expected, result)
		}
	}
}

func isExpected(result string, expectedurl string, expectedparams []string) bool {
	// URLでスタートするか
	if !strings.HasPrefix(result, expectedurl) {
		log.Printf("Has No Url = %v\n", result)
		return false
	}

	// すべての予期した要素を含むか
	for _, expectedparam := range expectedparams {
		if !strings.Contains(result, expectedparam) {
			log.Printf("Has No Param = %v\n", expectedparam)
			return false
		}
	}

	return true
}
