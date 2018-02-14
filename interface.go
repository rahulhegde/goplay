package main

import (
	"fmt"
	"math"
)

/*
	Interface defines list of methods signature - that will be implemented
	by the concrete structure like Circle and Rectangle.

	Secondly - receiver can be either pointer or value. this cannot be used
	interchange-able. Pointer defined variable cannot use interface defined
	with pass-by-value receiver.

	Third - interface defined object cannot access structure directly alike
	interface.variable however this needs to be type-casted and then used.

	Fourth - how do i find the data-type of the interface (InterfaceConversionPlay)
	==> use <object>.(concrete-type-that-is-implementing-interface)
	Type Assertion (
	https://golang.org/ref/spec#Type_assertions
	https://stackoverflow.com/questions/38816843/explain-type-assertions-in-go)

	Fifth: Interface variable cannot be nil like a pointer cannot be nil.
 */
type Shape interface {
	Area() float64
	Name()	string
}

type Circle struct {
	Radius float64
}


func (cir *Circle) Area () float64 {
	return math.Pi * cir.Radius * cir.Radius
}

func (cir *Circle) Name() string {
	return "Circle"
}

type Rectangle struct {
	Breadth float64
	Length 	float64
}

func (rec Rectangle) Area() float64 {
	return rec.Breadth * rec.Length
}

func (cir Rectangle) Name() string {
	return "Rectangle"
}

type Square struct {
	Size float64
}

func InterfacePlay()  {
	fmt.Println("*** InterfacePlay ***")
	var shape Shape

	shape = &Circle{100}
	fmt.Println("Area of", shape.Name(), shape.Area())

	shape = Rectangle{Length:20, Breadth:40}
	shape.Area()
	fmt.Println("Area of", shape.Name(), shape.Area())
}


func InterfaceTypeAssertionPlay () {
	var shape Shape
	shape = &Circle{Radius:10}

	var ok bool
	// result in run-time panic with syntax: _ = shape.(Rectangle)
	// hence used below assignment syntax to get the evaluated type-assert
	_, ok = shape.(Rectangle)
	fmt.Println("Conversion of circle interface to rectangle is success", ok)

	_, ok = shape.(*Circle)
	fmt.Println("Conversion of circle interface to circle is success", ok)

	/*
		Square is not interface -> this results in compilation error i.e.

		./interface.go:XX: impossible type assertion:
		Square does not implement Shape (missing Area method)

		_, ok = shape.(Square)
		fmt.Println("Conversion of circle interface to square is success", ok)
	*/


	/*
		// static-type conversion is a strict check, circle variable is not interface
		var circle Circle
		_, ok = circle.(Rectangle)
		fmt.Println("Conversion of circle interface to rectangle is success", ok)
	*/
}

func TypeSwitchPlay()  {
	var shape Shape
	shape = &Circle{100}
	//shape = Rectangle{Breadth:100, Length:200}

	switch inter := shape.(type) {
	case *Circle:
		fmt.Printf("Circle found [%+v]", inter)
	case Rectangle:
		fmt.Printf("rectange found [%+v]", inter)
	default:
		fmt.Println("Interface type found")
	}
}
