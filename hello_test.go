package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Guilherme", "English")
		want := "Hello, Guilherme!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Luciana", "Spanish")
		want := "Hola, Luciana!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Luciana", "French")
		want := "Bonjour, Luciana!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		fmt.Printf("Got: %q, Lenght: %d\n", got, len(got))
		fmt.Printf("Want: %q, Lenght: %d\n", want, len(want))
		t.Errorf("got %q want %q", got, want)
	}
}
