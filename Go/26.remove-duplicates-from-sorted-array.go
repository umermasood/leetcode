package leetcode

func removeDuplicates(nums []int) int {
	// loop over the slice
	// store unique elements in map[int, true]
	// lookup num in the map
	// if num is not there, store it in the map and update nums[index] = num to update the nums

	lookup := make(map[int]bool)
	index := 0

	for _, num := range nums {
		if _, ok := lookup[num]; !ok {
			// store unique num in the lookup map
			lookup[num] = true
			nums[index] = num
			index++
		}
	}

	return index
}
