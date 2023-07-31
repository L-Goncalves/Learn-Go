package main

import "fmt"

type BankAccount struct {
	accountHolder string
	branchCode    int
	accountNumber int
	balance       float64
}

func (c *BankAccount) withdraw(withdrawValue float64) (string, float64) {
	canWithdraw := withdrawValue > 0 && withdrawValue < c.balance

	if canWithdraw {
		c.balance -= withdrawValue
		return "Withdraw Successful", c.balance
	}

	return "Balance Unsuficient to complete transaction", c.balance
}

func (c *BankAccount) addBalance(balanceAdded float64) (string, float64) {
	if balanceAdded > 0 {
		c.balance += balanceAdded
		return "Balance added Successfully", c.balance
	}

	return "An Error Occurred while attempting to add balance.", c.balance

}

func main() {

	var josueAccount *BankAccount = new(BankAccount)

	josueAccount.balance = 700

	josueAccount.withdraw(700)

	fmt.Println(josueAccount.addBalance(6500))
}
