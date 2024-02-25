package main

import (
	"fmt"
	"strings"
)

func isStrictlyPalindromic(n int) bool {
	for i := 2; i <= n-2; i++ {
		nStr := baseEncode(n, i)
		if !isPalindrome(nStr) {
			return false
		}
	}

	return true
}

func isPalindrome(s string) bool {
	start := 0
	end := len(s) - 1

	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}

	return true
}

func baseEncode(n, b int) string {
	res := strings.Builder{}

	for n != 0 {
		res.WriteString(fmt.Sprint((n % b)))
		n /= b
	}

	return reverse(res.String())
}

func reverse(s string) string {
	res := strings.Builder{}
	for i := len(s) - 1; i >= 0; i-- {
		res.WriteByte(s[i])
	}
	return res.String()
}
