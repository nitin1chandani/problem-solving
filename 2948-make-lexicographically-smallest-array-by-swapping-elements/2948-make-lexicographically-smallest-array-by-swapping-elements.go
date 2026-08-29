const MAX = 1_00_000

var MaxBuffer [MAX][2]int
var CompPtr [MAX]int

func lexicographicallySmallestArray(nums []int, limit int) []int {
	n := len(nums)
	sorted := MaxBuffer[:n]
	compPtr := CompPtr[:n]

	for idx, val := range nums {
		sorted[idx] = [2]int{val, idx}
	}

	slices.SortFunc(sorted, func(a, b [2]int) int {
		return a[0] - b[0]
	})

	compPtr[0] = 0
	for i, componentIdx := 0, 0; i < n; i++ {
		origIdx := sorted[i][1]
		nums[origIdx] = componentIdx

		if i < n-1 && sorted[i+1][0]-sorted[i][0] > limit {
			componentIdx++
			compPtr[componentIdx] = i + 1
		}
	}

	for i, idx := range nums {
		componentIdx := compPtr[idx]
		compPtr[idx]++
		nums[i] = sorted[componentIdx][0]
	}
	return nums
}