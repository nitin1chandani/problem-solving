func firstStableIndex(nums []int, k int) int {
    n := len(nums)

    prefixMax := make([]int, n)
    suffixMin := make([]int, n)

    prefixMax[0] = nums[0]
    suffixMin[n-1] = nums[n-1]

    for i := 1; i<n; i++{
        if prefixMax[i-1]<nums[i]{
            prefixMax[i] = nums[i]
        }else{
            prefixMax[i] = prefixMax[i-1]
        }
    }

    for i := n-2; i>=0; i--{
        if suffixMin[i+1]>nums[i]{
            suffixMin[i] = nums[i]
        }else{
            suffixMin[i] = suffixMin[i+1]
        }
    }

    for i := 0; i<n; i++{
        if prefixMax[i]-suffixMin[i] <= k{
            return i
        }
    }

    return -1
}