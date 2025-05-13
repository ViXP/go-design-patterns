package flyweight

// TileFactory implements the flyweight factory for the Tile construction (or the reusage of the existing Tiles).
type TileFactory struct {
	tiles []*Tile
}

// GetTileType returns the existing Tile, or creates the new Tile based on the provided character symbol.
func (tf *TileFactory) GetTileType(character rune) *Tile {
	for _, tile := range tf.tiles {
		if tile != nil && tile.Character == character {
			return tile
		}
	}
	newTileType := &Tile{character}
	tf.tiles = append(tf.tiles, newTileType)
	return newTileType
}

// NewTileFactory constructs TileFactory.
func NewTileFactory() *TileFactory {
	return &TileFactory{}
}
