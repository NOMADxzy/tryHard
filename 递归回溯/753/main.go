package main

import "fmt"

func findMaybeKey(arr []byte, pos int, k int, keyList *[]string) {
	if pos == len(arr) {
		*keyList = append(*keyList, string(arr))
		return
	}
	for i := 0; i < k; i++ {
		arr[pos] = byte('0' + i)
		findMaybeKey(arr, pos+1, k, keyList)
	}
}

func findNext(acc string, k int, visited map[string]bool, remain, n int) string {
	if remain == 0 {
		return acc
	}
	pre := acc[len(acc)-n+1:]
	for i := 0; i < k; i++ {
		newChar := string(byte('0' + i))
		newKey := pre + newChar
		if !visited[newKey] {
			visited[newKey] = true
			if s := findNext(acc+newChar, k, visited, remain-1, n); len(s) > 0 {
				return s
			}
			visited[newKey] = false
		}
	}
	return ""
}

func crackSafe(n int, k int) string {
	var maybeKeyList []string
	arr := make([]byte, n)
	findMaybeKey(arr, 0, k, &maybeKeyList)
	visited := map[string]bool{}
	start := maybeKeyList[0]
	visited[start] = true
	ans := findNext(start, k, visited, len(maybeKeyList)-1, n)
	return ans
}

func main() {
	fmt.Println(crackSafe(1, 2))
}
