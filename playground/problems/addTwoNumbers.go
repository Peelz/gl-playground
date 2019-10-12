package problems

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	var nodes = ListNode{}
	currentNode := &nodes

	for l1 != nil || l2 != nil || carry != 0 {
		v1 := 0
		v2 := 0
		// fmt.Println(l1.Next, l2.Next)
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		tmp := v1 + v2 + carry
		if tmp >= 10 {
			carry = 1
			tmp = tmp % 10
		} else {
			carry = 0
		}
		// fmt.Println("carry", carry, l1, l2)
		if l1 == nil && l2 == nil && carry == 0 {
			currentNode.Next = nil

		} else {
			currentNode.Next = &ListNode{}
		}
		currentNode.Val = tmp

		currentNode = currentNode.Next

	}
	return &nodes
}
