/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func traversal(node *TreeNode, currAns []int) []int {
    if(node == nil) { return currAns }
    currAns = traversal(node.Left, currAns)
    currAns = append(currAns, node.Val)
    currAns = traversal(node.Right, currAns)
    return currAns
}

func inorderTraversal(root *TreeNode) []int {
    return traversal(root, []int{})
}