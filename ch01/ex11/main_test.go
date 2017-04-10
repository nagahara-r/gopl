package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	os.Args = []string{"", "http://gopl.io", "http://google.com", "http://Baidu.com", "http://Google.co.in", "http://Tmall.com", "http://sohu.com", "http://live.com", "http://vk.com", "http://Instagram.com", "http://sina.com.cn", "http://360.cn", "http://jd.com", "http://google.de", "https://www.yandex.ru/"}
	main()
}

// $ sh runall.bash
// 0.28s    10732  http://google.com
// 3.37s        0  http://Instagram.com
// 3.39s       81  http://Baidu.com
// 3.48s    10440  http://google.de
// 3.49s    13252  http://Google.co.in
// 3.74s     4154  http://gopl.io
// Get http://jd.com: dial tcp 111.206.227.118:80: connect: network is unreacha
// ble
// Get http://www.tmall.com/: dial tcp 202.47.28.119:80: connect: network is un
// reachable
// Get http://www.sohu.com/: dial tcp 175.100.207.204:80: connect: network is u
// nreachable
// Get http://www.sina.com.cn/: dial tcp 14.0.35.48:80: connect: network is unr
// eachable
// Get https://www.360.cn: dial tcp: lookup www.360.cn: no such host
// Get https://login.live.com/login.srf?wa=wsignin1.0&rpsnv=13&ct=1491820720&rv
// er=6.7.6643.0&wp=MBI_SSL_SHARED&wreply=https:%2F%2Fmail.live.com%2Fdefault.a
// spx%3Frru%3Dinbox&lc=1033&id=64855&mkt=en-US&cbcxt=mai: dial tcp 131.253.61.
// 82:443: connect: network is unreachable
// Get https://m.vk.com/: dial tcp 95.213.11.133:443: i/o timeout

// Get https://www.yandex.ru/: read tcp 192.168.3.5:49507->5.255.255.5:443: rea
// d: connection reset by peer
// 124.90s elapsed
// 接続ができなかったサーバに対してはタイムアウトしています。
