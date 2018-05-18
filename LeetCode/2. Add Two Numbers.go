/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{0, nil}
    curr := dummy
    carry := false
    for (l1 != nil || l2 != nil || carry) {
        curr.Next = &ListNode{0, nil}
        curr = curr.Next
        if carry {
            curr.Val += 1
            carry = false
        }
        if l1 != nil {
            curr.Val += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            curr.Val += l2.Val
            l2 = l2.Next
        }
        if curr.Val > 9 {
            curr.Val -= 10
            carry = true
        }
    }
    return dummy.Next
}