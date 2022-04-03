package main

import "fmt"

//Insert variables declarations here based on activity

func tellMeTypes() {
	//insert code here to print out types of variables
	text := "The following is the account information"
	firstName := "Luke"
	lastName := "Skywalkter"
	age := 20
	weight := 73
	height := 1.72
	credit := 123.55
	accountName := "admin"
	accountPassword := "password"
	isSubscribed := true

	fmt.Printf("value: %v, type: %T \n", text, text)
	fmt.Printf("value: %v, type: %T \n", firstName, firstName)
	fmt.Printf("value: %v, type: %T \n", lastName, lastName)
	fmt.Printf("value: %v, type: %T \n", age, age)
	fmt.Printf("value: %v, type: %T \n", weight, weight)
	fmt.Printf("value: %v, type: %T \n", height, height)
	fmt.Printf("value: %v, type: %T \n", credit, credit)
	fmt.Printf("value: %v, type: %T \n", accountName, accountName)
	fmt.Printf("value: %v, type: %T \n", accountPassword, accountPassword)
	fmt.Printf("value: %v, type: %T \n", isSubscribed, isSubscribed)
}

func main() {
	tellMeTypes()
}
