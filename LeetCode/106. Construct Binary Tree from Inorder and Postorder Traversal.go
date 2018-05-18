/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
    if inorder == nil || len(inorder) == 0 ||
        postorder == nil || len(postorder) == 0 ||
        len(inorder) != len(postorder) {
        return nil
    }

    //num: inorder idx
    inMap := make(map[int]int)
    for idx, val := range(inorder) {
        inMap[val] = idx
    }

    return buildSubTree (inorder, 0, len(inorder) - 1, postorder, 0, len(postorder) - 1, inMap)
}

func buildSubTree (inorder []int, inStart int, inEnd int,
                   postorder []int, postStart int, postEnd int,
                   inMap map[int]int) *TreeNode {

    if postStart > postEnd {
        return nil
    }

    num := postorder[postEnd]
    root := TreeNode{Val: num}

    if postStart == postEnd {
        return &root
    }

    //[9, 3, 15,20,7]
    //[9, 15,7,20, 3]

    numPos := inMap[num]
    rTreeNum := inEnd - numPos
    postRTreeStart := postEnd - rTreeNum
    root.Left = buildSubTree(inorder, inStart, numPos - 1, postorder, postStart, postRTreeStart - 1, inMap)
    root.Right = buildSubTree(inorder, numPos + 1, inEnd, postorder, postRTreeStart, postEnd - 1, inMap)

    return &root
}