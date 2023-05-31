package five

import "math/rand"

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

// HandGenerator is a pseudo random hand generator, it is not crypto safe
type HandGenerator struct {
	deck []Card
}

// FiveHand do not reuse the slices
func (g *HandGenerator) FiveHand() []Card {
	rand.Shuffle(len(g.deck), func(i, j int) {
		g.deck[i], g.deck[j] = g.deck[j], g.deck[i]
	})
	return g.deck[:5]
}
