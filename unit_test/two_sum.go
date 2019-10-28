package unit_test

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		key := target - num
		if j, ok := m[key]; ok {
			return []int{j, i}
		}
		m[nums[i]] = i
	}
	return []int{}
}
