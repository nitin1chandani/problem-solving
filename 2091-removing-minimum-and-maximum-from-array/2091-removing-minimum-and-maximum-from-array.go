func minimumDeletions(nums []int) int {
    n := len(nums)

    if n==1{
        return 1
    }

    min := math.MaxInt
    minIndex := -1

    max := math.MinInt
    maxIndex := -1

    // find min and max
    for i, v := range nums{
        if v > max{
            max = v
            maxIndex = i
        }

        if v < min{
            min = v
            minIndex = i
        }
    }

    // now find min no. of operations
    distMinFromStart := minIndex+1
    distMinFromEnd := n - minIndex

    distMaxFromStart := maxIndex+1
    distMaxFromEnd := n-maxIndex

    result := 0

    if distMinFromStart<distMinFromEnd{
        result += distMinFromStart
    }else{
        result += distMinFromEnd
    }

    if distMaxFromStart<distMaxFromEnd{
        result+=distMaxFromStart
    }else{
        result+=distMaxFromEnd
    }

    if distMaxFromStart>distMinFromStart && distMaxFromStart<result{
        result = distMaxFromStart
    }

    if distMinFromStart>distMaxFromStart && distMinFromStart<result{
        result = distMinFromStart
    }

    if distMinFromEnd>distMaxFromEnd && distMinFromEnd<result{
        result = distMinFromEnd
    }

    if distMaxFromEnd>distMinFromEnd && distMaxFromEnd<result{
        result = distMaxFromEnd
    }
    
    return result


}