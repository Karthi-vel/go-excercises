package main

type deck []string

func main() {

	//cards := newDeck()
	//cards.print()
	//handCards, remainingCards := deal(cards, 13)
	//handCards.print()
	//remainingCards.print()
	//cards.WriteToFile("My_Cards.txt")
	cards := reafFromFile("My_Cards.txt")
	cards.print()
	
}

//func WriteFile(filename string, data []byte, perm fs.FileMode) error

//func Join(elems []string, sep string) string

