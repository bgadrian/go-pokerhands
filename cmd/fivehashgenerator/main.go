package main

import (
	"fmt"
	"github.com/bgadrian/go-pokerhands/five"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func destPath() string {
	//ex, err := os.Executable()
	//check(err)
	//return fmt.Sprintf("%s..%s%s",filepath.Dir(ex), os.PathSeparator, "generated")
	//
	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//check(err)
	//return fmt.Sprintf("%s%s%s%s", dir, string(os.PathSeparator), "generated", string(os.PathSeparator))

	dir, err := os.Getwd()
	check(err)
	return fmt.Sprintf("%s%s%s%s", dir, string(os.PathSeparator), "generated", string(os.PathSeparator))
}

func main() {
	fPath := destPath() + "five_scores_map.go"
	f, err := os.Create(fPath)
	check(err)
	log.Printf("writing in file: %s\n", fPath)

	defer func() {
		err = f.Close()
		check(err)
	}()
	_, err = f.WriteString(`package generated

// FiveRankMap contains all 300M+ poker hands as strings and their ranking according to https://suffe.cool/poker/7462.html
var FiveRankMap = map[uint32]int{
`)
	check(err)

	//generate all possible hands
	var generatedCount int
	deck := five.GenerateDeck()
	for a := 0; a < 52; a++ {
		c1 := deck[a]
		for b := 0; b < 52; b++ {
			if b == a {
				continue //cards are unique in a deck
			}
			c2 := deck[b]
			for c := 0; c < 52; c++ {
				if c == a || c == b {
					continue //cards are unique in a deck
				}
				c3 := deck[c]
				for d := 0; d < 52; d++ {
					if d == a || d == b || d == c {
						continue //cards are unique in a deck
					}
					c4 := deck[d]
					for e := 0; e < 52; e++ {
						if e == a || e == b || e == c || e == d {
							continue //cards are unique in a deck
						}
						c5 := deck[e]

						key := five.FiveHandHashFromStructs(five.FiveHand{c1, c2, c3, c4, c5})
						score := five.ComputeHandScore(
							five.ComputeCardScore(c1),
							five.ComputeCardScore(c2),
							five.ComputeCardScore(c3),
							five.ComputeCardScore(c4),
							five.ComputeCardScore(c5))
						_, err = f.WriteString(fmt.Sprintf("%d:%d,\n", key, score))
						check(err)
						generatedCount++
						if generatedCount%10000000 == 0 {
							log.Printf("... generating ... %d/300M \n", generatedCount)
						}
					}
				}
			}
		}
	}

	_, err = f.WriteString(`
}
`)
	check(err)
	err = f.Sync()
	check(err)
	log.Println("finished with success")
}
