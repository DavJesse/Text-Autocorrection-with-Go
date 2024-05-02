package main

func vowels(arr []string) []string {
	vowels := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
		'H': true,
		'h': true,
	}

	for i := 0; i < len(arr)-1; i++ {

		// Check for and correct errors following "A" and "a"
		if arr[i] == "A" || arr[i] == "a" {
			if i < len(arr)-1 {
				next := arr[i+1]

				if vowels[next[0]] || vowels[next[0]+32] {
					if arr[i] == "A" {
						arr[i] = "An"
					}
					if arr[i] == "a" {
						arr[i] = "an"
					}
				}
			}
		}

		// Check for and correct errors following "An" and "an"
		if arr[i] == "An" || arr[i] == "an" {
			if i < len(arr) {
				next := arr[i+1]

				if !(vowels[next[0]] || vowels[next[0]+32]) {
					if arr[i] == "An" {
						arr[i] = "A"
					}
					if arr[i] == "an" {
						arr[i] = "a"
					}
				}
			}
		}
	}
	return arr
}
