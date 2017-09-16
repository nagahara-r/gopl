// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習12-12: フィールドタグを拡張し、妥当性がチェックできるようにします。
// フィールドタグ動作確認用

// See page 348.

// Search is a demo of the params.Unpack function.
package main

import (
	"fmt"
	"log"
	"net/http"
)

//!+

import "github.com/naga718/golang-practice/ch12/ex12/params"

// search implements the /search URL endpoint.
func search(resp http.ResponseWriter, req *http.Request) {
	var data struct {
		ZipCode       int    `http:"zcode,zipcode"`
		MailAddress   string `http:"mail,mailaddress"`
		CreditCardNum string `http:"credit,creditcard"`
	}
	data.ZipCode = 1234567 // set default
	if err := params.Unpack(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest) // 400
		return
	}

	// ...rest of handler...
	fmt.Fprintf(resp, "Search: %+v\n", data)
}

//!-

func main() {
	http.HandleFunc("/search", search)
	log.Fatal(http.ListenAndServe(":12345", nil))
}

/*
//!+output
$ go build github.com/naga718/golang-practice/ch12/ex12
$ ./ex12 &
$ ./fetch "http://localhost:12345/search?zcode=1234567&mail=a@a.com&credit=1234567890123456"
Search: {ZipCode:1234567 MailAddress:a@a.com CreditCardNum:1234567890123456}
$ ./fetch "http://localhost:12345/search?zcode=1234567&mail=a@a.com&credit=123456789012345a"
credit: Invalid CreditCard Number
//!-output
*/
