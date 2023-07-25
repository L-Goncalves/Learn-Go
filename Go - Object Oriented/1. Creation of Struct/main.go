package main

import "fmt"

// It's a type struct
type BankAccount struct {
	accountHolder string
	branchCode    int
	accountNumber int
	balance       float64
}

func main() {
	// Creates a new bank account and prints it
	fmt.Println(BankAccount{"Lucas", 0001, 1234567, 76.000})
}
