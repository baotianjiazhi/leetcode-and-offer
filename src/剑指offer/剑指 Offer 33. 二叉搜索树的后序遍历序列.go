package main


func verifyPostorder(postorder []int) bool {
	if len(postorder) == 0 {
		return true
	}

	f := len(postorder)-1

	for i := 0; i < len(postorder); i++ {
		if postorder[i] >= postorder[len(postorder)-1] {
			f = i
		}

		if i > f && postorder[i] < postorder[len(postorder)-1] {
			return false
		}
	}

	return verifyPostorder(postorder[:f]) && verifyPostorder(postorder[f:len(postorder)-1])
}