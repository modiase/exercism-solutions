package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPreparationTime int) int {
    if avgPreparationTime == 0 {
        avgPreparationTime = 2
    }
    return len(layers) * avgPreparationTime;
}

const noodlesGramsPerLayer = 50
const sauceLitresPerLayer = 0.2
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodleLayerCount, sauceLayerCount := 0, 0
    for _, val := range(layers){
		switch val{
            case "noodles":
            	noodleLayerCount += 1
            case "sauce":
            	sauceLayerCount += 1
            default: {}
        }
    }
    return noodlesGramsPerLayer * noodleLayerCount, sauceLitresPerLayer * float64(sauceLayerCount)
    
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
    myList[len(myList)-1] = friendsList[len(friendsList)-1]
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) (scaled []float64) {
    scaled = make([]float64, len(quantities))
    for idx, val := range(quantities){
        scaled[idx] = val * float64(portions) / 2.0
    }
    return
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
