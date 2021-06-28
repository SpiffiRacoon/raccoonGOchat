package main

import (
	"fmt"
)

func main() {
	fmt.Println("raccoonGOchat - early build\n===================================\n",
		"1. Server\n2. client")
	fmt.Print(">> ")
	var text string
	fmt.Scanln(&text)
	fmt.Println(text)

	if text == "1" {
		fmt.Println("Please provide what port to listen to:")
		fmt.Print(">> ")
		var answer string
		fmt.Scanln(&answer)
		server(answer)
	} else if text == "2" {
		fmt.Println("Please provide host:port to dial to:")
		fmt.Print(">> ")
		var answer string
		fmt.Scanln(&answer)
		client(answer)
	}
}
