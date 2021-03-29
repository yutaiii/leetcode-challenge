/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    returnArr := make([]int, 0)
    for _, v := range l1 {
        returnArr = append(returnArr, v)
    }
    lenSum := len(l1) + len(l2) 
    for _, v2 := range l2 {
        for i, _ := range returnArr {
            // 最後以外
            if i + 1 < len(returnArr) {
                if v2 > returnArr[i + 1] {
                    
                }
            } else {
                // 最後
                
            }
        }
    }
    for i := 0; i < lenSum < i++ {

    }
    return returnArr
}