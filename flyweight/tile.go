package flyweight

import "fmt"

// Tile is a flyweight object that represents renderable item on the map.
type Tile struct {
	Character rune
}

// Render implements the ability of Tile to be rendered.
func (t *Tile) Render() {
	fmt.Print(string(t.Character))
}

// Interface implementation assertion.
var _ Renderer = &Tile{}
