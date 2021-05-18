package algorithms

import (
	ds "github.com/codewithji/go-algorithms/data_structures"
)

// Floyd's tortoise and hare algorithm
// You can use two pointers to detect whether there is a cycle in a linked list
// Benefit of not using a hashmap to store seen nodes, which would be O(N) space
func detectCycleInLinkedList(head *ds.ListNode) *ds.ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}

	slow, fast := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			break
		}
		// determines if there is an end to the linked list
		if fast.Next == nil || fast.Next.Next == nil {
			return nil
		}
	}
	// reset slow pointer to beginning
	slow = head

	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}
