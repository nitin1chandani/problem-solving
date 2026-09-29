// find the path from start to make remaining 0
func backtrack(path []int, start int, remaining int, arr []int, target int, result *[][]int){
    if remaining == 0{
        *result = append(*result, append([]int{}, path...))
        return
    }

    // choices
    for j := start; j<len(arr); j++{
        // make choice
        if arr[j]>remaining{
            continue
        }

        path = append(path, arr[j])

        backtrack(path, j, remaining-arr[j], arr, target, result)

        path = path[:len(path)-1]
    }
}


func combinationSum(candidates []int, target int) [][]int {
    result := make([][]int, 0)
    backtrack([]int{}, 0, target, candidates, target, &result)
    return result
}