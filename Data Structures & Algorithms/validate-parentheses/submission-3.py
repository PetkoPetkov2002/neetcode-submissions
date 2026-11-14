class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        dict_map = {')':'(',']':'[','}':'{'}
        for ch in s:
            if ch in dict_map:
                if not stack:
                    return False
                if dict_map.get(ch) == stack[-1]:
                    stack.pop()
                else:
                    return False
            else:
                stack.append(ch)
            
        return not stack