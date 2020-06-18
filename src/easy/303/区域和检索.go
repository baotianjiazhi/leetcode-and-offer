package main

type NumArray struct {
	data []int
}


func Constructor(nums []int) NumArray {
	return NumArray{ nums }
}


func (this *NumArray) SumRange(i int, j int) int {
	sum := 0
	for ; i<=j; i++ {
		sum += this.data[i]
	}

	return sum
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */