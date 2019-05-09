package main

import (
	fmt "fmt"
	"log"
	"reflect"

	proto "github.com/golang/protobuf/proto"
)

// testing protobuf
func ProtoBufPlay() {
	fmt.Println("ProtoBufPlay")
	person1 := Person{Id: 100, Email: "hegde.rahul@gmail.com"}
	person2 := Person{Id: 101, Email: "hegde.rahul@gmail.com"}
	book := &AddressBook{}
	book.Person = append(book.Person, &person1)
	book.Person = append(book.Person, &person2)
	fmt.Println("Address Book: ", book)
	data, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	fmt.Println("Reflection: ", reflect.TypeOf(data).ChanDir, ", data value: ", reflect.ValueOf(data))

	var bookcopy AddressBook
	proto.Unmarshal(data, &bookcopy)
	fmt.Println("Address Book Copy: ", bookcopy)

}
