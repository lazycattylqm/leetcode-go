package leetcode

func longestUnivaluePath(root *TreeNode) int {
	var result int = 0
	var calcPath func(root *TreeNode) int

	calcPath = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		value := root.Val
		leftNode := root.Left
		rightNode := root.Right
		leftPath := 0
		rightPath := 0
		leftMax := calcPath(leftNode)
		rightMax := calcPath(rightNode)
		if leftNode != nil && leftNode.Val == value {
			leftPath = leftMax + 1
		}
		if rightNode != nil && rightNode.Val == value {
			rightPath = rightMax + 1
		}
		result = max(result, leftPath+rightPath)
		return max(leftPath, rightPath)
	}

	path := calcPath(root)

	return max(result, path)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
