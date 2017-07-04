// Copyright © 2017 Yuki Nagahara

package bank_test

import (
	"testing"

	"github.com/naga718/golang-practice/ch09/ex01"
)

func TestWithDraw(t *testing.T) {
	tests := []struct {
		balance  int
		amount   int
		expected bool
	}{
		{300, 100, true},
		{100, 100, true},
		{0, 0, true},
		{0, 100, false},
		{400, 1000000, false},
		{1, 2, false},
	}

	for _, test := range tests {
		bank.Deposit(test.balance)

		result := bank.Withdraw(test.amount)

		if result != test.expected {
			t.Errorf("balance = %v, amount = %v, unexpected Withdraw = %v\n", test.balance, test.amount, result)
		}

		bank.Withdraw(bank.Balance())
	}
}
