package main

import "fmt"

//Insert variables declarations here based on activity

func tellMeTypes() {
	//insert code here to print out types of variables
	text := "The following is the account information"
	credit := 123.55

	fmt.Printf("%T %T", text, credit)

}

func main() {
	tellMeTypes()
}
