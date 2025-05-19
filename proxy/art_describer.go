package proxy

// ArtDescriber defines the common interface for accessing artwork metadata and data, implemented by both Artwork and
// ArtworkProxy types.
type ArtDescriber interface {
	Author() string
	Title() string
	Year() string
	Description() string
	LoadData() []byte
}
