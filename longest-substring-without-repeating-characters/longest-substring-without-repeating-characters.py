class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        l = res = 0
        charset = set()
        
        for r in range(len(s)):
            while s[r] in charset:
                charset.remove(s[l])
                l += 1
            charset.add(s[r])
            res = max(res, r-l+1)
        return res
            