package accounts

import (
	client "project/clients"
)

type BankAccount struct {
	AccountHolder client.AccountHolder
	BranchCode    int
	AccountNumber int
	Balance       float64
}

func (c *BankAccount) Withdraw(withdrawValue float64) (string, float64) {
	canWithdraw := withdrawValue > 0 && withdrawValue < c.Balance

	if canWithdraw {
		c.Balance -= withdrawValue
		return "Withdraw Successful", c.Balance
	}

	return "Balance Unsuficient to complete transaction", c.Balance
}

func (c *BankAccount) AddBalance(BalanceAdded float64) (string, float64) {
	if BalanceAdded > 0 {
		c.Balance += BalanceAdded
		return "Balance added Successfully", c.Balance
	}

	return "An Error Occurred while attempting to add Balance.", c.Balance

}

func (c *BankAccount) Transfer(value float64, account *BankAccount) bool {
	if value <= c.Balance && value > 0 {
		c.Balance -= value
		account.AddBalance(value)
		return true
	} else {
		return false
	}
}
