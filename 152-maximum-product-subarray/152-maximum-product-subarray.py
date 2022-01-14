class Solution(object):
    def maxProduct(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        res = max(nums)
        minnum, maxnum = 1, 1
        
        for n in nums:
            if n == 0:
                minnum, maxnum = 1, 1
            temp = maxnum * n
            
            maxnum = max(n * maxnum, n * minnum, n)
            minnum = min(n* minnum, temp, n)
            res = max(res, maxnum)
        
        return res