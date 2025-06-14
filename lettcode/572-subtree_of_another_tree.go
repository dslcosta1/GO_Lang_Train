/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isTheSameTree(root *TreeNode, subroot *TreeNode) bool {
	if root == nil && subroot == nil {
		return true
	}

	if root == nil || subroot == nil {
		return false
	}

	if root.Val == subroot.Val {
		checkRight := isTheSameTree(root.Right, subroot.Right)
		checkLeft := isTheSameTree(root.Left, subroot.Left)

		if checkRight && checkLeft {
			return true
		}
	}

	return false
}

func isTheSameTree2(root *TreeNode, subroot *TreeNode) bool {
	if root == nil {
		return false
	}

	if isTheSameTree(root, subroot) {
		return true
	}

	return isTheSameTree2(root.Right, subroot) || isTheSameTree2(root.Left, subroot)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	return isTheSameTree2(root, subRoot)
}