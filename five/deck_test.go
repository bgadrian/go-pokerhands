package five

import (
	"testing"
)

// test that all the cards are different, and it will also test the labels helpers
func TestGenerateDeck(t *testing.T) {
	d := GenerateDeck()
	if len(d) != 52 {
		t.Error("it should have 52 cards")
	}
	for i := range d {
		firstCard := d[i]

		_, err := CardRankLabel(firstCard.Rank)
		if err != nil {
			t.Error(err, "rank is not valid", firstCard)
		}
		_, err = CardSuitLabel(firstCard.Suit)
		if err != nil {
			t.Error(err, "suit is not valid", firstCard)
		}
		//this algorithm is O(n^n) complexity, but its just a test ...
		for j := i + 1; j < 52; j++ {
			secondCard := d[j]
			if firstCard == secondCard {
				t.Errorf("found duplicate card: %+v", secondCard)
			}
		}
	}
}

func TestNewHandGenerator(t *testing.T) {
	g := NewHandGenerator()
	for i := 0; i < 100; i++ {
		hand := g.FiveHand()

		if len(hand) != 5 {
			t.Error("should have 5 cards")
		}
	}
}
