class Solution(object):
    def monotoneIncreasingDigits(self, n):
        """
        :type n: int
        :rtype: int
        """
        s = str(n)
        for i in range(len(s)-1, 0, -1):
            if s[i] < s[i-1]:
                n -= (int(s[i:])+1)
                s = str(n)
        return n