package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var agePaul, ageJohn int = rand.Intn(110), rand.Intn(110)

	fmt.Println("John is", ageJohn, "years old")
	fmt.Println("Paul is", agePaul, "years old")

	// example if statement
	if agePaul > ageJohn {
		fmt.Println("Paul is older than John")
	} else {
		// another if/else structure
		if agePaul == ageJohn {
			fmt.Println("Paul and John have the same age")
		} else {
			fmt.Println("Paul is younger than John")
		}
	}

	// example if statement with else if
	if agePaul > ageJohn {
		fmt.Println("Paul is older than John")
	} else if agePaul == ageJohn {
		fmt.Println("Paul and John have the same age")
	} else {
		fmt.Println("Paul is younger than John")
	}

	// switch expression
	switch ageJohn {
	case 10:
		fmt.Println("John is 10 years old")
	case 20:
		fmt.Println("John is 20 years old")
	case 100:
		fmt.Println("John is 100 years old")
	default:
		fmt.Println("John has neither 10,20 nor 100 years old")
	}

	// switch statement; expression
	switch ageSum := ageJohn + agePaul; ageSum {
	case 10:
		fmt.Println("ageJohn + agePaul = 10")
	case 20, 30, 40: //*\label{switchMulti}
		fmt.Println("ageJohn + agePaul = 20 or 30 or 40")
	case 2 * agePaul:
		fmt.Println("ageJohn + agePaul = 2 times agePaul")
	}

	// switch (without statement, without expression)
	switch {
	case agePaul > ageJohn:
		fmt.Println("agePaul > ageJohn")
	case agePaul == ageJohn:
		fmt.Println("agePaul == ageJohn")
	case agePaul < ageJohn:
		fmt.Println("agePaul < ageJohn")
	}

	//for statement

	const emailToSend = 3
	emailSent := 0

	//for statement with single statement
	// it's basically while statement in another programming language
	for emailSent < emailToSend {
		fmt.Println("sending email..")
		emailSent++ //*\label{forWithSingle2}
	}
	fmt.Println("all email send")

	//For statement with a for clause
	for i := 0; i < emailToSend; i++ {
		fmt.Println("interation N", i)
	}

}
