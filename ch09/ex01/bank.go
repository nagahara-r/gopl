// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// Withdraw (引き出し) をbankに追加します。

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var updated = make(chan bool) // updated balance

// Deposit は預け入れを行います。
func Deposit(amount int) {
	deposits <- amount
	<-updated
}

// Withdraw は引き出しを行います。
func Withdraw(amount int) bool {
	deposits <- -amount
	return <-updated
}

// Balance は残高照会を行います。
func Balance() int { return <-balances }

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			if balance+amount < 0 {
				updated <- false
				continue
			}
			balance += amount
			updated <- true
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
