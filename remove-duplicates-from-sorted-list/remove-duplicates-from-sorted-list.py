# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution(object):
    def deleteDuplicates(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        first = head
        
        while first is not None:
            second = first.next
            while second is not None and first.val == second.val:
                second = second.next
                
            first.next = second
            first = second
        return head