class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        
        res = set()
        for i in range(len(nums)):
            res.add(nums[i])
        return not(len(res) == len(nums))