package strings

// WordFlipper reverses each word in string s without changing the order of the words.
func WordFlipper(s string) string {
	b := []byte(s)
	i, j := nextword(b, 0, 0)
	for i < len(b) && j < len(b) {
		nj := j
		b = reverse(b, i, j)
		i, j = nextword(b, nj+1, nj+1)
	}

	return string(b)
}

// nextword scans b and sets i to start-of-next-word and j to end-of-next-word.
func nextword(b []byte, i, j int) (int, int) {
	k := i
	for k < len(b) && string(b[k]) == " " {
		k++
	}
	ni := k
	for k < len(b) && string(b[k]) != " " {
		k++
	}
	nj := k-1

	return ni, nj
}

// reverse reverses the word starting and i and ending at j.
func reverse(b []byte, i, j int) []byte {
	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
	
	return b
}