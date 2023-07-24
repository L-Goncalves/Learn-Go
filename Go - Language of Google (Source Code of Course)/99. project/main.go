package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
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
	websites := []string{"https://httpstat.us/500", "https://www.alura.com.br", "https://www.caelum.com.br"}

	for _, website := range websites {
		testWebsite(website)
	}

	wait()
	fmt.Println("")

}

func wait() {
	time.Sleep(5 * time.Minute)
	startMonitoring()
}

func readCommand() int {
	var commandRead int
	fmt.Scan(&commandRead)

	return commandRead
}

func testWebsite(website string) {
	resp, _ := http.Get(website)

	if resp.StatusCode == 200 {
		fmt.Println("Website:", website, "has been successfully loaded")
	} else {
		fmt.Println("Website", website, "has problems. Status Code", resp.StatusCode)
	}
}
