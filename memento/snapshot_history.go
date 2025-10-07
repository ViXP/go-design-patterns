package memento

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

// SnapshotHistory is the concrete Caretaker that manages Snapshot mementos — aggregating, serializing, and restoring
// them as needed.
type SnapshotHistory struct {
	snapshots []Snapshot
}

// String implements Stringer interface for SnapshotHistory.
func (s *SnapshotHistory) String() string {
	var builder strings.Builder

	for i, shot := range s.snapshots {
		fmt.Fprintf(&builder, "Snapshot #%v: %v\n", i, shot.Serialized[len(shot.Serialized)-32:])
	}
	return builder.String()
}

// Load deserializes and returns the state from the specific Snapshot memento by index from the internal storage.
func (s *SnapshotHistory) Load(index int) *Forecast {
	if index < 0 || index >= len(s.snapshots) {
		panic(fmt.Sprintf("snapshot index %d out of range", index))
	}

	var forecast Forecast
	var snapshot Snapshot = s.snapshots[index]

	decoded, err := base64.StdEncoding.DecodeString(snapshot.Serialized)

	if err != nil {
		panic("snapshot decoding error!")
	}

	serErr := json.Unmarshal(decoded, &forecast)

	if serErr != nil {
		panic("deserialization error!")
	}

	return &forecast
}

// Save serializes and stores the specific Forecast state into the Snapshot memento inside the internal storage.
func (s *SnapshotHistory) Save(forecast *Forecast) {
	serialized, err := json.Marshal(forecast)
	if err != nil {
		panic("wrong snapshot data!")
	}
	s.snapshots = append(s.snapshots, *NewSnapshot(base64.StdEncoding.EncodeToString(serialized)))
}

// Interfaces implementation assertion.
var (
	_ Caretaker    = &SnapshotHistory{}
	_ fmt.Stringer = &SnapshotHistory{}
)
