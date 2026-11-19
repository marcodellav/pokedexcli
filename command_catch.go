package main

import (
	"fmt"
	"time"

	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemonInfo(name)
	if err != nil {
		return err
	}

	// Base XP is 20 to 306
	baseXp := pokemon.BaseExperience
	fmt.Printf("%s Base XP: %d\n", name, baseXp)
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	x := int(r.Float64() * r.Float64() * 306)

	if x > baseXp {
		fmt.Printf("%s was caught!\n", name)
		cfg.caughtPokemon[pokemon.Name] = pokemon

	} else {
		fmt.Printf("%s escaped!\n", name)
	}
	return nil
}
