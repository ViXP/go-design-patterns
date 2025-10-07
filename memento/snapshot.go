package memento

// Snapshot is a serialized memento capturing the state of a Forecast.
type Snapshot struct {
	Serialized string
}

// NewSnapshot is the Snapshot constructor.
func NewSnapshot(serialized string) *Snapshot {
	return &Snapshot{serialized}
}
