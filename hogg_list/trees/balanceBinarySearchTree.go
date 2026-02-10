package main
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func balanceBST(root *TreeNode) *TreeNode {
    // traverse the nodes with inorder traversal and put them in a slice 
    // in ascending order. 
    var nodes []int
    var inOrder func(root *TreeNode)  
    inOrder = func(root *TreeNode) {
        if root == nil {
            return 
        }
        inOrder(root.Left)
        nodes = append(nodes, root.Val)
        inOrder(root.Right)
    }

    inOrder(root)

    // Use the middle value as the root of the bst, then create a bst from the nodes
    // return the bst
    var constructBST func(int, int) *TreeNode 
    constructBST = func(l, r int) *TreeNode {
        if l > r {
            return nil
        }
    
    mid := (l + r) / 2
    balancedBST := &TreeNode{Val: nodes[mid]}
    root.Left = constructBST(l, mid - 1)
    root.Right = constructBST(mid + 1, r)
    return balancedBST
  }
  return constructBST(0, len(nodes) - 1)
}
