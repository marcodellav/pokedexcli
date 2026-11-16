package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/marcodellav/pokedexcli/internal/pokeapi"
)

func startRepl(c *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		commandArg := ""
		if len(words) > 1 {
			commandArg = words[1]
		}

		command, exists := c.commands[commandName]
		if exists {
			err := command.callback(c, commandArg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type config struct {
	commands      map[string]cliCommand
	URL           string
	Next          *string
	Previous      *string
	pokeapiClient pokeapi.Client
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
	args        string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
			args:        "",
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
			args:        "",
		},
		"map": {
			name:        "map",
			description: "Show 20 location areas",
			callback:    commandMap,
			args:        "",
		},
		"mapb": {
			name:        "mapb",
			description: "Show previous 20 location areas",
			callback:    commandMapb,
			args:        "",
		},
		"explore": {
			name:        "explore",
			description: "Show pokemon in this area",
			callback:    commandExplore,
			args:        "name or id",
		},
		"catch": {
			name:        "catch",
			description: "throw a Pokeball at a Pokemon",
			callback:    commandCatch,
			args:        "pokemon",
		},
	}
}
