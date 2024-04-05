package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	mark := map[int]bool{}
	visited := map[*Node]*Node{}
	var dfs func(node *Node) *Node

	getNode := func(node *Node) *Node {
		if _, ok := visited[node]; !ok {
			visited[node] = &Node{node.Val, []*Node{}}
		}
		return visited[node]
	}

	dfs = func(node *Node) *Node {
		newNode := getNode(node)
		mark[node.Val] = true
		for _, nextNode := range node.Neighbors {
			newNode.Neighbors = append(newNode.Neighbors, getNode(nextNode))
			if !mark[nextNode.Val] {
				dfs(nextNode)
			}
		}
		return newNode
	}
	ans := dfs(node)
	return ans
}
