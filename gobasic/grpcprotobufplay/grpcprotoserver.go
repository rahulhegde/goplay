package grpcprotobufplay

import (
	context "context"
	"fmt"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

// server is used to implement helloworld.GreeterServer.
type AddressBookServer struct {
	UnimplementedAddressBookApiServer
}

// SayHello implements helloworld.GreeterServer
func (s *AddressBookServer) SayHello(ctx context.Context, in *AddressBookRequest) (*AddressBookResponse, error) {
	log.Printf("server: received: %v", in)
	return &AddressBookResponse{Response: "completes grpc request-response test"}, nil
}

func Server(channel chan string) {
	lis, _ := net.Listen("tcp", "localhost:50051")
	grpcServer := grpc.NewServer()
	RegisterAddressBookApiServer(grpcServer, &AddressBookServer{})
	fmt.Println("server started listening: ", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	channel <- "server processing complete"
}
