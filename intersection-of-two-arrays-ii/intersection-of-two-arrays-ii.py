class Solution(object):
    def intersect(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: List[int]
        """
        c = Counter(nums1)
        output = []
        for n in nums2:
            if c[n] > 0:
                output.append(n)
                c[n] -= 1
        return output
            