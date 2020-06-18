package main

func fourSumCount(A []int, B []int, C []int, D []int) int {
	res_A_B := map[int]int{}
	sum := 0
	for _, v := range A {
		for _, j := range B {
			res_A_B[v+j] += 1
		}
	}

	for _, v := range C {
		for _, j := range D {
			if v, ok := res_A_B[-(v+j)]; ok {
				sum += v
			}
		}
	}


	return sum
 }
