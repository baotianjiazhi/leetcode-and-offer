package main


var result [][]int
func permute(nums []int) [][]int {
	result = [][]int{}
	pathNum := []int{}
	used := make([]bool, len(nums))
	backtrack(pathNum, nums, used)

	return result
 }

func backtrack(pathNum, nums []int, used []bool) {
	if len(pathNum) == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, pathNum)
		result  = append(result, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			pathNum = append(pathNum, nums[i])
			backtrack(pathNum, nums, used)
			pathNum = pathNum[:len(pathNum)-1]
			used[i] = false
		}
	}
}
