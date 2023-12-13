package main

import "fmt"

type MyTreeNode struct {
	Val      int
	Children []*MyTreeNode
	isFile   bool
}

type Pair struct {
	node  *MyTreeNode
	layer int
}

func getNextLayer(input string, i int) int {
	i++
	layer := 0
	for i < len(input) && input[i] == '\t' {
		layer++
		i++
	}
	return layer
}

func findDeepest(root *MyTreeNode, acc int, best *int) {
	if root.Children == nil {
		if root.isFile && acc+root.Val > *best {
			*best = acc + root.Val
		}
		return
	}
	for _, child := range root.Children {
		findDeepest(child, acc+root.Val+1, best)
	}
}

func lengthLongestPath(input string) int {
	var stack []*Pair
	var l, r int
	stack = append(stack, &Pair{&MyTreeNode{-1, nil, false}, -1})
	var nextLayer int
	for r < len(input) {
		if r == 0 {
			nextLayer = 0
		} else {
			nextLayer = getNextLayer(input, l)
			r = l + 1 + nextLayer
			l = r
		}

		isFile := false
		for r < len(input) && input[r] != '\n' {
			if input[r] == '.' {
				isFile = true
			}
			r++
		}
		node := &MyTreeNode{r - l, nil, isFile}
		l = r
		if nextLayer > stack[len(stack)-1].layer+1 {
			fmt.Println("error")
		}
		if nextLayer < stack[len(stack)-1].layer {
			for stack[len(stack)-1].layer > nextLayer {
				curLayer := stack[len(stack)-1].layer
				var i int
				for i = len(stack) - 1; i >= 0 && stack[i].layer == curLayer; i-- {
				}
				for len(stack)-1 > i {
					stack[i].node.Children = append(stack[i].node.Children, stack[len(stack)-1].node)
					stack = stack[:len(stack)-1]
				}
			}
		}
		stack = append(stack, &Pair{node: node, layer: nextLayer})
	}
	for stack[len(stack)-1].layer > -1 {
		curLayer := stack[len(stack)-1].layer
		var i int
		for i = len(stack) - 1; i >= 0 && stack[i].layer == curLayer; i-- {
		}
		for len(stack)-1 > i {
			stack[i].node.Children = append(stack[i].node.Children, stack[len(stack)-1].node)
			stack = stack[:len(stack)-1]
		}
	}
	root := stack[0].node
	ans := 0
	findDeepest(root, 0, &ans)
	return ans
}

func main() {
	fmt.Println(lengthLongestPath("file1.txt\nfile2.txt\nlongfile.txt"))
}
