package main

import "testing"

func TestVowels(t *testing.T) {
	testSlc := []string{"There", "it", "was.", "A", "amazing", "rock!"}
	got := vowels(testSlc)
	expected := []string{"There", "it", "was.", "An", "amazing", "rock!"}

	for i, sg := range got {
		for j, se := range got {
			if i == j && sg != se {
				t.Errorf("Got: %s", got[i])
				t.Errorf("Expected: %s", expected[j])
				t.Errorf("TestVowels Failed!")
				t.FailNow()
			}
		}
	}
}
