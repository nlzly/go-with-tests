package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() { messages <- "test message" }()

	msg := <-messages

	fmt.Println(msg)
}
