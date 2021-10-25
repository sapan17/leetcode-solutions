class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        arr1 = [0] * 26
        arr2 = [0] * 26
        
        if len(s) != len(t):
            return False
        
        for i in range(len(s)):
            arr1[ord(s[i]) - 97] += 1
            arr2[ord(t[i]) - 97] += 1
            
        return arr1 == arr2