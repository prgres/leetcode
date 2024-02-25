package main

func maximumStrongPairXor(nums []int) int {
	return bruteForce(nums)
}

func bruteForce(nums []int) int {
	xorMax := 0

	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			x := nums[i]
			y := nums[j]

			if !ifStrongPair(x, y) {
				continue
			}

			xorMax = max(x^y, xorMax)
		}
	}
	return xorMax
}

func ifStrongPair(x, y int) bool {
	return abs(y-x) <= min(x, y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
