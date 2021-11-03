/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    fast := head
    slow := head
    cycle := false
    
    for fast != nil && fast.Next != nil{
        fast = fast.Next.Next
        slow = slow.Next
        
        if slow == fast{
            cycle = true
            break
        }
    }
    
    if cycle == false{
        return nil
    }
    
    slow = head
    for slow != fast{
        slow = slow.Next
        fast = fast.Next
    }
    return fast
    
}