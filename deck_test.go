package main

import "testing"

func TestNewDeck(t *testing.T){

	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected the length of deck is 16, but got %v", len(d))
	}
}