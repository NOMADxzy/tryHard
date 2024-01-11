package main

func minOperations(boxes string) []int {
	leftCnt, rightCnt := 0, 0
	times := 0
	for i := 0; i < len(boxes); i++ {
		if boxes[i] == '1' {
			rightCnt++
			times += i
		}
	}
	ans := make([]int, len(boxes))
	if boxes[0] == '1' {
		rightCnt--
	}
	for i := 0; i < len(boxes); i++ {
		ans[i] = times
		times -= rightCnt

		if boxes[i] == '1' {
			leftCnt++
		}
		times += leftCnt
		if i+1 < len(boxes) && boxes[i+1] == '1' {
			rightCnt--
		}
	}
	return ans
}
