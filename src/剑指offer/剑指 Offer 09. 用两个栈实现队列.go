package main

type CQueue struct {
	instack  []int
	outstack []int
}

func Constructor() CQueue {
	return CQueue{
		instack: []int{},
		outstack: []int{},
	}
}

func (this *CQueue) AppendTail(value int) {
	this.instack = append(this.instack, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.instack) == 0 && len(this.outstack) == 0 {
		return -1
	}

	if len(this.outstack) == 0 {
		for len(this.instack) > 0 {
			index := len(this.instack) - 1
			this.outstack = append(this.outstack, this.instack[index])
			this.instack = this.instack[:index]
		}
	}

	res := this.outstack[len(this.outstack)-1]
	this.outstack = this.outstack[:len(this.outstack)-1]
	return res
}
