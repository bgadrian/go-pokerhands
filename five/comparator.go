package five

import (
	"math"
	"strings"
)

type FiveHand []Card

// built for our CLI/debug
func (h FiveHand) String() string {
	b := strings.Builder{}
	b.Grow(10 + 4)
	for i := range h {
		b.WriteString(h[i].UI())
		b.WriteRune(' ')
	}
	return b.String()
}

// TODO add error handling?
func ParseFiveHand(s string) FiveHand {
	var res []Card
	for i := 0; i < len(s); i += 2 {
		c, _ := ParseCard(s[i : i+2])
		res = append(res, c)
	}
	return res
}

type FiveHandScored struct {
	Winner bool
	Hand   FiveHand
	Rank   CactusHandRank
}

// ComputeWinnerHands returns the winner hands indeces from the input slice
// contains more  elements if its a tie
func ComputeWinnerHands(input []FiveHand) []FiveHandScored {
	var result []FiveHandScored
	winnerRank := CactusHandRank(math.MaxInt)

	for index := range input {
		_ = input[index][4] //help the compiler, avoid panic checks for each element access

		rank := ComputeHandScore(
			ComputeCardScore(input[index][0]),
			ComputeCardScore(input[index][1]),
			ComputeCardScore(input[index][2]),
			ComputeCardScore(input[index][3]),
			ComputeCardScore(input[index][4])) //TODO avoid panic here if is not valid?
		//smaller is better
		// hands from 1 (strongest) to 7,462 (weakest) so a hand with value 2302 beats a hand with value 3402.
		if rank < winnerRank {
			winnerRank = rank
		}
		result = append(result, FiveHandScored{
			Hand: input[index],
			Rank: rank,
		})
	}
	for i := range result {
		result[i].Winner = result[i].Rank == winnerRank
	}
	return result
}
