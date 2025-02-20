package main

import (
	"bufio"
	"fmt"
	"os"
	"shhh/pkg/prompt"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf(prompt.Get())

		_, err := r.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occured while trying to read the command: %v\n", err)
			os.Exit(1)
		}

	}

}
