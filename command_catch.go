package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	const threshhold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	//fmt.Println(pokemon.BaseExperience, randNum, threshhold)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if randNum > threshhold {
		return fmt.Errorf("failed to catch %s", pokemonName)
	}

	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("%s was caught!\n", pokemonName)

	return nil
}
