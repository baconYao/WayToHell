# Statement

Write an algorithm to determine if a number `num` is a happy number.

We use the following process to check if a given number is a happy number:

- Starting with the given number `num`, replace the number with the sum of the squares of its digits.
- Repeat the process until:
  - The number equals 1, which will depict that the given number `num` is a happy number.
  - It enters a cycle, which will depict that the given number `num` is not a happy number.

Return TRUE if `num` is a happy number, and FALSE if not.

# Constraints:

- 1 <= `num` <= 2^31 - 1

# Example:

1. Input: `23`, Return `True`
```
    2^2 + 3^2 = 13
    1^2 + 3^2 = 10
    1^2 + 0^2 = 1
    True, it's happy number
```
2. Input: `2`, Return `False`
```
    2^2 = 4
    4^2 = 16
    The next elements in the sequencewill be as follows: 16 -> 37 -> 58 -> 89 -> 145 ->42 -> 20 -> 4As we've encountered 4 before, itindicates that there exists a cycle,and 2 is not a happy number.
```
