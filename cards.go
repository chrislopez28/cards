package cards

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const (
	Club    Suit = "C"
	Diamond Suit = "D"
	Heart   Suit = "H"
	Spade   Suit = "S"
	NoSuit  Suit = "_"
)

const (
	Ace   CardValue = "A"
	Two   CardValue = "2"
	Three CardValue = "3"
	Four  CardValue = "4"
	Five  CardValue = "5"
	Six   CardValue = "6"
	Seven CardValue = "7"
	Eight CardValue = "8"
	Nine  CardValue = "9"
	Ten   CardValue = "T"
	Jack  CardValue = "J"
	Queen CardValue = "Q"
	King  CardValue = "K"
	Joker CardValue = "★"
)

type Suit string

type CardValue string

// Card -----------------------------------------------------------------------

type Card struct {
	Suit  Suit
	Value CardValue
}

func (c Card) PrintCard() {
	fmt.Println(c.String())
}

// String returns a short representation of the card (e.g. AH for Ace of Hearts, 2C for Two of Clubs)
func (c Card) String() string {
	return fmt.Sprintf("%v%v", c.Value, c.Suit)
}

// Card Slices ----------------------------------------------------------------

func IsCardStackEmpty(cs []Card) bool {
	return len(cs) == 0
}

func InsertCard(c Card, cs []Card) ([]Card, error) {
	cs = append(cs, c)

	return cs, nil
}

func InsertCardBottom(c Card, cs []Card) ([]Card, error) {
	cs = append([]Card{c}, cs...)

	return cs, nil
}

func TakeCard(cs []Card) (Card, []Card, error) {
	if len(cs) == 0 {
		return Card{}, cs, errors.New("no cards left")
	}

	var c Card
	c, cs = (cs)[len(cs)-1], (cs)[:(len(cs)-1)]

	return c, cs, nil
}

// Shuffle implements the Fisher-Yates shuffle
func Shuffle(cs []Card) []Card {
	rand.Seed(time.Now().UnixNano())

	for i := len(cs) - 1; i > 0; i-- {
		j := rand.Intn(i)
		cs[i], cs[j] = cs[j], cs[i]
	}

	return cs
}

// Deck -----------------------------------------------------------------------

type Deck []Card

func LoadDeckN(numberDecks int) Deck {
	suits := [4]Suit{Club, Diamond, Heart, Spade}
	values := [13]CardValue{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}
	var newDeck Deck

	for i := 0; i < numberDecks; i++ {
		for _, suit := range suits {
			for _, value := range values {
				newDeck = append(newDeck, Card{suit, value})
			}
		}
	}

	return newDeck
}

func LoadDeck() Deck {
	return LoadDeckN(1)
}

func (d *Deck) AddJokerN(n int) {
	var err error

	for i := 0; i < n; i++ {
		*d, err = InsertCard(Card{Suit: NoSuit, Value: Joker}, *d)

		if err != nil {
			fmt.Print(err)
			return
		}
	}
}

func (d Deck) PrintCards() {
	for _, card := range d {
		card.PrintCard()
	}
}

func (d Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())

	for i := len(d) - 1; i > 0; i-- {
		j := rand.Intn(i)
		d[i], d[j] = d[j], d[i]
	}
}

func (d *Deck) DealCard() Card {
	var c Card
	c, *d = (*d)[len(*d)-1], (*d)[0:(len(*d)-1)]

	return c
}

func (d *Deck) DealCardBottom() Card {
	var c Card
	c, *d = (*d)[0], (*d)[1:]

	return c
}

func (d *Deck) DealCards(n int) []Card {
	var x []Card

	x, *d = (*d)[len(*d)-n:], (*d)[:len(*d)-n]

	return x
}

func (d *Deck) DealCardsBottom(n int) []Card {
	var x []Card

	x, *d = (*d)[0:n-1], (*d)[n:]

	return x
}
