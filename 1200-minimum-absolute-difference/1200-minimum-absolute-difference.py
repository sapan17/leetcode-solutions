class Solution(object):
    def minimumAbsDifference(self, arr):
        """
        :type arr: List[int]
        :rtype: List[List[int]]
        """
        arr.sort()
        min_diff = float('inf')
        res = []
        for i in range(len(arr)-1):
            if arr[i+1] - arr[i] < min_diff:
                min_diff = arr[i+1] - arr[i]
        
        for i in range(len(arr)-1):
            if arr[i+1] - arr[i] == min_diff:
                res.append([arr[i],arr[i+1]])
        
        return res
        