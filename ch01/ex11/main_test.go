package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	os.Args = []string{"", "http://gopl.io", "http://google.com", "http://Baidu.com", "http://Google.co.in", "http://Tmall.com", "http://sohu.com", "http://live.com", "http://vk.com", "http://Instagram.com", "http://sina.com.cn", "http://360.cn", "http://jd.com", "http://google.de", "https://www.yandex.ru/"}
	main()
}

// 0.19s    10684  http://google.com
// 0.25s       81  http://Baidu.com
// 0.37s     4154  http://gopl.io
// 0.39s    14530  http://Google.co.in
// 0.63s   106686  http://jd.com
// 1.23s    10400  http://google.de
// 1.38s   224073  http://Tmall.com
// 1.55s   598787  http://sina.com.cn
// 1.56s   431300  http://sohu.com
// 2.43s    15648  http://live.com
// 2.85s     6563  http://vk.com
// 3.00s   278879  http://360.cn
// 4.20s        0  http://Instagram.com
// 4.20s elapsed
// PASS
// coverage: 75.0% of statements
// ok  	github.com/naga718/golang-practice/ch01/ex11	4.216s
// Success: Tests passed.
// instagramのサイトが応答していない？
// 0バイトを読み込んでいます
