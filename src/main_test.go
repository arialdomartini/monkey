package main

import "testing"

func TestPass(t *testing.T) {
    actual := greet()
    expected := "Hello, world!"

    if actual != expected {
        t.Errorf("Got %s but expected %s", actual, expected)
    }
}