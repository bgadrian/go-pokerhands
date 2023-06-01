package main

import (
	"github.com/bgadrian/go-pokerhands/five"
	"log"
)

func main() {
	g := five.NewHandGenerator()
	handsCount := 8

	var allHands []five.FiveHand
	for i := 0; i < handsCount; i++ {
		allHands = append(allHands, g.FiveHand())
	}

	results := five.ComputeWinnerHands(allHands)
	labels := map[bool]string{
		true:  "WON",
		false: "lost",
	}
	for i := range results {
		category := five.ComputeCategoryForHandRank(results[i].Rank)
		log.Printf("%s\t%s\t%s\n",
			results[i].Hand,
			five.HandRankCategoriesLabels[category],
			labels[results[i].Winner],
		)
	}
}
