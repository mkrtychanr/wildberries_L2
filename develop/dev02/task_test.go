package main

import "testing"

func TestSimple(t *testing.T) {
	got, _ := RLERevert("a4bc2d5e3")
	want := "aaaabccdddddeee"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestError(t *testing.T) {
	_, err := RLERevert("45aaa")
	if err == nil {
		t.Errorf("Expect error but it isn't")
	}
}
