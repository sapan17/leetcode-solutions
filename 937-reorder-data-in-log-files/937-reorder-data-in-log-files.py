class Solution(object):
    def reorderLogFiles(self, logs):
        """
        :type logs: List[str]
        :rtype: List[str]
        """
        def sorting(log):
            left, right = log.split(" ", 1)
            if right[0].isalpha():
                return(0, right, left)
            else:
                return(1,)
        return sorted(logs, key=sorting)