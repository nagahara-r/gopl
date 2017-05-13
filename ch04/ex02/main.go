// Copyright © 2017 Yuki Nagahara
// 標準入力の内容からSHA Sumを出力します

package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	l := flag.Bool("384", false, "SHA384 Support")
	ll := flag.Bool("512", false, "SHA512 Support")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	if *l {
		fmt.Printf("%x\n", sha512.Sum384([]byte(scanner.Text())))
	} else if *ll {
		fmt.Printf("%x\n", sha512.Sum512([]byte(scanner.Text())))
	} else {
		fmt.Printf("%x\n", sha256.Sum256([]byte(scanner.Text())))
	}
}
