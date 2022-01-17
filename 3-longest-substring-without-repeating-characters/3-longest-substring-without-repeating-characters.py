class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        l = res = 0
        charset = set()
        resLen = 0
        for r in range(len(s)):
            while s[r] in charset:
                charset.remove(s[l])
                l += 1
            charset.add(s[r])
            resLen = max(r-l+1, resLen)
        return resLen