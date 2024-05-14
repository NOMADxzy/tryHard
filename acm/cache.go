package main

import (
	"fmt"
	"time"
)

func printFunc1(x int, chan1, chan2 chan interface{}) {
	for x < 100 {
		_ = <-chan2
		fmt.Println(x)
		chan1 <- struct{}{}
		x += 2
	}
}
func printFunc2(x int, chan1, chan2 chan interface{}) {
	for x < 100 {
		_ = <-chan1
		fmt.Println(x)
		chan2 <- struct{}{}
		x += 2
	}
}

func main() {
	chan1 := make(chan interface{}, 1)
	chan2 := make(chan interface{}, 1)
	chan2 <- struct{}{}

	go printFunc1(0, chan1, chan2)
	go printFunc2(1, chan1, chan2)
	time.Sleep(10 * time.Second)
}
