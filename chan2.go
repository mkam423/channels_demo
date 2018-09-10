// Example sets message receiver to be ready before message is set.

package main2

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
