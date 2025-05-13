// Package flyweight implement the Flyweight design pattern.
// The flyweight object Tile is used by the TileFactory to optimize the memory usage for the storing of the renderable
// Tile items. It encapsulates the logic to create new or return the existing Tile for the client to define and render
// the optimized GameMap.
package flyweight

func Run() {
	tilesStore := NewTileFactory()

	// Utilize flyweight pattern to create/reuse the Tiles
	tree := tilesStore.GetTileType('T')
	water := tilesStore.GetTileType('~')
	air := tilesStore.GetTileType(' ')
	ground := tilesStore.GetTileType('.')
	bridge := tilesStore.GetTileType('=')
	wall := tilesStore.GetTileType('#')
	roof := tilesStore.GetTileType('#')

	// Construct the map
	towerMap := NewGameMap(
		10, 10, air, air, air, air, air, air, air, air, air, air, air, roof, roof, roof, air, air, air, air, air, air,
		air, wall, wall, wall, air, air, air, wall, wall, air, air, wall, wall, wall, roof, roof, wall, wall, wall, air, air,
		wall, wall, wall, wall, wall, wall, wall, wall, air, air, wall, wall, wall, air, air, wall, wall, wall, air, air,
		wall, wall, wall, air, air, wall, wall, wall, air, water, water, water, water, bridge, bridge, water, water, water,
		water, ground, tree, ground, ground, bridge, bridge, ground, ground, ground, ground, ground, ground, ground,
		ground, ground, ground, ground, ground, tree, ground,
	)
	towerMap.Render()
}
