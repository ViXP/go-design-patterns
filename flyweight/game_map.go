package flyweight

import "fmt"

// GameMap represents the map of the game that could be rendered.
type GameMap struct {
	width, height uint8
	tiles         [][]*Tile
}

// Render implements the ability of the GameMap to be rendered.
func (gm *GameMap) Render() {
	for i := range gm.width {
		for j := range gm.height {
			gm.tiles[i][j].Render()
		}
		fmt.Printf("\n")
	}
}

// NewGameMap is the GameMap constructor based on the provided boundaries and the list of Tiles.
func NewGameMap(width, height uint8, tiles ...*Tile) *GameMap {
	remappedTiles := make([][]*Tile, height)

	var tileIndex uint8 = 0
	for i := range height {
		remappedTiles[i] = make([]*Tile, width)

		for j := range width {
			remappedTiles[i][j] = tiles[tileIndex]
			tileIndex++
		}
	}

	return &GameMap{width: width, height: height, tiles: remappedTiles}
}

// Interface implementation assertion.
var _ Renderer = &GameMap{}
