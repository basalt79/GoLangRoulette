package game

import (
	"fmt"
	"game/gameType"
	"util"
)

//type Game struct {
//	money int
//	player string
//
//}

func Start(player string, money int) int {
	fmt.Println(player, money)

	var funcMap = map[string]interface{}{
		"1": gameType.Straight,
	}

	for {
		if money == 0 {
			fmt.Println("Game is Over for.")
		} else {
			betType := util.ConvertToInteger(ConsolePrompt("Choose your Bet Type: "))
			fmt.Println(betType)
			betMoney := util.ConvertToInteger(ConsolePrompt("Choose your Money to bring in: "))

			if betMoney == 0 || betMoney > money {
				fmt.Println("Invalid Bet.")
			}

			funcMap["1"].(func())()
		}

	}

	return money

}
