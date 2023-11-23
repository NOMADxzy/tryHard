package main

import "fmt"

func goNext(target string, tarPos int, pattern string, pos int) bool {
	if tarPos == len(target) && pos == len(pattern) {
		return true
	} else if tarPos != len(target) && pos != len(pattern) {
		if pattern[pos] >= 'a' && pattern[pos] <= 'z' {
			if target[tarPos] == pattern[pos] && goNext(target, tarPos+1, pattern, pos+1) {
				return true
			}
		} else if pattern[pos] == '.' {
			if goNext(target, tarPos+1, pattern, pos+1) {
				return true
			}
		} else if pattern[pos] == '*' {
			if pos == 0 {
				return false
			}
			c := pattern[pos-1]
			for p := tarPos; p < len(target) && (c == '.' || target[p] == c); p++ {
				if goNext(target, p+1, pattern, pos+1) {
					return true
				}
			}
			return false
		}
		if pos+1 < len(pattern) && pattern[pos+1] == '*' &&
			(goNext(target, tarPos, pattern, pos+2) || (pattern[pos] == '.' || target[tarPos] == pattern[pos]) && goNext(target, tarPos+1, pattern, pos+2)) {
			return true
		}
	}
	if len(target) == tarPos {
		if pos+1 < len(pattern) && pattern[pos+1] == '*' && goNext(target, tarPos, pattern, pos+2) {
			return true
		}
	}
	return false
}

func isMatch(s string, p string) bool {
	newP := p[:1]
	for i := 1; i < len(p); i++ {
		if p[i] == '*' && p[i-1] == '*' {
			continue
		}
		newP += string(p[i])
	}
	return goNext(s, 0, newP, 0)
}

func main() {
	fmt.Println(isMatch("aaa", "ab*a*c*a"))
}
