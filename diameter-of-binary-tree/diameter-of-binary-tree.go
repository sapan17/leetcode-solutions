/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil{
        return 0
    }
    return max(diameterOfBinaryTree(root.Left), max(diameterOfBinaryTree(root.Right),(length(root.Left) + length(root.Right))))
}

func length(root *TreeNode) int{
    if root == nil{
        return 0
    }
    return 1 + max(length(root.Left), length(root.Right))
}

func max(a,b int) int{
    if a > b{
        return a
    }
    return b
}