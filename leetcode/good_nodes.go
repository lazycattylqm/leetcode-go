package leetcode

func goodNodes(root *TreeNode) int {
	return goodNodesHelper(root, root.Val)

}

func goodNodesHelper(root *TreeNode, val int) int {
	if root == nil {
		return 0
	}
	var count int
	if root.Val >= val {
		count = 1
		val = root.Val
	}
	return count + goodNodesHelper(root.Left, val) + goodNodesHelper(root.Right, val)

}
