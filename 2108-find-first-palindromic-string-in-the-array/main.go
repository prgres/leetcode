package main

func firstPalindrome(words []string) string {
	for _, w := range words {
		if checkPalindrome(w) {
			return w
		}
	}

	return ""
}

func checkPalindrome(word string) bool {
	startPtr := 0
	endPtr := len(word) - 1
	for startPtr < endPtr {
		if word[startPtr] != word[endPtr] {
			return false
		}
		startPtr++
		endPtr--
	}
	return true
}
