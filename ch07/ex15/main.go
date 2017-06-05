// Copyright © 2017 Yuki Nagahara
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"gopl.io/ch7/eval"
)

// 練習7-15 標準入力から単一の式を読み込み、その式内の変数に対する値を
// ユーザに問い合わせるプログラム

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("%v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(b))
	fmt.Printf("式内の変数を定義してください。\n定義終了後、C-dを押してください。\n(var)=(value)\n例: A=1.20\n")

	env := eval.Env{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		envs := strings.Split(scanner.Text(), "=")
		if len(envs) != 2 {
			fmt.Printf("invalid var def: %v\n", scanner.Text())
			fmt.Println("Try again!")
			continue
		}

		v, err := strconv.ParseFloat(envs[1], 64)
		if err != nil {
			fmt.Printf("%v\n", err)
			fmt.Println("Try again!")
			continue
		}
		env[eval.Var(envs[0])] = v
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	expr, err := eval.Parse(string(b))
	if err != nil {
		log.Fatalf("%v", err)
	}

	got := fmt.Sprintf("%.6g", expr.Eval(env))
	fmt.Printf("\n%v == %s\n", string(b), got)

}
