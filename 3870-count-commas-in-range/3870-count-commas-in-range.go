func countCommas(n int) int {
    // 1,000 1,001 10,000
    // 20,000,
    // 10,000,00
    if n<1000{
        return 0
    }

    // 1000 -> 10,000,00

    // 10,000,00
    // 500,000

    count := 0
    
    // for i:=1001; i<=n && i<=999999; i++{
    //     c
    // }

    if n<1000000{
        count = n-1000+1
    }else{
       count = n-1000+2 
    }

    return count

}