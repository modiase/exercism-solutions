package jedlik

import (
    "fmt"
)

// Drive the car one step
func (c* Car) Drive() {
    if (c.battery < c.batteryDrain){
        return
    }
    c.battery -= c.batteryDrain
    c.distance += c.speed
}

// Show the current distance travelled
func (c* Car) DisplayDistance() string {
    return fmt.Sprintf("Driven %d meters", c.distance)
}

// Show the current remaining battery
func (c* Car) DisplayBattery() string{
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

// Returns true if the car can complete the track; otherwise, false.
func (c* Car) CanFinish(trackDistance int) bool {
    return ((trackDistance + c.speed - 1) / c.speed) <= (c.battery / c.batteryDrain)
}
