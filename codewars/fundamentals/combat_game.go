package fundamentals

func Combat(health, damage float64) float64 {
	if damage > health {
		return 0.0
	}

	return health - damage
}
