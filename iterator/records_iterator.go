package iterator

// RecordsIterator represents the iterator that allows to traverse through jukebox collection manually or sequentially.
type RecordsIterator struct {
	currentRec int
	jukebox    *Jukebox
}

// Next returns the next Record in jukebox's collection and stops previously played record.
func (ri *RecordsIterator) Next() *Record {
	if ri.currentRec > -1 {
		ri.jukebox.Records[ri.currentRec].Stop()
	}
	if ri.currentRec+1 < len(ri.jukebox.Records) {
		ri.currentRec += 1
		return ri.jukebox.Records[ri.currentRec]
	} else {
		return nil
	}
}

// Prev returns the previous Record in jukebox's collection and stops current record.
func (ri *RecordsIterator) Prev() *Record {
	if ri.currentRec > -1 {
		ri.jukebox.Records[ri.currentRec].Stop()
	}
	if ri.currentRec > 0 {
		ri.currentRec -= 1
		return ri.jukebox.Records[ri.currentRec]
	} else {
		return nil
	}
}

// Sequentially returns the channel of Records, that allows to traverse through collection sequentially.
func (ri *RecordsIterator) Sequentially() <-chan *Record {
	sequence := make(chan *Record)

	go func(s chan<- *Record, iter *RecordsIterator) {
		defer close(s)

		for i, r := range iter.jukebox.Records {
			if iter.currentRec > -1 {
				iter.jukebox.Records[iter.currentRec].Stop()
			}
			iter.currentRec = i
			s <- r
		}
	}(sequence, ri)

	return sequence
}

// NewRecordsIterator constructs the RecordsIterator.
func NewRecordsIterator(juke *Jukebox) *RecordsIterator {
	return &RecordsIterator{-1, juke}
}
