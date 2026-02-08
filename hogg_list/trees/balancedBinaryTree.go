package main

type TreeNode struct {
	Left *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
    balanced := true

    var dfs func(root *TreeNode) int
    dfs = func(root *TreeNode) int {
        if root == nil {
            return 0
        }

        leftHeight := dfs(root.Left)
        rightHeight := dfs(root.Right)

        if abs(leftHeight - rightHeight) > 1 {
            balanced = false
            return 0
        }

        return 1 + max(leftHeight, rightHeight)
    }

    dfs(root)
    return balanced
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
