package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestExample1(t *testing.T) {
	assert.Equal(t, countPairs([]int{-1, 1, 2, 3, 1}, 2), 3)
}

func TestExample2(t *testing.T) {
	assert.Equal(t, countPairs([]int{-6, 2, 5, -2, -7, -1, 3}, -2), 10)
}
