//func(path, i) - build the path or subset from i ownward
// basecase/complete case record every path as subset
// choices choose any element from index i ownward
// choose make arr[j] to path
// recurse f(path, j+1)
//undo 


func backtrack(nums, path []int, i int, result *[][]int) {
    *result = append(*result, append([]int{}, path...))

    // choices
    for j:=i; j<len(nums); j++{
        //choose
        path = append(path, nums[j])

        //recurse
        backtrack(nums, path, j+1, result)

        //undo
        path = path[:len(path)-1]
    }
}

func subsets(nums []int) [][]int {
    path := make([]int, 0)
    result := make([][]int, 0)
    backtrack(nums, path, 0, &result)
    return result
}