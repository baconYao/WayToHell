class Solution:
    def sortColors(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        start, current, end = 0, 0, len(nums) - 1
        while current <= end:
            if nums[current] == 0:
                nums[current], nums[start] = nums[start], nums[current]
                start += 1
                current += 1
            elif nums[current] == 1:
                current += 1
            else:
                nums[current], nums[end] = nums[end], nums[current]
                end -= 1
        return nums
