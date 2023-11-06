package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Transaction struct {
	name   string
	time   int
	amount int
	city   string
	string string
}

func invalidTransactions(transactions []string) []string {
	m := len(transactions)
	trans := make([]*Transaction, m)
	for i, transaction := range transactions {
		elems := strings.Split(transaction, ",")
		time, _ := strconv.Atoi(elems[1])
		amount, _ := strconv.Atoi(elems[2])
		trans[i] = &Transaction{elems[0], time, amount, elems[3], transaction}
	}
	//startMap := map[string]int{}
	var ans []string
	invalidMap := map[int]bool{}
	for i := 0; i < m; i++ {
		if trans[i].amount > 1000 {
			invalidMap[i] = true
		}
		for j := i + 1; j < m; j++ {
			if trans[i].name == trans[j].name && trans[i].city != trans[j].city &&
				((trans[i].time-trans[j].time <= 60 && trans[i].time-trans[j].time >= 0) || (trans[j].time-trans[i].time <= 60 && trans[j].time-trans[i].time >= 0)) {
				invalidMap[i] = true
				invalidMap[j] = true
			}
		}

	}
	for i := 0; i < m; i++ {
		if invalidMap[i] {
			ans = append(ans, trans[i].string)
		}
	}

	return ans
}

func main() {
	Transactions := []string{"alice,20,800,mtv", "alice,50,100,mtv", "alice,51,100,frankfurt", "alice,80,100,frankfurt",
		"alice,90,100,beijing", "bob,51,100,frankfurt"}
	fmt.Println(invalidTransactions(Transactions))
}
