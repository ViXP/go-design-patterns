package memento

import "fmt"

// WeatherMonitor is the specific Originator, responsable for the retrieval of the current weather Forecast from the
// external forecast providers. It is also allowed to Print the current Forecast, save Forecast snapshot through the
// Caretaker and restore the Forecast state from the snapshot.
type WeatherMonitor struct {
	currentMeasures *Forecast
	provider        ForecastProvider
	history         Caretaker
}

// Update mutates the current forecast in the Weather Monitor.
func (w *WeatherMonitor) Update(forecast *Forecast) {
	w.currentMeasures = forecast
}

// Retrieve is fetching the current weather forecast from the external ForecastProvider and stores it as the current
// Forecast state.
func (w *WeatherMonitor) Retrieve() *Forecast {
	w.Update(w.provider.Retrieve())
	return w.currentMeasures
}

// RestoreSnapshot updates the current Forecast from the Snapshot, retrieved from the Caretaker.
func (w *WeatherMonitor) RestoreSnapshot(index int) {
	w.Update(w.history.Load(index))
}

// TakeSnapshot saves current Forecast as the new Snapshot memento, with help of the Caretaker.
func (w *WeatherMonitor) TakeSnapshot() {
	w.history.Save(w.currentMeasures)
}

// Print allows to print the current forecast state and the Caretaker state to screen.
func (w *WeatherMonitor) Print() {
	fmt.Println("---")
	fmt.Println("Current weather forecast:")
	fmt.Println("---")
	fmt.Print(w.currentMeasures)
	fmt.Println("---")
	fmt.Print(w.history)
}

// NewWeatherMonitor constructs the new WeatherMonitor structure.
func NewWeatherMonitor(provider ForecastProvider) *WeatherMonitor {
	return &WeatherMonitor{nil, provider, &SnapshotHistory{}}
}

// Interfaces implementation assertion.
var _ Originator = &WeatherMonitor{}
