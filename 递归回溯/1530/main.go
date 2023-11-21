package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNext(root *TreeNode, distance int) (int, []int) {
	if root.Left == nil && root.Right == nil {
		dist := make([]int, distance+1)
		dist[0] = 1
		return 0, dist
	}
	cnt := 0
	var distArrL, distArrR []int
	var leftCnt, rightCnt int
	if root.Left != nil {
		leftCnt, distArrL = countNext(root.Left, distance)
	}
	if root.Right != nil {
		rightCnt, distArrR = countNext(root.Right, distance)
	}
	cnt += leftCnt + rightCnt
	if distArrL != nil && distArrR != nil {
		for i := 0; i <= distance-2; i++ {
			for j := 0; i+j <= distance-2; j++ {
				cnt += distArrL[i] * distArrR[j]
			}
		}
	}
	newDistArr := make([]int, distance+1)
	for i := 1; i <= distance; i++ {
		if distArrL != nil {
			newDistArr[i] += distArrL[i-1]
		}
		if distArrR != nil {
			newDistArr[i] += distArrR[i-1]
		}
	}
	return cnt, newDistArr
}

func countPairs(root *TreeNode, distance int) int {
	if distance == 1 {
		return 0
	}
	cnt, _ := countNext(root, distance)
	//fmt.Println(arr)
	return cnt
}

func main() {
	root := &TreeNode{1, &TreeNode{2, nil, &TreeNode{4, nil, nil}}, &TreeNode{3, nil, nil}}
	fmt.Println(countPairs(root, 1))
}
