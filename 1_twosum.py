# coding: utf-8

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        for i, num in enumerate(nums):
            if (target - num) in nums:
                if i != nums.index(target - num):
                    return i, nums.index(target - num)