package main

import (
	"fmt"
	"project/accounts"
)

func main() {
	exampleAccount := accounts.BankAccount{}
	exampleAccount.Addbalance(100)

	fmt.Println(exampleAccount.GetBalance())

}
