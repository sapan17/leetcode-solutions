class Solution(object):
    def minFlipsMonoIncr(self, s):
        """
        :type s: str
        :rtype: int
        """
        n = len(s)
        count1, count0 = 0, 0
        for i in range(n):
            if s[i] == "1":
                count1 += 1
            else:
                count0 += 1
            count0 = min(count0, count1)
        return count0
            
            
        