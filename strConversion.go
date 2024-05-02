package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hex(arr []string) []string {
	for i := 0; i < len(arr); i++ {
		if strings.Contains(arr[i], "(hex)") {
			if arr[i] == "(hex)" {
				if !isPunct(arr[i-1][0]) {
				deci, err := strconv.ParseInt(arr[i-1], 16, 64)
				if err != nil {
					fmt.Printf("%q is not a valid hexadecimal variable\n", arr[i-1])
					continue
				}
				arr[i-1] = strconv.Itoa(int(deci))
				arr = append(arr[:i], arr[i+1:]...)
			} else {
				continue
			}
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
		if strings.Contains(arr[i], "(bin)") {
			if arr[i] == "(bin)" {
				if !isPunct(arr[i-1][0]) {
					binry, err := strconv.ParseInt(arr[i-1], 2, 16)
					if err != nil {
						fmt.Printf("%q is not a valid binary variable\n", arr[i-1])
						continue
					}
					arr[i-1] = strconv.Itoa(int(binry))
					arr = append(arr[:i], arr[i+1:]...)
				} else {
					continue
				}
				
			} else {
				fmt.Printf("%q can't compute, wrong format\n", arr[i])
				continue
			}
		}
	}
	return arr
}

