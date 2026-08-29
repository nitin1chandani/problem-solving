package main

import "fmt"

func lexGreaterPermutation(s string, target string) string {
	count := make([]int, 26)

	for _, ch := range s {
		count[ch-'a']++
	}

	result := make([]rune, 0, len(s))

	// Match target as long as possible.
	for _, ch := range target {
		idx := ch - 'a'

		if count[idx] == 0 {
			break
		}

		result = append(result, ch)
		count[idx]--
	}

	// We could not match the entire target.
	// Try making the current position greater.
	if len(result) < len(target) {
		pos := len(result)
		targetIdx := target[pos] - 'a'

		for i := targetIdx + 1; i < 26; i++ {
			if count[i] > 0 {
				result = append(result, rune(i)+'a')
				count[i]--

				// Fill remaining characters in ascending order.
				for j := 0; j < 26; j++ {
					for count[j] > 0 {
						result = append(result, rune(j)+'a')
						count[j]--
					}
				}

				return string(result)
			}
		}
	}

	// Entire target was matched.
	// Backtrack and increase an earlier character.
	for i := len(result) - 1; i >= 0; i-- {
		// Put the current character back.
		count[result[i]-'a']++

		// Find the smallest character greater than it.
		for j := result[i] - 'a' + 1; j < 26; j++ {
			if count[j] > 0 {
				result[i] = rune(j) + 'a'
				count[j]--

				// Remove the old suffix.
				result = result[:i+1]

				// Fill suffix with smallest available characters.
				for k := 0; k < 26; k++ {
					for count[k] > 0 {
						result = append(result, rune(k)+'a')
						count[k]--
					}
				}

				return string(result)
			}
		}
	}

	return ""
}

