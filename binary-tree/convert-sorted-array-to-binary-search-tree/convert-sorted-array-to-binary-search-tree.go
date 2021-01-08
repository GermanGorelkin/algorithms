package convert_sorted_array_to_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	input: [0,1,2,3,4,5]
	expected: [3,1,5,0,2,4]

					3
			1 				5
		0		2		4
*/
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := (len(nums) - 1) / 2
	node := &TreeNode{}
	node.Val = nums[mid]
	node.Left = sortedArrayToBST(nums[:mid])
	node.Right = sortedArrayToBST(nums[mid+1:])

	return node
}
