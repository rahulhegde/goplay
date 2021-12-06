package grpcprotobufplay

import (
	context "context"
	"fmt"
	"time"

	grpc "google.golang.org/grpc"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func constructAddressBook(addrbook *AddressBook) {
	// declare three person to be added to address books - nameN
	p1 := &Person{Name: "name1", Id: 1, Email: "name1@email.com"}
	p1number1 := &Person_PhoneNumber{Number: "name1:101", Type: Person_CELL}
	p1number2 := &Person_PhoneNumber{Number: "name1:102", Type: Person_HOME}
	p1number3 := &Person_PhoneNumber{Number: "name1:103", Type: Person_WORK}
	lastUpdated1 := &timestamppb.Timestamp{Seconds: 100, Nanos: 100}
	p1.LastUpdated = lastUpdated1
	p1.Phones = append(p1.Phones, p1number1, p1number2, p1number3)

	p2 := &Person{Name: "name2", Id: 2, Email: "name2@email.com"}
	p2number1 := &Person_PhoneNumber{Number: "name2:201", Type: Person_WORK}
	p2number2 := &Person_PhoneNumber{Number: "name2:202", Type: Person_WORK}
	p2number3 := &Person_PhoneNumber{Number: "name2:203", Type: Person_CELL}
	lastUpdated2 := &timestamppb.Timestamp{Seconds: 200, Nanos: 200}
	p2.LastUpdated = lastUpdated2
	p2.Phones = append(p2.Phones, p2number1, p2number2, p2number3)

	p3 := &Person{Name: "name3", Id: 3, Email: "name3@email.com"}
	p3number1 := &Person_PhoneNumber{Number: "name3:301", Type: Person_HOME}
	p3number2 := &Person_PhoneNumber{Number: "name3:302", Type: Person_HOME}
	p3number3 := &Person_PhoneNumber{Number: "name3:303", Type: Person_WORK}
	lastUpdated3 := &timestamppb.Timestamp{Seconds: 300, Nanos: 300}
	p3.LastUpdated = lastUpdated3
	p3.Phones = append(p3.Phones, p3number1, p3number2, p3number3)

	addrbook.Person = append(addrbook.Person, p1, p2, p3)
}

func Client(channel chan string) {
	addrbook := AddressBook{}
	constructAddressBook(&addrbook)
	addressbookRequest := &AddressBookRequest{AddressBook: &addrbook}

	// construct the address book request serialized bytes
	requestBytes, _ := proto.Marshal(addressbookRequest)
	fmt.Println("client: address book request bytes: ", requestBytes)

	// grpc client implementation (type connection)  used for RPC communication, can be http OR anything.
	conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer conn.Close()
	client := NewAddressBookApiClient(conn)

	// build the context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	addressbookResponse, _ := client.SayHello(ctx, addressbookRequest)
	fmt.Println("client: address book response: ", addressbookResponse)

	channel <- "client processing complete"
}
