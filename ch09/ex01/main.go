package main

import (
	"fmt"

	bank "github.com/naga718/golang-practice/ch09/ex01/bank"
)

func main() {
	bank.Deposit(300)
	fmt.Printf("Balance = %v\n", bank.Balance())

	fmt.Printf("Withdraw(300) = %v\n", bank.Withdraw(300))
	fmt.Printf("Withdraw(300) = %v\n", bank.Withdraw(300))
}
