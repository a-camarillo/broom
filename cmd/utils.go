package cmd

func pop(s []string, i int) []string {
	copy(s[i:], s[i+1:])
	s[len(s)-1] = ""
	s = s[:len(s)-1]
	return s
}
