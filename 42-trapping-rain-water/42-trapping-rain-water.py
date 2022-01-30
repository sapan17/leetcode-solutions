class Solution(object):
    def trap(self, height):
        """
        :type height: List[int]
        :rtype: int
        """
        maxleft, maxright, water, l = 0, 0, 0, 0
        r = len(height) - 1
        
        while l < r:
            if height[l] < height[r]:
                maxleft = max(maxleft, height[l])
                water += maxleft - height[l]
                l += 1
            
            else:
                maxright = max(maxright, height[r])
                water += maxright - height[r]
                r -= 1
        return water