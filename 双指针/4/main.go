package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)

	i, j := 0, 0
	for i < m && j < n && i+j < (m+n-1)/2 {
		if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	for i < m && i+j < (m+n-1)/2 {
		i++
	}
	for j < n && i+j < (m+n-1)/2 {
		j++
	}
	if i == m {
		if (m+n)%2 == 0 {
			return float64(nums2[j]+nums2[j+1]) / 2
		}
		return float64(nums2[j])
	} else if j == n {
		if (m+n)%2 == 0 {
			return float64(nums1[i]+nums1[i+1]) / 2
		}
		return float64(nums1[i])
	} else if (m+n)%2 == 0 {
		m1, m2 := nums1[i], nums2[j]
		if m1 > m2 {
			m1, m2 = m2, m1
		}
		if i+1 < m && nums1[i+1] < m2 {
			m2 = nums1[i+1]
		}
		if j+1 < n && nums2[j+1] < m2 {
			m2 = nums2[j+1]
		}

		return float64(m1+m2) / 2
	} else {
		ans := float64(nums1[i])
		if nums2[j] < nums1[i] {
			ans = float64(nums2[j])
		}
		return ans
	}
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{-1, 3}))
}
