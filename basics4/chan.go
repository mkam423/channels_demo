// Example demonstrates specifying send/receive channels in parameter.
//  This allows for enhanced type safety in functional programming.

package main

import "fmt"

func send(messages chan<- string, msg string) {
	messages <- msg
}

func receive(messages <-chan string) {
	for {
		select {
		case msg := <-messages:
			fmt.Println(msg)
		}
	}
}

func main() {
	messages := make(chan string) // Explicitly declare buffer size of channel.

	go receive(messages)
	send(messages, "hello")
	fmt.Scanln()
}
