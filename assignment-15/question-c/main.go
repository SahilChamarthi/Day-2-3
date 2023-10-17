package main

import (
	"fmt"
	"time"
)

func sender(ch chan int) {

	go func() {
		for i := 1; i <= 5; i++ { // sending channel in loop
			ch <- i
		}
	}()

}

func receiver(ch chan int) {

	go func() {
		for num := range ch {
			fmt.Println("receiving channel", num)
		}
	}()

}

func main() {

	ch := make(chan int, 5) // Buffered channel

	for i := 1; i <= 3; i++ {
		go sender(ch)
	}

	for i := 1; i <= 2; i++ {
		go receiver(ch)
	}

	time.Sleep(5 * time.Second)

	fmt.Println("end of main")

}
