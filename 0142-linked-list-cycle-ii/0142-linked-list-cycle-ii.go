func detectCycle(head *ListNode) *ListNode {
	arr := []*ListNode{}
	for head != nil {
		count := 0
		arr = append(arr, head)
		for _, key := range arr {
			if key == head.Next {
				fmt.Println(count)
				return key
			}
			count++
		}
		head = head.Next
	}
	return nil
}