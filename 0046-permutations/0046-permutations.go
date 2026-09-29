// f(path) -  generate all permutations using the elements not yet present in path
// base case - if len(path)==len(nums) add it in result 
//choose any unused element from nums
// make choice - add to path
// recurse - f(path)
//undo

func backtrack(path, nums []int, result *[][]int, used []bool){
    if len(path)==len(nums){
        *result = append(*result, append([]int{}, path...))
        return
    }

    for i := 0; i<len(nums); i++{
        if used[i]{
            continue
        }

        path = append(path, nums[i])
        used[i] = true

        backtrack(path, nums, result, used)

        path = path[:len(path)-1]
        used[i] = false
    }
}

func permute(nums []int) [][]int {
    result := make([][]int, 0)
    used := make([]bool, len(nums))
    
    backtrack([]int{}, nums, &result, used)
    return result
}