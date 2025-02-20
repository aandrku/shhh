package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	prompt := "fancy prompt: "

	for {
		fmt.Printf(prompt)

		raw, err := r.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occured while trytin to read the command: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("your typed %s", raw)

	}

}
