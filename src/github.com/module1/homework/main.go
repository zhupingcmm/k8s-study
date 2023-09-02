package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	messages := make(chan int, 10)
	done := make(chan bool)

	defer close(messages)
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-done:
				fmt.Println("child process interrupt ...")
			default:
				fmt.Printf("Consumer: %d\n", <-messages)
			}
		}
	}()

	go func() {
		ticker := time.NewTicker(2 * time.Second)
		for _ = range ticker.C {
			num := rand.Intn(100)
			messages <- rand.Intn(num)
			select {
			case messages <- num:
				fmt.Printf("producer: %d\n", num)
			default:
				fmt.Println("messages queue is full")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)

	close(done)
	time.Sleep(time.Second)
	fmt.Println("main process exit")

}
