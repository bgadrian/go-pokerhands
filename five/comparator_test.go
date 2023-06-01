package five

import (
	"reflect"
	"testing"
)

func TestComputeWinnerHands(t *testing.T) {
	tests := []struct {
		name string
		args []FiveHand
		want []int
	}{
		{name: "trey beats pair", args: []FiveHand{
			ParseFiveHand("2c3cQdQsQh"),
			ParseFiveHand("2c3cJdKs2h"),
		}, want: []int{0}},
		{name: "four kind beats flush", args: []FiveHand{
			ParseFiveHand("2c3cJcKc8c"),
			ParseFiveHand("2cQcQdQsQh"),
		}, want: []int{1}},
		{name: "four kind beats four kind", args: []FiveHand{
			ParseFiveHand("AcJsJhJdJc"),
			ParseFiveHand("2cQcQdQsQh"),
		}, want: []int{1}},
		{name: "full house beats straight and pairs", args: []FiveHand{
			ParseFiveHand("8c9cTdJsQh"),
			ParseFiveHand("2c3cQdQsQh"),
			ParseFiveHand("2c2h2dKcKh"), //<<<<
			ParseFiveHand("2c3cJdKs2h"),
		}, want: []int{2}},

		{name: "pairs beats high card", args: []FiveHand{
			ParseFiveHand("2c2hQd7s8h"), //<<<<
			ParseFiveHand("Ac3cJdKs2h"),
			ParseFiveHand("2d2sQd7s8h"), //<<<<
		}, want: []int{0, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeWinnerHands(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ComputeWinnerHands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseFiveHand(t *testing.T) {
	tests := []struct {
		args string
		want FiveHand
	}{
		{"2dKcTsQhJh", FiveHand{
			Card{DIAMOND, Deuce},
			Card{CLUB, King},
			Card{SPADE, Ten},
			Card{HEART, Queen},
			Card{HEART, Jack},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.args, func(t *testing.T) {
			if got := ParseFiveHand(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseFiveHand() = %v, want %v", got, tt.want)
			}
		})
	}
}
