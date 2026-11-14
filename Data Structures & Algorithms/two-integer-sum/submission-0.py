class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        targets = {}
        answer = []
        
        for i in range(len(nums)):
            if nums[i] in targets:
                print("this executed")
                answer.append(targets.get(nums[i]))
                answer.append(i)
                
            targets[target-nums[i]] = i
        return answer