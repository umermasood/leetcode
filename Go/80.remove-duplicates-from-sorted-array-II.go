package leetcode

func removeDuplicates2(nums []int) int {
	// keys store unique nums
	// vals store count of unique nums

	// nums = [0,0,1,1,1,1,2,3,3]

	// index := 0

	// for n in nums
	// lookup n in map
	// if n is not in the map, we add it to the map and set count = 1
	// nums[index] = n
	// index++
	// else n is already in the map
	// if count == 1, we make count++
	// nums[index] = n
	// index++
	// else means count > 2, so we count++ again
	// ---------------------------------------------------------------

	// k,v pair of n,c where n is unique num and c is its count
	lookup := make(map[int]int)

	index := 0

	for _, n := range nums {
		// if n is not in the map
		if count, ok := lookup[n]; !ok {
			// add n to the map and set count = 1
			lookup[n] = 1
			nums[index] = n
			index++
		} else {
			// n is already in the map
			if count == 1 {
				lookup[n] = 2
				nums[index] = n
				index++
			}
		}
	}
	return index
}
