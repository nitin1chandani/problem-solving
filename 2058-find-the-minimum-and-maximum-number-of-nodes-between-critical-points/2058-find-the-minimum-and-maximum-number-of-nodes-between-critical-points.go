func nodesBetweenCriticalPoints(head *ListNode) []int {
	var prev *ListNode
	prev = nil

	var iterator *ListNode
	count := 0
	iterator = head

	list := make([]int, 0)

	for iterator != nil {
		if iterator == head {
			prev = iterator
			iterator = iterator.Next
			continue
		}

		if iterator.Next != nil {
			if iterator.Val > iterator.Next.Val && iterator.Val > prev.Val {
				list = append(list, count)
			}

			if iterator.Val < iterator.Next.Val && iterator.Val < prev.Val {
				list = append(list, count)
			}
		}

		prev = iterator
		iterator = iterator.Next
		count++
	}

	if len(list) < 2 {
		return []int{-1, -1}
	}

	minDist := list[1] - list[0]

	for i := 2; i < len(list); i++ {
		dist := list[i] - list[i-1]

		if dist < minDist {
			minDist = dist
		}
	}

	return []int{
		minDist,
		list[len(list)-1] - list[0],
	}
}