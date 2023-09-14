package main

import (
	"fmt"
	"sort"
)

func find(potions []int, spell int64, success int64) int {
	if int64(potions[0])*spell >= success {
		return len(potions)
	} else if int64(potions[len(potions)-1])*spell < success {
		return 0
	} else {
		l, r := 0, len(potions)-1
		for l < r {
			if l+1 == r {
				break
			} else {
				mid := (l + r) / 2
				if int64(potions[mid])*spell < success {
					l = mid
				} else {
					r = mid
				}
			}
		}
		return len(potions) - 1 - l
	}
}

func successfulPairs(spells []int, potions []int, success int64) []int {
	validPotionNums := make([]int, len(spells))

	sort.Slice(potions, func(i, j int) bool {
		return potions[i] < potions[j]
	})
	for i := 0; i < len(spells); i++ {
		validPotionNums[i] = find(potions, int64(spells[i]), success)
	}
	return validPotionNums
}

func main() {
	spells := []int{5, 1, 3}
	potions := []int{5, 4, 1, 2, 3}
	fmt.Println(successfulPairs(spells, potions, 1))
}
