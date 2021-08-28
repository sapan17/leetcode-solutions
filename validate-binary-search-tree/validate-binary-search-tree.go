/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    return isValid(root, nil, nil)
}

func isValid(node *TreeNode, min *int, max *int) bool{
    if node == nil{
        return true
    }
    value := node.Val
    
    if min != nil{
        if value <= *min{
            return false
        }
    }
    if max != nil{
        if value >= *max{
            return false
        }
    }
    return isValid(node.Left, min, &value) && isValid(node.Right, &value, max)
}