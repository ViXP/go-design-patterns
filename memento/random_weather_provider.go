package memento

import "math/rand"

// RandomWeatherProvider implements ForecastProvider interface and generates random forecast.
type RandomWeatherProvider struct{}

// Retrieve creates the random weather Forecast.
func (r *RandomWeatherProvider) Retrieve() *Forecast {
	temp := -60.00 + rand.Float32()*120.00
	hum := rand.Float32() * 100.00
	wind := rand.Float32() * 407
	pres := rand.Intn(1083-870) + 870
	prec := Precipitation{byte(rand.Intn(5))}

	return &Forecast{temp, hum, wind, pres, prec}
}

// Interfaces implementation assertion.
var _ ForecastProvider = &RandomWeatherProvider{}
