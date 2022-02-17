class Solution(object):
    def findDuplicates(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        count = Counter(nums)
        res = []
        for key, value in count.items():
            if value > 1:
                res.append(key)
        return res