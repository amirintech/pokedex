package main

import (
	"bufio"
	"fmt"
	"os"
)

func repl() {
	fmt.Println("Welcome to Pokedex CLI!")

	for {
		fmt.Print("> ")

		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		cmdName := sc.Text()

		cmd := getCommand(cmdName)
		if cmd == nil {
			fmt.Printf("Command \"%s\" is not supported.\n", cmdName)
			fmt.Printf("Use \"help\" to view supported commands.\n")
		} else {
			cmd.execute()
		}
	}
}
