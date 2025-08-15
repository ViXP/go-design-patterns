package iterator

import "fmt"

// Player is a general interface that defines the playing abilities.
type Player interface {
	Stop()
	Play()
}

// Record represents the vinyl record that could be played separately of within Jukebox.
type Record struct {
	artist   string
	title    string
	duration int
	Playing  bool
}

// Stop sets Record's playing state to false.
func (r *Record) Stop() {
	r.Playing = false
}

// Play sets Record's playing state to true.
func (r *Record) Play() {
	r.Playing = true
}

// String formats Record information in human readable format.
func (r *Record) String() string {
	return fmt.Sprintf("%s: %s (length: %v)\n", r.artist, r.title, r.duration)
}

// NewRecord construct the new Record.
func NewRecord(artist, title string, duration int) *Record {
	return &Record{artist, title, duration, false}
}

// Interfaces implementation assertion.
var (
	_ fmt.Stringer = &Record{}
	_ Player       = &Record{}
)
