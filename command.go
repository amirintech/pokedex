package main

import (
	"fmt"
	"os"
)

type command struct {
	name        string
	description string
	execute     func() error
}

var commands map[string]command

func init() {
	commands = map[string]command{
		"help": {
			name:        "help",
			description: "prints this menu",
			execute:     handleHelp,
		},
		"exit": {
			name:        "exit",
			description: "exits pokedex CLI",
			execute:     handleExit,
		},
	}
}

func getCommand(name string) *command {
	cmd, ok := commands[name]
	if !ok {
		return nil
	}

	return &cmd
}

func handleHelp() error {
	fmt.Println("  help:")
	fmt.Printf("    %s\n\n", commands["help"].description)

	for k, v := range commands {
		if k != "help" {
			fmt.Printf("  %s:\n", k)
			fmt.Printf("    %s\n\n", v.description)
		}
	}

	return nil
}

func handleExit() error {
	fmt.Println("See you!")
	os.Exit(0)
	return nil
}
