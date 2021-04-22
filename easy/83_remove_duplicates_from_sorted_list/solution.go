/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	res := head
	if head.Val == head.Next.Val {
		res = deleteDuplicates(head.Next)
	}

	// TODO 後方の数字の再起をどうするか
	return res
}