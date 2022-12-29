package cards

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Suit string
type CardValue string

const (
	Club    Suit = "C"
	Diamond Suit = "D"
	Heart   Suit = "H"
	Spade   Suit = "S"
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
)

type Card struct {
	Suit  Suit
	Value CardValue
}

type Deck []Card
type Hand []Card

type Cards interface {
	PrintCards()
}

func (d Deck) PrintCards() {
	for _, card := range d {
		card.PrintCard()
	}
}
func (h Hand) PrintCards() {
	for _, card := range h {
		card.PrintCard()
	}
}

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

func (c Card) PrintCard() {
	fmt.Printf("%v%v\n", c.Value, c.Suit)
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

func IsCardStackEmpty(cs []Card) bool {
	if len(cs) == 0 {
		return true
	} else {
		return false
	}
}

func TakeCard(cs []Card) (Card, []Card, error) {
	if len(cs) == 0 {
		return Card{}, cs, errors.New("no cards left")
	}

	var c Card
	c, cs = (cs)[len(cs)-1], (cs)[:(len(cs)-1)]

	return c, cs, nil
}

func InsertCardBottom(c Card, cs []Card) ([]Card, error) {
	cs = append([]Card{c}, cs...)

	return cs, nil
}

// func main() {
// 	d := LoadDeck()

// 	// d[0].PrintCard()
// 	d.Shuffle()

// 	// d[3].PrintCard()

// 	h := d.TakeCard()
// 	fmt.Println(len(d))
// 	h.PrintCard()

// 	hand := d.TakeCards(1)

// 	for _, card := range hand {
// 		card.PrintCard()
// 	}
// 	fmt.Println(len(d))
// }
