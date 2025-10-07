// Package memento implements the Memento design pattern.
// The Originator, WeatherMonitor, retrieves current weather forecasts from an external source (represented by
// RandomWeatherProvider, which implements the ForecastProvider interface) and can create a Snapshot (the Memento) of
// its current state by delegating to a Caretaker — the SnapshotHistory.
// The Caretaker is responsible for serializing and deserializing snapshots (mementos) from its internal storage and
// providing them back to the Originator when needed.
// The Snapshot remains opaque to the outside world due to base64-encoded state serialization.
package memento

import (
	"fmt"
	"strconv"
	"strings"
)

func Run() {
	// Data provider construction
	provider := RandomWeatherProvider{}

	// Originator construction
	monitor := NewWeatherMonitor(&provider)

	// First state retrieval
	monitor.Retrieve()
	monitor.Print()

	for {
		fmt.Println("---")
		fmt.Println("Commands: S - save forecast snapshot; R - restore forecast snapshot; Q - quit; N - get next forecast")
		fmt.Println("---")
		var command string
		fmt.Scan(&command)

		switch strings.ToUpper(command) {
		case "Q":
			return
		case "S":
			monitor.TakeSnapshot()
			monitor.Print()
		case "R":
			var index string
			fmt.Print("Which snapshot number? ")
			fmt.Scan(&index)
			i, err := strconv.Atoi(index)

			if err == nil {
				monitor.RestoreSnapshot(i)
				monitor.Print()
			}
		default:
			monitor.Retrieve()
			monitor.Print()
		}
	}
}
