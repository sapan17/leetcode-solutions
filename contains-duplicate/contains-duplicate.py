class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        num = set()
        
        for i in range(len(nums)):
            num.add(nums[i])
        
        return not(len(num) == len(nums))