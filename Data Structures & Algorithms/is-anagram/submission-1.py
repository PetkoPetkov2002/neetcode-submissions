class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        
        if len(s) != len(t):
            return False  

        freq = {}
        for ch in s:
            freq[ch] = freq.get(ch,0) +1
        
        for char in t:

            if char in freq:
                freq[char]-=1
                if freq[char] == 0:
                    del freq[char]
        
        if len(freq) == 0:
            return True
        else:
            return False

