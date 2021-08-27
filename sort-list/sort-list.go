/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil{
        return head
    }
    dummy := &ListNode{}
    slow := head
    fast := head
    prev := dummy
    
    for fast != nil && fast.Next != nil{
        prev = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    
    prev.Next = nil
    head1 := sortList(head)
    head2 := sortList(slow)
    prev = dummy
    
    for head1 != nil && head2 != nil{
        if head1.Val <= head2.Val{
            prev.Next = head1
            head1 = head1.Next
        } else{
            prev.Next = head2
            head2 = head2.Next
        }
        prev = prev.Next
    }
    if head1 != nil{
        prev.Next = head1
    }
    if head2 != nil{
        prev.Next = head2
    }
    return dummy.Next
}