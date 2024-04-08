package main

import (
	"fmt"
	"os"
)

type command struct {
	name        string
	description string
	execute     func(conf *config) error
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
		"map": {
			name:        "map",
			description: "prints the the next page of locations",
			execute:     handleMap,
		},
		"mapb": {
			name:        "mapb",
			description: "prints the the previous page of locations",
			execute:     handleMapb,
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

func handleHelp(conf *config) error {
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

func handleExit(conf *config) error {
	fmt.Println("See you!")
	os.Exit(0)
	return nil
}

func handleMap(conf *config) error {
	return execMap(conf.next, conf)
}

func handleMapb(conf *config) error {
	return execMap(conf.previous, conf)
}

func execMap(url *string, conf *config) error {
	var locations []location
	if url == nil || len(locations) == 0 {
		res, err := getLocations(url)
		if err != nil {
			return err
		}
		conf.next = res.Next
		conf.previous = res.Previous
		locations = res.Results
		conf.cache.add(defaultUrl, locations)
	} else {
		locations = conf.cache.get(*url)
	}

	for _, loc := range locations {
		fmt.Printf("  - %s\n", loc.Name)
	}
	fmt.Println()

	return nil
}
