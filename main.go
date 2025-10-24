package main

import (
	"fmt"
	"time"
)

func sayHello(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Hello from goroutine!"
}

func printMessage(channel <-chan string) {
	println("Waiting to receive message from channel...")
	mes := <-channel
	fmt.Println(mes)
}
func main() {

	channel := make(chan string)
	println("Starting goroutine...")
	go sayHello(channel)
	println("Waiting for message...")
	printMessage(channel)
	println("Message received, exiting.")

	channel2 := make(chan int)
	go func() {
		println("Goroutine started.")

		for i := 0; i < 5; i++ {
			channel2 <- i
		}
		close(channel2)
	}()

	for val := range channel2 {
		fmt.Println("Received value:", val)
	}
}
