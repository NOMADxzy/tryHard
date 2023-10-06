package main

import (
	"fmt"
	"sort"
)

func smallestTrimmedNumbers(nums []string, queries [][]int) []int {
	sLen := len(nums[0])
	for i := 0; i < len(queries); i++ {
		queries[i] = append(queries[i], i)
	}
	sort.Slice(queries, func(i, j int) bool {
		return queries[i][1]*100+queries[i][0] < queries[j][1]*100+queries[j][0]
	})
	ans := make([]int, len(queries))

	maxLen := queries[len(queries)-1][1]
	qIdx := 0
	numsIdx := make([]int, len(nums))
	for i := 0; i < len(numsIdx); i++ {
		numsIdx[i] = i
	}
	for i := sLen - 1; i > sLen-maxLen-1; i-- {
		buckets := make([][]int, 10)
		for j := 0; j < len(numsIdx); j++ {
			bucketIdx := int(nums[numsIdx[j]][i] - '0')
			buckets[bucketIdx] = append(buckets[bucketIdx], numsIdx[j])
		}

		idx := 0
		for k := 0; k < 10; k++ {
			bkt := buckets[k]
			for l := 0; l < len(bkt); l++ {
				numsIdx[idx] = bkt[l]
				if qIdx < len(queries) {
					for qIdx < len(queries) && queries[qIdx][1] == sLen-i && queries[qIdx][0]-1 == idx {
						ans[queries[qIdx][2]] = bkt[l]
						qIdx++
					}
				}
				idx++
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(smallestTrimmedNumbers([]string{"87", "70"}, [][]int{{2, 1}, {1, 2}, {1, 2}, {1, 1}}))
}
