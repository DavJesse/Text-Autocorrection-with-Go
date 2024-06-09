package main

func punct(arr []string) []string {
	for i := 0; i < len(arr); i++ {
		// As long as there's no empty string, move punctuation marks appearing at the beginning of the second and subsequent words
		if i > 0 && arr[i] != "" {
			if isPunct(arr[i][0]) {
				arr[i-1] += string(arr[i][0])
				arr[i] = arr[i][1:]

				// Purge empty strings
				if arr[i] == "" {
					arr = append(arr[:i], arr[i+1:]...)
					i--
				}
			}
		}
		// Re-run function in case the punctuation mark is misplaced in other areas of sentence
		if isPunct(arr[i][0]) {
			arr = punct(arr)
		}

	}
	return arr
}

func apostrophe(arr []string) []string {
	for i := 0; i < len(arr); i++ {
		// in words that end with a quotation mark, move the appostrophe to next word
		if isQuote(arr[i][len(arr[i])-1]) {
			if i != len(arr)-1 {
				emptSlc := string(arr[i][len(arr[i])-1])
				arr[i] = arr[i][:len(arr[i])-1]
				arr[i+1] = emptSlc + arr[i+1]
				break
			}
		}
	}
	return arr
}

func isPunct(ch byte) bool {
	punctMarks := ".,!?:;'\""

	for _, mark := range punctMarks {
		if rune(ch) == mark {
			return true
		}
	}
	return false
}

func isQuote(ch byte) bool {
	quoteMarks := "'\""

	for _, quote := range quoteMarks {
		if rune(ch) == quote {
			return true
		}
	}
	return false
}
