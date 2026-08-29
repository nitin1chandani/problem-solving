func buildMaxPalindrome(freq []int, mid rune, prefix string) string {
	left := prefix

	for i := 25; i >= 0; i-- {
		for freq[i] > 0 {
			left += string(rune('a' + i))
			freq[i]--
		}
	}

	right := ""
	for i := len(left) - 1; i >= 0; i-- {
		right += string(left[i])
	}

	if mid != 0 {
		return left + string(mid) + right
	}

	return left + right
}

func lexPalindromicPermutation(s string, target string) string {
	n := len(s)

	// Frequency of characters in s.
	freq := make([]int, 26)

	for _, ch := range s {
		freq[ch-'a']++
	}

	// Find middle character.
	var mid rune
	oddCount := 0

	for i := 0; i < 26; i++ {
		if freq[i]%2 == 1 {
			oddCount++
			mid = rune('a' + i)
		}

		if oddCount > 1 {
			return ""
		}
	}

	// Only half of each frequency is needed.
	for i := 0; i < 26; i++ {
		freq[i] /= 2
	}

	halfLen := n / 2

	/*
		We construct the left half greedily.

		For every position:
			a -> z

		Temporarily put that character in the prefix.

		Then ask:

		"What is the LARGEST palindrome I can make
		from this prefix?"

		If even that largest palindrome <= target,
		this prefix can never work.

		If it is > target, this character is safe.
	*/
	prefix := ""

	for pos := 0; pos < halfLen; pos++ {
		found := false

		for c := 0; c < 26; c++ {
			if freq[c] == 0 {
				continue
			}

			// Choose this character.
			freq[c]--
			prefix += string(rune('a' + c))

			// Make a copy because buildMaxPalindrome consumes freq.
			tempFreq := make([]int, 26)
			copy(tempFreq, freq)

			maxPalindrome := buildMaxPalindrome(
				tempFreq,
				mid,
				prefix,
			)

			if maxPalindrome > target {
				// This is the smallest character that can work.
				found = true
				break
			}

			// This character cannot possibly produce
			// a palindrome > target.
			prefix = prefix[:len(prefix)-1]
			freq[c]++
		}

		if !found {
			return ""
		}
	}

	// At this point prefix is the lexicographically smallest
	// valid left half.
	result := prefix

	right := ""
	for i := len(prefix) - 1; i >= 0; i-- {
		right += string(prefix[i])
	}

	if mid != 0 {
		result += string(mid)
	}

	result += right

	if result > target {
		return result
	}

	return ""
}