/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
    
    var queue []*TreeNode
    var means []float64

    if root == nil {
        return means
    }

    numberOfElementsInLevel := 1
    queue = append(queue, root)
    
    for len(queue) != 0 {
        count := 0
        sum := 0
        for i:=0; i<numberOfElementsInLevel; i++ {
            first := queue[0]
            queue = queue[1:]
            sum += first.Val

            if first.Left != nil {
                queue = append(queue, first.Left)
                count++
            }

            if first.Right != nil {
                queue = append(queue, first.Right)
                count++
            }
        }
        means = append(means, float64(sum)/float64(numberOfElementsInLevel))
        numberOfElementsInLevel = count
    }

    return means
}