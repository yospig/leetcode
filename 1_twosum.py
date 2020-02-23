# coding: utf-8

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        check = {}
        for i, num in enumerate(nums):
            j = target - num
            if j in check:
                return [check[j],i]
            else:
                check[num] = i