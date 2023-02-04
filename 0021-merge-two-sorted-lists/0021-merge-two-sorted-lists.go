func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil || list2 == nil {
		if list1 != nil {
			return list1
		} else if list2 != nil {
			return list2
		}
		return nil
	}

	var root *ListNode
	var prev *ListNode

	if list1.Val <= list2.Val {
		root = list1
		list1 = list1.Next
	} else {
		root = list2
		list2 = list2.Next
	}
	prev = root
	
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			prev.Next = list1
			prev = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			prev = list2
			list2 = list2.Next
		}
	}

	if list1 == nil {
		prev.Next = list2
	} else {
		prev.Next = list1
	}

	return root
}