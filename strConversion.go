package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func hex(arr []string) []string {
	for i := 0; i < len(arr); i++ {

		// Identify '(hex)' string and convert the preceding string
		if strings.Contains(arr[i], "(hex)") {
			if arr[i] == "(hex)" {
				deci, err := strconv.ParseInt(arr[i-1], 16, 64)
				if err != nil {
					fmt.Printf("%q is not a valid hexadecimal variable\n", arr[i-1])
					os.Exit(0)
				}
				arr[i-1] = strconv.Itoa(int(deci))

				// Purge '(hex)' string
				arr = append(arr[:i], arr[i+1:]...)

				// Identify incorrectly formated '(bin)' strings
			} else {
				fmt.Printf("%q can't commpute, wrong format\n", arr[i])
				continue
			}
		}
	}
	return arr
}

func bin(arr []string) []string {
	for i := 0; i < len(arr); i++ {

		// Identify '(bin)' string and convert the preceding string
		if strings.Contains(arr[i], "(bin)") {
			if arr[i] == "(bin)" {

				binry, err := strconv.ParseInt(arr[i-1], 2, 16)
				if err != nil {
					fmt.Printf("%q is not a valid binary entry\n", arr[i-1])
					os.Exit(0)
				}
				arr[i-1] = strconv.Itoa(int(binry))

				// Purge '(bin)' string
				arr = append(arr[:i], arr[i+1:]...)

				// Identify incorrectly formated '(bin)' strings
			} else {
				fmt.Printf("%q can't compute, wrong format\n", arr[i])
				continue
			}
		}
	}
	return arr
}
