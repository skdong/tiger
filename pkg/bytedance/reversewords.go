package bytedance

func ReverseWords(s string) string {
	return reverseWords(s)
}

func reverseWords(s string) string {
	words := []string{}
	reverse := ""
	word := ""
	for _, ch := range s {
		if ch <= 'z' && ch >= 'A' {
			word += string(ch)
		} else {
			if len(word) > 0 {
				words = append(words, word)
				word = ""
			}
		}
	}
	for i := len(words) - 1; i > -1; i-- {
		reverse += words[i]
	}
	return reverse
}
