package leetcode_102

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
}

// Entry function
func levelOrder(root *TreeNode) [][]int {
    var ans [][]int
    if(root == nil) {return ans}
    curr := []*TreeNode{root}
    for len(curr) > 0 {
        var next []*TreeNode
        var a []int
        for _,e := range curr {
            a = append(a, e.Val)
            if(e.Left != nil) {next = append(next, e.Left)}
            if(e.Right != nil) {next = append(next, e.Right)}
        }
        ans = append(ans, a)
        curr = next
    }
    return ans
}
