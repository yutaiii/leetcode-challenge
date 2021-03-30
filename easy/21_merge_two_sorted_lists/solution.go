/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// hint: https://leetcode.com/problems/merge-two-sorted-lists/discuss/167437/go-solution
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var newList = &ListNode{}
	var out = newList
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			newList.Next = l1
			// l1を1つ進める
			l1 = l1.Next
			newList = newList.Next
			// fmt.Printf("if newList %+v", newList)
			fmt.Printf("if out %+v", out)
		} else {
			newList.Next = l2
			// l2を1つ進める
			l2 = l2.Next
			newList = newList.Next
			fmt.Printf("else out %+v", out)
		}
	}
	// 値はどこに保存されている？

	fmt.Printf("newList %+v", newList)
	fmt.Printf("out %+v", out)

	// 最後に残ったものを格納
	if l1 != nil {
		newList.Next = l1
	} else if l2 != nil {
		newList.Next = l2
	}
	return out.Next
}