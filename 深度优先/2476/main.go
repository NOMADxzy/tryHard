package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func midOrder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	midOrder(root.Left, arr)
	*arr = append(*arr, root.Val)
	midOrder(root.Right, arr)
}

func closestNodes(root *TreeNode, queries []int) [][]int {
	var arr []int
	midOrder(root, &arr)
	m := len(queries)
	ans := make([][]int, len(queries))
	idxs := make([]int, m)
	for i := 0; i < len(queries); i++ {
		idxs[i] = i
	}
	sort.Slice(idxs, func(i, j int) bool {
		return queries[idxs[i]] < queries[idxs[j]]
	})
	left, right := 0, len(arr)-1
	for i := 0; i < m; i++ {
		idx := idxs[i]
		target := queries[idx]
		if arr[left] > target {
			ans[idx] = []int{-1, arr[left]}
		} else if arr[right] < target {
			ans[idx] = []int{arr[right], -1}
		} else if arr[right] == target {
			ans[idx] = []int{arr[right], arr[right]}
		} else {
			for left < right {
				mid := (left + right) / 2
				if arr[mid] <= target {
					left = mid + 1
				} else if arr[mid] > target {
					right = mid
				}
			}
			left--
			if arr[left] == target {
				ans[idx] = []int{arr[left], arr[left]}
			} else if arr[left] < target {
				ans[idx] = []int{arr[left], arr[right]}
			}
		}
		right = len(arr) - 1
	}
	return ans
}

func main() {
	root := &TreeNode{6,
		&TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{4, nil, nil}},
		&TreeNode{13,
			&TreeNode{9, nil, nil},
			&TreeNode{15, &TreeNode{4, nil, nil}, nil}}}
	fmt.Println(closestNodes(root, []int{2, 5, 16}))
}
