package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return (kind == "car") || (kind == "truck")
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    const trailer = " is clearly the better choice."
	if (option1 < option2){ 
        return option1 + trailer;
    }
    return option2 + trailer;
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if (age < 3){
        return 0.8 * originalPrice
    }
    if (age < 10){
        return 0.7 * originalPrice
    }
	return 0.5 * originalPrice
}
