package main

import (
	"fmt"
	"game"
	"util"
)

func main() {
	fmt.Println("I'm going to be a ROULETTE!")

	name := util.ConsolePrompt("Type your name: ")
	fmt.Println(name)

	money := util.ConsolePrompt("How much Money you want to bring in: ")
	geld := util.ConvertToInteger(money)

	game := game.Start(name, geld)
	println(game)
}
