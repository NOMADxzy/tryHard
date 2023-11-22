package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findNext(root *TreeNode, startV, destV int) ([]byte, []byte) {
	var s1, s2 []byte
	if root.Val == startV {
		s1 = append(s1, 'O')
	} else if root.Val == destV {
		s2 = append(s2, 'O')
	}
	if root.Left != nil {
		ls1, ls2 := findNext(root.Left, startV, destV)
		if len(ls1) > 0 && len(ls2) > 0 {
			return ls1, ls2
		}
		if len(ls1) > 0 {
			s1 = append(ls1, 'U')
		}
		if len(ls2) > 0 {
			s2 = append(ls2, 'L')
		}
	}
	if root.Right != nil {
		rs1, rs2 := findNext(root.Right, startV, destV)
		if len(rs1) > 0 && len(rs2) > 0 {
			return rs1, rs2
		}
		if len(rs1) > 0 {
			s1 = append(rs1, 'U')
		}
		if len(rs2) > 0 {
			s2 = append(rs2, 'R')
		}
	}
	return s1, s2
}

func getDirections(root *TreeNode, startValue int, destValue int) string {
	s1, s2 := findNext(root, startValue, destValue)
	s := make([]byte, len(s1)+len(s2)-2)
	s1 = s1[1:]
	s2 = s2[1:]
	copy(s, s1)
	idx := len(s1)
	for i := len(s2) - 1; i >= 0; i-- {
		s[idx] = s2[i]
		idx++
	}
	return string(s)
}

func main() {
	root := &TreeNode{5, &TreeNode{1, &TreeNode{3, nil, nil}, nil},
		&TreeNode{2, &TreeNode{6, nil, nil}, &TreeNode{4, nil, nil}}}
	fmt.Println(getDirections(root, 3, 6))
}
