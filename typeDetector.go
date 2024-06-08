package main

func isNumeric(s string) bool {
	var status bool

	for _, r := range s {
		if r >= '0' && r <= '9' {
			status = true
		}
	}
	return status
}
