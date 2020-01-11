package algorithm

func TwoSum(nums []int, target int) []int {
	dic := make(map[int]int) // num=>index
	for k, v := range nums {
		dic[v] = k
	}
	for i, v := range nums {
		if j, ok := dic[target-v]; ok {
			return []int{i, j}
		}
	}
	return []int{}
}
