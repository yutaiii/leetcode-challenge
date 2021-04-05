/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// hint: https://leetcode.com/problems/merge-two-sorted-lists/discuss/167437/go-solution
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 空
	var newList = &ListNode{}
	// 空　参照渡し
	var out = newList
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			newList.Next = l1
			// l1を1つ進める
			l1 = l1.Next
			// newListをl1にする
			newList = newList.Next
			fmt.Println("=========if==========")
			fmt.Printf("newList %+v \n", newList)
			fmt.Printf("out.Next %+v \n", out.Next)
		} else {
			newList.Next = l2
			// l2を1つ進める
			l2 = l2.Next
			newList = newList.Next
			fmt.Println("=========else==========")
			fmt.Printf("newList %+v \n", newList)
			fmt.Printf("out.Next %+v \n", out.Next)
		}
	}

	// 最後に残ったものを格納
	if l1 != nil {
		newList.Next = l1
	} else if l2 != nil {
		newList.Next = l2
	}
	return out.Next
}