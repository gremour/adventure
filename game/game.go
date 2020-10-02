package game

import "fmt"

var g *Game

// Game ...
type Game struct {
	CurLoc *Location
	Locs   map[string]*Location
}

// New creates ready to use game.
func New() *Game {
	g = &Game{
		Locs: make(map[string]*Location),
	}

	g.Locs["meadow"] = &Location{
		Title:       "Поляна",
		Description: "Солнечная поляна со скамейкой.",
		Exits: map[Dir]string{
			North: "house",
			East:  "church",
		},
	}

	g.Locs["house"] = &Location{
		Title:       "Дом",
		Description: "Кирпичный дом с деревянной дверью.",
		Exits: map[Dir]string{
			South: "meadow",
		},
	}

	g.Locs["church"] = &Location{
		Title:       "Церковь",
		Description: "Покосившаяся деревянная церковь.",
		Exits: map[Dir]string{
			West: "meadow",
		},
	}

	g.CurLoc = g.Locs["meadow"]

	return g
}

// MoveNorth ...
func (g *Game) MoveNorth() {
	exit := g.CurLoc.Exits[North]
	if exit == "" {
		fmt.Println("Нет выхода в этом направлении.")
		return
	}
	loc := g.Locs[exit]
	g.CurLoc = loc
}

// MoveSouth ...
func (g *Game) MoveSouth() {
	exit := g.CurLoc.Exits[South]
	if exit == "" {
		fmt.Println("Нет выхода в этом направлении.")
		return
	}
	loc := g.Locs[exit]
	g.CurLoc = loc
}

// MoveEast ...
func (g *Game) MoveEast() {
	exit := g.CurLoc.Exits[East]
	if exit == "" {
		fmt.Println("Нет выхода в этом направлении.")
		return
	}
	loc := g.Locs[exit]
	g.CurLoc = loc
}

// MoveWest ...
func (g *Game) MoveWest() {
	exit := g.CurLoc.Exits[West]
	if exit == "" {
		fmt.Println("Нет выхода в этом направлении.")
		return
	}
	loc := g.Locs[exit]
	g.CurLoc = loc
}
