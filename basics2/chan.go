// Example will work b/c channels act as pipes between concurrent goroutines
//  Ping message is blocked from sending until there is a receiving end of the channel.
package main

import "fmt"

func main() {
	/** Alternative ways to declare variables **/
	//var messages chan string
	//messages = make(chan string)

	//var messages chan string = make(chan string)

	messages := make(chan string)      // Channels are unbuffered by default.
	go func() { messages <- "ping" }() // Go Routine + Anonymous function (Closures)

	msg := <-messages
	fmt.Println(msg)
}
