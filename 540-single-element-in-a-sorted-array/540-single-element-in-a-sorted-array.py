class Solution(object):
    def singleNonDuplicate(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        count = Counter(nums)
        for key, values in count.items():
            if values == 1:
                return key
        return -1