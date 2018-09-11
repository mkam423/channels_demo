// Example demonstrates unbuffered channels.
//  Example will not work b/c receiver is not set when message is sent.

package main

import "fmt"

func main() {
	/** Alternative ways to declare variables **/
	//var messages chan string
	//messages = make(chan string)

	//var messages chan string = make(chan string)

	messages := make(chan string) // Channels are unbuffered by default.
	messages <- "ping"

	msg := <-messages
	fmt.Println(msg)
}
