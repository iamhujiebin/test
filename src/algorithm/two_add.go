package algorithm

type ListNode struct {
	Val  int
	Next *ListNode
}

//letcode中限制，这个方法比下面的还快？！
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l1
	}
	if l2 == nil {
		return l2
	}
	var p, q *ListNode
	p, q = l1, l2
	var carry int
	node := &ListNode{
		Val: -1,
	}
	head := node
	for p != nil && q != nil {
		n := new(ListNode)
		pv, qv := p.Val, q.Val
		v := pv + qv + carry
		carry = 0
		if v > 9 {
			v = v % 10
			carry = 1
		}
		n.Val = v
		node.Next = n
		node = node.Next
		p, q = p.Next, q.Next
	}
	for p != nil {
		v := p.Val + carry
		carry = 0
		if v > 9 {
			v = v % 10
			carry = 1
		}
		node.Next = &ListNode{
			Val: v,
		}
		node = node.Next
		p = p.Next
	}
	for q != nil {
		v := q.Val + carry
		carry = 0
		if v > 9 {
			v = v % 10
			carry = 1
		}
		node.Next = &ListNode{
			Val: v,
		}
		node = node.Next
		q = q.Next
	}
	if carry > 0 {
		node.Next = &ListNode{
			Val: carry,
		}
	}
	return head.Next
}

func AddTwoNumbersUp(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l1
	}
	if l2 == nil {
		return l2
	}
	var p, q *ListNode
	p, q = l1, l2
	var carry int
	node := &ListNode{
		Val: -1,
	}
	head := node
	for p != nil || q != nil {
		var pv, qv int
		if p != nil {
			pv = p.Val
		}
		if q != nil {
			qv = q.Val
		}
		n := new(ListNode)
		v := pv + qv + carry
		carry = 0
		if v > 9 {
			v = v % 10
			carry = 1
		}
		n.Val = v
		node.Next = n
		node = node.Next
		if p != nil {
			p = p.Next
		} else {
			p = nil
		}
		if q != nil {
			q = q.Next
		} else {
			q = nil
		}
	}
	if carry > 0 {
		node.Next = &ListNode{
			Val: carry,
		}
	}
	return head.Next
}
