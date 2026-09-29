// f(i, path) - generate all combinations of size k from start i
// basecase len(path)==2 append in result
// choices - choose any number from start ownward
// make choice - add arr[j] to path
// recurse f(j+1, path)
// undo

func backtrack(start, k int, arr, path []int, result *[][]int){
    if len(path) == k{
        *result = append(*result, append([]int{}, path...))
        return
    }

    for j := start; j<len(arr); j++{
        //choose
        path = append(path, arr[j])

        //recurse
        backtrack(j+1, k, arr, path, result)

        //undo
        path = path[:len(path)-1]
    }

}


func combine(n int, k int) [][]int {
    arr := make([]int, n)

    for i := 0; i<n; i++{
        arr[i] = i+1
    }

    result := make([][]int, 0)

    backtrack(0, k, arr, []int{}, &result)

    return result
}