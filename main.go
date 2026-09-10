package main

func main() {
	s := "https://pokeapi.co/api/v2/location-area"
	c := &config{
		commands: getCommands(),
		URL:      &s,
		Next:     nil,
		Previous: nil,
	}
	startRepl(c)
}
