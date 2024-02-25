package main

func countPairs(nums []int, target int) int {
	return ver2(nums, target)
	// return ver1(nums, target)
	// return brute(nums, target)
}

func brute(nums []int, target int) int {
	counter := 0
	if len(nums) <= 1 {
		return counter
	}

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] < target {
				counter++
			}
		}
	}

	return counter
}

func ver2(nums []int, target int) int {
	nums = quickSort(nums)
	counter := 0
	//
	// if len(nums) <= 1 {
	// 	return counter
	// }

	leftPtr := 0
	rightPtr := len(nums) - 1

	sum := nums[leftPtr] + nums[rightPtr]

	for leftPtr < rightPtr {
		if sum < target {
			counter += rightPtr - leftPtr
			sum -= nums[leftPtr]
			leftPtr++
			sum += nums[leftPtr]
			continue
		}

		sum -= nums[rightPtr]
		rightPtr--
		sum += nums[rightPtr]
	}

	return counter
}

func ver1(nums []int, target int) int {
	nums = quickSort(nums)
	counter := 0
	leftPtr := 0
	rightPtr := len(nums) - 1

	for leftPtr < rightPtr {
		if nums[leftPtr]+nums[rightPtr] < target {
			counter += rightPtr - leftPtr
			leftPtr++
			continue
		}

		rightPtr--
	}

	return counter
}

func quickSort(nums []int) []int {
	return quickSortSort(nums, 0, len(nums)-1)
}

func quickSortSort(nums []int, lowIdx int, highIdx int) []int {
	if lowIdx < highIdx {
		var pivot int
		pivot, nums = quickSortPartition(nums, lowIdx, highIdx)
		nums = quickSortSort(nums, lowIdx, pivot-1)
		nums = quickSortSort(nums, pivot+1, highIdx)
	}

	return nums
}

func quickSortPartition(nums []int, lowIdx int, highIdx int) (int, []int) {
	pivot := nums[highIdx]
	i := lowIdx - 1
	for j := lowIdx; j < highIdx; j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[i+1], nums[highIdx] = nums[highIdx], nums[i+1]

	return i + 1, nums
}
