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
	fmt.Println("The Address of the variable is", &command)
	fmt.Println("The Chosen command was", command)
}
