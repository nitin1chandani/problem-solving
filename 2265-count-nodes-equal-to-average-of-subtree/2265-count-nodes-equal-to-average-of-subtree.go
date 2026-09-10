func solve(node *TreeNode, result *int) (int, int) {
    if node == nil {
        return 0, 0
    }

    leftSum, leftCount := solve(node.Left, result)
    rightSum, rightCount := solve(node.Right, result)

    sum := node.Val + leftSum + rightSum
    count := 1 + leftCount + rightCount

    avg := sum / count

    if avg == node.Val {
        *result++
    }

    return sum, count
}

func averageOfSubtree(root *TreeNode) int {
    result := 0

    solve(root, &result)

    return result
}