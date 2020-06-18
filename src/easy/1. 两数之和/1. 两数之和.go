package main


func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for k, v := range nums {
		if i, ok := m[target - v]; ok {
			return []int{k, i}
		}

		m[v] = k
	}

	return nil
}