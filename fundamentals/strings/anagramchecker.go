package strings

import "strings"

// AnagramChecker compares two input strings and returns true if one can be created by rearranging the other's characters.
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