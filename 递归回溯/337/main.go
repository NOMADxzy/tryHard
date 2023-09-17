package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 递归超时
func find(root *TreeNode, choosed bool, choose map[*TreeNode]int, nochoose map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if choosed && choose[root] > 0 {
		return choose[root]
	}
	if !choosed && nochoose[root] > 0 {
		return nochoose[root]
	}
	leftNot := find(root.Left, false, choose, nochoose)
	rightNot := find(root.Right, false, choose, nochoose)
	if choosed {
		money := leftNot + rightNot + root.Val
		choose[root] = money
		return money
	} else {
		leftYes := find(root.Left, true, choose, nochoose)
		rightYes := find(root.Right, true, choose, nochoose)
		money := max(leftYes, leftNot) + max(rightYes, rightNot)
		nochoose[root] = money
		return money
	}
}

// 地址作为哈希表健值
func postOrder(node *TreeNode, choose map[*TreeNode]int, nochoose map[*TreeNode]int) {
	if node.Left != nil {
		postOrder(node.Left, choose, nochoose)
	}
	if node.Right != nil {
		postOrder(node.Right, choose, nochoose)
	}
	leftNot := nochoose[node.Left]
	rightNot := nochoose[node.Right]
	leftYes := choose[node.Left]
	rightYes := choose[node.Right]

	choose[node] = node.Val + leftNot + rightNot
	nochoose[node] = max(leftYes, leftNot) + max(rightYes, rightNot)

}

func rob(root *TreeNode) int {
	choose := map[*TreeNode]int{}
	nochoose := map[*TreeNode]int{}
	choose[nil] = 0
	nochoose[nil] = 0
	postOrder(root, choose, nochoose)
	return max(choose[root], nochoose[root])
	//return max(find(root, true, choose, nochoose), find(root, false, choose, nochoose))
}

func main() {
	root := &TreeNode{3,
		&TreeNode{2, nil, &TreeNode{3, nil, nil}},
		&TreeNode{3, nil, &TreeNode{1, nil, nil}}}
	//root := &TreeNode{3,
	//	&TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}},
	//	&TreeNode{5, nil, &TreeNode{1, nil, nil}}}
	fmt.Println(rob(root))
}
