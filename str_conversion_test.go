package main

import "testing"

func TestHex(t *testing.T) {
	testSlc := []string{"1E", "(hex)", "files", "were", "added"}
	got := hex(testSlc)
	expected := []string{"30", "files", "were", "added"}

	for i, sg := range got {
		for j, se := range got {
			if i == j && sg != se {
				t.Errorf("Got: %s", got[i])
				t.Errorf("Expected: %s", expected[j])
				t.Errorf("TestHex Failed!")
				t.FailNow()
			}
		}
	}
}
