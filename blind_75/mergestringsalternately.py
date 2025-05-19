class Solution(object):
    def mergeAlternately(self, word1, word2):
        """
        :type word1: str
        :type word2: str
        :rtype: str
        """
        resultString = ""
        length1, length2 = len(word1), len(word2)
        minLength = min(length1, length2)
        
        for i in range(minLength):
            resultString += word1[i] + word2[i]
        
        if length1 > minLength:
            resultString += word1[minLength:]
        elif length2 > minLength:
            resultString += word2[minLength:]
        
        return resultString

