package five

import (
	"sort"
)

var (
	/*RankPrimes Central to the Cactus Kev algorithm is the idea of associating a prime number with each card rank:
	The beauty of this system is that if you multiply the prime values of the rank of each card in your hand, you get a unique product, regardless of the order of the five cards. In my above example, the King High Straight hand will always generate a product value of 14,535,931. Since multiplication is one of the fastest calculations a computer can make, we have shaved hundreds of milliseconds off our time had we been forced to sort each hand before evaluation.
	*/
	RankPrimes = map[CardRank]int{
		Deuce: 2,
		Trey:  3,
		Four:  5,
		Five:  7,
		Six:   11,
		Seven: 13,
		Eight: 17,
		Nine:  19,
		Ten:   23,
		Jack:  29,
		Queen: 31,
		King:  37,
		Ace:   41,
	}
)

/*
CactusCardScore is made up of four bytes.  The high-order
///   bytes are used to hold the rank bit pattern, whereas
///   the low-order bytes hold the suit/rank/prime value
///   of the card.
///
///   +--------+--------+--------+--------+
///   |xxxbbbbb|bbbbbbbb|cdhsrrrr|xxpppppp|
///   +--------+--------+--------+--------+
///
///   p = RankPrimes = prime number of rank (deuce=2,trey=3,four=5,five=7,...,ace=41)
///   r = CardRank = rank of card (deuce=0,trey=1,four=2,five=3,...,ace=12)
///   cdhs = CardSuit = suit of card
///   b = bit turned on depending on rank of card

00001000 00000000 01001011 00100101    King of Diamonds
00000000 00001000 00010011 00000111    Five of Spades
00000010 00000000 10001001 00011101    Jack of Clubs
*/
type CactusCardScore int

const (
	bitsSizeRankPrimes = 8
)

// ComputeCardScore creates a single int containing the cards rank, suit and its prime value
// misbehaves or panics if Rank/Suit are not valid values TODO refactor to implement Open principle
func ComputeCardScore(c Card) CactusCardScore {
	res := RankPrimes[c.Rank]
	res |= int(c.Rank) << bitsSizeRankPrimes //add the rank. Its most significant bits, the others will be rewritten
	res |= int(c.Suit)                       //suit values are already shifted
	res |= 1 << (16 + c.Rank)

	return CactusCardScore(res)
}

/*
	CactusHandRank the Cactus Kev evaluator orders

hands from 1 (strongest) to 7,462 (weakest) so a hand with value 2302 beats a hand with value 3402.
*/
type CactusHandRank int

// ComputeHandScore using Cactus Card Scores, is basically adding each card value
// see https://suffe.cool/poker/evaluator.html
func ComputeHandScore(c1, c2, c3, c4, c5 CactusCardScore) CactusHandRank {
	index := int((c1 | c2 | c3 | c4 | c5) >> 16)

	/* check for Flushes and StraightFlushes
	 */
	//0xF000 is a mask that clears the Primes values leaving only suits
	isFlush := (c1 & c2 & c3 & c4 & c5 & 0xF000) != 0
	if isFlush {
		return CactusHandRank(flushes[index])
	}
	/* check for Straights and HighCard hands
	 */
	scoreUniqueCards := unique5[index]
	if scoreUniqueCards != 0 {
		return CactusHandRank(scoreUniqueCards)
	}

	/* let's do it the hard way */
	primeNumbersMultiplication := int((c1 & 0xFF) * (c2 & 0xFF) * (c3 & 0xFF) * (c4 & 0xFF) * (c5 & 0xFF))
	//var found bool
	//for i := range products {
	//	if products[i] == primeNumbersMultiplication {
	//		found = true
	//		index = i
	//		break
	//	}
	//}
	index = sort.SearchInts(products, primeNumbersMultiplication)
	//index, found = sort.Find(4887, func(i int) int {
	//	if products[i] < primeNumbersMultiplication {
	//		return -1
	//	}
	//	if products[i] == primeNumbersMultiplication {
	//		return 0
	//	}
	//	return 1
	//})
	//if !found {
	//	panic(fmt.Sprintf("not found in products: %d", primeNumbersMultiplication))
	//}
	return CactusHandRank(values[index])
}
