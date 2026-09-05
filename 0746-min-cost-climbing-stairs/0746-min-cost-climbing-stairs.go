var dp []int

func solve(cost []int, i, n int) int{
    if i>=n{
        return 0
    }

    if dp[i] != -1{
        return dp[i]
    }

    dp[i] = cost[i] + min(solve(cost, i+2, n), solve(cost, i+1, n))
    return dp[i]
}


func minCostClimbingStairs(cost []int) int {
    n := len(cost)
    dp = make([]int , n+1)

    for i := range dp{
        dp[i] = -1
    }

    return min(solve(cost, 0, n), solve(cost, 1, n))
}