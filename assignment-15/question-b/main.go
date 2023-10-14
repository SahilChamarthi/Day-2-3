package main

import (
	"fmt"
	"time"
)

func sender(ch chan int, cap int) {
	for i := 1; i <= cap; i++ { // sending channel in loop
		ch <- i
	}
}

func receiver(ch chan int) {
	for num := range ch {
		fmt.Println("receiving channel", num)
	}
}

func main() {

	cap := 3

	ch := make(chan int, cap) // Buffered channel

	go sender(ch, cap)
	go receiver(ch)

	time.Sleep(1000 * time.Millisecond)

	fmt.Println("end of main")

}
