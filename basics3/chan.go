// Example demonstrates channels acting as pipes pipes between concurrent goroutines
//  Receiver is setup before message is sent.
//  Example executes successfully.

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	messages := make(chan string) // Channels are unbuffered by default.
	var wg sync.WaitGroup
	var counter = 0

	go func(wg *sync.WaitGroup) {
		for {
			select {
			case msg := <-messages:
				fmt.Println(msg)
				wg.Done()
			}
		}
	}(&wg)

	for i := range time.Tick(1 * time.Second) {
		messages <- fmt.Sprintf("ping. %v", i)
		wg.Add(1)
		counter++

		if counter == 5 {
			break
		}
	}

	wg.Wait()
}
