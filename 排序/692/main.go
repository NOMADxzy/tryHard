package main

import (
	"fmt"
	"sort"
)

func upCheck(arr []string, i int, kMap map[string]int) {
	f := (i - 1) / 2
	if i >= 0 && (kMap[arr[i]] < kMap[arr[f]] || kMap[arr[i]] == kMap[arr[f]] && arr[i] > arr[f]) {
		arr[i], arr[f] = arr[f], arr[i]
		upCheck(arr, f, kMap)
	}
}

func downCheck(arr []string, i int, kMap map[string]int) {
	left := 2*i + 1
	right := 2*i + 2
	smallest := i
	if left < len(arr) && (kMap[arr[left]] < kMap[arr[smallest]] || kMap[arr[left]] == kMap[arr[smallest]] && arr[left] > arr[smallest]) {
		smallest = left
	}
	// 容易smallest写成i
	if right < len(arr) && (kMap[arr[right]] < kMap[arr[smallest]] || kMap[arr[right]] == kMap[arr[smallest]] && arr[right] > arr[smallest]) {
		smallest = right
	}
	if i != smallest {
		arr[i], arr[smallest] = arr[smallest], arr[i]
		downCheck(arr, smallest, kMap) // 容易smallest写成i
	}
}

func topKFrequent(words []string, k int) []string {
	var queue []string
	kMap := map[string]int{}
	for _, word := range words {
		kMap[word]++
	}
	for s, cnt := range kMap {
		if len(queue) < k {
			queue = append(queue, s)
			upCheck(queue, len(queue)-1, kMap)
		} else if cnt > kMap[queue[0]] || cnt == kMap[queue[0]] && s < queue[0] {
			cnts := make([]int, len(queue))
			for i := 0; i < len(queue); i++ {
				cnts[i] = kMap[queue[i]]
			}
			queue[0] = s
			//queue = queue[:len(queue)-1]
			downCheck(queue, 0, kMap)
			//
			//queue = append(queue, s)
			//upCheck(queue, len(queue)-1, kMap)
		}
	}
	sort.Slice(queue, func(i, j int) bool {
		return kMap[queue[i]] > kMap[queue[j]] || kMap[queue[i]] == kMap[queue[j]] && queue[i] < queue[j]
	})
	return queue
}

func main() {
	words := []string{"glarko", "zlfiwwb", "nsfspyox", "pwqvwmlgri", "qggx", "qrkgmliewc", "zskaqzwo", "zskaqzwo", "ijy", "htpvnmozay", "jqrlad", "ccjel", "qrkgmliewc", "qkjzgws", "fqizrrnmif", "jqrlad", "nbuorw", "qrkgmliewc", "htpvnmozay", "nftk", "glarko", "hdemkfr", "axyak", "hdemkfr", "nsfspyox", "nsfspyox", "qrkgmliewc", "nftk", "nftk", "ccjel", "qrkgmliewc", "ocgjsu", "ijy", "glarko", "nbuorw", "nsfspyox", "qkjzgws", "qkjzgws", "fqizrrnmif", "pwqvwmlgri", "nftk", "qrkgmliewc", "jqrlad", "nftk", "zskaqzwo", "glarko", "nsfspyox", "zlfiwwb", "hwlvqgkdbo", "htpvnmozay", "nsfspyox", "zskaqzwo", "htpvnmozay", "zskaqzwo", "nbuorw", "qkjzgws", "zlfiwwb", "pwqvwmlgri", "zskaqzwo", "qengse", "glarko", "qkjzgws", "pwqvwmlgri", "fqizrrnmif", "nbuorw", "nftk", "ijy", "hdemkfr", "nftk", "qkjzgws", "jqrlad", "nftk", "ccjel", "qggx", "ijy", "qengse", "nftk", "htpvnmozay", "qengse", "eonrg", "qengse", "fqizrrnmif", "hwlvqgkdbo", "qengse", "qengse", "qggx", "qkjzgws", "qggx", "pwqvwmlgri", "htpvnmozay", "qrkgmliewc", "qengse", "fqizrrnmif", "qkjzgws", "qengse", "nftk", "htpvnmozay", "qggx", "zlfiwwb", "bwp", "ocgjsu", "qrkgmliewc", "ccjel", "hdemkfr", "nsfspyox", "hdemkfr", "qggx", "zlfiwwb", "nsfspyox", "ijy", "qkjzgws", "fqizrrnmif", "qkjzgws", "qrkgmliewc", "glarko", "hdemkfr", "pwqvwmlgri"}
	fmt.Println(topKFrequent(words, 14))
	//fmt.Println("nbuorw" > "qggx")
}
