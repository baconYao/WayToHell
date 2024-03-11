# Statement

Given an integer list, nums, find the maximum values in all the contiguous subarrays (windows) of size w.

Note: If the window size is greater than the array size, we consider the entire array as a single window.

# Constraints:

- 1 <= `arr.length` <= 10 ^ 3
- -10 ^ 4 <= `arr[i]` <= 10 ^ 4
- 1 <= `w`

# Example

## Ex 1

Input
```
nums: [-4,2,-5,3,6]
window size: 3
```

Output
```
[2,3,6]
```

## Ex 2

Input
```
nums: [1,2,3,4,5,6]
window size: 6
```

Output
```
[6]
```

## Ex 3

Input
```
nums: [1,2,3,4,5,6,7,8,9,10]
window size: 4
```

Output
```
[4,5,6,7,8,9,10]
```


# Reference

- [239. Sliding Window Maximum](https://leetcode.com/problems/sliding-window-maximum/description/)
