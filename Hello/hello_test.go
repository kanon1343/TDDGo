package main

import "testing"

func TestHello(t *testing.T){
	got := Hello("kanon")
	want := "Hello, kanon"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}