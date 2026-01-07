package fundamentals

func Number(busStops [][2]int) (peopleLeft int) {
	onBus, offBus := 0, 0

	for _, people := range busStops {
		peopleLeft += people[0] - people[1]
	}

	return
}
