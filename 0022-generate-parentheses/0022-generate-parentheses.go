
func backtrack(n, open, close int, result *[]string, path []byte){
    if len(path) == 2*n{
        *result = append(*result, string(path))
        return
    }

    // choices

    // make choice
    if open<n{
        path = append(path, '(')
        backtrack(n, open+1, close, result, path)
        path = path[:len(path)-1]
    }

    if close < open{
        path = append(path, ')')
        backtrack(n, open, close+1, result, path)
        path = path[:len(path)-1]
    }
}

func generateParenthesis(n int) []string {
    result := make([]string, 0)

    backtrack(n, 0, 0, &result, []byte{})

    return result
}