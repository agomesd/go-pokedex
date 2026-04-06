package utils

func TryCatchPokemon(roll float64, baseExp int) bool {

	normalized := float64(baseExp) / float64(300)
	catchChance := 1.0 - float64(normalized)
	if catchChance > roll {
		return true
	} else {
		return false
	}
}
