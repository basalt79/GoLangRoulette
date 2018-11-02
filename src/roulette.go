package main

import (
	"fmt"
	"game"
	"util"
)

func main() {
	fmt.Println("I'm going to be a ROULETTE!")

	name := game.ConsolePrompt("Type your name: ")
	fmt.Println(name)

	money := game.ConsolePrompt("How much Money you want to bring in: ")
	geld := util.ConvertToInteger(money)

	game := game.Start(name, geld)
	println(game)
}
