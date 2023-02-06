func preorderTraversal(node *Node, arr *[]int) {
	if node != nil {

		*arr = append(*arr, node.Val)

		for _, j := range node.Children {
			preorderTraversal(j, arr)
		}
	}
}

func preorder(root *Node) []int {

	arr := []int{}

	preorderTraversal(root, &arr)

	return arr
}