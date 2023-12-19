package main

import "fmt"

func predictPartyVictory(senate string) string {
	seq := make([]byte, len(senate))
	for i := 0; i < len(senate); i++ {
		seq[i] = senate[i]
	}
	for {
		outGame := make([]bool, len(seq))
		outR, outD := 0, 0
		for i := 0; i < len(seq); i++ {
			if seq[i] == 'R' {
				if outR > 0 {
					outGame[i] = true
					outR--
				} else {
					outD++
				}
			} else {
				if outD > 0 {
					outGame[i] = true
					outD--
				} else {
					outR++
				}
			}
		}
		for i := 0; i < len(seq) && (outR > 0 || outD > 0); i++ {
			if !outGame[i] {
				if outR > 0 && seq[i] == 'R' {
					outR--
					outGame[i] = true
				}
				if outD > 0 && seq[i] == 'D' {
					outD--
					outGame[i] = true
				}
			}
		}
		if outD > 0 {
			return "Radiant"
		} else if outR > 0 {
			return "Dire"
		} else {
			var newSeq []byte
			for i := 0; i < len(seq); i++ {
				if !outGame[i] {
					newSeq = append(newSeq, seq[i])
				}
			}
			seq = newSeq
		}
	}
}

func main() {
	fmt.Println(predictPartyVictory("RDD"))
}
