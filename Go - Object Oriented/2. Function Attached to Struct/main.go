package main

import "fmt"

type BankAccount struct {
	accountHolder string
	branchCode    int
	accountNumber int
	balance       float64
}

func (c *BankAccount) withdraw(withdrawValue float64) string {
	canWithdraw := withdrawValue > 0 && withdrawValue < c.balance

	if canWithdraw {
		c.balance -= withdrawValue
		return "Withdraw Successful"
	}

	return "Balance Unsuficient to complete transaction"
}

func main() {

	var josueAccount *BankAccount = new(BankAccount)

	josueAccount.balance = 700

	josueAccount.withdraw(600)

	fmt.Println(josueAccount.balance)
}
