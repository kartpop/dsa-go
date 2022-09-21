package main

import (
	"fmt"
	"strings"
)

// StringReverser reverses a string's characters and returns the resultant string.
func StringReverser(s string) string {
	sb := []byte(s) // strings are immutable, first convert string to slice of bytes
	n := len(sb)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		sb[i], sb[j] = sb[j], sb[i]
	}
	return string(sb)
}

// AnagramChecker compares two input strings to check if one can be created by rearranging the other's characters.
// The function is case insensitive and ignores white spaces.
func AnagramChecker(s1 string, s2 string) bool {
	chrcnt := map[string]int{}
	for _, r := range s1 {
		if r != ' ' {
			chrcnt[strings.ToLower(string(r))] += 1
		}
	}
	for _, r := range s2 {
		if r != ' ' {
			chrcnt[strings.ToLower(string(r))] -= 1
		}
	}
	for _, v := range chrcnt {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	// StringReverser test
	s1 := "hey there"
	fmt.Printf("orig: '" + s1 + "' reversed: '" + StringReverser(s1) + "'\n")

	// AnagramChecker test
	anatest := map[string]bool{}
	anatest["waiter; water"] = !AnagramChecker("waiter", "water")
	anatest["Dormitory; Dirty room"] = AnagramChecker("Dormitory", "Dirty room")
	anatest["Slot machines; Cash lost in me"] = AnagramChecker("Slot machines", "Cash lost in me")
	for k, v := range anatest {
		if v {
			fmt.Printf("%s: Pass \n", k)
		} else {
			fmt.Printf("%s: Fail \n", k)
		}

	}
}
