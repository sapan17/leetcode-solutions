class Solution(object):
    def maximumUnits(self, boxTypes, truckSize):
        """
        :type boxTypes: List[List[int]]
        :type truckSize: int
        :rtype: int
        """
        boxTypes.sort(key= lambda x:x[1], reverse = True)
        res = 0
        for i in range(len(boxTypes)):
            if truckSize > boxTypes[i][0]:
                truckSize -= boxTypes[i][0]
                res += boxTypes[i][1] * boxTypes[i][0]
            else:
                res += boxTypes[i][1] * truckSize
                truckSize = 0
        return res