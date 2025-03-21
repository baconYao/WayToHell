class Solution:
    def circularArrayLoop(self, nums: List[int]) -> bool:
        for i in range(len(nums)):
            # True if positive else False
            pre_direction = nums[i] > 0
            slow_i = fast_i = i

            while True:
                # Move slow_i
                slow_i = self._next_idx(slow_i, nums[slow_i], len(nums))
                if self._is_not_cycle(pre_direction, nums[slow_i], len(nums)):
                    break
                # Move fast_i
                fast_i = self._next_idx(fast_i, nums[fast_i], len(nums))
                if self._is_not_cycle(pre_direction, nums[fast_i], len(nums)):
                    break
                # Move fast_i again
                fast_i = self._next_idx(fast_i, nums[fast_i], len(nums))
                if self._is_not_cycle(pre_direction, nums[fast_i], len(nums)):
                    break
                # The direction has changed
                if slow_i == fast_i:
                    return True

        return False

    def _next_idx(self, index: int, value: int, size: int) -> int:
        result = (index + value) % size
        if result < 0:
            result = result + size
        return result

    def _is_not_cycle(
        self, pre_direction: bool, curr_value: int, size_of_nums: int
    ) -> bool:
        # True if positive else False
        curr_direction = curr_value > 0
        if pre_direction != curr_direction or curr_value % size_of_nums == 0:
            return True
        return False
