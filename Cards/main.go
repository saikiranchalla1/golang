package main

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// fmt.Println(card)
	// cards := deck{"Ace of Diamonds", newCard()}

	// cards = append(cards, "Six of Spades")

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// cards := newDeck()
	// cards.print()
	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()
	// greeting := "Hi, There"

	// fmt.Println([]byte(greeting))
	// fmt.Println(cards.toString())

	// cards.saveToFile("mycards")

	// cards := newDeckFromFile("mycards")
	// cards.print()

	// cards.shuffle()
	// cards.print()

	//var card string = "Ace of Spades"
	// cards := newDeck()
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()

	// greeting := "Hi, There"
	// fmt.Println([]byte(greeting))
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
