package gameType

import (
	"fmt"
	"util"
)

func Straight() bool {
	fmt.Println("Straight was called")
	bet := util.ConvertToInteger(util.ConsolePrompt("Choose your number: "))
	if bet >= 0 && bet <= 36 {
		fmt.Println(fmt.Sprintf("You choose <%v> no more bets please!", bet))
		spin := util.SpinRoulette()
		return spin == bet
	} else {
		fmt.Println("Invalid Number, please choose a number between 0 and 36.")
		return Straight()
	}
	return false
}
