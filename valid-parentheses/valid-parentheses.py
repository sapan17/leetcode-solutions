class Solution(object):
    def isValid(self, s):
        """
        :type s: str
        :rtype: bool
        """
        stack = []
        CloseToOpen = {"]":"[", "}":"{", ")":"("}
        for r in s:
            if r in CloseToOpen:
                if stack and stack[-1] == CloseToOpen[r]:
                    stack.pop()
                else:
                    return False
            else:
                stack.append(r)
        return True if not stack else False
            
        