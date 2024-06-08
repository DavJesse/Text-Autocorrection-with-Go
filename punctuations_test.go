package main

import "testing"

func TestPunct(t *testing.T) {
	testSlc := []string{"I", "was", "sitting", "over", "there", ",and", "then", "BAMM", "!!"}
	got := cap(testSlc)
	expected := []string{"I", "was", "sitting", "over", "there,", "and", "then", "BAMM!!"}

	for i, sg := range got {
		for j, se := range got {
			if i == j && sg != se {
				t.Errorf("Got: %s", got[i])
				t.Errorf("Expected: %s", expected[j])
				t.Errorf("TestCap Failed!")
				t.FailNow()
			}
		}
	}
}
