# Poker hands comparator

Contains few packages for a 5cards poker hand comparator based on Cactus 2006 algorithm https://suffe.cool/poker/evaluator.html


functional TODO
* reduce generator code complexity, avoid having 3 arrays and binary shifting and use sort to reach https://suffe.cool/poker/7462.html (single array)
* (try) to reduce memory footprint by switching to a trie instead of map lookups

non-functional TODO
* create a .docker
* create a github builder for CLI with most common platform binaries as release
* create a http server instead of CLI