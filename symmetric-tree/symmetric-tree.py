# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def isSymmetric(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        return self.dfs(root, root)
    
    def dfs(self, root1, root2):
        
        if root1 == None and root2 == None:
            return True
        if root1 == None or root2 == None:
            return False
        return root1.val == root2.val and self.dfs(root1.left , root2.right) and self.dfs(root2.right , root1.left)
    