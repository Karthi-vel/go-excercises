package main

import (
	"fmt"
	"os"
	"strings"
)

func newDeck() deck {
	suits := []string{"Hears", "Clubs", "Diamonds", "Spades"}
	numbers := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	d := []string{}
	for _, suit := range suits {
		for _, num := range numbers {
			d = append(d, num + " of " + suit)
		}
	}
	return d
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i," " ,card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}


//func WriteFile(filename string, data []byte, perm fs.FileMode) error

//func Join(elems []string, sep string) string
func (d deck) toString() string {
	return strings.Join(d, ",")
}
func (d deck) WriteToFile(filename string) {
	os.WriteFile(filename, []byte(d.toString()), 0777)
}

// func Split(s, sep string) []string
//func ReadFile(name string) ([]byte, error)
func reafFromFile(filename string) deck {
	fb,_ := os.ReadFile(filename)
	return strings.Split(string(fb), ",")
}