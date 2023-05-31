package five

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func ComputeCategoryForHandScore(s CactusHandRank) HandRankCategory {
	switch {
	case s > 6185:
		return HIGH_CARD
	case s > 3325:
		return ONE_PAIR
	case s > 2467:
		return TWO_PAIR
	case s > 1609:
		return THREE_OF_A_KIND
	case s > 1599:
		return STRAIGHT
	case s > 322:
		return FLUSH
	case s > 166:
		return FULL_HOUSE
	case s > 10:
		return FOUR_OF_A_KIND
	default:
		return STRAIGHT_FLUSH
	}
}

var uniqueHandsCounts = map[HandRankCategory]int{
	HIGH_CARD:       1277,
	ONE_PAIR:        2860,
	TWO_PAIR:        858,
	THREE_OF_A_KIND: 858,
	STRAIGHT:        10,
	FLUSH:           1277,
	FULL_HOUSE:      156,
	FOUR_OF_A_KIND:  156,
	STRAIGHT_FLUSH:  10,
}

func TestCategoriesComputeHandScore(t *testing.T) {
	//go trough all unique possible hands
	deck := GenerateDeck()
	foundCounts := map[HandRankCategory]int{}
	for a := 0; a < 48; a++ {
		for b := a + 1; b < 49; b++ {
			for c := b + 1; c < 50; c++ {
				for d := c + 1; d < 51; d++ {
					for e := d + 1; e < 52; e++ {
						c1 := ComputeCardScore(deck[a])
						c2 := ComputeCardScore(deck[b])
						c3 := ComputeCardScore(deck[c])
						c4 := ComputeCardScore(deck[d])
						c5 := ComputeCardScore(deck[e])

						score := ComputeHandScore(c1, c2, c3, c4, c5)
						category := ComputeCategoryForHandScore(score)
						foundCounts[category]++
					}
				}
			}
		}
	}

	diff := cmp.Diff(uniqueHandsCounts, foundCounts)
	if diff != "" {
		t.Error(diff)
	}
}

func TestComputeCardScore(t *testing.T) {
	tests := []struct {
		args Card
		want CactusCardScore
	}{
		{args: Card{Rank: King, Suit: DIAMOND}, want: 134236965},
		{args: Card{Rank: Five, Suit: SPADE}, want: 529159},
		{args: Card{Rank: Jack, Suit: CLUB}, want: 33589533},
	}
	for _, tt := range tests {
		t.Run(tt.args.String(), func(t *testing.T) {
			if got := ComputeCardScore(tt.args); got != tt.want {
				t.Errorf("ComputeCardScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

// see https://suffe.cool/poker/7462.html
func TestComputeHandScore(t *testing.T) {
	type args struct {
		c1 CactusCardScore
		c2 CactusCardScore
		c3 CactusCardScore
		c4 CactusCardScore
		c5 CactusCardScore
	}
	tests := []struct {
		name string
		args args
		want CactusHandRank
	}{
		//test flushes array
		{name: "Six-High Straight Flush", args: args{
			c1: ComputeCardScore(Card{Rank: Six, Suit: DIAMOND}),
			c2: ComputeCardScore(Card{Rank: Five, Suit: DIAMOND}),
			c3: ComputeCardScore(Card{Rank: Four, Suit: DIAMOND}),
			c4: ComputeCardScore(Card{Rank: Trey, Suit: DIAMOND}),
			c5: ComputeCardScore(Card{Rank: Deuce, Suit: DIAMOND}),
		}, want: 9},
		{name: "Six-High Straight Flush Other Order", args: args{
			c5: ComputeCardScore(Card{Rank: Six, Suit: DIAMOND}),
			c3: ComputeCardScore(Card{Rank: Five, Suit: DIAMOND}),
			c1: ComputeCardScore(Card{Rank: Four, Suit: DIAMOND}),
			c4: ComputeCardScore(Card{Rank: Trey, Suit: DIAMOND}),
			c2: ComputeCardScore(Card{Rank: Deuce, Suit: DIAMOND}),
		}, want: 9},
		{name: "King-High Flush K J 9 3 2", args: args{
			c1: ComputeCardScore(Card{Rank: King, Suit: DIAMOND}),
			c2: ComputeCardScore(Card{Rank: Jack, Suit: DIAMOND}),
			c3: ComputeCardScore(Card{Rank: Nine, Suit: DIAMOND}),
			c4: ComputeCardScore(Card{Rank: Trey, Suit: DIAMOND}),
			c5: ComputeCardScore(Card{Rank: Deuce, Suit: DIAMOND}),
		}, want: 983},
		//test unique array
		{name: "Q J T 9 8 Queen-High Straight", args: args{
			c1: ComputeCardScore(Card{Rank: Queen, Suit: DIAMOND}),
			c2: ComputeCardScore(Card{Rank: Jack, Suit: SPADE}),
			c3: ComputeCardScore(Card{Rank: Ten, Suit: CLUB}),
			c4: ComputeCardScore(Card{Rank: Nine, Suit: SPADE}),
			c5: ComputeCardScore(Card{Rank: Eight, Suit: HEART}),
		}, want: 1602},
		{name: "A K 9 8 6  Ace-High", args: args{
			c1: ComputeCardScore(Card{Rank: Ace, Suit: DIAMOND}),
			c2: ComputeCardScore(Card{Rank: King, Suit: SPADE}),
			c3: ComputeCardScore(Card{Rank: Six, Suit: CLUB}),
			c4: ComputeCardScore(Card{Rank: Nine, Suit: SPADE}),
			c5: ComputeCardScore(Card{Rank: Eight, Suit: HEART}),
		}, want: 6295},

		//others
		{name: "K K K A Q  Three Kings", args: args{
			c1: ComputeCardScore(Card{Rank: King, Suit: DIAMOND}),
			c2: ComputeCardScore(Card{Rank: Ace, Suit: SPADE}),
			c3: ComputeCardScore(Card{Rank: King, Suit: CLUB}),
			c4: ComputeCardScore(Card{Rank: Queen, Suit: SPADE}),
			c5: ComputeCardScore(Card{Rank: King, Suit: HEART}),
		}, want: 1676},
		{name: "Q Q T 8 7   Pair of Queens", args: args{
			c1: ComputeCardScore(Card{Rank: Ten, Suit: DIAMOND}),
			c2: ComputeCardScore(Card{Rank: Queen, Suit: SPADE}),
			c3: ComputeCardScore(Card{Rank: Eight, Suit: CLUB}),
			c4: ComputeCardScore(Card{Rank: Queen, Suit: SPADE}),
			c5: ComputeCardScore(Card{Rank: Seven, Suit: HEART}),
		}, want: 3909},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeHandScore(tt.args.c1, tt.args.c2, tt.args.c3, tt.args.c4, tt.args.c5); got != tt.want {
				t.Errorf("ComputeHandScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
