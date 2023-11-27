package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	var queue []*Node
	var ans [][]int

	queue = append(queue, root)
	for len(queue) > 0 {
		var layer []int
		var next_q []*Node
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			layer = append(layer, node.Val)
			for _, child := range node.Children {
				next_q = append(next_q, child)
			}
		}
		ans = append(ans, layer)
		queue = next_q
	}
	return ans
}

//
//func main() {
//	fmt.Println(levelOrder("1"))
//}
