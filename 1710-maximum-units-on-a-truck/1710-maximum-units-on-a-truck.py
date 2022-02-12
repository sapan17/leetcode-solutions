class Solution:
    def maximumUnits(self, boxTypes: List[List[int]], truckSize: int) -> int:
        boxTypes.sort(key= lambda x:x[1])
        boxTypes.reverse()
        res = 0
        for i in range(len(boxTypes)):
            if truckSize > boxTypes[i][0]:
                truckSize -= boxTypes[i][0]
                res += boxTypes[i][1] * boxTypes[i][0]
            else:
                res += boxTypes[i][1] * truckSize
                truckSize = 0
        return res
        