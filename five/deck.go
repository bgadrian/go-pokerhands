package five

import (
	"math/rand"
	"time"
)

func GenerateDeck() []Card {
	res := make([]Card, 0, 52)
	for r := Deuce; r <= Ace; r++ {
		for s := SPADE; s <= CLUB; s <<= 1 {
			res = append(res, Card{Suit: s, Rank: r})
		}
	}
	return res
}

func NewHandGenerator() *HandGenerator {
	return &HandGenerator{deck: GenerateDeck()}
}

// HandGenerator is a pseudo random hand generator, it is not crypto safe and not performant (does memory allocs)
type HandGenerator struct {
	deck []Card
}

// FiveHand returns a random set of 5 distinct cards
func (g *HandGenerator) FiveHand() FiveHand {
	rand.Seed(time.Now().Unix()) //force a seed for some envs and force more entropy

	rand.Shuffle(len(g.deck), func(i, j int) {
		g.deck[i], g.deck[j] = g.deck[j], g.deck[i]
	})
	//to avoid slice as a pointer to be reused, we copy the cards
	result := make(FiveHand, 5)
	copy(result, g.deck) //will only copy 5
	return result
}
