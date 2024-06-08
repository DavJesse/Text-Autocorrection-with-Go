package main

import (
	"fmt"
	"os"
	"strings"
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
	//strSlce = recursiveCleanUp(strSlce) //re-run the cap, low, up, bin, and hex functions

	// stitch string slice and write updated string to result.txt
	finalStr := strings.Join(strSlce, " ") + "\n"
	finalByte := []byte(finalStr)
	os.WriteFile(args[1], finalByte, 0o644)
}
