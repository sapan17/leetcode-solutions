class Solution(object):
    def findMaxLength(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        count = 0
        maxlen = 0
        table = {0 : 0}
        
        for index, num in enumerate(nums,1):
            if num == 0:
                count -= 1
            else:
                count += 1
                
            if count in table:
                maxlen = max(maxlen, index - table[count])
            else:
                table[count] = index
        return maxlen