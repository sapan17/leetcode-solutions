class Solution(object):
    def longestConsecutive(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        if not nums:
            return 0
        nums.sort()
        
        cur_streak = 1
        long_streak = 1
        
        for i in range(1, len(nums)):
            if nums[i] != nums[i-1]:
                if nums[i] == nums[i-1]+1:
                    cur_streak += 1
                else:
                    long_streak = max(long_streak, cur_streak)
                    cur_streak = 1
        return max(long_streak, cur_streak)