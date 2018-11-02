package game

import (
	"bufio"
	"fmt"
	"os"
)

func ConsolePrompt(sentence string) string {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(sentence)
	text, _ := reader.ReadString('\n')

	if len(text) == 1 {
		return ConsolePrompt(sentence)
	}

	text2 := text[0 : len(text)-1]
	return text2

}
