# Statement

An input array, `nums` containing non-zero integers, is given, where the value at each index represents the number of places to skip forward (if the value is positive) or backward (if the value is negative). When skipping forward or backward, wrap around if you reach either end of the array. For this reason, we are calling it a circular array. Determine if this circular array has a cycle. A cycle is a sequence of indices in the circular array characterized by the following:

- The same set of indices is repeated when the sequence is traversed in accordance with the aforementioned rules.
- The length of the sequence is at least two.
- The loop must be in a single direction, forward or backward.

It should be noted that a cycle in the array does not have to originate at the beginning. A cycle can begin from any point in the array.

# Constraints:

Let n be the number of nodes in a linked list.

- 1 <= nums.length <= 10^4
- -5000 <= nums[i] <= 5000
- nums[i] != 0
