package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestExample1(t *testing.T) {
	assert.Equal(t, firstPalindrome([]string{"abc", "car", "ada", "racecar"}), "ada")
}

func TestExample2(t *testing.T) {
	assert.Equal(t, firstPalindrome([]string{"notapalindrome", "racecar"}), "racecar")
}

func Test7(t *testing.T) {
	assert.Equal(t, firstPalindrome([]string{
		"cqllrtyhw", "swwisru", "gpzmbders", "wqibjuqvs", "pp", "usewxryy", "ybqfuh", "hqwwqftgyu", "jggmatpk",
	}), "pp")
}
