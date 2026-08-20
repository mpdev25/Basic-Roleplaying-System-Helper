package dice
import (
	math/rand/v2
)

func diceRoller (count, sides int) {
	results := make([]int, count)
	for i := 0; i < count; i++ {
		results[i] = rand.Intn(sides) + 1
	}
	return results
}
