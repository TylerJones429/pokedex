package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Pokemon in pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
