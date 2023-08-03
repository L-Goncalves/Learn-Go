package accounts

import (
	client "project/clients"
)

type BankAccount struct {
	AccountHolder client.AccountHolder
	BranchCode    int
	AccountNumber int
	balance       float64
}

func (c *BankAccount) Withdraw(withdrawValue float64) (string, float64) {
	canWithdraw := withdrawValue > 0 && withdrawValue < c.balance

	if canWithdraw {
		c.balance -= withdrawValue
		return "Withdraw Successful", c.balance
	}

	return "balance Unsuficient to complete transaction", c.balance
}

func (c *BankAccount) Addbalance(balanceAdded float64) (string, float64) {
	if balanceAdded > 0 {
		c.balance += balanceAdded
		return "balance added Successfully", c.balance
	}

	return "An Error Occurred while attempting to add balance.", c.balance

}

func (c *BankAccount) Transfer(value float64, account *BankAccount) bool {
	if value <= c.balance && value > 0 {
		c.balance -= value
		account.Addbalance(value)
		return true
	} else {
		return false
	}
}

func (c *BankAccount) GetBalance() float64 {
	return c.balance
}
