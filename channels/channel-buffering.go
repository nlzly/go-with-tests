package main

import "fmt"

func main() {

	messages := make(chan string, 2)

	messages <- "nosfer"
	messages <- "atu"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
