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


class BetterSolution:
    def circularArrayLoop(self, nums: List[int]) -> bool:
        n, visited = len(nums), set()  # 獲取數組長度，初始化已訪問集合
        for i in range(n):  # 遍歷每個元素作為起點
            if i not in visited:
                local_s = set()  # 當前起點的本地集合
                while True:
                    # 如果當前索引已在本地集合中，表示形成循環
                    if i in local_s:
                        return True

                    # 如果當前索引已在全域已訪問集合中，跳出當前循環
                    if i in visited:
                        break

                    # 將當前索引標記為已訪問
                    visited.add(i)
                    local_s.add(i)

                    # 計算下一個索引
                    prev, i = i, (i + nums[i]) % n

                    # 如果下一個索引和當前索引相同，或方向不一致，中斷
                    if prev == i or (nums[i] > 0) != (nums[prev] > 0):
                        break

        return False
