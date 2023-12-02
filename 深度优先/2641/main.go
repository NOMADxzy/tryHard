package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goNextLayer(root *TreeNode, curLayer int, layerSum, childSum []int, parents map[*TreeNode]*TreeNode) {
	if len(layerSum) <= curLayer {
		layerSum = append(layerSum, 0)
	}
	layerSum[curLayer] += root.Val
	if root.Left != nil {
		parents[root.Left] = root
		goNextLayer(root.Left, curLayer+1, layerSum, childSum, parents)
	}
	if root.Right != nil {
		parents[root.Right] = root
		goNextLayer(root.Right, curLayer+1, layerSum, childSum, parents)
	}
}

func getSumNext(root *TreeNode, curLayer int, layerSum []int, parents map[*TreeNode]*TreeNode) {
	if root == nil {
		return
	}
	sum := layerSum[curLayer] - root.Val
	if parents[root].Left != nil && parents[root].Right != nil {
		if parents[root].Left.Val == root.Val {

		}
	}
}

func replaceValueInTree(root *TreeNode) *TreeNode {

	root.Val = 0

	q := []*TreeNode{root}

	for len(q) > 0 {

		tmp := q

		q = nil

		nextLevelSum := 0 // 下一层的节点值之和

		for _, node := range tmp {

			if node.Left != nil {

				q = append(q, node.Left)

				nextLevelSum += node.Left.Val

			}

			if node.Right != nil {

				q = append(q, node.Right)

				nextLevelSum += node.Right.Val

			}

		}

		// 再次遍历，更新下一层的节点值

		for _, node := range tmp {

			childrenSum := 0 // node 左右儿子的节点值之和

			if node.Left != nil {

				childrenSum += node.Left.Val

			}

			if node.Right != nil {

				childrenSum += node.Right.Val

			}

			// 更新 node 左右儿子的节点值

			if node.Left != nil {

				node.Left.Val = nextLevelSum - childrenSum

			}

			if node.Right != nil {

				node.Right.Val = nextLevelSum - childrenSum

			}

		}

	}

	return root

}
