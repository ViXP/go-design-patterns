package memento

import "fmt"

// Precipitation is the stringifiable part of the Forecast.
type Precipitation struct {
	State byte
}

// Precipitation state enums.
const (
	Clear byte = iota
	Cloudy
	Rain
	ThunderStorm
	Snow
	SnowStorm
)

// String implements Stringer interface for the Precipitation.
func (p *Precipitation) String() string {
	switch p.State {
	case Cloudy:
		return "Cloudy"
	case Rain:
		return "Raining"
	case ThunderStorm:
		return "Thunderstorm"
	case Snow:
		return "Snowing"
	case SnowStorm:
		return "Blizzard"
	default:
		return "Clear"
	}
}

// Forecast is the changeable state, that represents the specific weather forecast (and is the subject of memento).
type Forecast struct {
	Temperature   float32
	Humidity      float32
	WindSpeed     float32
	Pressure      int
	Precipitation Precipitation
}

// String implements Stringer interface for the Forecast.
func (f *Forecast) String() string {
	return fmt.Sprintf(
		"Temperature: %.2f°C\nHumidity level: %.1f%%\nSpeed of wind: %.2fkm/h\nAtmospheric pressure: %d millibars\nPrecipitation: %s\n",
		f.Temperature, f.Humidity, f.WindSpeed, f.Pressure, f.Precipitation.String())
}

// Interfaces implementation assertion.
var (
	_ fmt.Stringer = &Forecast{}
	_ fmt.Stringer = &Precipitation{}
)
