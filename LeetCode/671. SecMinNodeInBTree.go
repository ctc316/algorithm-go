/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findMinNode(node *TreeNode, minArr []int) {
    if(node != nil) {
        for i, e := range minArr {
            if node.Val == e {
                break
            } else if e == -1 {
                minArr[i] = node.Val
                break
            } else if node.Val < e {
                moving_val := minArr[i]
                minArr[i] = node.Val
                // push back
                size := len(minArr)
                for i:= i+1; i < size; i++ {
                    tmp := minArr[i]
		            minArr[i] = moving_val
                    moving_val = tmp
	            }
                break
            }
        }
        findMinNode(node.Left, minArr)
        findMinNode(node.Right, minArr)
    }
}

func findSecondMinimumValue(root *TreeNode) int {
    minArr := []int{-1, -1}
    findMinNode(root, minArr)
    return minArr[1]
}