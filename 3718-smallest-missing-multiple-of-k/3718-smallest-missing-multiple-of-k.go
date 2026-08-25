func missingMultiple(nums []int, k int) int {
    set := make(map[int]bool)

    for _, v := range nums{
        if v%k==0{
            set[v] = true
        }
    }

    n := len(set)
    result := -1

    for i:=1; i<=n+1; i++{
        curr := k*i
        if !set[curr]{
            result = curr 
            break
        }
    }

    return result

}