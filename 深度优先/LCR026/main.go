package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	nodes := []*ListNode{}
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	i, j := 0, len(nodes)-1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i == j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil
}

//func reorderList1(head *ListNode) {
//	var changeNext func(head *ListNode, last bool) *ListNode
//	h1 := &ListNode{0,nil}
//	h2 := &ListNode{0,nil}
//	isFirst := true
//	for p := head; p != nil ; p=p.Next {
//		if isFirst{
//			h1.Next =
//		}
//	}
//}
//
//func main() {
//	fmt.Println(reorderList())
//}
