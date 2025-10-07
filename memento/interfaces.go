package memento

// ForecastProvider defines a common interface for components that can supply weather forecast data.
type ForecastProvider interface {
	Retrieve() *Forecast
}

// Snapshoter defines the ability to create and restore snapshots of internal state.
type Snapshoter interface {
	TakeSnapshot()
	RestoreSnapshot(index int)
}

// Originator represents an entity whose state can be saved and restored using snapshots (mementos).
type Originator interface {
	Update(*Forecast)
	Print()
	Snapshoter
	ForecastProvider
}

// Caretaker defines a common interface for objects responsible for storing and restoring snapshots (mementos).
type Caretaker interface {
	Save(forecast *Forecast)
	Load(index int) *Forecast
}
