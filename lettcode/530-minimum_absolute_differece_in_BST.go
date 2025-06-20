/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func maxNode(root *TreeNode) int {
    if root.Right == nil {
        return root.Val
    } 

    return maxNode(root.Right)
}

func minNode(root *TreeNode) int {
    if root.Left == nil {
        return root.Val
    } 

    return minNode(root.Left)
}


func min(a int, b int) int {
    if a < b {
        return a
    }
    
    return b
}


func getMinimumDifference(root *TreeNode) int {
    leftMin := 1 << 32 - 1
    minimumLeft := 1 << 32 - 1
    rightMin := 1 << 32 - 1
    minimumRight := 1 << 32 - 1


    if (root.Left != nil) {
        leftMin = root.Val - maxNode(root.Left)
        minimumLeft = getMinimumDifference(root.Left)
    }

    if (root.Right != nil) {
        rightMin = minNode(root.Right) - root.Val
        minimumRight = getMinimumDifference(root.Right)
    }

    return min(min(leftMin, minimumLeft), min(rightMin, minimumRight))
}