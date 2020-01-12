package algorithm

import (
	"testing"
)

func TestTwoAdd(t *testing.T) {
	n1 := &ListNode{Val: 2}
	n2 := &ListNode{Val: 4}
	n3 := &ListNode{Val: 3}
	n1.Next = n2
	n2.Next = n3

	n4 := &ListNode{Val: 5}
	n5 := &ListNode{Val: 6}
	n6 := &ListNode{Val: 4}
	n4.Next = n5
	n5.Next = n6

	target := AddTwoNumbers(n1, n4)
	if target == nil {
		t.Log("target is nil")
		return
	}
	var n int
	for target != nil {
		n++
		t.Log(target.Val)
		target = target.Next
	}
}

func TestTwoAddUp(t *testing.T) {
	n1 := &ListNode{Val: 2}
	n2 := &ListNode{Val: 4}
	n3 := &ListNode{Val: 3}
	n1.Next = n2
	n2.Next = n3

	n4 := &ListNode{Val: 5}
	n5 := &ListNode{Val: 6}
	n6 := &ListNode{Val: 4}
	n4.Next = n5
	n5.Next = n6

	target := AddTwoNumbersUp(n1, n4)
	if target == nil {
		t.Log("target is nil")
		return
	}
	var n int
	for target != nil {
		n++
		t.Log(target.Val)
		target = target.Next
	}
}

func TestTwoSum(t *testing.T) {
	arr := TwoSum([]int{1, 2, 3, 4}, 5)
	t.Log(arr)
}

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Log(LengthOfLongestSubstring("ababcabcdef"))
	t.Log(LengthOfLongestSubstring2("ababcabcdef"))
	t.Log(LengthOfLongestSubStringTarget2("ababc"))
}

func TestFindMedianSortedArrays(t *testing.T) {
	t.Log(FindMedianSortedArrays([]int{1, 2, 2, 5, 6}, []int{3, 4, 4}))
	t.Log(FindMedianSortedArrays2([]int{1, 2, 2, 5, 6}, []int{3, 4, 4}))
}

func TestLongestPalindrome(t *testing.T) {
	t.Log(LongestPalindrome("ddddexabcbaxeeffff"))
	t.Log(LongestPalindrome("aabbaa"))
}
