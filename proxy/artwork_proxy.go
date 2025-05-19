package proxy

import (
	"fmt"
	"strconv"
)

// ArtworkProxy is the lightweight proxy of the original Artwork.
type ArtworkProxy struct {
	original    *Artwork
	title       string
	author      string
	year        uint16
	description string
}

// Title returns the commonly known name of the artwork.
func (ap *ArtworkProxy) Title() string {
	return ap.title
}

// Author returns the name of the author.
func (ap *ArtworkProxy) Author() string {
	return ap.author
}

// Year returns the year of creation.
func (ap *ArtworkProxy) Year() string {
	return strconv.Itoa(int(ap.year))
}

// Description returns a short textual explanation of the artwork.
func (ap *ArtworkProxy) Description() string {
	return ap.description
}

// LoadData delegates the heavy data loading process to the original artwork.
func (ap *ArtworkProxy) LoadData() []byte {
	if ap.original == nil {
		ap.original = &Artwork{title: ap.title, author: ap.author, year: ap.year, description: ap.description}
	}
	return ap.original.LoadData()
}

// String returns the aggregated human readable metadata of the artwork.
func (ap *ArtworkProxy) String() string {
	return fmt.Sprintf(
		"The picture: %s\nBy: %s\nYear: %s\nDescription: %s\n", ap.Title(), ap.Author(), ap.Year(), ap.Description(),
	)
}

// Interfaces implementation assertion.
var (
	_ ArtDescriber = &ArtworkProxy{}
	_ fmt.Stringer = &ArtworkProxy{}
)
