/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
   var prev *ListNode = nil
    
    for head != nil{
        curr := head
        head = head.Next
        curr.Next = prev
        prev = curr
    }
    return prev
}