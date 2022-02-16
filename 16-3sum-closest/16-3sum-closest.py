class Solution(object):
    def threeSumClosest(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        nums.sort()
        diff = float('inf')
        
        for i in range(len(nums)-1):
            lo, hi = i + 1, len(nums)-1
            while lo < hi:
                targetsum = nums[i] + nums[lo] + nums[hi]
                
                if abs(targetsum-target) < abs(diff):
                    diff = target - targetsum
                
                if targetsum > target:
                    hi -= 1
                    
                if targetsum < target:
                    lo += 1
                    
                if targetsum == target:
                    break
        return target - diff