package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Sum(root *TreeNode, sMap map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	sMap[root] = root.Val + Sum(root.Left, sMap) + Sum(root.Right, sMap)
	return sMap[root]
}

func getDelta(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func maxProduct(root *TreeNode) int {
	MOD := 1000000007
	sMap := map[*TreeNode]int{}
	Sum(root, sMap)
	total := sMap[root]
	delta := total
	for _, sum := range sMap {
		if d := getDelta(sum, total-sum); d < delta {
			delta = d
		}
	}
	val1 := (total - delta) / 2
	return (val1 % MOD) * ((val1 + delta) % MOD) % MOD
}

func main() {
	root := &TreeNode{10,
		&TreeNode{7, nil, &TreeNode{10, nil, nil}},
		&TreeNode{6, nil, &TreeNode{10, nil, nil}}}
	fmt.Println(maxProduct(root))
}
