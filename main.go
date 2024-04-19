package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// Extract arguments types in command line, if not possible, print error message
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Please Type Two Arguments")
		return
	}
	// Read from sample.txt and write to result.txt
	fileContent, _ := os.ReadFile(args[0])
	os.WriteFile(args[1], fileContent, 0o644)
	strSlce := strings.Fields(string(fileContent))

	strSlce = cap(strSlce)        // Solve "(cap)" & "(cap, int)"
	strSlce = low(strSlce)        // Solve "(low)" & "(low, int)"
	strSlce = up(strSlce)         // Solve "(up)" & "(up, int)"
	strSlce = hex(strSlce)        // Solve "(hex)"
	strSlce = bin(strSlce)        // Solve "(bin)"
	strSlce = punct(strSlce)      // solve punctuation marks
	strSlce = vowels(strSlce)     // Adjust articles "a" and "an" before vowels
	strSlce = apostrophe(strSlce) // solve single quotation marks

	// stitch string slice and write updated string to result.txt
	finalStr := strings.Join(strSlce, " ") + "\n"
	finalByte := []byte(finalStr)
	os.WriteFile(args[1], finalByte, 0o644)
}

func cap(arr []string) []string {
	for i := 0; i < len(arr); i++ {
		if strings.Contains(arr[i], "(cap") {
			if strings.Contains(arr[i], "(cap,") {
				numArr, _ := strconv.Atoi(arr[i+1][:len(arr[i+1])-1])
				// Loop through and capitalize the words preceding "(cap, int)"
				for j := i - numArr; j < i; j++ {
					arr[j] = strings.Title(arr[j])
				}
				// Purge "(cap, int)"
				arr = append(arr[:i], arr[i+2:]...)
				
			} else if strings.Contains(arr[i], "(cap)") {
				arr[i-1] = strings.Title(arr[i-1])
				// Purge "(cap)"
				arr = append(arr[:i], arr[i+1:]...)
			}
		}
	}
	return arr
}

func low(arr []string) []string {
	for i := 0; i < len(arr); i++ {
		if strings.Contains(arr[i], "(low") {
			if strings.Contains(arr[i], "(low,") {
				numArr, _ := strconv.Atoi(arr[i+1][:len(arr[i+1])-1])

				// Loop through and lowercase the words preceding "(low, int)"
				for j := i - numArr; j < i; j++ {
					arr[j] = strings.ToLower(arr[j])
				}
				// Purge "(low, int)"
				arr = append(arr[:i], arr[i+2:]...)
			} else if strings.Contains(arr[i], "(low)") {
				arr[i-1] = strings.ToLower(arr[i-1])
				// Purge "(low)"
				arr = append(arr[:i], arr[i+1:]...)
			}
		}
	}
	return arr
}

func up(arr []string) []string {
	for i := 0; i < len(arr); i++ {
		if strings.Contains(arr[i], "(up") {
			if strings.Contains(arr[i], "(up,") {
				numArr, _ := strconv.Atoi(arr[i+1][:len(arr[i+1])-1])

				// Loop through and capitalize the words preceding "(up, int)"
				for j := i - numArr; j < i; j++ {
					arr[j] = strings.ToUpper(arr[j])
				}
				// Purge "(up, int)"
				arr = append(arr[:i], arr[i+2:]...)
			} else if strings.Contains(arr[i], "(up)") {
				arr[i-1] = strings.ToUpper(arr[i-1])
				// Purge "(up)"
				arr = append(arr[:i], arr[i+1:]...)
			}
		}
	}
	return arr
}

func hex(arr []string) []string {
	for i := 0; i < len(arr); i++ {
		if strings.Contains(arr[i], "(hex)") {
			deci, _ := strconv.ParseInt(arr[i-1], 16, 64)
			arr[i-1] = strconv.Itoa(int(deci))
			arr = append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}

func bin(arr []string) []string {
	for i := 0; i < len(arr); i++ {
		if arr[i] == "(bin)" {
			binry, _ := strconv.ParseInt(arr[i-1], 2, 16)
			arr[i-1] = strconv.Itoa(int(binry))
			arr = append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}

func punct(arr []string) []string {
	// for _, v := range punctMarks {
	for i := 0; i < len(arr); i++ {

		// As long as there's no empty string, move punctuation marks appearing at the beginning of the second and subsequent words
		if i > 0 && arr[i] != "" {
			// if arr[i][0] == v {
			if unicode.IsPunct(rune(arr[i][0])) {
				if arr[i][0] == '.' || arr[i][0] == ',' || arr[i][0] == '!' || arr[i][0] == '?' || arr[i][0] == ':' || arr[i][0] == ';' || arr[i][0] == '\'' {
					arr[i-1] += string(arr[i][0])
					arr[i] = arr[i][1:]

					// Purge empty strings
					if arr[i] == "" {
						arr = append(arr[:i], arr[i+1:]...)
						i--
					}
				}

			}
		}
		// Re-run function in case the punctuation mark is misplaced in other areas of sentence
		// if arr[i][0] == v {
		if unicode.IsPunct(rune(arr[i][0])) {
			if arr[i][0] == '.' || arr[i][0] == ',' || arr[i][0] == '!' || arr[i][0] == '?' || arr[i][0] == ':' || arr[i][0] == ';' || arr[i][0] == '\'' {
				arr = punct(arr)
			}
		}
	}
	return arr
}

func apostrophe(arr []string) []string {
	// Establish labrary of quotation marks to target
	quotMarks := []byte{39, '"'}

	for _, v := range quotMarks {
		for i := 0; i < len(arr); i++ {
			// in words that end with a quotation mark, move the appostrophe to next word
			if arr[i][len(arr[i])-1] == v {
				if i != len(arr)-1 {
					emptSlc := string(v)
					arr[i] = arr[i][:len(arr[i])-1]
					arr[i+1] = emptSlc + arr[i+1]
					break
				}
			}
		}
	}
	return arr
}

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

		//Check for and correct errors following "A" and "a"
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

		//Check for and correct errors following "An" and "an"
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