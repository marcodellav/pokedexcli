package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	if _, ok := cfg.caughtPokemon[name]; !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	pokemon := cfg.caughtPokemon[name]

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	fmt.Println("  -hp:", pokemon.Stats[0].BaseStat)
	fmt.Println("  -attack:", pokemon.Stats[1].BaseStat)
	fmt.Println("  -defense:", pokemon.Stats[2].BaseStat)
	fmt.Println("  -special-attack:", pokemon.Stats[3].BaseStat)
	fmt.Println("  -special-defense:", pokemon.Stats[4].BaseStat)
	fmt.Println("  -speed:", pokemon.Stats[5].BaseStat)
	fmt.Println("Types:")
	for t := range pokemon.Types {
		fmt.Println("  -type:", pokemon.Types[t].Type.Name)
	}
	return nil
}

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for p := range cfg.caughtPokemon {
		fmt.Printf("- %s\n", cfg.caughtPokemon[p].Name)
	}
	return nil
}
