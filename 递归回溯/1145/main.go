package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func canNotSureWin(player, nextPlayer []*TreeNode, neighbors map[*TreeNode][]*TreeNode, keyString1, keyString2 []byte, hist map[string]int, preOver bool) bool {
	keyString := string(keyString1) + string(keyString2)
	if val := hist[keyString]; val != 0 {
		return val > 0
	}
	//var valids []*TreeNode
	ans := true
	acted := false
loop:
	for _, node := range player {
		for _, nNode := range neighbors[node] {
			if keyString1[nNode.Val-1] == '0' && keyString2[nNode.Val-1] == '0' {
				acted = true
				keyString1[nNode.Val-1] = '1'
				if canNotSureWin(nextPlayer, append(player, nNode), neighbors, keyString2, keyString1, hist, false) {
					ans = false
					keyString1[nNode.Val-1] = '0'
					break loop
				}
				keyString1[nNode.Val-1] = '0'
			}
		}
	}
	if !acted {
		if preOver {
			return len(player) <= len(nextPlayer)
		}
		if canNotSureWin(nextPlayer, player, neighbors, keyString2, keyString1, hist, true) {
			ans = false
		}
	}
	if ans {
		hist[keyString] = 1
	} else {
		hist[keyString] = -1
	}
	return ans
}

func record(root *TreeNode, neighbors map[*TreeNode][]*TreeNode, fa *TreeNode) {
	if fa != nil {
		neighbors[root] = append(neighbors[root], fa)
	}
	if root.Left != nil {
		neighbors[root] = append(neighbors[root], root.Left)
		record(root.Left, neighbors, root)
	}
	if root.Right != nil {
		neighbors[root] = append(neighbors[root], root.Right)
		record(root.Right, neighbors, root)
	}
}

func btreeGameWinningMove1(root *TreeNode, n int, x int) bool {
	neighbors := map[*TreeNode][]*TreeNode{}

	record(root, neighbors, nil)

	var player, nextPlayer []*TreeNode
	var nodes []*TreeNode
	var node1 *TreeNode
	find(root, x, &nodes, &node1)
	player = append(player, node1)

	keyString1 := make([]byte, n)
	keyString2 := make([]byte, n)
	for i := 0; i < n; i++ {
		keyString1[i] = '0'
		keyString2[i] = '0'
	}
	keyString1[node1.Val-1] = '1'
	hist := map[string]int{}

	for _, node := range nodes {
		if keyString1[node.Val-1] == '0' {
			keyString2[node.Val-1] = '1'
			if canNotSureWin(player, append(nextPlayer, node), neighbors, keyString1, keyString2, hist, false) {
				return true
			}
			keyString2[node.Val-1] = '0'
		}
	}
	return false
}

func find(root *TreeNode, val int, nodes *[]*TreeNode, target **TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	*nodes = append(*nodes, root)
	if root.Val == val {
		*target = root
	}
	r1 := find(root.Left, val, nodes, target)
	r2 := find(root.Right, val, nodes, target)
	if r1 != nil {
		return r1
	} else {
		return r2
	}
}

func cntNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return cntNodes(root.Left) + cntNodes(root.Right) + 1
}

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	//neighbors := map[*TreeNode][]*TreeNode{}
	//record(root, neighbors, nil)

	var nodes []*TreeNode
	var node1 *TreeNode
	find(root, x, &nodes, &node1)

	cntL, cntR, cntF := 0, 0, 0

	cntL = cntNodes(node1.Left)
	cntR = cntNodes(node1.Right)
	cntF = n - cntL - cntR - 1

	maxCnt := 0
	if cntL > maxCnt {
		maxCnt = cntL
	}
	if cntR > maxCnt {
		maxCnt = cntR
	}
	if cntF > maxCnt {
		maxCnt = cntF
	}
	return maxCnt > n/2
}

func main() {
	root := &TreeNode{1,
		&TreeNode{2,
			&TreeNode{4, &TreeNode{8, nil, nil}, &TreeNode{9, nil, nil}},
			&TreeNode{5, &TreeNode{10, nil, nil}, &TreeNode{11, nil, nil}}},
		&TreeNode{3, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(btreeGameWinningMove(root, 11, 3))
}
