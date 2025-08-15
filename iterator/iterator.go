// Package iterator implements the Iterator design pattern.
// The Jukebox aggregates a collection of vinyl Records.
// Its Load method returns a RecordsIterator, which can traverse through collection manually using the Next and Prev
// methods, or sequentially using a goroutine and channel via the Sequentially method.
package iterator

import "fmt"

func Run() {
	// Defining the records
	records := []*Record{
		NewRecord("Michael Jackson", "Thriller", 42),
		NewRecord("Queen", "Greatest Hits", 58),
		NewRecord("The Beatles", "Abbey Road", 183),
		NewRecord("Prince", "Purple Rain", 44),
	}
	// Setting up the jukebox
	jukebox := NewJukebox(records)

	// Constructing the records iterator
	sequencer := jukebox.Load()

	// Play the records one by one
	for {
		r := sequencer.Next()
		if r == nil {
			break
		}
		fmt.Print(r)
		r.Play()
	}

	fmt.Println()

	// Play the records with channel based sequential iterator
	for r := range sequencer.Sequentially() {
		fmt.Print(r)
		r.Play()
	}
}
