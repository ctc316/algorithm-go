/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var currVal int
var inited  bool

func recurBST(root *TreeNode) bool {
    if root == nil {return true}
    if !recurBST(root.Left) {return false}
    if !inited {
        inited = true
    } else if root.Val <= currVal {return false}
    currVal = root.Val
    if !recurBST(root.Right) {return false}
    return true
}
func isValidBST(root *TreeNode) bool {
    result := recurBST(root)
    currVal = 0
    inited = false
    return result
}