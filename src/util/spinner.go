package util

import (
	"fmt"
	"math/rand"
)

func SpinRoulette() int {
	fmt.Println("###################################")
	fmt.Println("Roulette spinning....\nand it is: ")
	min := 0
	max := 36
	result := rand.Intn(max-min) + min
	fmt.Println(result)
	fmt.Println("###################################")
	return result
}
