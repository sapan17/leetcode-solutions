/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil && l2 == nil {
        return nil
    }
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    l3:=&ListNode{}
    if l1.Val < l2.Val{
        l3.Val = l1.Val
        l1 = l1.Next
    } else{
        l3.Val = l2.Val
        l2 = l2.Next
    }
    l3.Next = mergeTwoLists(l1,l2)
    return l3
}