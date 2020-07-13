package main


func intersect(nums1 []int, nums2 []int) []int {
	m := map [int]int{}
	res := []int{}
	for _, v := range nums1 {
		m[v]++
	}

	for _, v2 := range nums2 {
		if k, _ := m[v2]; k > 0 {
			m[v2]--
			res = append(res, v2)
			}
		}

	return res
}