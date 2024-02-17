package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	var midOrder func(root *TreeNode, list *[]int)
	midOrder = func(root *TreeNode, list *[]int) {
		if root == nil {
			return
		}
		midOrder(root.Left, list)
		*list = append(*list, root.Val)
		midOrder(root.Right, list)
	}
	var list1, list2 []int
	midOrder(root1, &list1)
	midOrder(root2, &list2)
	ans := make([]int, len(list1)+len(list2))
	var i, j, idx int
	for i < len(list1) || j < len(list2) {
		if i < len(list1) && (j == len(list2) || list1[i] < list2[j]) {
			ans[idx] = list1[i]
			i++
		} else {
			ans[idx] = list2[j]
			j++
		}
		idx++
	}
	return ans
}

func main() {
	root1 := &TreeNode{2, &TreeNode{0, nil, nil}, &TreeNode{4, nil, nil}}
	root2 := &TreeNode{3, &TreeNode{1, nil, nil}, &TreeNode{5, nil, nil}}
	fmt.Println(getAllElements(root1, root2))
}
