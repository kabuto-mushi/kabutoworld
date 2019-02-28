package main

import "testing"

func TestBugs(t *testing.T) {

	if false {
		t.Errorf("this should never happen")
	}
}
