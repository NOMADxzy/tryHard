package main

import "fmt"

func findSeq(mark int, idxses *[][]int, acc []int) {
	if mark == 0 {
		tmp := make([]int, len(acc))
		copy(tmp, acc)
		*idxses = append(*idxses, acc)
		return
	}
	mask := 1
	for i := 0; i < 4; i++ {
		if mark&mask > 0 {
			findSeq(mark-mask, idxses, append(acc, i))
		}
		mask *= 2
	}
}

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	var idxses [][]int
	findSeq(15, &idxses, []int{})
	ps := [][]int{p1, p2, p3, p4}
	getDist := func(a, b []int) int {
		return (a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1])
	}
	for _, idxs := range idxses {
		a, b, c, d := ps[idxs[0]], ps[idxs[1]], ps[idxs[2]], ps[idxs[3]]
		d1, d2, d3, d4 := getDist(a, b), getDist(b, c), getDist(c, d), getDist(d, a)
		d5, d6 := getDist(a, c), getDist(b, d)
		if d1 != 0 && d1 == d2 && d3 == d4 && d2 == d3 && d5 == d6 && d5 != 0 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(validSquare([]int{0, 0}, []int{1, 1}, []int{1, 0}, []int{0, 1}))
}
