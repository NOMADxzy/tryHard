package main

import (
	"container/heap"
	"fmt"
)

type BigHeap [][]int

func (e *BigHeap) Push(x any) {
	*e = append(*e, x.([]int))
}

func (e *BigHeap) Pop() any {
	res := (*e)[len(*e)-1]
	*e = (*e)[:len(*e)-1]
	return res
}

func (e BigHeap) Len() int { return len(e) }
func (e BigHeap) Less(i, j int) bool { // 2/5 2/3
	di0, di1 := (e[i][0]+1)*e[i][1]-(e[i][1]+1)*e[i][0], e[i][1]*(e[i][1]+1) // 3 30
	dj0, dj1 := (e[j][0]+1)*e[j][1]-(e[j][1]+1)*e[j][0], e[j][1]*(e[j][1]+1) // 1 12
	return di0*dj1 > dj0*di1
	//return e[i][0] < e[j][0]
}
func (e BigHeap) Swap(i, j int) { e[i], e[j] = e[j], e[i] }

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	hp := &BigHeap{}
	ans := float64(0)
	for _, class := range classes {
		hp.Push(class)
	}
	heap.Init(hp)
	for extraStudents > 0 {
		class := heap.Pop(hp).([]int)
		class[0]++
		class[1]++
		heap.Push(hp, class)
		extraStudents--
	}
	for len(*hp) > 0 {
		class := hp.Pop().([]int)
		ans += float64(class[0]) / float64(class[1])
	}
	return ans / float64(len(classes))
}

func main() {
	fmt.Println(maxAverageRatio([][]int{{2, 2}, {3, 5}, {1, 2}}, 2))
}
