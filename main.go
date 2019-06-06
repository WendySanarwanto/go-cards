package main

import "fmt"

var aNumber int

func main() {
	// Declaring a string variable in go
	var aceSpadeCard string = "Ace of Spades"

	// Alternative way to declare and assign a variable
	heartCard := "Ace of Hearts"
	fmt.Println(heartCard)

	// Re-assigning a new value to the heartCard
	heartCard = "Five of Hearts"

	// Integer var
	// var aNumber int
	aNumber = 10

	fmt.Println(aceSpadeCard)
	fmt.Println(heartCard)
	fmt.Println(aNumber)
}