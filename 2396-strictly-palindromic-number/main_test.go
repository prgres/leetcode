package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestExample1(t *testing.T) {
	assert.Equal(t, isStrictlyPalindromic(9), false)
}

func TestExample2(t *testing.T) {
	assert.Equal(t, isStrictlyPalindromic(4), false)
}
