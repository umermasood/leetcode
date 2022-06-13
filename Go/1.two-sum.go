package leetcode

func TwoSum(nums []int, target int) []int {
	lookup := make(map[int]int)
	for k, v := range nums {
		if val, ok := lookup[target-v]; ok {
			return []int{val, k}
		}
		lookup[v] = k
	}
	return nil
}
