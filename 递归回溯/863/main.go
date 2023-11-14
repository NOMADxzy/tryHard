package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findKdist(root *TreeNode, k int, res *[]int) {
	if root == nil {
		return
	}
	if k == 0 {
		*res = append(*res, root.Val)
		return
	}
	findKdist(root.Left, k-1, res)
	findKdist(root.Right, k-1, res)
}

func findTarget(root *TreeNode, target int, faMap map[*TreeNode]*TreeNode) *TreeNode {
	if root.Val == target {
		return root
	}
	if root.Left != nil {
		faMap[root.Left] = root
		if n := findTarget(root.Left, target, faMap); n != nil {
			return n
		}
	}
	if root.Right != nil {
		faMap[root.Right] = root
		if n := findTarget(root.Right, target, faMap); n != nil {
			return n
		}
	}
	return nil
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	faMap := map[*TreeNode]*TreeNode{}
	node := findTarget(root, target.Val, faMap)
	var pre *TreeNode
	var ans []int
	findKdist(node, k, &ans)                       //先考虑node下面的
	for i := 1; i < k && faMap[node] != nil; i++ { //再考虑node祖先的另一子树
		pre = node
		node = faMap[node]
		if node.Left == pre {
			findKdist(node.Right, k-1-i, &ans)
		} else {
			findKdist(node.Left, k-1-i, &ans)
		}
	}
	if k > 0 && faMap[node] != nil { //最后考虑node顶级祖先，不能和1重复
		ans = append(ans, faMap[node].Val)
	}
	return ans
}

func main() {
	//root := &TreeNode{3,
	//	&TreeNode{5,
	//		&TreeNode{6, nil, nil},
	//		&TreeNode{2,
	//			&TreeNode{7, nil, nil},
	//			&TreeNode{4, nil, nil}}},
	//	&TreeNode{1,
	//		&TreeNode{0, nil, nil},
	//		&TreeNode{8, nil, nil}}}
	root := &TreeNode{0, &TreeNode{1,
		&TreeNode{3, nil, nil}, &TreeNode{2, nil, nil}}, nil}
	fmt.Println(distanceK(root, &TreeNode{2, nil, nil}, 1))
}
