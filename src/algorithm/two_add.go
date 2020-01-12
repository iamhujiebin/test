package algorithm

import (
	"fmt"
	"sort"
)

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

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	len := len(nums)
	if len <= 0 {
		return 0
	}
	mid := int(len / 2)
	two := false
	if len%2 == 0 {
		two = true
	}
	if !two {
		return float64(nums[mid])
	}
	return (float64(nums[mid-1]) + float64(nums[mid])) / 2
}

func FindMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	count := len(nums1) + len(nums2)
	if count <= 0 {
		return 0
	}
	done := false
	var nums []int
	var n, n2 int //记录中位数，最多两位
	var i int
	mid := int(count / 2)
	two := false
	if count%2 == 0 {
		two = true
	}
	p, q := 0, 0
	for p < len(nums1) && q < len(nums2) {
		if nums1[p] < nums2[q] {
			nums = append(nums, nums1[p])
			p++
		} else {
			nums = append(nums, nums2[q])
			q++
		}
		if i == mid {
			if two {
				n, n2 = nums[i-1], nums[i]
			} else {
				n = nums[i]
			}
			done = true
			i++
			break
		}
		i++
	}
	if !done {
		for p < len(nums1) {
			nums = append(nums, nums1[p])
			p++
			if i == mid {
				if two {
					n, n2 = nums[i-1], nums[i]
				} else {
					n = nums[i]
				}
				done = true
				i++
				break
			}
			i++
		}
	}
	if !done {
		for q < len(nums2) {
			nums = append(nums, nums2[q])
			q++
			if i == mid {
				if two {
					n, n2 = nums[i-1], nums[i]
				} else {
					n = nums[i]
				}
				done = true
				i++
				break
			}
			i++
		}
	}
	fmt.Println(nums)
	if two {
		return float64(n+n2) / 2
	} else {
		return float64(n)
	}
}

func FindMedianSortedArraysTarget(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		return FindMedianSortedArraysTarget(nums2, nums1)
	}
	iMin, iMax, halfLen := 0, m, (n+m+1)/2
	for iMin <= iMax {
		i := iMin + (iMax-iMin)/2
		j := halfLen - i
		if i > iMin && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else if i < iMax && nums2[j-1] > nums1[i] {
			iMin = i + 1
		} else {
			leftMax := 0
			if i == 0 {
				leftMax = nums2[j-1]
			} else if j == 0 {
				leftMax = nums1[i-1]
			} else {
				leftMax = max(nums1[i-1], nums2[j-1])
			}
			if (m+n)%2 != 0 {
				return float64(leftMax)
			}
			rightMin := 0
			if i == m {
				rightMin = nums2[j]
			} else if j == n {
				rightMin = nums1[i]
			} else {
				rightMin = min(nums1[i], nums2[j])
			}
			return float64(leftMax+rightMin) / 2.0
		}
	}
	return 0.0
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
