package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()
	if len(cards) != 16 {
		t.Errorf("expected 16 card but got %v", len(cards))
	}

	if cards[0] != "Ace of Spades" {
		t.Errorf("First card expected as Ace of Spades but got %v", cards[0])
	}
	if cards[len(cards)-1] != "Four of Clubs" {

		t.Errorf("Last card expected as Four of ClubsSpades but got %v", cards[len(cards)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	os.Remove("_testingFile")
	d := newDeck()
	d.saveFile("_testingFile")
	getDeckFromFile := readFile("_testingFile")
	if len(getDeckFromFile) != 16 {
		t.Errorf("expected 16 card but got %v", len(getDeckFromFile))
	}
	os.Remove("_testingFile")
}
