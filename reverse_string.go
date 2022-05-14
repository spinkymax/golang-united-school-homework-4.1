package reverse_string

func ReverseString(input string) (output string) {
	var s = []byte(input)
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-i-1; j++ {
			s[j], s[j+1] = s[j+1], s[j]
		}
}
	return (string(s))


}
