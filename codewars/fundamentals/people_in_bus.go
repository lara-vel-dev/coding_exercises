package fundamentals

func Number(busStops [][2]int) int {
	onBus, offBus := 0, 0

	for _, people := range busStops {
		onBus += people[0]
		offBus += people[1]
	}

	return onBus - offBus
}
