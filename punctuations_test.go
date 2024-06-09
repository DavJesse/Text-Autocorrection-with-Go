package main

import "testing"

func TestPunct(t *testing.T) {
	// Declare utility variables
	testSlc := []string{"I", "was", "sitting", "over", "there", ",and", "then", "BAMM", "!!"}
	got := punct(testSlc)
	expected := []string{"I", "was", "sitting", "over", "there,", "and", "then", "BAMM!!"}

	// Compare individual elements in 'got' and 'expected' slices
	for i, sg := range got {
		for j, se := range got {
			if i == j && sg != se {
				t.Errorf("Got: %s", got[i])
				t.Errorf("Expected: %s", expected[j])
				t.Errorf("TestPunct Failed!")
				t.FailNow()
			}
		}
	}
}

func TestApostrophe(t *testing.T) {
	// Declare utility variables
	testSlc := []string{"I", "am", "exactly", "how", "they", "describe", "me:'", "awesome'"}
	got := apostrophe(testSlc)
	expected := []string{"I", "am", "exactly", "how", "they", "describe", "me:", "'awesome'"}

	// Compare individual elements in 'got' and 'expected' slices
	for i, sg := range got {
		for j, se := range got {
			if i == j && sg != se {
				t.Errorf("Got: %s", got[i])
				t.Errorf("Expected: %s", expected[j])
				t.Errorf("TestApostrophe Failed!")
				t.FailNow()
			}
		}
	}
}

func TestIsQuote(t *testing.T) {
	// Declare utility variables
	testSlc := []byte{'\'', '"'}
	expected := true
	var got bool

	// Compare 'got' and 'expected'
	for _, r := range testSlc {
		if !isQuote(r) {
			got = isQuote(r)
			t.Errorf("Got: %t", got)
			t.Errorf("Expected: %t", expected)
			t.Errorf("TestIsQuote Failed!")
			t.FailNow()
		}
	}
}
