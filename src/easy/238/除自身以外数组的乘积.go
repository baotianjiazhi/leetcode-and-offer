package main
import "fmt"


func main() {

	nums := []int{1, 2, 3, 4}
	productExceptSelf(nums)
}
func productExceptSelf(nums []int) []int {
	length := len(nums)
	L, R, res := make([]int, length), make([]int, length), make([]int, length)

	L[0] = 1
	for i := 1; i < length; i ++ {
		L[i] = L[i - 1] * nums[i - 1]
	}

	R[length - 1] = 1
	for i := length - 2; i >= 0; i -- {
		R[i] = R[i + 1] * nums[i+1]
	}

	fmt.Println(R, L)
	for i := 0; i < length; i ++ {
		res[i] = R[i] * L[i]
	}

	return res
}
