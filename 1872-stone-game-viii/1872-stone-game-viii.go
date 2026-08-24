func stoneGameVIII(stones []int) int {
    n := len(stones)
    prefix := make([]int , n)

    prefix[0] = stones[0]

    for i:=1; i<n; i++{
        prefix[i] = stones[i] + prefix[i-1]
    }

    best := prefix[n-1]

    for i:=n-2; i>=1; i--{
        best = max(best, prefix[i]-best)
    }

    return best
}