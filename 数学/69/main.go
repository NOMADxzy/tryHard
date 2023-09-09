package main

import "fmt"

func find(left int, right int, target int) int {
	if left == right {
		if right*right > target {
			return right - 1
		}
		return left
	}
	mid := (left + right) / 2
	if mid*mid == target {
		return mid
	} else if mid*mid > target {
		if (mid-1)*(mid-1) > target {
			return find(left, mid-1, target)
		} else {
			return mid - 1
		}
	} else {
		if (mid+1)*(mid+1) < target {
			return find(mid+1, right, target)
		} else if (mid+1)*(mid+1) == target {
			return mid + 1
		} else {
			return mid
		}

	}
}

func mySqrt(x int) int {
	return find(1, x, x)
}

func main() {
	fmt.Println(mySqrt(9))
}
