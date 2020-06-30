package main

type MinStack struct {
	nums[] int
	minnums[] int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		[]int{},
		[]int{},
	}
}


func (this *MinStack) Push(x int)  {
	this.nums = append(this.nums, x)
	if len(this.minnums) == 0 {
		this.minnums = append(this.minnums, x)
	} else if this.minnums[len(this.minnums)-1] < x {
		this.minnums = append(this.minnums, this.minnums[len(this.minnums)-1])
	} else {
		this.minnums = append(this.minnums, x)
	}
}


func (this *MinStack) Pop()  {
	this.nums = this.nums[:len(this.nums)-1]
	this.minnums = this.minnums[:len(this.minnums)-1]
}


func (this *MinStack) Top() int {
	return this.nums[len(this.nums)-1]
}


func (this *MinStack) Min() int {
	return this.minnums[len(this.minnums)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */