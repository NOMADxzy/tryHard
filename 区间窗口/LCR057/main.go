package main

import "fmt"

// 哈希表
func containsNearbyAlmostDuplicate1(nums []int, k int, t int) bool {
	cnts := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if i >= k+1 {
			cnts[nums[i-k-1]]--
			if cnts[nums[i-k-1]] == 0 {
				delete(cnts, nums[i-k-1])
			}
		}
		for v, _ := range cnts {
			delta := nums[i] - v
			if delta < 0 {
				delta = -delta
			}
			if delta <= t {
				return true
			}
		}
		cnts[nums[i]]++
	}
	return false
}

type AVLNode struct {
	Left, Right *AVLNode // 表示指向左孩子和右孩子
	Val         int      // 结点存储数据
	Height      int      // 记录这个结点此时的高度
}

// RightRotate LL型，顺时针旋转，右旋
func (avlNode *AVLNode) RightRotate() *AVLNode {
	//旋转
	headNode := avlNode.Left
	avlNode.Left = headNode.Right
	headNode.Right = avlNode

	// 更新高度
	// 先更新 avlNode，后更新 headNode，自下而上
	avlNode.Height = Max(avlNode.Left.GetHeight(), avlNode.Right.GetHeight()) + 1
	headNode.Height = Max(headNode.Left.GetHeight(), headNode.Right.GetHeight()) + 1

	return headNode
}

// LeftRotate RR型，逆时针旋转，左旋
func (avlNode *AVLNode) LeftRotate() *AVLNode {
	//旋转
	headNode := avlNode.Right
	avlNode.Right = headNode.Left
	headNode.Left = avlNode

	// 更新高度
	// 先更新 avlNode，后更新 headNode，自下而上
	avlNode.Height = Max(avlNode.Left.GetHeight(), avlNode.Right.GetHeight()) + 1
	headNode.Height = Max(headNode.Left.GetHeight(), headNode.Right.GetHeight()) + 1

	return headNode
}

// LeftThenRightRotate LR型，先逆时针旋转再顺时针旋转，先左旋，在右旋
func (avlNode *AVLNode) LeftThenRightRotate() *AVLNode {
	// 先左孩子进行左旋
	avlNode.Left = avlNode.Left.LeftRotate()
	// 然后自己进行右旋
	return avlNode.RightRotate()
}

// RightThenLeftRotate RL型，先顺时针旋转再逆时针旋转，先右旋，再左旋
func (avlNode *AVLNode) RightThenLeftRotate() *AVLNode {
	// 先右孩子进行右旋
	avlNode.Right = avlNode.Right.RightRotate()
	// 然后自己进行右旋
	return avlNode.LeftRotate()
}

// adjust 调整AVL树的平衡
func (avlNode *AVLNode) adjust() *AVLNode {
	if avlNode.Right.GetHeight()-avlNode.Left.GetHeight() == 2 {
		// 如果右子树的高度比左子树的高度大于2，R型
		if avlNode.Right.Right.GetHeight() > avlNode.Right.Left.GetHeight() {
			// 如果 avlNode.Right 的右子树的高度大于 avlNode.Right 的左子树高度，RR型
			// 那么就直接对 avlNode 进行左旋
			avlNode = avlNode.LeftRotate()
		} else {
			// 否则先对 avlNode.Right进行右旋转，然后对 avlNode 进行左旋，RL型
			avlNode = avlNode.RightThenLeftRotate()
		}
	} else if avlNode.Right.GetHeight()-avlNode.Left.GetHeight() == -2 {
		// 如果左子树的高度比右子树的高度大2
		if avlNode.Left.Left.GetHeight() > avlNode.Left.Right.GetHeight() {
			// 如果 avlNode.Left 的左子树高度大于 avlNode.Left 的右子树高度，LL型
			// 那么就直接对 avlNode 进行右旋
			avlNode = avlNode.RightRotate()
		} else {
			// 否则先对 avlNode.Left 进行左旋，然后对 avlNode 进行右旋，LR型
			avlNode = avlNode.LeftThenRightRotate()
		}
	}

	return avlNode
}

// NewAVLNode 新建一个结点
func NewAVLNode(val int) *AVLNode {
	return &AVLNode{
		Left:   nil,
		Right:  nil,
		Val:    val,
		Height: 1,
	}
}

// Compare 比较两个值
func Compare(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	} else {
		return 0
	}
}

// GetData 获取结点数据
func (avlNode *AVLNode) GetData() int {
	return avlNode.Val
}

// SetData 设置结点数据
func (avlNode *AVLNode) SetData(val int) {
	if avlNode == nil {
		return
	}
	avlNode.Val = val
}

