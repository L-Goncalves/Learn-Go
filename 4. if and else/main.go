package main

import "fmt"

func main() {

	name := "Lucas"
	version := 1.1

	fmt.Println("Hello,", name)
	fmt.Println("This Program is at version", version)

	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit")

	var command int
	fmt.Scan(&command)

	if command == 1 {
		fmt.Println("Monitoring...")
	} else if command == 2 {
		fmt.Println("Showing Logs...")
	} else if command == 0 {
		fmt.Println("Exiting..")
	} else {
		fmt.Println("I don't know this command")
	}
}
