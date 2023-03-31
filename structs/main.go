package main

import "fmt"

type Info struct {
	Email      string
	Address    string
	IsVerified bool
}

// Person struct with Name, Age and Addresses. Starting struct name and fields with Capital letters so that they can accessed from anywhere making them public
type Person struct {
	Name       string
	Age        int32
	PersonInfo *Info
}

func main() {

	raviInfo := &Info{
		Email:      "ravi.somani@druva.com",
		Address:    "14/534, Vasant Colony",
		IsVerified: true,
	}

	ravi := &Person{
		Name:       "Ravi Somani",
		Age:        23,
		PersonInfo: raviInfo,
	}

	fmt.Printf("Person %+v\n", ravi)
	fmt.Printf("Ravi Info %+v\n", ravi.PersonInfo)

}
