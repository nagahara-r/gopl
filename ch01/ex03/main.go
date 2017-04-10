package main

import (
	"os"
	"strings"
)

func echoWithIndexLoop() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	//fmt.Println(s)
}

func echoWithRangeLoop() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	//fmt.Println(s)
}

func echoWithJoin() {
	strings.Join(os.Args[1:], " ")
	//fmt.Println(strings.Join(os.Args[1:], " "))
}
