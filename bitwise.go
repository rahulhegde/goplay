package main

import "fmt"

type Flag int

const (
	EAT Flag = 1
	DRINK Flag = 2
	SLEEP Flag = 4
)

type PersonStatus struct {
	status Flag
}

func (person *PersonStatus) SetEat() {
	person.status = person.status | EAT
}

func (person *PersonStatus)  UnsetEat() {
	person.status = person.status & ^EAT
}

func (person *PersonStatus) SetDrink() {
	person.status = person.status | DRINK
}

func (person *PersonStatus)  UnsetDrink() {
	person.status = person.status & ^DRINK
}



func BitwiseCheckPlay() {
	fmt.Println("*** BitwiseCheckPlay ***")
	person := PersonStatus{}
	fmt.Println("person: ", person)
	person.SetDrink()
	person.SetEat()
	fmt.Println("person: ", person)
	person.UnsetDrink()
	fmt.Println("person: ", person)
	person.UnsetEat()
	fmt.Println("person: ", person)
}
