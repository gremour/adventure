package game

import (
	"fmt"
	"strings"
)

// Location ...
type Location struct {
	Title       string
	Description string
	Exits       map[Dir]string
}

// Dir ...
type Dir int

const (
	// Here ...
	Here Dir = iota
	// North ...
	North Dir = iota
	// South ...
	South Dir = iota
	// East ...
	East Dir = iota
	// West ...
	West Dir = iota
)

var dirToText = map[Dir]string{
	Here:  "здесе",
	North: "севере",
	South: "юге",
	East:  "востоке",
	West:  "западе",
}

// Show ...
func (l *Location) Show() {
	fmt.Println("===", l.Title, "===")
	fmt.Println(l.Description)
	exits := make([]string, 0)
	for dir, exit := range l.Exits {
		locName := g.Locs[exit].Title
		exits = append(exits, fmt.Sprintf("%v на %v", locName, dirToText[dir]))
	}
	fmt.Printf("Выходы: %v.\n", strings.Join(exits, ", "))
}
