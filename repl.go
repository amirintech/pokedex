package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type config struct {
	next     *string
	previous *string
	cache    *cache
}

func repl() {
	conf := &config{
		cache: newCache(time.Minute * 10),
	}
	fmt.Println("Welcome to Pokedex CLI!")

	for {
		fmt.Print("> ")

		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		cmdName := strings.TrimSpace(sc.Text())

		cmd := getCommand(cmdName)
		if cmd == nil {
			fmt.Printf("Command \"%s\" is not supported.\n", cmdName)
			fmt.Printf("Use \"help\" to view supported commands.\n")
			continue
		}

		if err := cmd.execute(conf); err != nil {
			fmt.Printf("falied to execute command \"%s\"\n", cmdName)
			fmt.Println(err.Error())
		}
	}
}
