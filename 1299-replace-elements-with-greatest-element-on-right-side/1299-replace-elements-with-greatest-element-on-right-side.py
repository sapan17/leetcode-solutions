class Solution(object):
    def replaceElements(self, arr):
        """
        :type arr: List[int]
        :rtype: List[int]
        """
        rightmax = -1
        for i in range(len(arr) - 1, -1, -1):
            maxnum = max(rightmax, arr[i])
            arr[i] = rightmax
            rightmax = maxnum
        return arr