package card

import (
	"fmt"
	"math/rand"
)

// Card represents a card in a deck with a suit, value, an image path, and an indication if it's an Ace.
type Card struct {
	Suit      string
	Value     int
	ImagePath string
	IsAce     bool
}

// NewCard creates a new card with a given suit and value.
func NewCard(suit string, value int) *Card {
	var rank string
	isAce := false
	switch value {
	case 11:
		rank = "jack"
	case 12:
		rank = "queen"
	case 13:
		rank = "king"
	case 1:
		rank = "ace"
		isAce = true
	default:
		rank = fmt.Sprintf("%d", value)
	}
	imagePath := fmt.Sprintf("card_%s_of_%s.png", rank, suit)
	return &Card{Suit: suit, Value: value, ImagePath: imagePath, IsAce: isAce}
}

// Deck represents a collection of cards with its own random generator.
type Deck struct {
	Cards []*Card // Changed from 'cards' to 'Cards' to export it
	rng   *rand.Rand
}

// NewDeck creates and returns a new deck with shuffled cards.
func NewDeck() *Deck {
	deck := &Deck{
		Cards: make([]*Card, 0, 52),                   // Allocate slice for cards
		rng:   rand.New(rand.NewSource(rand.Int63())), // Initialize a local random generator
	}
	suits := []string{"hearts", "diamonds", "clubs", "spades"}
	for _, suit := range suits {
		for value := 1; value <= 13; value++ {
			deck.Cards = append(deck.Cards, NewCard(suit, value))
		}
	}
	deck.Shuffle() // Optionally shuffle here
	return deck
}

// Shuffle randomizes the order of cards in the deck.
func (d *Deck) Shuffle() {
	d.rng.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}
