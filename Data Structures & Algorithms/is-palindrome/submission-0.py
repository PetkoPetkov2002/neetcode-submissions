
import string
class Solution:
    def isPalindrome(self, s: str) -> bool:
        new_string = "".join(char.lower() for char in s if char.isalnum())
        new_s = new_string
        j = len(new_s)-1
        for i in range(len(new_s)):
            if j<i:
                return True
            if new_s[i] != new_s[j]:
                return False
        
            j-=1
        return True
