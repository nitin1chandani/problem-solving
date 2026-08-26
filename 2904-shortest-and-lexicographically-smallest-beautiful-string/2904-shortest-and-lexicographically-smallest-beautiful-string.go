func countOf1(s string) int {
	count := 0
	for _, v := range s {
		if v == '1' {
			count++
		}
	}

	return count
}

func shortestBeautifulSubstring(s string, k int) string {
	n := len(s)

	for i := k; i <= n; i++ {
		sbs := ""
		for j := 0; j <= n-i; j++ {
			curr := s[j : j+i]
			if sbs == "" && countOf1(curr) == k {
				sbs = curr
				if len(sbs) == k {
					return sbs
				}
			}

			if curr < sbs && countOf1(curr) == k {
				sbs = curr
			}
		}
		if sbs != "" {
			return sbs
		}
	}
    return ""
}