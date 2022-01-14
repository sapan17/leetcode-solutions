class Solution(object):
    def maxProduct(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        res = max(nums)
        maxNum, minNum = 1, 1
        
        for n in nums:
            if n == 0:
                maxNum, minNum = 1, 1
                continue
            temp = n * maxNum
            maxNum = max(n*maxNum, n*minNum, n)
            minNum = min(temp, n*minNum, n)
            res = max(res, maxNum)
        return res
            