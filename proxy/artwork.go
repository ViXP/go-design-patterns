package proxy

import (
	"fmt"
	"os"
	"strconv"
)

// Artwork is the original heavy data model that represents the painting.
type Artwork struct {
	title       string
	author      string
	year        uint16
	description string
	data        []byte
}

// Title returns the commonly known name of the artwork.
func (a *Artwork) Title() string {
	return a.title
}

// Author returns the name of the author.
func (a *Artwork) Author() string {
	return a.author
}

// Year returns the year of creation.
func (a *Artwork) Year() string {
	return strconv.Itoa(int(a.year))
}

// Description returns a short textual explanation of the artwork.
func (a *Artwork) Description() string {
	return a.description
}

// LoadData reads the image data from filesystem and stores it into the internal data of the artwork.
func (a *Artwork) LoadData() []byte {
	if len(a.data) == 0 {
		rawData, err := os.ReadFile(a.title + ".png")
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return nil
		}
		a.data = rawData
	}

	return a.data
}

// String returns the aggregated human readable metadata of the artwork.
func (a *Artwork) String() string {
	return fmt.Sprintf(
		"The picture: %s\nBy: %s\nYear: %s\nDescription: %s\n", a.Title(), a.Author(), a.Year(), a.Description(),
	)
}

// Interfaces implementation assertion.
var (
	_ ArtDescriber = &Artwork{}
	_ fmt.Stringer = &Artwork{}
)
