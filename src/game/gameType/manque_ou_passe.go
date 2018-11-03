package gameType

import (
	"fmt"
	"util"
)

func ManqueOuPasse() bool {
	fmt.Println("ManqueOuPasse was called")

	bet := util.ConvertToInteger(util.ConsolePrompt("{} EUR on (1)Manque or (2)Passe?: "))
	spin := util.SpinRoulette()

	if bet == 1 || bet == 2 {
		if bet == 1 && util.InBetween(spin, 1, 18) {
			return true
		}
		if bet == 2 && util.InBetween(spin, 19, 36) {
			return true
		}
		return false
	} else {
		fmt.Println("Invalid choice, please select (1)Manque or (2)Passe.")
		return ManqueOuPasse()
	}
	return false
}
