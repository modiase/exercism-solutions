package speed
type Car struct {
	battery int
    batteryDrain int
    distance int
    speed int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{speed: speed, batteryDrain: batteryDrain, distance: 0, battery: 100 }
}

type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{ distance: distance }
}


func Drive(car Car) Car {
    if (car.battery >= car.batteryDrain){
			car.battery -= car.batteryDrain;
    		car.distance += car.speed;
        }
    return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	return ((track.distance - car.distance) * car.batteryDrain) / (car.speed) <= car.battery

}
