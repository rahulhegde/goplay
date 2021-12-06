package grpcprotobufplay

import (
	"fmt"
	"time"
)

func GRPCProtoBufPlay() {
	fmt.Println("*** GRPC Protobuffer Play ***")

	c := make(chan string, 2)
	go Server(c)

	// hard coded sleep to get enough time for server to start
	time.Sleep(5 * time.Second)

	go Client(c)

	// wait until client receives response
	value := <-c
	fmt.Println("sent: ", value)
}
