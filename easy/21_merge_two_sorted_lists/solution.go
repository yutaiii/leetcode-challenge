/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	type ListNode struct {
		Val  int
		Next *ListNode
	}

	// わからん...
	fmt.Printf("l1.Val %+v", l1.Val)
	fmt.Println("=================")
	fmt.Printf("l1.Next %+v", l1.Next)
	fmt.Printf("l1.Val %+v", l1.Next.Val)
	fmt.Println("=================")
	fmt.Printf("l1.Next %+v", l1.Next.Next)
	return l1
}