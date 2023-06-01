# Poker hands comparator

We have an infinite methods of comparing poker hands, we can rely on complex algorithms, patter recognition, but I like to think at scale. To scale you need a simple logic. The possible number of hands is a bounded list, so I've used the  Cactus 2006 algorithm https://suffe.cool/poker/evaluator.html which ranks all unique possible hands, making it easier to compare them. The original algorithms optimises the runtime and memory by using bit operations and predefined lists of ranks. 

Poker hands combinations (if order does not matter) = 2.598.960 unique hands
To speed the lookup and avoid runtime sorting we can compute all permutations (order matters): 311.875.200 hands


The code is NOT production ready, it does memory allocations (custom structs) in favor of source code usuability. 

Limitations
* "normal" poker rules
* 52 cards deck, no jokers
* no assertions on input (duplicate cards, malformed)

## How to run

Using docker 
```bash
docker build --tag=poker42:latest .

docker run poker42:latest
```

Using local Go
```bash
./run.sh
```

## Runtime optimisiation (WIP / Not finished)

This program is intended to be used in Cloud, so we have something the original author did not had, memory(RAM).
We can use this software as a shared process/microservice
So we go one step forward, we use his unique hand ranking algorithm and generate all hands and store their rank in a huge map, sacrificing memory for latency.

This way the processing of each input hands score will be O(1)


A string of 5 cards, 2 letters each = 10 bytes in ASCII
Value is an integer score = 4 bytes
Used memory for combinations = 36Mb (+extra for internals) = this does not help much as we would need sorting, probably Cactus algorithm is faster than this.

Used memory for permutations = 4.3Gb (+extra for internals)

We reduce memory footprint of the generated map by either
  * use a 32bit key that is a murmur3 hash of the hand (and ensure there are no collisions)  = 1.8Gb
  * use a trie instead of map lookups (size depends on the implementation)

 Generation

To generate the map we rely on go:generate functionality, but because the code is simple we do not use `text/template` instead a simple stdout. Template is also a not compatible with our hand generation which can be more optimal as a stream to file method.


## Roadmap

functional TODO
* reduce generator code complexity, avoid having 3 arrays and binary shifting and use sort to reach https://suffe.cool/poker/7462.html (single array) - we only need all hands rankings in a single list
* create benchmarks reduce memory allocations (from strings to structs)

non-functional TODO
* create a http server instead of CLI https://goswagger.io/
  * for best performance keep a minimal websocket / http2 without encoding like json
  * for best interopability use swagger/json
* create a github builder for server with most common platform binaries as release


### References

https://en.wikipedia.org/wiki/List_of_poker_hands
https://github.com/MarcFletcher/XPokerEval/blob/main/XPokerEval.CactusKev.CSharp.Test/Program.cs
https://www.codingthewheel.com/archives/poker-hand-evaluator-roundup/#cactus_kev_senzee
https://suffe.cool/poker/evaluator.html
https://suffe.cool/poker/7462.html

