// Example demonstrates buffered channels that allow receiver not to be set when message is sent.

package main

import "fmt"

func main() {
	/** Alternative ways to declare variables **/
	//var messages chan string
	//messages = make(chan string)

	//var messages chan string = make(chan string)

	messages := make(chan string, 1) // Explicitly declare buffer size of channel.
	messages <- "ping"

	msg := <-messages
	fmt.Println(msg)
}
