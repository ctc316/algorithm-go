/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type queueNode struct{
    tnode *TreeNode
    prev *queueNode
}

type Queue struct {
    head *queueNode
    tail *queueNode
    Length int
}

func (this *Queue) Push(tnode *TreeNode) {
    if(tnode == nil) {return}
    newNode := &queueNode{tnode: tnode, prev: nil} 
    if(this.head == nil) {
        this.head = newNode
        this.tail = newNode
    } else {
        this.head.prev = newNode 
        this.head = this.head.prev
    }
    this.Length += 1
}

func (this *Queue) Pop() *TreeNode {
    if(this.tail != nil) {
        retNode := this.tail.tnode
        this.tail = this.tail.prev
        this.Length -= 1
        if this.Length == 0 {this.head = nil}
        return retNode
    }
    return nil
}

func (this *Queue) Print() {
    fmt.Print("[")
    for node:= this.tail; node != nil; node = node.prev {
        fmt.Print(node.tnode.Val,", ")
    }
    fmt.Print("]\n")
}

func findBottomLeftValue(root *TreeNode) int {
    /* Custom Queue */
    queue := Queue{}
    queue.Push(root)
    var curr *TreeNode
    for queue.Length > 0 {
        // queue.Print()
        curr = queue.Pop()
        queue.Push(curr.Right)
        queue.Push(curr.Left)
    }
    return curr.Val

    /***   Standard List   ***/
	// queue := list.New()
	// queue.PushFront(root)
	// lastVal := 0
	// for e := queue.Front(); e != nil; e = e.Next() {
	// tnode := e.Value.(*TreeNode)
	// lastVal = tnode.Val
	// if tnode.Right != nil {queue.PushBack(tnode.Right)}
	// if tnode.Left != nil {queue.PushBack(tnode.Left)}
	// }
	// return lastVal

    /***   Standard array (slice)   ***/
    // curr := root
    // queue := []*TreeNode{curr}
    // for len(queue) > 0 {
    //     curr = queue[0]
    //     queue = queue[1:]
    //     if curr.Right != nil {queue = append(queue, curr.Right)}
    //     if curr.Left != nil  {queue = append(queue, curr.Left)}
    // }
    // return curr.Val
}