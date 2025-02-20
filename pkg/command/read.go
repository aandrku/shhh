package command

import (
	"bufio"
	"fmt"
	"os"
	"shhh/pkg/prompt"
	"strings"
)

// Read function reads user command, trims redundant spaces
// and handles line continuation.
//
// It returns a string which is a command provided by the user.
func Read() string {
	r := bufio.NewReader(os.Stdin)
	builder := strings.Builder{}

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Printf("shhh encounterd an error while trying to process the command: %v", err)
		}

		line = strings.Trim(line, " \n")
		last := len(line) - 1

		if line[last] != '\\' {
			builder.WriteString(line)
			return builder.String()
		}

		line = line[:last]
		builder.WriteString(line)
		fmt.Printf(prompt.Get2())
	}

}
