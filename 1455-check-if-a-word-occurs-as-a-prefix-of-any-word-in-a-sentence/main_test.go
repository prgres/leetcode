package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestExample1(t *testing.T) {
	assert.Equal(t, 4, isPrefixOfWord("i love eating burger", "burg"))
}

func TestExample2(t *testing.T) {
	assert.Equal(t, 2, isPrefixOfWord("this problem is an easy problem", "pro"))
}

func TestExample3(t *testing.T) {
	assert.Equal(t, -1, isPrefixOfWord("i am tired", "you"))
}

func Test8(t *testing.T) {
	assert.Equal(t, -1, isPrefixOfWord("hellohello hellohellohello", "ell"))
}

func Test9(t *testing.T) {
	assert.Equal(t, -1, isPrefixOfWord("i dumb", "jiezcqu"))
}

func Test11(t *testing.T) {
	assert.Equal(t, -1, isPrefixOfWord("jonathan corona they", "cywxvea"))
}

func Test40(t *testing.T) {
	assert.Equal(t, 4, isPrefixOfWord("love errichto jonathan dumb", "dumb"))
}
