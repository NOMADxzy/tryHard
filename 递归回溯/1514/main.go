package main

import (
	"container/heap"
	"fmt"
)

// go 语言中，不提供有限队列数据结构，需要结合heap自己实现
type he []int

var prob [][]float64
var start int

func (h he) Len() int              { return len(h) }
func (h he) Less(i, j int) bool    { return prob[start][h[i]] > prob[start][h[j]] }
func (h he) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *he) Push(v interface{})   { *h = append(*h, v.(int)) }
func (h *he) Pop() (v interface{}) { a := *h; *h, v = a[1:], a[0]; return }

func maxProbability1(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	start = start_node
	prob = make([][]float64, n)
	for i := 0; i < n; i++ {
		prob[i] = make([]float64, n)
	}
	for i, edge := range edges {
		prob[edge[0]][edge[1]] = succProb[i]
		prob[edge[1]][edge[0]] = succProb[i]
	}
	S := map[int]bool{}
	S[start_node] = true

	hp := &he{}
	for j := 0; j < n; j++ {
		if !S[j] {
			hp.Push(j)
		}
	}
	heap.Init(hp)

	for i := 1; i < n; i++ {
		best_node := hp.Pop().(int)
		if prob[start_node][best_node] == 0 || best_node == end_node {
			break
		}
		S[best_node] = true
		for j := 0; j < n; j++ {
			if !S[j] && prob[start_node][best_node]*prob[best_node][j] > prob[start_node][j] {
				v := prob[start_node][best_node] * prob[best_node][j]
				prob[start_node][j] = v
				prob[j][start_node] = v
			}
		}
		heap.Init(hp)
	}
	return prob[start_node][end_node]
}

type Node struct {
	x        int
	y        int
	distance float64
}

// go 语言中，不提供有限队列数据结构，需要结合heap自己实现
type hp []Node

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].distance > h[j].distance }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(Node)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	graph := make([][]Node, n)

	for i, ev := range edges {
		graph[ev[0]] = append(graph[ev[0]], Node{x: ev[1], distance: succProb[i]})
		graph[ev[1]] = append(graph[ev[1]], Node{x: ev[0], distance: succProb[i]})
	}

	res := dijkstra(n, start, graph)

	return res[end]
}

func dijkstra(n int, start int, graph [][]Node) []float64 {
	res := make([]float64, n)

	q := &hp{}
	heap.Init(q)
	heap.Push(q, Node{x: start, distance: 1})
	res[start] = 1

	for q.Len() > 0 {
		cur := heap.Pop(q).(Node)

		if cur.distance < res[cur.x] {
			continue
		}

		for _, gv := range graph[cur.x] {
			x := gv.x
			d := gv.distance
			if d*cur.distance > res[x] {
				res[x] = d * cur.distance
				heap.Push(q, Node{x: x, distance: res[x]})
			}
		}
	}

	return res
}

func main() {
	edges := [][]int{{0, 1}, {1, 2}, {0, 2}}
	fmt.Println(maxProbability(3, edges, []float64{0.5, 0.5, 0.2}, 0, 2))
}
