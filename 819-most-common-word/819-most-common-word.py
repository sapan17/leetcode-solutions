class Solution(object):
    def mostCommonWord(self, paragraph, banned):
        """
        :type paragraph: str
        :type banned: List[str]
        :rtype: str
        """
        
        normal = ''.join([c.lower() if  c.isalnum() else ' ' for c in paragraph])
                
        words = normal.split()
        word_count = defaultdict(int)
        banned_words = set(banned)
        for word in words:
            if word not in banned:
                word_count[word] += 1
                
        return max(word_count.items(), key=operator.itemgetter(1))[0]