package strings

import "testing"

func Test_Reverser(t *testing.T) {
	s := "heY there"
	e := "ereht Yeh"
	r := Reverser(s)
	if e != r {
		t.Error("incorrect result: expected: ereht Yeh, got: ", r)
	}
}

