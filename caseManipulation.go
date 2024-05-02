package main

import (
	"fmt"
	"strconv"
	"strings"
)

func cap(arr []string) []string {
	for i := 0; i < len(arr); i++ {
		if strings.Contains(arr[i], "(cap") {
			if isPunct(arr[i][0]) {
				fmt.Printf("%q can't compute, wrong format\n", arr[i-1]+" "+arr[i])
			} else {
				if strings.Contains(arr[i], "(cap,") {
					numArr, err := strconv.Atoi(arr[i+1][:len(arr[i+1])-1])
					if err != nil {
						fmt.Printf("%q can't compute, wrong format\n", arr[i]+" "+arr[i+1])
					}
					// Loop through and capitalize the words preceding "(cap, int)"
					for j := i - numArr; j < i; j++ {
						if j >= 0 {
							arr[j] = strings.ToUpper(arr[j][:1]) + arr[j][1:]
						} else {
							for j := 0; j < i; j++ {
								arr[j] = strings.ToUpper(arr[j][:1]) + arr[j][1:]
							}
							fmt.Printf("%q is out of range\n", arr[i]+" "+arr[i+1])
							break
						}
					}
					// Purge "(cap, int)"
					arr = append(arr[:i], arr[i+2:]...)

				} else if strings.Contains(arr[i], "(cap)") {
					if arr[i] == "(cap)" {
						arr[i-1] = strings.ToUpper(arr[i-1][:1]) + arr[i-1][1:]
						// Purge "(cap)"
						arr = append(arr[:i], arr[i+1:]...)
					} else {
						fmt.Printf("%q can't compute, wrong format\n", arr[i])
					}
				}
			}
		}
	}
	return arr
}

func low(arr []string) []string {
	for i := 0; i < len(arr); i++ {
		if strings.Contains(arr[i], "(low") {
			if strings.Contains(arr[i], "(low,") {
				numArr, err := strconv.Atoi(arr[i+1][:len(arr[i+1])-1])
				if err != nil {
					fmt.Printf("%q can't compute, wrong format\n", arr[i]+" "+arr[i+1])
				}

				// Loop through and lowercase the words preceding "(low, int)"
				for j := i - numArr; j < i; j++ {
					if j >= 0 {
						arr[j] = strings.ToLower(arr[j])
					} else {
						for j := 0; j < i; j++ {
							arr[j] = strings.ToLower(arr[j])
						}
						fmt.Printf("%q is out of range\n", arr[i]+" "+arr[i+1])
						break
					}
				}
				// Purge "(low, int)"
				arr = append(arr[:i], arr[i+2:]...)
			} else if strings.Contains(arr[i], "(low)") {
				if arr[i] == "(low)" {
					arr[i-1] = strings.ToLower(arr[i-1])
					// Purge "(low)"
					arr = append(arr[:i], arr[i+1:]...)
				} else {
					fmt.Printf("%q cant't compute, wrong format\n", arr[i])
				}
			}
		}
	}
	return arr
}

func up(arr []string) []string {
	for i := 0; i < len(arr); i++ {
		if strings.Contains(arr[i], "(up") {
			if strings.Contains(arr[i], "(up,") {
				numArr, err := strconv.Atoi(arr[i+1][:len(arr[i+1])-1])
				if err != nil {
					fmt.Printf("%q can't compute, wrong format\n", arr[i]+" "+arr[i+1])
				}
				// Loop through and capitalize the words preceding "(up, int)"
				for j := i - numArr; j < i; j++ {
					if j >= 0 {
						arr[j] = strings.ToUpper(arr[j])
					} else {
						for j := 0; j < i; j++ {
							arr[j] = strings.ToUpper(arr[j])
						}
						fmt.Printf("%q is out of range\n", arr[i]+" "+arr[i+1])
						break
					}
				}
				// Purge "(up, int)"
				arr = append(arr[:i], arr[i+2:]...)
			} else if strings.Contains(arr[i], "(up)") {
				if arr[i] == "(up)" {
					arr[i-1] = strings.ToUpper(arr[i-1])
					// Purge "(up)"
					arr = append(arr[:i], arr[i+1:]...)
				} else {
					fmt.Printf("%q can't compute, wrong format\n", arr[i])
				}
			}
		}
	}
	return arr
}
