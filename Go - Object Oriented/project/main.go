package main

import "fmt"

type BankAccount struct {
	accountHolder string
	branchCode    int
	accountNumber int
	balance       float64
}

func main() {

	fmt.Println(BankAccount{})

}
