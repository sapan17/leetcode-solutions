class Solution(object):
    def compareVersion(self, version1, version2):
        """
        :type version1: str
        :type version2: str
        :rtype: int
        """
        nums1 = version1.split(".")
        nums2 = version2.split(".")
        m, n = len(nums1), len(nums2)
        
        for i in range(max(m,n)):
            int1 = int(nums1[i]) if i < m else 0
            int2 = int(nums2[i]) if i < n else 0
            if int1 != int2:
                return 1 if int1 > int2 else -1
        return 0
            