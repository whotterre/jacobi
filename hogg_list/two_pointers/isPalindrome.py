
class Solution:
    def isPalindrome(self, s: str) -> bool:
        lower_case = "".join(char.lower() for char in s if char.isalnum())
        # Two pointers
        i = 0
        j = len(lower_case) - 1

        while i < j :
            if lower_case[i] != lower_case[j]:
                return False
            i += 1
            j -= 1
        return True