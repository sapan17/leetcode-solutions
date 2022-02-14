class Solution(object):
    def findLeastNumOfUniqueInts(self, arr, k):
        """
        :type arr: List[int]
        :type k: int
        :rtype: int
        """
        
        c = Counter(arr)
        s = sorted(arr,key = lambda x:(c[x],x))
        return len(set(s[k:]))