package main

import "testing"

func TestCap(t *testing.T) {
	// Declare utility variables
	testSlc := []string{"Welcome", "to", "the", "Brooklyn", "bridge", "(cap)"}
	got := cap(testSlc)
	expected := []string{"Welcome", "to", "the", "Brooklyn", "Bridge"}

	// Compare individual elements from 'got' and 'expected' slices
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
	// Declare utility variables
	testSlc := []string{"I", "should", "stop", "SHOUTING", "(low)"}
	got := low(testSlc)
	expected := []string{"I", "should", "stop", "shouting"}

	// Compare individual elements from 'got' and 'expected' slices
	for i, sg := range got {
		for j, se := range got {
			if i == j && sg != se {
				t.Errorf("Got: %s", got[i])
				t.Errorf("Expected: %s", expected[j])
				t.Errorf("TestLow Failed!")
				t.FailNow()
			}
		}
	}
}

func TestUp(t *testing.T) {
	// Declare utility variables
	testSlc := []string{"Ready,", "set,", "go", "(up)", "!"}
	got := up(testSlc)
	expected := []string{"Ready,", "set,", "GO", "!"}

	// Compare individual elements from 'got' and 'expected' slices
	for i, sg := range got {
		for j, se := range got {
			if i == j && sg != se {
				t.Errorf("Got: %s", got[i])
				t.Errorf("Expected: %s", expected[j])
				t.Errorf("TestUp Failed!")
				t.FailNow()
			}
		}
	}
}
