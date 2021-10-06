package main

import "fmt"

func main() {

	//Variable

	//variable declaration without an initializer
	var emptyNumber int
	emptyNumber = 100

	//variable declaration without an initializer and multiple identifier
	var day, month, year int
	day, month, year = 10, 11, 1999

	//variable declaration with an initializer and type
	var parkNumber int = 29

	//variable declaration with multiple initializer
	var roomNumber, floorNumber int = 1, 2

	//variable declaration without type
	var firstName = "Bagus"

	//short variable declaration
	//Warning : you cannot use the value nil for a short variable declaration.
	//Warning : short variable declaration cannot be used outside functions !

	//short variable declaration with one identifier
	lastName := "Bagus"

	//short variable declaration with multiple identifier
	address, region := "Singaraja", "Bali"
	fmt.Println(emptyNumber, day, month, year, parkNumber, roomNumber, floorNumber, firstName, lastName, address, region)

	//Constant

	//typed constant example
	const version string = "1.3.2"

	//untyped constant example
	//An untyped constant has no type , default type , and limits
	const versionSemantic = "1.3.2.2"

	fmt.Println(version, versionSemantic)
	//example of untyped constant has no type
	const occupancyLimit = 12

	var occupancyLimit1 uint8
	var occupancyLimit2 int64
	var occupancyLimit3 float32

	// assign our untyped const to an uint8 variable
	occupancyLimit1 = occupancyLimit
	// assign our untyped const to an int64 variable
	occupancyLimit2 = occupancyLimit
	// assign our untyped const to a float32 variable
	occupancyLimit3 = occupancyLimit

	//Our program compiles. The value of our constant can be put into variables of different types!
	fmt.Println(occupancyLimit1, occupancyLimit2, occupancyLimit3)

	//An untyped constant has a default type that is defined by the value assigned
	//to it at compilation. In our example, occupancyLimit has a default type of int
	//. An int cannot be assigned to a string variable.

	//const occupancyLimit5 = 12
	//
	//var occupancyLimit6 string
	//
	//this code will produce error
	//occupancyLimit5 = occupancyLimit6
	//
	//fmt.Println(occupancyLimit5)

}
