// Example demonstrates specifying send/receive channels in parameter.
//  This allows for enhanced type safety in functional programming.

package main

import "fmt"

func send(messages chan<- string, msg string) {
	messages <- msg
}

func receive(messages <-chan string) {
	msg := <-messages
	fmt.Println(msg)
}

func main() {
	messages := make(chan string, 1) // Explicitly declare buffer size of channel.

	send(messages, "hello")
	receive(messages)
}
