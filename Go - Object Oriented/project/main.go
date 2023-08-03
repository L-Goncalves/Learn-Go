package main

import (
	"fmt"
	"project/accounts"
)

func main() {

	josueAccount := accounts.BankAccount{AccountHolder: "Josue", Balance: 300}
	jorgeAccount := accounts.BankAccount{AccountHolder: "Jorginho", Balance: 200}
	josueAccount.Balance = 300

	status := jorgeAccount.Transfer(-200, &josueAccount)
	fmt.Println(status)
	fmt.Println(jorgeAccount)
	fmt.Println(josueAccount)
}
