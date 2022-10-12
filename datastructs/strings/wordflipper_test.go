package strings

import "testing"

func Test_WordFlipper(t *testing.T) {
	data := []struct{
		name string
		input string
		output string
	}{
		{"test1", "Hello how are you", "olleH woh era uoy"},
		{"test2", "Kartik", "kitraK"},
		{"test3", "sihT si eno llams pets rof ...", "This is one small step for ..."},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T){
			res := WordFlipper(d.input)
			if res != d.output {
				t.Errorf("%s -> Expected: %s, Got: %s", d.name, d.output, res)
			}
		})
	}
}