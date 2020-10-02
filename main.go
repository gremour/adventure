package main

import (
	"github.com/gremour/adventure/command"
	"github.com/gremour/adventure/game"
)

func main() {
	g := game.New()
	c := command.Commander{
		NotFoundPrompt: "Неизвестная команда, '?' для списка команд.",
		Handlers: map[string]func(){
			"с":      g.MoveNorth,
			"север":  g.MoveNorth,
			"ю":      g.MoveSouth,
			"юг":     g.MoveSouth,
			"в":      g.MoveEast,
			"восток": g.MoveEast,
			"з":      g.MoveWest,
			"запад":  g.MoveWest,
		},
	}

	for {
		g.CurLoc.Show()
		c.Input()
	}
}
