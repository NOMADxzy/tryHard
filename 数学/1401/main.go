package main

import "fmt"

func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
	xsmaller, xbigger, ysmaller, ybigger := false, false, false, false
	cnt := 0
	if xCenter > x2 {
		xbigger = true
		cnt++
	}
	if xCenter < x1 {
		xsmaller = true
		cnt++
	}
	if yCenter > y2 {
		ybigger = true
		cnt++
	}
	if yCenter < y1 {
		ysmaller = true
		cnt++
	}
	if cnt == 0 {
		return true
	} else {
		px, py := xCenter, yCenter
		if xsmaller {
			px = x1
		}
		if xbigger {
			px = x2
		}
		if ysmaller {
			py = y1
		}
		if ybigger {
			py = y2
		}
		return (xCenter-px)*(xCenter-px)+(yCenter-py)*(yCenter-py) <= radius*radius
	}
}

func main() {
	fmt.Println(checkOverlap(1, 0, 0, 1, -1, 3, 1))
}