// GetRight 获取结点的右孩子结点
func (avlNode *AVLNode) GetRight() *AVLNode {
	if avlNode == nil {
		return nil
	}
	return avlNode.Right
}

// GetLeft 获取结点的左孩子结点
func (avlNode *AVLNode) GetLeft() *AVLNode {
	if avlNode == nil {
		return nil
	}
	return avlNode.Left
}

// GetHeight 获取结点的高度
func (avlNode *AVLNode) GetHeight() int {
	if avlNode == nil {
		return 0
	}
	return avlNode.Height
}

// Max 比较两个子树高度的大小
func Max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

// Find 查找指定值
func (avlNode *AVLNode) Find(val int) *AVLNode {
	var found *AVLNode
	// 调用比较函数比较两个结点的值的大小
	switch Compare(val, avlNode.Val) {
	case -1:
		found = avlNode.Left.Find(val)
	case 1:
		found = avlNode.Right.Find(val)
	case 0:
		return avlNode
	}

	return found
}

// FindMin 查找最小值
func (avlNode *AVLNode) FindMin() *AVLNode { // 递归写法
	var found *AVLNode

	if avlNode.Left != nil {
		found = avlNode.Left.FindMin()
	} else {
		found = avlNode
	}

	return found
}

// FindMax 查找最大值
func (avlNode *AVLNode) FindMax() *AVLNode {
	var found *AVLNode

	if avlNode.Right != nil {
		found = avlNode.Right.FindMax()
	} else {
		found = avlNode
	}

	return found
}

// FindLowerBound 查找大于等于目标值的最小值
func (avlNode *AVLNode) lowerBound(val int) *AVLNode {
	if avlNode == nil {
		return nil
	}
	var found *AVLNode
	// 调用比较函数比较两个结点的值的大小
	switch Compare(val, avlNode.Val) {
	case -1:
		found = avlNode.Left.lowerBound(val)
		if found != nil {
			return found
		}
	case 1:
		return avlNode.Right.lowerBound(val)
	case 0:
		return avlNode
	}
	return avlNode
}

// Insert 插入数据
// 因为没有定义结点的parent指针，所以插入数据就只能递归的插入，因为我要调整树的平衡和高度
func (avlNode *AVLNode) Insert(val int) *AVLNode {
	if avlNode == nil {
		return NewAVLNode(val)
	}

	switch Compare(val, avlNode.Val) {
	case -1:
		// 如果value 小于 avlNode.Data 那么就在avlNode的左子树上去插入value
		avlNode.Left = avlNode.Left.Insert(val)
		avlNode = avlNode.adjust() // 自动调整平衡
	case 1:
		avlNode.Right = avlNode.Right.Insert(val)
		avlNode = avlNode.adjust()
	case 0:
		fmt.Print("数据已经存在")
	}
	// 修改结点的高度
	avlNode.Height = Max(avlNode.Left.GetHeight(), avlNode.Right.GetHeight()) + 1

	return avlNode
}

// Delete 删除数据
func (avlNode *AVLNode) Delete(val int) *AVLNode {
	// 没有找到匹配的数据
	if avlNode == nil {
		//fmt.Println("不存在", value)
		return nil
	}

	switch Compare(val, avlNode.Val) {
	case 1:
		avlNode.Right = avlNode.Right.Delete(val)
	case -1:
		avlNode.Left = avlNode.Left.Delete(val)
	case 0:
		// 找到数据，删除结点
		if avlNode.Left != nil && avlNode.Right != nil { // 结点有左孩子和右孩子
			avlNode.Val = avlNode.Right.FindMin().Val
			avlNode.Right = avlNode.Right.Delete(avlNode.Val)
		} else if avlNode.Left != nil { // 结点只有左孩子，无右孩子
			avlNode = avlNode.Left
		} else { // 结点只有右孩子或者无孩子
			avlNode = avlNode.Right
		}

	}

	// 自动调整平衡, 如果avlNode!=nil说明执行了对avlNode 的某个子树执行了删除结点，那么就需要重新调整树的平衡
	if avlNode != nil {
		avlNode.Height = Max(avlNode.Left.GetHeight(), avlNode.Right.GetHeight()) + 1
		avlNode = avlNode.adjust() // 自动平衡
	}

	return avlNode
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	var root *AVLNode
	for i, v := range nums {
		node := root.lowerBound(v - t)
		if node != nil && node.GetData() <= v+t {
			return true
		}
		root = root.Insert(v)
		if i >= k {
			root = root.Delete(nums[i-k])
		}
	}
	return false
}

func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3))
}
