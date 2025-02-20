package repl

import (
	"fmt"
	"shhh/pkg/command"
	"shhh/pkg/prompt"
)

func Loop() {
	for {
		fmt.Printf(prompt.Get())

		raw := command.Read()

		fmt.Printf("command: %s\n", raw)
	}
}
