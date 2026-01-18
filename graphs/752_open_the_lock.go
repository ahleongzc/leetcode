package graphs

type lock string

func (l lock) decrease(index int) lock {
	lockCopy := []byte(l)
	if lockCopy[index] == '0' {
		lockCopy[index] = '9'
	} else {
		lockCopy[index]--
	}
	return lock(lockCopy)
}

func (l lock) increase(index int) lock {
	lockCopy := []byte(l)
	if lockCopy[index] == '9' {
		lockCopy[index] = '0'
	} else {
		lockCopy[index]++
	}
	return lock(lockCopy)
}

func openLock(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}

	explored := make(map[lock]struct{})
	invalid := make(map[lock]struct{})

	for _, deadend := range deadends {
		invalid[lock(deadend)] = struct{}{}
	}

	if _, ok := invalid["0000"]; ok {
		return -1
	}

	queue := []lock{}
	queue = append(queue, "0000")
	explored["0000"] = struct{}{}

	turns := 0

	for len(queue) != 0 {
		currLevelLength := len(queue)

		for _ = range currLevelLength {
			curr := queue[0]
			queue = queue[1:]

			for i := range 4 {
				increasedLock := curr.increase(i)
				decreasedLock := curr.decrease(i)

				if string(increasedLock) == target {
					return turns + 1
				}

				if string(decreasedLock) == target {
					return turns + 1
				}

				_, isIncreasedInvalid := invalid[increasedLock]
				_, isIncreasedExplored := explored[increasedLock]

				_, isDecreasedInvalid := invalid[decreasedLock]
				_, isDecreasedExplored := explored[decreasedLock]

				if !isIncreasedInvalid && !isIncreasedExplored {
					queue = append(queue, increasedLock)
					explored[increasedLock] = struct{}{}
				}

				if !isDecreasedInvalid && !isDecreasedExplored {
					queue = append(queue, decreasedLock)
					explored[decreasedLock] = struct{}{}
				}
			}
		}

		turns++
	}

	return -1
}
