package main

import "testing"

func TestCap(t *testing.T) {
	testSlc := []string{"Welcome", "to", "the", "Brooklyn", "bridge", "(cap)"}
	got := cap(testSlc)
	expected := []string{"Welcome", "to", "the", "Brooklyn", "Bridge"}

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

func TestLow(t *testing.T) {

}
