"""
# Definition for a Node.
class Node:
    def __init__(self, x, next=None, random=None):
        self.val = int(x)
        self.next = next
        self.random = random
"""

class Solution(object):
    
    def __init__(self):
        self.visitedHash = {}
    
    def copyRandomList(self, head):
        """
        :type head: Node
        :rtype: Node
        """
        if head  == None:
            return None
        
        if head in self.visitedHash:
            return self.visitedHash[head]
        
        node = Node(head.val, None, None)
        self.visitedHash[head] = node
        
        node.next = self.copyRandomList(head.next)
        node.random = self.copyRandomList(head.random)
        
        return node
        
        