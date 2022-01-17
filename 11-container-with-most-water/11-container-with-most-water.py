class Solution(object):
    def maxArea(self, height):
        """
        :type height: List[int]
        :rtype: int
        """
        l, r = 0, len(height) - 1
        
        maxArea = 0
        while l < r:
            Area = min(height[l], height[r]) * (r-l)
            maxArea = max(maxArea, Area)
            
            if height[l] > height[r]:
                r -= 1
            else:
                l += 1
        
        return maxArea