package main

import (
	"fmt"
	"time"
)

// Routine player spawing a new thread
func GoRoutinePlay() {
	go GoRoutinePrinter("Rahul", 100)
	GoRoutinePrinter("Hegde", 200)
}

func GoRoutinePrinter(name string, d time.Duration) {
	for i := 0; i < 2; i++ {
		fmt.Println("name: ", name)
		time.Sleep(d * time.Millisecond)
	}
}

// channel responder sending back what is sent to u
func ChannelResponder(test int, channel chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("sending: ", test, i)
		channel <- test
		fmt.Println("sent: ", test)
	}
}

// channel responder sending back what is sent to u
func ChannelReceiver(channel chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("receiving: ", i)
		value := <-channel
		fmt.Println("received: ", value)
		time.Sleep(1 * time.Second)
	}
}

func GoChannelPlay() {
	c := make(chan int, 1)
	random := 3210
	go ChannelResponder(random, c)
	go ChannelReceiver(c)
	fmt.Println("time: ", time.Now().String())
	time.Sleep(10 * time.Second)
	//time.Tick()
	//value := <- c
	//fmt.Println("Received: ", value)
}
