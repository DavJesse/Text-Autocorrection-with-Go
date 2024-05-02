package main

func recursiveCleanUp(strSlce []string) []string {
	keyWords := []string{"(hex)", "(bin)", "(cap)", "(cap,", "(up)", "(up,", "(low)", "(low,"}
	for i := 0; i < len(strSlce); i++ {
		for _, word := range keyWords {
			if strSlce[i] == word {
				strSlce = cap(strSlce) // Solve "(cap)" & "(cap, int)"
				strSlce = low(strSlce) // Solve "(low)" & "(low, int)"
				strSlce = up(strSlce)  // Solve "(up)" & "(up, int)"
				strSlce = hex(strSlce) // Solve "(hex)"
				strSlce = bin(strSlce)
			}
		}
	}
	return strSlce
}
