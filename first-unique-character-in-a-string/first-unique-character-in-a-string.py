class Solution:
    def firstUniqChar(self, s: str) -> int:
        n = len(s)
        
        arr1 = [0]*26
        for i in range(n):
            arr1[ord(s[i]) - ord('a')] += 1
        for i in range(n):
            if arr1[ord(s[i]) - ord('a')] == 1:
                return i
        return -1