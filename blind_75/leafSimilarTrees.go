package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// Helper function to perform depth-first tree traversal
func findLeaves(rootNode *TreeNode, leafList *[]int) {
    // Base case: If root is nil
    if rootNode == nil {
        return 
    }
    // If we have reached a leaf node, append it to the leaflist
    if rootNode.Left == nil && rootNode.Right == nil {
        *leafList = append(*leafList, rootNode.Val)
        return
    }
    // Recursively find the leaves 
    findLeaves(rootNode.Left, leafList)
    findLeaves(rootNode.Right, leafList)
}


 func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    // Maintain slices to hold result of traversal
    leavesOfTreeOne := make([]int, 0)
    leavesOfTreeTwo := make([]int, 0)
    // Get the leaves recursively and add to the helper slices
    findLeaves(root1, &leavesOfTreeOne)
    findLeaves(root2, &leavesOfTreeTwo)

    
    // If the lengths of the slices differs, then they're not leaf similar
    if len(leavesOfTreeOne) != len(leavesOfTreeTwo) {
        return false
    }
    // Compare the content of both slices
    for i := 0; i < len(leavesOfTreeOne); i++ {
        if leavesOfTreeOne[i] != leavesOfTreeTwo[i] {
            return false
        }
    }

    return true




}
