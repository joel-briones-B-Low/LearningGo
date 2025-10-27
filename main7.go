package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello7(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Hello from goroutine!"
}

func printMessage7(channel <-chan string) {
	println("Waiting to receive message from channel...")
	mes := <-channel
	fmt.Println(mes)
}
func main7() {

	channel := make(chan string)
	println("Starting goroutine...")
	go sayHello7(channel)
	println("Waiting for message...")
	printMessage7(channel)
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

	var contador int
	var mu sync.RWMutex

	// writter
	go func() {

		for i := 0; i < 6; i++ {
			if i == 3 {

			}
			mu.Lock()
			contador++
			mu.Unlock()
			time.Sleep(50 * time.Millisecond)
		}
	}()

	// reader
	for i := 0; i < 5; i++ {
		go func(id int) {
			for j := 0; j < 5; j++ {

				mu.RLock()
				fmt.Printf("Goroutine %d: contador = %d\n", id, contador)
				mu.RUnlock()
				time.Sleep(30 * time.Millisecond)
			}
		}(i)
	}
	time.Sleep(2 * time.Second)
}
