package Algorithms

func reverseWords(s string) string {
	inWord := true
	l := 0
	sl := []byte(s)

	for i := range sl {
		if s[i] == ' ' && inWord {
			reverseWord(sl[l:i])
			inWord = false
		} else if s[i] != ' ' && !inWord {
			l = i
			inWord = true
		} else if i == len(sl)-1 {
			reverseWord(sl[l : i+1])
		}
	}

	return string(sl)
}

func reverseWord(s []byte) {
	last := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		s[i], s[last-i] = s[last-i], s[i]
	}
}
