package params

import (
	"log"
	"math"
	"strings"
	"testing"
)

func TestPack(t *testing.T) {
	url := "http://localhost:12345/search"

	type S struct {
		Lavels     []string `http:"l"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}
	tests := []struct {
		str      S
		expected []string
	}{
		{
			S{[]string{"a", "b", "c"}, 100, true},
			[]string{"l=a", "l=b", "l=c", "max=100", "x=true"},
		}, {
			S{[]string{}, math.MaxInt32, true},
			[]string{"max=2147483647", "x=true"},
		}, {
			S{nil, math.MinInt32, false},
			[]string{"max=-2147483648", "x=false"},
		},
	}

	for _, test := range tests {
		result := Pack(url, &test.str)

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
