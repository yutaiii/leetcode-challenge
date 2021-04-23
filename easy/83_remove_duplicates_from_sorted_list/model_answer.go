/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	// 参照渡し
	curr := head
	temp := curr

	for temp != nil {
		// temp.Val == curr.Valは絶対一緒じゃない？
		for temp != nil && temp.Val == curr.Val {
			// tempを1つ進める
			temp = temp.Next
		}
		curr.Next = temp
		curr = temp
	}
	return head
}

// Make use of a temp pointer to skip duplicate nodes.
// After skipping of nodes, link the previous pointer to it's new node.