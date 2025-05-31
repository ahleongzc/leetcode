package binarysearch

func MinEatingSpeed(piles []int, h int) int {

	var timeTakenToFinishBananas = func(piles []int, speed int) int {
		time := 0
		for _, pile := range piles {
			time += pile / speed
			if pile%speed != 0 {
				time += 1
			}
		}

		return time
	}

	startSpeed := 1
	endSpeed := 0

	for _, pile := range piles {
		endSpeed = max(endSpeed, pile)
	}

	minEatingSpeed := endSpeed

	for startSpeed <= endSpeed {
		speed := (startSpeed + endSpeed) / 2
		time := timeTakenToFinishBananas(piles, speed)

		if time <= h {
			minEatingSpeed = speed
			endSpeed = speed - 1
		} else {
			startSpeed = speed + 1
		}
	}

	return minEatingSpeed
}
