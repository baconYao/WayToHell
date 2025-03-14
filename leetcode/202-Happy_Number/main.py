class Solution:
    def isHappy(self, n: int) -> bool:
        slow_p = n
        fast_p = self._sum_of_squared_digits(n)

        if fast_p == 1:
            return True
        while slow_p != fast_p:
            fast_p = self._sum_of_squared_digits(self._sum_of_squared_digits(fast_p))
            if fast_p == 1:
                return True
            slow_p = self._sum_of_squared_digits(slow_p)
        return False

    # Worse Performance
    # def _sum_of_squared_digits(self, num: int) -> int:
    #     return sum(map(lambda x: int(x) ** 2, str(num)))

    # Better performance
    def _sum_of_squared_digits(self, number: int) -> int:
        total_sum = 0
        while number > 0:
            number, digit = divmod(number, 10)
            total_sum += digit**2
        return total_sum
