/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {

	var nodeMap = make(map[*ListNode]struct{})

	if head == nil {
		return false
	}
	nodeMap[head] = struct{}{}

	next := head.Next

	for next != nil {
		_, ok := nodeMap[next]

		if ok {
			return true
		}
		nodeMap[next] = struct{}{}
		next = next.Next
	}

	return false
}