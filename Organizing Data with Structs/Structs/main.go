package main

import "fmt"

//embedding another struct
type contactInfo struct {
	email   string
	zipCode int
}

//Defning a struct
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// //Declaring Structs
	// //1st method
	// alex := person{"Alex", "Standall"}

	// //2nd method
	// hannah := person{firstName: "Hannah", lastName: "Baker"}

	// //3rd method
	// var justin person
	// justin.firstName = "Justin"
	// justin.lastName = "Foley"

	// //Different ways of Printing
	// fmt.Println(alex)
	// fmt.Println(hannah)
	// fmt.Printf("%+v ", justin)

	var narissa person
	narissa.firstName = "Narissa"
	narissa.lastName = "Flower"
	narissa.contactInfo.email = "narissaf@gmail.com"
	narissa.contactInfo.zipCode = 45128

	fmt.Println(narissa)
	narissaPointer := &narissa
	narissaPointer.updateName("Jenny")
	narissa.print()
}

//Creating Reciever functions for structure

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v ", p)
}
