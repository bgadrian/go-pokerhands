package five

import (
	"errors"
)

var ErrInvalid = errors.New("invalid value")

// HandRankCategory see https://en.wikipedia.org/wiki/List_of_poker_hands
type HandRankCategory int

const (
	UNKNOWN HandRankCategory = iota
	STRAIGHT_FLUSH
	FOUR_OF_A_KIND
	FULL_HOUSE
	FLUSH
	STRAIGHT
	THREE_OF_A_KIND
	TWO_PAIR
	ONE_PAIR
	HIGH_CARD
)

type Card struct {
	Suit CardSuit
	Rank CardRank
}

// String implements Stringer interface
func (c Card) String() string {
	rl, _ := CardRankLabel(c.Rank)
	sl, _ := CardSuitLabel(c.Suit)
	return string([]rune{rl, sl})
}

func ParseCard(in string) (Card, error) {
	//TODO sanitize, check for UTF8?
	if len(in) != 2 {
		return Card{}, ErrInvalid
	}
	rank, err := CardRankValue(rune(in[0]))
	if err != nil {
		return Card{}, ErrInvalid
	}
	suit, err := CardSuitValue(rune(in[1]))
	if err != nil {
		return Card{}, ErrInvalid
	}
	return Card{Rank: rank, Suit: suit}, nil
}

// CardSuit internal value for a Card CardSuit
// its values are obtained by shifting by 1 bit, to be used in the mask
type CardSuit int

const (
	CLUB    CardSuit = 0x8000
	DIAMOND CardSuit = 0x4000
	HEART   CardSuit = 0x2000
	SPADE   CardSuit = 0x1000
)

var (
	//cardSuitLabels maps from CardRank to a human friendly value
	cardSuitLabels = map[CardSuit]rune{
		CLUB:    'c',
		DIAMOND: 'd',
		HEART:   'h',
		SPADE:   's',
	}
	//cardRankFromLabel maps from its label to CardRank value
	cardSuitFromLabel = map[rune]CardSuit{
		cardSuitLabels[CLUB]:    CLUB,
		cardSuitLabels[DIAMOND]: DIAMOND,
		cardSuitLabels[HEART]:   HEART,
		cardSuitLabels[SPADE]:   SPADE,
	}
)

// CardSuitLabel returns ErrInvalid or the ranks friendly label
func CardSuitLabel(s CardSuit) (rune, error) {
	suit, exists := cardSuitLabels[s]
	if !exists {
		return 0, ErrInvalid
	}
	return suit, nil
}

// CardSuitValue returns ErrInvalid or the ranks internal value
func CardSuitValue(r rune) (CardSuit, error) {
	suit, exists := cardSuitFromLabel[r]
	if !exists {
		return 0, ErrInvalid
	}
	return suit, nil
}

// CardRank internal numeric value of a card rank
type CardRank int

const (
	Deuce CardRank = iota
	Trey
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten //8
	Jack
	Queen //10
	King
	Ace
)

var (
	//cardRankLabels maps from CardRank to a human friendly value
	cardRankLabels = []rune{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}
	//cardRankFromLabel maps from its label to CardRank value
	cardRankFromLabel = map[rune]CardRank{
		cardRankLabels[Deuce]: Deuce,
		cardRankLabels[Trey]:  Trey,
		cardRankLabels[Four]:  Four,
		cardRankLabels[Five]:  Five,
		cardRankLabels[Six]:   Six,
		cardRankLabels[Seven]: Seven,
		cardRankLabels[Eight]: Eight,
		cardRankLabels[Nine]:  Nine,
		cardRankLabels[Ten]:   Ten,
		cardRankLabels[Jack]:  Jack,
		cardRankLabels[Queen]: Queen,
		cardRankLabels[King]:  King,
		cardRankLabels[Ace]:   Ace,
	}
)

// CardRankLabel returns ErrInvalid or the ranks friendly label
func CardRankLabel(r CardRank) (rune, error) {
	if r < 0 || r >= CardRank(len(cardRankLabels)) {
		return 'X', ErrInvalid
	}
	return cardRankLabels[int(r)], nil
}

// CardRankValue returns ErrInvalid or the ranks internal value
func CardRankValue(r rune) (CardRank, error) {
	rank, exists := cardRankFromLabel[r]
	if !exists {
		return 0, ErrInvalid
	}
	return rank, nil
}
