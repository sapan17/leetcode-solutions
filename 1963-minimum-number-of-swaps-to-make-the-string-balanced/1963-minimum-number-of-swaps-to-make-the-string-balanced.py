class Solution(object):
    def minSwaps(self, s):
        """
        :type s: str
        :rtype: int
        """
        close = 0
        closemax = 0
        for i in range(len(s)):
            if s[i] == '[':
                close -= 1
            if s[i] == ']':
                close += 1
            closemax = max(close, closemax)
        
        return (closemax+1)/2