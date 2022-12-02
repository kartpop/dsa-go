package strings

// Reverser reverses a string's characters and returns the resultant string.
func Reverser(s string) string {
	sb := []byte(s) // strings are immutable, first convert string to slice of bytes
	n := len(sb)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		sb[i], sb[j] = sb[j], sb[i]
	}
	
	return string(sb)
}
