// Example demonstrates usage of select statements for receipt of message from multiple channels
// Scenario simulated is a restaurant where orders can be made at driveThrough or to sitDown
//   Requests occur asynchronously with one common point being the handler of these requests.

package main

import (
	"fmt"
	"time"
)

func sitDownOrder(c chan<- string) {
	for range time.Tick(time.Second * 2) {
		c <- "For Here"
	}
}

func driveThroughOrder(d chan<- string) {
	for range time.Tick(time.Second * 1) {
		d <- "To Go"
	}
}

func handler(c <-chan string, d <-chan string) {
	for {
		select {
		// Handle Req for sitDown
		case sitDown := <-c:
			fmt.Println(sitDown)
		// Handle Req for takeOut
		case takeOut := <-d:
			fmt.Println(takeOut)
		}
	}
}

func main() {
	var c = make(chan string)
	var d = make(chan string)

	go sitDownOrder(c)
	go driveThroughOrder(d)
	go handler(c, d)

	fmt.Scanln()

}
