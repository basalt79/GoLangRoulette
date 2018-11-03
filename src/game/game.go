package game

import (
	"fmt"
	"game/gameType"
	"util"
)

func Start(player string, money int) int {
	fmt.Println(player, money)

	var funcMap = map[int]interface{}{
		1: gameType.Straight,
		2: gameType.ManqueOuPasse,
	}

	for {
		if money == 0 {
			fmt.Println("Game is Over for.")
			break
		} else {
			betType := util.ConvertToInteger(util.ConsolePrompt("Choose your Bet Type: "))
			fmt.Println(betType)
			betMoney := util.ConvertToInteger(util.ConsolePrompt("Choose your Money to bring in: "))

			if betMoney == 0 || betMoney > money {
				fmt.Println("Invalid Bet.")
			}

			if funcMap[betType].(func() bool)() {
				amountToAdd := win(betMoney, betType)
				fmt.Println(fmt.Sprintf("You won, your win <%v> will be added to <%v>.", amountToAdd, money))
				money += betMoney
				fmt.Println(fmt.Sprintf("You have now <%v>. money", money))
			} else {
				fmt.Println(fmt.Sprintf("You loose, your bet <%v> will be reduced from <%v>.", betMoney, money))
				money -= betMoney
				fmt.Println(fmt.Sprintf("You have now <%v>. money", money))
			}
		}

	}

	return money

}

func win(amount int, betType int) int {
	if betType == 1 {
		return amount * 35 // straight
	}
	if util.InBetween(betType, 2, 4) {
		return amount // manque/passe - red/black - even/odd
	}
	return amount * 2 // column - dozen bets
}
