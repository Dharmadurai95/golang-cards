package main

func main() {

	// cards := newDeck()
	cards := readFile("cardsFile")
	cards.shuffle()
	cards.print()
	//if you want pick first three element [0:3] 0 index (include)start 3 ending index (exclude) 0,1,2

	//deal get some cards
	// hand, remaining := deal(cards, 3)

	//save file
	// err := cards.saveFile("cardsFile")
	// fmt.Println(err)

}
