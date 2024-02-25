package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestExample1(t *testing.T) {
	assert.Equal(t, maximumStrongPairXor([]int{1, 2, 3, 4, 5}), 7)
}

func TestExample2(t *testing.T) {
	assert.Equal(t, maximumStrongPairXor([]int{10, 100}), 0)
}

func TestExample3(t *testing.T) {
	assert.Equal(t, maximumStrongPairXor([]int{5, 6, 25, 30}), 7)
}

func Test735(t *testing.T) {
	assert.Equal(t, maximumStrongPairXor([]int{1, 2, 2, 1, 2}), 3)
}

func BenchmarkBruteForce(b *testing.B) {
	bruteForce([]int{1, 2, 3, 4, 5})
}
