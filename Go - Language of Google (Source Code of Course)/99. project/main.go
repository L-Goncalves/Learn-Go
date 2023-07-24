package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const delay = 5 // In Minutes

func main() {
	readWebsitesFromFile()
	showIntroduction()

	for {
		showMenu()
		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			showLogs()
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

	fmt.Println("")
	fmt.Println("Waiting for " + strconv.Itoa(delay) + " Minutes to verify again..")
	fmt.Println("")
	wait()

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
	resp, err := http.Get(website)

	if err != nil {
		fmt.Println("An error has occured at the moment of testing / reaching website", website, err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Website:", website, "has been successfully loaded")
		registerLog(website, true)
	} else {
		fmt.Println("Website", website, "has problems. Status Code", resp.StatusCode)
		registerLog(website, false)
	}
}

func readWebsitesFromFile() []string {
	var websites []string

	archive, err := os.Open("websites.txt")

	if err != nil {
		fmt.Println("An error has occured", err)
	}

	reader := bufio.NewReader(archive)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		websites = append(websites, line)

		if err == io.EOF {
			break
		}
	}

	archive.Close()

	return websites

}

func registerLog(website string, status bool) {
	archive, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	archive.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + website + " - ONLINE: " + strconv.FormatBool(status) + "\n")

	archive.Close()
}

func showLogs() {
	archive, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("An error has ocurred while opening logs:", err)
	}

	fmt.Println(string(archive))
}
