package main

import "fmt"

func main() {
	go func() {
		for {
			fmt.Println("sub thread is running")
		}
	}()

	for {
		fmt.Println("main thread is running")
	}
}
