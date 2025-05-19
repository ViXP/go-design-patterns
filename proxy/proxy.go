// Package proxy implements the Proxy design pattern.
// ArtworkProxy is a lightweight stand-in for the real Artwork object. It implements the shared ArtDescriber interface
// and provides access to basic metadata such as title, author, year, and description. The proxy defers loading the
// heavy scan data until it is explicitly requested, delegating that responsibility to the underlying Artwork instance
// (lazy loading).
package proxy

import "fmt"

func Run() {
	painting := ArtworkProxy{
		title: "La Mappa dell'Inferno", year: 1481, author: "Sandro Botticelli",
		description: "A detailed vertical map of Hell, based on Dante's Inferno",
	}

	// Print the proxy data of the painting
	fmt.Print(&painting)

	// Loading the data from file and storing it into the original object
	painting.LoadData()
}
