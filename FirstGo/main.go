package main

import "fmt"

func main() {
	card := "hello cards"
	cards := []string{"one", "two", "three"}
	for _, card := range cards {
		fmt.Println(card)
	}
	fmt.Println(card)
	fmt.Println(newCard())
}
func newCard() string {
	return "Ace of spades"
}
