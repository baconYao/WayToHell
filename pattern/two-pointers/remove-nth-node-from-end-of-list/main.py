'''
Given a singly linked list, remove the N th node from the end of the list and
return its head.

Constraints:
    - The number of nodes in the list is k.
    - 1 ≤ k ≤ 10^4
    - −10^3 ≤ Node.value ≤ 10^3 ≤ 103
    - 1 ≤ n ≤ Number of nodes in the list
'''


def remove_nth_last_node(head, n):
    if n == 0 or head is None:
        return head

    front = rear = head

    remaining_n = n
    # Move rear pointer n nodes
    while rear.next is not None and remaining_n > 0:
        rear = rear.next
        remaining_n -= 1

    if remaining_n != 0:
        return head

    # Move front and rear pointer until end
    while rear.next is not None:
        front = front.next
        rear = rear.next

    # Update the previous node's next pointer to skip the candidate node
    front.next = front.next.next

    return head
