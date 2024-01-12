package main

import (
	"container/heap"
	"fmt"
)

type Node struct {
	x        int
	distance float64
}

// go 语言中，不提供有限队列数据结构，需要结合heap自己实现
type hp []Node

//var probVals []float64

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].distance > h[j].distance }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(Node)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	graph := make([][]Node, n)

	for i, ev := range edges {
		graph[ev[0]] = append(graph[ev[0]], Node{x: ev[1], distance: succProb[i]})
		graph[ev[1]] = append(graph[ev[1]], Node{x: ev[0], distance: succProb[i]})
	}

	hp := &hp{}
	heap.Push(hp, Node{start_node, 1})
	heap.Init(hp)

	probVals := make([]float64, n)
	probVals[start_node] = 1

	for i := 1; i < n && hp.Len() > 0; i++ {
		best_node := heap.Pop(hp).(Node).x
		if probVals[best_node] == 0 || best_node == end_node {
			break
		}
		for _, node := range graph[best_node] {
			x, v := node.x, node.distance
			newV := probVals[best_node] * v
			if newV > probVals[x] {
				probVals[x] = newV
				heap.Push(hp, Node{x, newV})
			}
		}
	}
	return probVals[end_node]
}

func main() {
	edges := [][]int{{0, 3}, {1, 7}, {1, 2}, {0, 9}}
	fmt.Println(maxProbability(10, edges, []float64{0.31, 0.9, 0.86, 0.36}, 2, 3))
}

// go 语言中，不提供有限队列数据结构，需要结合heap自己实现
