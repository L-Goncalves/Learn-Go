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

func (c *BankAccount) transfer(value float64, account *BankAccount) bool {
	if value <= c.balance && value > 0 {
		c.balance -= value
		account.addBalance(value)
		return true
	} else {
		return false
	}
}

func main() {

	josueAccount := BankAccount{accountHolder: "Josue", balance: 300}
	jorgeAccount := BankAccount{accountHolder: "Jorginho", balance: 200}
	josueAccount.balance = 300

	status := jorgeAccount.transfer(-200, &josueAccount)
	fmt.Println(status)
	fmt.Println(jorgeAccount)
	fmt.Println(josueAccount)
}
