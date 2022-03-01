package lasagna

func PreparationTime(layers []string, avgTime int) int {
	if avgTime == 0 {
		return 2 * len(layers)
	}

	return avgTime * len(layers)
}

func Quantities(layers []string) (noodles int, sauce float64) {
	for _, value := range layers {
		switch value {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}

	return
}

func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len((myList))-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaled := make([]float64, len(quantities))

	for index, value := range quantities {
		scaled[index] = value * (float64(portions) / float64(2))
	}

	return scaled
}
