package main
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// If the value of p and q are less than the root's value
	// Recursively check the left subtree
	if p.Val < root.Val && q.Val < root.Val {
        return lowestCommonAncestor(root.Left, p, q)
    }
	// Otherwise, if the value of the root is less than the values of p and q
	// recursively check the right subtree
     else if root.Val < p.Val && root.Val < q.Val {
        return lowestCommonAncestor(root.Right, p, q)
    } else {
    	// Else, considering the root node is the LCA of all nodes by default, return it
        return root
    }
}
