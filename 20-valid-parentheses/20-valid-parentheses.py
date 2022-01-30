class Solution(object):
    def isValid(self, s):
        """
        :type s: str
        :rtype: bool
        """
        
        stack = []
        OpentoClose = {"}":"{", ")":"(", "]":"["}
        
        for c in s:
            if c in OpentoClose:
                if stack and stack[-1] == OpentoClose[c]:
                    stack.pop()
                else:
                    return False
            else:
                stack.append(c)
        return True if not stack else False