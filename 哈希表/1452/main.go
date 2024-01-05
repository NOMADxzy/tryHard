package main

import "fmt"

func peopleIndexes(favoriteCompanies [][]string) []int {
	whereMap := map[string][]int{}
	for i, company := range favoriteCompanies {
		for _, s := range company {
			whereMap[s] = append(whereMap[s], i)
		}
	}
	var ans []int

	tmp := make([]int, len(favoriteCompanies))
loop:
	for i := 0; i < len(favoriteCompanies); i++ {
		clear(tmp)
		for _, s := range favoriteCompanies[i] {
			places := whereMap[s]
			for _, place := range places {
				tmp[place]++
			}
		}
		for k := 0; k < len(tmp); k++ {
			if tmp[k] == len(favoriteCompanies[i]) && k != i {
				continue loop
			}
		}
		ans = append(ans, i)
	}
	return ans
}

func main() {
	fmt.Println(peopleIndexes([][]string{{"leetcode", "google", "facebook"}, {"google", "microsoft"}, {"google", "facebook"}, {"google"}, {"amazon"}}))
}
