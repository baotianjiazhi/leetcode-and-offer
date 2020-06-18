package main

import "fmt"

/**
堆排序的时间复杂度是O(N*lgN)

堆必须是完全二叉树，大根堆的parent>=child，小根堆相反

parent = (i-1)/2
c1 = 2i+1
c2 = 2i+2

  */

// 必须保证子树都是堆的情况下才能建成堆
func heapify(list []int, n int, i int)  {
	if (i >= n) {
		return
	}
	c1 := 2 * i + 1
	c2 := 2 * i + 2
	max := i
	if (c1 < n && list[c1] > list[i]) {
		max = c1
 	}

	if (c2 < n && list[c2] > list[max]) {
		max = c2
	}

	if (max != i) {
		swap(list, max, i)
		heapify(list, n, max)
	}
}

// 将乱序的数组构建成堆
func Build_heap(list []int, n int) {
	last_node := n-1
	parent := (last_node - 1) / 2
	for i := parent; i >= 0; i-- {
		heapify(list, n, i)
	}
}

func heap_sort(list []int, n int) {
	Build_heap(list, n)
	for i := n-1; i>=0; i-- {
		swap(list, i, 0)
		heapify(list, i, 0)
	}
}

func swap(list []int, max, i int) {
	temp := list[i]
	list[i] = list[max]
	list[max] = temp
}

func main() {
	tree := []int{5, 10, 1, 4, 3, 2}
	n := 6
	heap_sort(tree, n)

	for i := 0;i < n;i++ {
		fmt.Println(tree[i])
	}
}