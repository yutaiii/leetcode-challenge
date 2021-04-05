/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// hint: https://leetcode.com/problems/merge-two-sorted-lists/discuss/360472/Golang-0ms-recursive-solution
// memo: シンプルでこれが一番いいかも
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		// put next value
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	// put next value
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}
