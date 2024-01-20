package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create new type of type 'deck'
// which is slice of string

type deck []string //why we use custom type becuase we can attache custom methods and functionality

func (d deck) print() {
	for i, v := range d {
		fmt.Println(i, v)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamond", "Hearts", "Clubs"}
	cardValue := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, cardVal := range cardValue {
			cards = append(cards, cardVal+" of "+suit)
		}
	}
	return cards
}

func deal(cards deck, handSize int) (deck, deck) {
	return cards[:handSize], cards[handSize:]
}

func (cards deck) toString() string {
	return strings.Join([]string(cards), ":")
}
func (cards deck) saveFile(filename string) error {
	err := os.WriteFile(filename, []byte(cards.toString()), 0666)
	if err != nil {
		return err
	}
	return nil
}

func readFile(fileName string) deck {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return deck(strings.Split(string(data), ":"))
}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for index, _ := range d {
		randomIndex := r.Intn(len(d) - 1)
		d[index], d[randomIndex] = d[randomIndex], d[index]
	}
}
