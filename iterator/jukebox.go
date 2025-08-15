package iterator

// Jukebox represents the aggregator of vinyl Records, that could Load them in RecordsIterator.
type Jukebox struct {
	Records []*Record
}

// Load returns the RecordIterator for the Jukebox.
func (j *Jukebox) Load() *RecordsIterator {
	return NewRecordsIterator(j)
}

// NewJukebox constructs the new Jukebox.
func NewJukebox(records []*Record) *Jukebox {
	return &Jukebox{records}
}
