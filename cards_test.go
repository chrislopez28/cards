package cards

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDealCard(t *testing.T) {
	d := LoadDeck()

	expectedCard := Card{Suit: Spade, Value: King}
	expectedNewTopCard := Card{Suit: Spade, Value: Queen}

	if d[len(d)-1] != expectedCard {
		t.Errorf("Expected top card to be King of Spades")
	}

	hand := d.DealCard()

	if hand != expectedCard {
		t.Errorf("Expected taken card not the same as the actual taken card.")
	}

	if len(d) != 51 {
		t.Errorf("Expected deck to be 51 cards after taking the card.")
	}

	if d[len(d)-1] != expectedNewTopCard {
		t.Errorf("Expected top card to be Queen of Spades after taking the card.")
	}
}

func TestDealCardBottom(t *testing.T) {
	d := LoadDeck()

	expectedCard := Card{Suit: Club, Value: Ace}
	expectedNewBottomCard := Card{Suit: Club, Value: Two}

	if d[0] != expectedCard {
		t.Errorf("Expected bottom card to be Ace of Clubs")
	}

	hand := d.DealCardBottom()

	if hand != expectedCard {
		t.Errorf("Expected taken card not the same as the actual taken card.")
	}

	if len(d) != 51 {
		t.Errorf("Expected deck to be 51 cards after taking the card.")
	}

	if d[0] != expectedNewBottomCard {
		t.Errorf("Expected top card to be Ace of Clubs after taking the card.")
	}
}

func TestInsertJokerN(t *testing.T) {
	d := LoadDeck()
	d.AddJokerN(1)

	expectedCard := Card{Suit: NoSuit, Value: Joker}

	actualCard := d[len(d)-1]

	if !cmp.Equal(expectedCard, actualCard) {
		t.Errorf("Expected top card to be %s. Got %s instead", expectedCard.String(), actualCard.String())
	}

	if len(d) != 53 {
		t.Errorf("Expected deck length to be 53. Got %v instead.", len(d))
	}

	d.AddJokerN(2)

	if len(d) != 55 {
		t.Errorf("Expected deck length to be 55. Got %v instead.", len(d))
	}
}

func TestInsertCardBottom(t *testing.T) {
	c1 := Card{Suit: Club, Value: Ace}
	c2 := Card{Suit: Club, Value: Two}
	c3 := Card{Suit: Club, Value: Three}

	cs := []Card{c2, c3}

	if !cmp.Equal(cs[0], c2) {
		t.Errorf("Card slice not initialized correctly.")
	}

	cs, err := InsertCardBottom(c1, cs)

	if err != nil {
		t.Errorf("InsertCardBottom error: %s", err)
	}

	if !cmp.Equal(cs[0], c1) {
		t.Errorf("Card not inserted to the bottom of card slice")
	}
}
