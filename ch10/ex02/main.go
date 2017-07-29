// Copyright © 2017 Yuki Nagahara

package main

import (
	"log"
	"os"

	"github.com/naga718/golang-practice/ch10/ex02/archive"
	_ "github.com/naga718/golang-practice/ch10/ex02/tar"
	_ "github.com/naga718/golang-practice/ch10/ex02/zip"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("args < 3")
	}
	err := archive.Unzip(os.Args[1], os.Args[2])
	if err != nil {
		log.Printf("%v", err)
		return
	}

	log.Printf("done.")
}
