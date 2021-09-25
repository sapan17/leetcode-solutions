/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var preorder []int
func inorderTraversal(root *TreeNode) []int {

    preorder = preorder[:0]
    helper(root)
    return preorder
}

func helper(root *TreeNode){
    if root != nil{
        helper(root.Left)
        preorder = append(preorder,root.Val)
        helper(root.Right)
    }
}