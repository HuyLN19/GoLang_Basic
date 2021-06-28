package main

import (
	"fmt"
	"time"
)

func send(c chan int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("send %v to channel\n", i)
		c <- i // send i to channel
	}
}

func receive(c chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Millisecond)
		s := <-c // receive channel
		fmt.Println(s)
	}
}

func main() {
	c := make(chan int)
	go send(c)
	go receive(c)
	time.Sleep(time.Second)
	fmt.Println("end")
}
