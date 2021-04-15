// https://leetcode.com/problems/remove-duplicates-from-sorted-list/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	var res *ListNode
	if head.Val == head.Next.Val {
		res = deleteDuplicates(head.Next)
	}

	// wip
}