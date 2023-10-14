package main

import (
	"fmt"
	"time"
)

func sender(ch chan int) {
	for i := 1; i <= 5; i++ { // sending channel in loop
		ch <- i
	}
}

func receiver(ch chan int) {
	for num := range ch {
		fmt.Println("receiving channel", num)
	}
}

func main() {

	ch := make(chan int) // unbuffered channel

	go sender(ch)
	go receiver(ch)

	time.Sleep(1000 * time.Millisecond)

	fmt.Println("end of main")

}
