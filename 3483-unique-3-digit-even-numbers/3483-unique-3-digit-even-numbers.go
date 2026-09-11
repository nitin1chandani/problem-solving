func totalNumbers(digits []int) int {
	var freq [10]int = [10]int{}
	for _, d := range digits { freq[d]++ }
	var overall int = 0
	for d := 0; d <= 8; d += 2 {
		if freq[d] == 0 { continue }
		freq[d]--
		var distinct, duplicates int = 0, 0
		for digit := 0; digit < 10; digit++ {
			if freq[digit] > 0 {
				distinct++
				if freq[digit] >= 2 { duplicates++ }
			}
		}
		overall += distinct * (distinct - 1) + duplicates
		if freq[0] > 0 {
			var invalid int = distinct - 1
			if freq[0] >= 2 { invalid++ }
			overall -= invalid
		}
		freq[d]++
	}
	return overall
}