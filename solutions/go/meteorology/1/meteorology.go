package meteorology

import (
    "fmt"
)

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = iota
	Fahrenheit
)

// Add a String method to the TemperatureUnit type
func (u *TemperatureUnit) String() string{
	switch *u{
        case Celsius:
        	return "°C"
    	case Fahrenheit:
        	return "°F"
    	default:
        	panic("Not set")
    }
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}
func (t *Temperature) String() string{
    return fmt.Sprintf("%d %s", (*t).degree, (*t).unit.String())
}
// Add a String method to the Temperature type

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = iota
	MilesPerHour 
)

// Add a String method to SpeedUnit
func (u *SpeedUnit) String() string{
	switch *u{
        case KmPerHour:
        	return "km/h"
    	case MilesPerHour:
        	return "mph"
    	default:
        	panic("Not set")
    }
}
type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (s *Speed) String() string{
	return fmt.Sprintf("%d %s", (*s).magnitude, (*s).unit.String())
}
type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (m *MeteorologyData) String() string{
	return fmt.Sprintf("%s: %s, Wind %s at %s, %d%% Humidity", (*m).location, (*m).temperature.String(), (*m).windDirection, (*m).windSpeed.String(), (*m).humidity)
}