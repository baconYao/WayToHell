# Statement

Given two strings, `str1` and `str2`, find the shortest substring in `str1` such that `str2` is a subsequence of that substring.

A substring is defined as a contiguous sequence of characters within a string. A subsequence is a sequence that can be derived from another sequence by deleting zero or more elements without changing the order of the remaining elements.

Let’s say you have the following two strings:
```
str1 = "abbcb"

str2 = "ac"
```

In this example, `"abbc"` is a substring of `str1`, from which we can derive `str2` simply by deleting both the instances of the character `b`. Therefore, `str2` is a subsequence of this substring. Since this substring is the shortest among all the substrings in which `str2` is present as a subsequence, the function should return this substring, that is, `"abbc"`.

- If there is no substring in str1 that covers all characters in str2, return an empty string.
- If there are multiple minimum-length substrings that meet the subsequence requirement, return the one with the left-most starting index.

# Constraints:

- 1 <= `str1.length.length` <= 2 * 10 ^ 3
- 1 <= `str2.length.length` <= 100
- `str1` and `str2` consist of uppercase and lowercase English letters.


# Reference

- [727. Minimum Window Substring](https://leetcode.com/problems/minimum-window-subsequence/description/)

## Similar
- [76. Minimum Window Substring](https://leetcode.com/problems/minimum-window-substring/description/)
