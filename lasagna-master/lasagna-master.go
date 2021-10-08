package lasagna

func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		timePerLayer = 2
	}
	pt := len(layers) * timePerLayer
	return pt
}

func Quantities(layers []string) (int, float64) {
	var n int
	var s float64
	for _, l := range layers {
		switch l {
		case "noodles":
			n += 50
		case "sauce":
			s += 0.2
		}
	}
	return n, s
}

func AddSecretIngredient(friendsList []string, myList []string) []string {
	secretIndex := len(friendsList) - 1
	myList = append(myList, friendsList[secretIndex])
	return myList
}

func ScaleRecipe(quantities []float64, numberOfPortions int) []float64 {
	scaledQuantities := make([]float64, len(quantities))
	extraAmount := float64(numberOfPortions) / 2
	for i := range scaledQuantities {
		scaledQuantities[i] = quantities[i] * extraAmount
	}
	return scaledQuantities
}
