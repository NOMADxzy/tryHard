package main

import (
	"fmt"
	"strconv"
)

func choose(price []int, special [][]int, needs []int, having []int, cache map[string]int) int {
	keyString := ""
	for i := 0; i < len(having); i++ {
		keyString += strconv.Itoa(having[i])
	}
	if cache[keyString] > 0 {
		return cache[keyString]
	}
	min := 0
	last := true
	for i := 0; i < len(special); i++ {
		var j int
		new_having := make([]int, len(having))
		for j = 0; j < len(price); j++ {
			if having[j]+special[i][j] > needs[j] {
				break
			} else {
				new_having[j] = having[j] + special[i][j]
			}
		}
		if j == len(price) {
			last = false
			p := special[i][len(price)] + choose(price, special, needs, new_having, cache)
			if min == 0 || p < min {
				min = p
			}
		}
	}
	if last {
		for i := 0; i < len(having); i++ {
			min += price[i] * (needs[i] - having[i])
		}
	}
	cache[keyString] = min
	return min
}

func shoppingOffers(price []int, special [][]int, needs []int) int {

	for i := 0; i < len(special); i++ {
		origin := 0
		for j := 0; j < len(price); j++ {
			origin += price[j] * special[i][j]
		}
		if special[i][len(price)] > origin {
			special[i][len(price)] = origin
		}
	}
	having := make([]int, len(needs))
	cache := map[string]int{}

	return choose(price, special, needs, having, cache)
}

func main() {
	price := []int{1, 1, 1}
	special := [][]int{{1, 1, 0, 0}, {2, 2, 1, 9}}
	needs := []int{1, 1, 0}
	fmt.Println(shoppingOffers(price, special, needs))
}
