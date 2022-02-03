class Solution(object):
    def threeSumClosest(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        nums.sort()
        diff = float("inf")
        for i in range(len(nums)):
            lo, high = i+1, len(nums)-1
            while lo < high:
                threesum = nums[i] + nums[lo] + nums[high]
                if abs(target-threesum) < abs(diff):
                    diff = target - threesum
                if threesum > target:
                    high -= 1
                else: 
                    lo += 1
            if diff == 0:
                break
        return target - diff 
                
            
            