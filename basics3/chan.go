// Example demonstrates channels acting as pipes pipes between concurrent goroutines
//  Receiver is setup before message is sent.
//  Example executes successfully.

package main

import "fmt"

func main() {
	/** Alternative ways to declare variables **/
	//var messages chan string
	//messages = make(chan string)

	//var messages chan string = make(chan string)

	messages := make(chan string) // Channels are unbuffered by default.

	// Go Routine + Anonymous function (Closures)
	go func() {
		for {
			select {
			case msg := <-messages:
				fmt.Println(msg)
			}
		}
	}()

	messages <- "ping"
}
