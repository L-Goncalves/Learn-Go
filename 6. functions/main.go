package main

import (
	"fmt"
	"os"
)

func main() {

	showIntroduction()
	showMenu()
	command := readCommand()

	switch command {
	case 1:
		fmt.Println("Monitoring...")
	case 2:
		fmt.Println("Showing Logs...")
	case 0:
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("I don't know this command.")
		os.Exit(-1)
	}

}

func showIntroduction() {
	name := "Lucas"
	version := 1.1

	fmt.Println("Hello,", name)
	fmt.Println("This Program is at version", version)
}

func showMenu() {
	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit")

}

func readCommand() int {
	var commandRead int
	fmt.Scan(&commandRead)

	return commandRead
}
