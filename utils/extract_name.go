package utils

// ExtractNameAndNumber extracts the name and number from a username
func ExtractNameAndNumber(username string) (string, string) {
	// Find the last index of a non-digit character from the end of the string
	index := len(username) - 1
	for ; index >= 0; index-- {
		if !Isdigit(username[index]) {
			break
		}
	}
	name := username[:index+1]
	numStr := username[index+1:]
	return name, numStr
}

// Isdigit checks if a byte is a digit
func Isdigit(b byte) bool {
	return '0' <= b && b <= '9'
}
