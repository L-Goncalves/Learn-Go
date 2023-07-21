package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	showIntroduction()

	for {
		showMenu()
		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
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

func startMonitoring() {
	fmt.Println("Monitoring...")
	website := "https://httpstat.us/200"

	resp, _ := http.Get(website)

	if resp.StatusCode == 200 {
		fmt.Println("Website:", website, "has been successfully loaded")
	} else {
		fmt.Println("Website", website, "has problems. Status Code", resp.StatusCode)
	}
}

func readCommand() int {
	var commandRead int
	fmt.Scan(&commandRead)

	return commandRead
}
