// Example demonstrates closing channels.
//  Channels are not usually closed, but if they are, it is done by the sender.

package main

import "fmt"

func send(messages chan<- string, msg string) {
	messages <- msg
	close(messages)
}

func receive(messages <-chan string) {
	for {
		select {
		case msg, ok := <-messages:
			if !ok {
				fmt.Println("Channel closed")
				return
			}

			fmt.Println(msg)
		}
	}
}

func main() {
	messages := make(chan string, 1) // Explicitly declare buffer size of channel.

	send(messages, "hello")
	go receive(messages)
	fmt.Scanln()
}
