package main

import (
	"fmt"

	"github.com/c-bata/go-prompt"
)

func main() {
	fmt.Printf("Easy DNS CLI\n\n")
	t := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(">>> "),
		prompt.OptionInputTextColor(prompt.Yellow),
	)
	t.Run()

}
