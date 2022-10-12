package strings

import "testing"

func Test_AnagramChecker(t *testing.T) {
	data := []struct{
		name string
		s1 string
		s2 string
		isAnagram bool
	}{
		{"test1", "waiter", "water", false},
		{"test2", "Dormitory", "Dirty room", true},
		{"test3", "Slot machines", "Cash lost in me", true},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T){
			res := AnagramChecker(d.s1, d.s2)
			if res != d.isAnagram {
				t.Errorf("s1: %s, s2: %s: Expected %t, got %t", d.s1, d.s2, d.isAnagram, res)
			}
		})
	}
}