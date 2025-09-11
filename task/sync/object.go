package sync

import (
	"fmt"
	"time"
)

type Shape interface {
	Area()
	Perimeter()
	Shape1
}

type Shape1 interface {
	Area1()
}
type Person struct {
	age  int
	name string
	Boy
	Shape1
}

type AType struct {
}

func (AType) Area1() {

}

func (person Person) printE() {
	fmt.Printf("p=%v\n", person)
}

type Boy struct {
	bage  int
	bname string
}

type Employee struct {
	Person
	//name string
	EmployeeId int
}

func (e Employee) printE() {
	e.Person.name = "zhangsan"
	e.bname = "lisi"
	e.Person.printE()
	fmt.Printf("e=%v\n", e)
}

type Rectangle struct {
	len   int
	width int
	AType
}

type Circle struct {
	len   int
	width int
	AType
}

func (r Rectangle) Area() {
	r.width = 10
	r.len = 11
	fmt.Printf("R=%v\n", r)
}
func (r Rectangle) Perimeter() {
	r.width = 10
	r.len = 11
	fmt.Printf("R=%v\n", r)
}

func (c *Circle) Area() {
	c.width = 10
	c.len = 11
	fmt.Printf("C=%v\n", c)
}

func (c Circle) Perimeter() {
	c.width = 10
	c.len = 11
	fmt.Printf("C=%v\n", c)
}

func call() {
	c := &Circle{}
	r := &Rectangle{}
	p := Person{}
	e := Employee{Person: p}
	c.Area()
	r.Area()
	e.printE()
	p.printE()

	fmt.Printf("C=%v\n", c)
	fmt.Printf("R=%v\n", r)
	fmt.Printf("E=%v\n", e)

	//many call
	//one interface call
	var ic Shape = c
	ic1 := Shape(r)
	ic1.Area()
	ic.Area()
	switch ic.(type) {
	case *Circle:
		fmt.Printf("ic type is %v\n", "*Circle")
	}

	test := []Shape{ic, ic1}
	go func([]Shape) {
		for _, b := range test {
			switch b.(type) {
			case *Rectangle:

				fmt.Printf("ic type is %v\n", "*Rectangle")
				b.(*Rectangle).Area()
				if b, ok := b.(*Rectangle); ok {
					b.Area()
				}
			case *Circle:
				fmt.Printf("ic type is %v\n", "*Circle")
			case Rectangle:
				fmt.Printf("ic type is %v\n", "Rectangle")
			default:
				fmt.Printf("ic type is %v\n", "nil")
			}
		}
	}(test)

	time.Sleep(10 * time.Second)
}
