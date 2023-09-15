package main

import "fmt"

type StockSpanner struct {
	stack []int
	cover []int
}

func Constructor() StockSpanner {
	return StockSpanner{[]int{}, []int{}}
}

func (this *StockSpanner) Next(price int) int {
	if len(this.stack) == 0 || price < this.stack[len(this.stack)-1] {
		this.stack = append(this.stack, price)
		this.cover = append(this.cover, 1)
		return 1
	} else {
		cnt := 1
		for len(this.stack) > 0 && this.stack[len(this.stack)-1] <= price {
			this.stack = this.stack[:len(this.stack)-1]
			cnt += this.cover[len(this.cover)-1]
			this.cover = this.cover[:len(this.cover)-1]
		}
		this.stack = append(this.stack, price)
		this.cover = append(this.cover, cnt)
		return cnt
	}
}

func main() {
	ss := Constructor()
	prices := []int{100, 80, 60, 70, 60, 75, 85}

	for i := 0; i < len(prices); i++ {
		fmt.Println(ss.Next(prices[i]))
	}
}
