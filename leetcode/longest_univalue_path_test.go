package leetcode

import "testing"

func TestUnipathCase_1(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val:  5,
			Left: nil,
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	result := longestUnivaluePath(root)
	if result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}

}
