package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inOrder(root *TreeNode, res []int, n int) int {
	if root == nil {
		return n
	} else {
		n = inOrder(root.Left, res, n)
		//res = append(res, root.Val)
		res[n] = root.Val
		n += 1
		n = inOrder(root.Right, res, n)
		return n
	}

}

func check(root *TreeNode) (int, int) {
	res := make([]int, 1000, 1000)
	n := inOrder(root, res, 0)

	val1, val2 := 0, 0
	pos := -1
	ErrTime := 0
	for i := 0; i < n-1; i++ {
		if ErrTime == 1 {
			if res[i] >= res[i+1] {
				val2 = res[i+1]
				ErrTime += 1
				break
			}
		}
		if res[i] >= res[i+1] {
			val1 = res[i]
			pos = i
			ErrTime += 1
		}
	}
	if ErrTime == 1 {
		val2 = res[pos+1]
	}
	return val1, val2
}

func change(root *TreeNode, val1 int, val2 int) {
	if root == nil {
		return
	} else {
		change(root.Left, val1, val2)
		if root.Val == val1 {
			root.Val = val2
		} else if root.Val == val2 {
			root.Val = val1
		}
		change(root.Right, val1, val2)
	}

}

func recoverTree(root *TreeNode) {
	val1, val2 := check(root)
	change(root, val1, val2)
}

func main() {
	tree := &TreeNode{1, &TreeNode{3, nil, &TreeNode{2, nil, nil}}, nil}
	recoverTree(tree)
	print(tree.Val)
}
