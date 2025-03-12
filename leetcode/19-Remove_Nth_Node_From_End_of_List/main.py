from typing import Optional


# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:  # noqa:501
        # Set a dummy node to point to head
        dummy = ListNode(0, head)
        left = dummy
        # Set right pointer to head
        right = head
        # Move the right pointer Nth steps forward
        for _ in range(n):
            right = right.next

        # Move left and right pointers until the right pointer reaches the
        # last node.
        # At this point, the left pointer will be pointing to the node behind
        # the nth last node.
        while right:
            right = right.next
            left = left.next

        # Relink the left node to the node next to left's next node
        left.next = left.next.next
        return dummy.next
